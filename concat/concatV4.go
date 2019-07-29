package concat

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

func insertPrefixV4(seq *sequence3, ctx *context3) {
	for i := len(*seq.sequence); i > 0; i-- {
		prefix := (*seq.sequence)[:i]
		ps := &prefixStore3{
			sequences: make([]*sequence3, 1),
			lock:      sync.Mutex{},
		}
		ps.sequences[0] = seq
		actual, loaded := ctx.prefixes.LoadOrStore(prefix, ps)
		if loaded {
			ps = actual.(*prefixStore3)
			ps.lock.Lock()
			ps.sequences = append(ps.sequences, seq)
			ps.lock.Unlock()
		}
	}
}

func linkSequenceV4(seq *sequence3, ctx *context3) {
	// 找到第一个有前缀匹配的，且不是子字符串的
	notFound := true
	prefixLength := 1
	var ps *prefixStore3

	for notFound {
		suffix := (*seq.sequence)[prefixLength:]
		value, exist := ctx.prefixes.Load(suffix)
		for ; !exist && prefixLength < len(*seq.sequence); prefixLength++ {
			suffix = (*seq.sequence)[prefixLength:]
			value, exist = ctx.prefixes.Load(suffix)
		}
		if !exist {
			// 若还是不存在，则这个节点是路径的结束
			seq.next = &path{
				sfxLen: 0,
				dest:   []*sequence3{},
			}
			return
		}
		if prefixLength > 1 {
			// 进入for循环才找到，退出时会加1，这里改回去
			prefixLength--
		}
		ps = value.(*prefixStore3)
		notFound = false
	}

	if ps == nil {
		// 这里不应该的
		fmt.Println("不应该ps是nil")
		os.Exit(1)
	}

	seq.next = &path{
		sfxLen: len(*seq.sequence) - prefixLength,
		dest:   ps.sequences,
	}
}

func ConcatV4(sequences []string) *string {
	ctx := &context3{
		nodes:    sync.Map{},
		prefixes: sync.Map{},
	}

	// 完成计数
	doneCnt := int64(0)

	// 保存到map中，并移除子字符串
	fmt.Print("装载中...")
	for i := 0; i < len(sequences); i++ {
		seq3 := &sequence3{
			sequence: &sequences[i],
			next:     nil,
		}
		ctx.nodes.Store(seq3, true)
	}
	fmt.Println("装载完成")

	fmt.Print("移除子字符串，已删除: 0")
	doneCnt = 0
	fmtLock := sync.Mutex{}
	ctx.nodes.Range(func(key, value interface{}) bool {
		ctx.wg.Add(1)
		seq3 := key.(*sequence3)
		go func(seq *sequence3) {
			defer func() {
				atomic.AddInt64(&doneCnt, 1)
				fmtLock.Lock()
				fmt.Printf("\r移除子字符串，已完成: %d", doneCnt)
				fmtLock.Unlock()
				ctx.wg.Done()
			}()
			ctx.nodes.Range(func(key, value interface{}) bool {
				s2 := key.(*sequence3)
				if s2 == seq3 {
					return true
				}
				if strings.Contains(*seq3.sequence, *s2.sequence) {
					ctx.nodes.Delete(s2)
				} else if strings.Contains(*s2.sequence, *seq3.sequence) {
					ctx.nodes.Delete(seq3)
					return false
				}
				return true
			})
		}(seq3)
		return true
	})
	ctx.wg.Wait()
	fmt.Println()

	fmt.Print("构建前缀表，已完成： 0")
	// 用于控制goroutine数量
	done := make(chan bool, runtime.NumCPU()*10)
	doneCnt = 0
	// 构建前缀表
	ctx.nodes.Range(func(key, value interface{}) bool {
		done <- true
		ctx.wg.Add(1)
		temp := key.(*sequence3)
		go func(seq *sequence3) {
			defer func() {
				atomic.AddInt64(&doneCnt, 1)
				fmt.Printf("\r构建前缀表，已完成： %d", doneCnt)
				ctx.wg.Done()
				<-done
			}()
			insertPrefixV4(seq, ctx)
		}(temp)
		return true
	})
	fmt.Println()
	nodeCnt := doneCnt

	fmt.Println("根据前缀表连接节点...")
	ProgressBar(0, 0, PROGRESS_BEGIN)
	doneCnt = 0
	// 将排除了部分前缀的剩余图节点连接起来，只要前后缀匹配能连接就连起来
	ctx.nodes.Range(func(key, value interface{}) bool {
		ctx.wg.Add(1)
		seq := key.(*sequence3)
		go func(s *sequence3) {
			defer ctx.wg.Done()
			linkSequenceV4(seq, ctx)
			atomic.AddInt64(&doneCnt, 1)
			ProgressBar(int(doneCnt), int(nodeCnt), PROGRESS_UPDATE)
		}(seq)
		return true
	})
	ctx.wg.Wait()
	ProgressBar(0, 0, PROGRESS_FINISH)

	// 原本的Map不需要了，让Go垃圾回收。回收资源以应对接下来的遍历
	ctx.prefixes.Range(func(key, value interface{}) bool {
		ctx.prefixes.Delete(key)
		return true
	})

	doneCnt = 0
	{
		fmt.Println("广度遍历连接...")
		ProgressBar(0, 0, PROGRESS_BEGIN)
	}
	distinctResults := &sync.Map{}
	longestResult := 0
	lrLock := sync.Mutex{}
	ctx.nodes.Range(func(key, value interface{}) bool {
		keySeq := key.(*sequence3)
		ctx.wg.Add(1)
		go func(seq *sequence3) {
			defer func() {
				atomic.AddInt64(&doneCnt, 1)
				ProgressBar(int(doneCnt), int(nodeCnt), PROGRESS_UPDATE)
				ctx.wg.Done()
			}()

			result := breathFirstSearch(seq, ctx, uint64(nodeCnt/4))
			{
				lrLock.Lock()
				for i := 0; i < len(result); i++ {
					if longestResult < len(*result[i]) {
						longestResult = len(*result[i])
					}
				}
				lrLock.Unlock()
			}

			// 与前面找到的结果相比较，查看是否有子字符串的，有的话就去掉。若没有，则加入到结果集中
			for i := 0; i < len(result); i++ {
				oneResult := result[i]
				// 仅在长度够的情况下插入
				if longestResult>>1 < len(*oneResult) {
					isSubString := false
					distinctResults.Range(func(key, value interface{}) bool {
						s := key.(*string)
						if strings.Contains(*s, *oneResult) {
							//这是原本的结果中包含result串的情况，不需要继续插入了
							isSubString = true
							return false
						} else if strings.Contains(*oneResult, *s) {
							distinctResults.Delete(s)
						}
						return true
					})

					if !isSubString {
						distinctResults.Store(oneResult, true)
					}
				}
			}
		}(keySeq)

		return true
	})
	ctx.wg.Wait()
	ProgressBar(0, 0, PROGRESS_FINISH)

	fmt.Println("连接完成，检查结果正确性...")
	doneCnt = 0
	distinctResults.Range(func(key, value interface{}) bool {
		str := key.(*string)
		ctx.wg.Add(1)
		go func(s *string) {
			defer func() {
				atomic.AddInt64(&doneCnt, 1)
				fmtLock.Lock()
				fmt.Printf("\r已完成 %d", doneCnt)
				fmtLock.Unlock()
				ctx.wg.Done()
			}()
			for j := 0; j < len(sequences); j++ {
				if !strings.Contains(*s, sequences[j]) {
					distinctResults.Delete(s)
					return
				}
			}
		}(str)
		return true
	})
	ctx.wg.Wait()
	fmt.Println()

	results := make([]*string, 0)
	var shortest *string
	distinctResults.Range(func(key, value interface{}) bool {
		s := key.(*string)
		results = append(results, s)
		if shortest == nil {
			shortest = s
		} else if len(*shortest) > len(*s) {
			shortest = s
		}
		return true
	})

	if len(results) == 0 {
		fmt.Println("错误，没有结果")
		return new(string)
	}
	fmt.Println("构造完成")
	return shortest
}
