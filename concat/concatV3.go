package concat

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type path struct {
	// 后缀长度
	sfxLen int
	dest   []*sequence3
}

type sequence3 struct {
	sequence *string
	// 本字符串最长后缀能够匹配的前缀的路径
	next *path
}

type prefixStore3 struct {
	sequences []*sequence3
	lock      sync.Mutex
}

type context3 struct {
	// 不包含前缀和后缀子字符串的节点集合，但是会包含中间的子字符串
	nodes sync.Map
	// 前缀表
	prefixes sync.Map
	wg       sync.WaitGroup
}

func printContext3(ctx *context3) {
	fmt.Println("==================prefixes==================")
	ctx.prefixes.Range(func(key, value interface{}) bool {
		ps := value.(*prefixStore3)
		fmt.Print(key, ":")
		for i := 0; i < len(ps.sequences); i++ {
			fmt.Print(*ps.sequences[i].sequence, " ")
		}
		fmt.Println()
		return true
	})

	fmt.Println("==================nodes====================")
	ctx.nodes.Range(func(key, value interface{}) bool {
		seq := key.(*sequence3)
		fmt.Print(*seq.sequence, " ", seq.next.sfxLen, " ")
		for i := 0; i < len(seq.next.dest); i++ {
			fmt.Print(*seq.next.dest[i].sequence, " ")
		}
		fmt.Println()
		return true
	})
}

func printPath(path []*sequence3) {
	fmt.Println("==========================================================")
	fmt.Println("path length:", len(path))
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print(*path[i].sequence, " ")
	}
	fmt.Println()

	fmt.Print(*path[len(path)-1].sequence)
	for i := len(path) - 2; i >= 0; i-- {
		fmt.Print((*path[i].sequence)[path[i+1].next.sfxLen:])
	}
	fmt.Println()
}

func outputToFile(ctx *context3) {
	file, _ := os.Create("context3.txt")
	writer := bufio.NewWriter(file)

	outputs := make([]string, 0)
	ctx.prefixes.Range(func(key, value interface{}) bool {
		prefix := key.(string)
		seqs := value.(*prefixStore3)
		output := prefix + " : "
		for i := 0; i < len(seqs.sequences); i++ {
			output += *seqs.sequences[i].sequence + " "
		}
		output += "\n"
		outputs = append(outputs, output)
		return true
	})
	sort.Strings(outputs)
	for i := 0; i < len(outputs); i++ {
		writer.WriteString(outputs[i])
	}
	file.Close()
}

func insertPrefix(str *string, ctx *context3) {
	// 通知结束
	defer ctx.wg.Done()

	_, ok := ctx.prefixes.Load(*str)
	if ok {
		// 子字符串，不需要插入
		return
	}

	seq3 := &sequence3{
		sequence: str,
		next:     nil,
	}

	ctx.nodes.Store(seq3, true)

	for i := len(*str); i > 0; i-- {
		prefix := (*str)[:i]
		ps := &prefixStore3{
			sequences: make([]*sequence3, 1),
			lock:      sync.Mutex{},
		}
		ps.sequences[0] = seq3
		actual, loaded := ctx.prefixes.LoadOrStore(prefix, ps)
		if loaded {
			ps = actual.(*prefixStore3)
			inserted := false
			ps.lock.Lock()
			for j := 0; j < len(ps.sequences); j++ {
				target := ps.sequences[j]
				// 检查target与seq的前缀关系
				if len(*target.sequence) < len(*seq3.sequence) && *target.sequence == (*seq3.sequence)[:len(*target.sequence)] {
					// target是seq的前缀
					ps.sequences[j] = seq3
					// 删除target，因为是子字符串
					ctx.nodes.Delete(target)
					inserted = true
					break
				} else if len(*target.sequence) >= len(*seq3.sequence) && (*target.sequence)[:len(*seq3.sequence)] == *seq3.sequence {
					// seq是target的前缀，也因此删除
					ctx.nodes.Delete(seq3)
					inserted = true
					break
				}
			}
			if !inserted {
				ps.sequences = append(ps.sequences, seq3)
			}
			ps.lock.Unlock()
		}
	}
}

func linkSequence(seq *sequence3, ctx *context3) {
	defer ctx.wg.Done()

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
		// 检查是否有子字符串，若有，则删除
		ps.lock.Lock()
		// 这里无需检查ps是否还在这map中，因为如果是删除，则必定ps.sequences的长度为0
		for i := 0; i < len(ps.sequences); i++ {
			// 判断整条字符串是否是suffix，若是，则说明这个字符串是后缀，我们不能连接
			if *ps.sequences[i].sequence == suffix {
				// 同时需要删除nodes
				ctx.nodes.Delete(ps.sequences[i])
				ps.sequences = append(ps.sequences[0:i], ps.sequences[i+1:]...)
				i--
			}
		}
		if len(ps.sequences) == 0 {
			// 若没有了，则删除，并继续
			ctx.prefixes.Delete(suffix)
		} else {
			// 还剩下的话，就说明找到了最长的
			notFound = false
		}
		ps.lock.Unlock()
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

// 广度优先遍历寻找下一个数组
func breathFirstSearch(seq *sequence3, ctx *context3, queueInitialLength uint64) []*string {
	// 每次广度遍历使用的参数
	type bfsArg struct {
		// 本次处理的seq
		seq3 *sequence3
		// 连接了seq3之后的字符串
		concat string
	}

	// itemQueue 为保存了待处理seq3的队列
	// concatQueue 保存了处理对应seq3时候使用的字符串
	// concat 连接了seq之后的字符串
	// 返回seq是否为路径结束点。如果seq3无法再添加子节点，就是路径的结束点了
	addNextToQueue := func(seq3 *sequence3, itemQueue *CasQueue, concat *string) (isEnd bool) {
		if seq3.next == nil || len(seq3.next.dest) == 0 {
			return true
		}
		isEnd = true
		for i := 0; i < len(seq3.next.dest); i++ {
			next := seq3.next.dest[i]
			// 不能是没有在nodes中的
			_, contain := ctx.nodes.Load(next)
			if !contain {
				continue
			}

			// 不能包含这个字符串
			if strings.Contains(*concat, *next.sequence) {
				// 若是因为是里面的子字符串而退出，则不能让cyclic为true
				continue
			}

			// 检查通过，加入到队列中
			isEnd = false
			newConcat := *concat + (*next.sequence)[seq3.next.sfxLen:]
			arg := &bfsArg{
				seq3:   next,
				concat: newConcat,
			}
			itemQueue.Put(arg)
		}
		return isEnd
	}

	t := time.Duration(1)
	itemQueue := NewQueue(queueInitialLength, t)
	// 放入初始后续节点
	for i := 0; i < len(seq.next.dest); i++ {
		arg := &bfsArg{
			seq3:   seq.next.dest[i],
			concat: *seq.sequence + (*seq.next.dest[i].sequence)[seq.next.sfxLen:],
		}
		itemQueue.Put(arg)
	}

	results := make([]*string, 0)
	item, ok, _ := itemQueue.Get()
	for ok {
		arg := item.(*bfsArg)
		isEnd := addNextToQueue(arg.seq3, itemQueue, &arg.concat)
		if isEnd {
			// cur是路径结束点，且不是因为产生回路而结束，遍历完成，添加结果
			results = append(results, &arg.concat)
		}
		item, ok, _ = itemQueue.Get()
	}
	return results
}

func ConcatV3(sequences []string, minimumLength int) *string {
	ctx := &context3{
		nodes:    sync.Map{},
		prefixes: sync.Map{},
		wg:       sync.WaitGroup{},
	}

	// 完成计数
	doneCnt := int64(0)

	fmt.Println("构建前缀表...")
	done := make(chan bool, runtime.NumCPU()*10)
	ProgressBar(0, len(sequences), PROGRESS_BEGIN)
	// 构建前缀表
	for i := 0; i < len(sequences); i++ {
		done <- true
		ctx.wg.Add(1)
		go func(index int) {
			insertPrefix(&sequences[index], ctx)
			<-done
			atomic.AddInt64(&doneCnt, 1)
			ProgressBar(int(doneCnt), len(sequences), PROGRESS_UPDATE)
		}(i)
	}
	ctx.wg.Wait()
	ProgressBar(0, 0, PROGRESS_FINISH)

	fmt.Println("根据前缀表连接节点...")
	fmt.Print("已完成 0 个节点")
	doneCnt = 0
	// 将排除了部分前缀的剩余图节点连接起来，只要前后缀匹配能连接就连起来
	ctx.nodes.Range(func(key, value interface{}) bool {
		ctx.wg.Add(1)
		seq := key.(*sequence3)
		go func(s *sequence3) {
			linkSequence(seq, ctx)
			atomic.AddInt64(&doneCnt, 1)
			fmt.Print("\r已完成", doneCnt, "个节点")
		}(seq)
		return true
	})
	ctx.wg.Wait()
	fmt.Println()

	// 原本的Map不需要了，让Go垃圾回收。回收资源以应对接下来的遍历
	ctx.prefixes.Range(func(key, value interface{}) bool {
		ctx.prefixes.Delete(key)
		return true
	})

	fmt.Print("计算除去前缀和后缀子字符串的剩余字符串数量...")
	nodeCnt := 0
	ctx.nodes.Range(func(key, value interface{}) bool {
		nodeCnt++
		return true
	})
	fmt.Println(nodeCnt, "条")

	//fmt.Print("移除子字符串，已删除: 0")
	//doneCnt = 0
	//fmtLock := sync.Mutex{}
	//ctx.nodes.Range(func(key, value interface{}) bool {
	//	ctx.wg.Add(1)
	//	seq3 := key.(*sequence3)
	//	go func(seq *sequence3) {
	//		defer func() {
	//			atomic.AddInt64(&doneCnt, 1)
	//			fmtLock.Lock()
	//			fmt.Printf("\r移除子字符串，已删除: %d", doneCnt)
	//			fmtLock.Unlock()
	//			ctx.wg.Done()
	//		}()
	//		ctx.nodes.Range(func(key, value interface{}) bool {
	//			s2 := key.(*sequence3)
	//			if s2 == seq3 {
	//				return true
	//			}
	//			if strings.Contains(*seq3.sequence, *s2.sequence) {
	//				atomic.AddInt64(&doneCnt, 1)
	//				ctx.nodes.Delete(s2)
	//			} else if strings.Contains(*s2.sequence, *seq3.sequence) {
	//				ctx.nodes.Delete(seq3)
	//				atomic.AddInt64(&doneCnt, 1)
	//				return false
	//			}
	//			return true
	//		})
	//	}(seq3)
	//	return true
	//})
	//ctx.wg.Wait()
	//fmt.Println()

	//fmt.Print("计算剩余字符串数量...")
	//nodeCnt = 0
	//ctx.nodes.Range(func(key, value interface{}) bool {
	//	nodeCnt++
	//	return true
	//})
	//fmt.Println(nodeCnt, "条")

	doneCnt = 0
	{
		fmt.Println("广度遍历连接...")
		ProgressBar(0, nodeCnt, PROGRESS_BEGIN)
	}
	distinctResults := make([]*string, 0)
	drRWLock := sync.RWMutex{}
	longestResult := 0
	lrLock := sync.Mutex{}
	ctx.nodes.Range(func(key, value interface{}) bool {
		keySeq := key.(*sequence3)
		ctx.wg.Add(1)
		go func(seq *sequence3) {
			defer func() {
				ctx.wg.Done()
				atomic.AddInt64(&doneCnt, 1)
				ProgressBar(int(doneCnt), nodeCnt, PROGRESS_UPDATE)
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
					drRWLock.RLock()
					hasSubString := false
					for i := 0; i < len(distinctResults); i++ {
					restart:
						s := distinctResults[i]
						if strings.Contains(*s, *oneResult) {
							//这是原本的结果中包含result串的情况，不需要继续插入了
							hasSubString = true
							drRWLock.RUnlock()
							break
						} else if strings.Contains(*oneResult, *s) {
							drRWLock.RUnlock()
							drRWLock.Lock()
							if s != distinctResults[i] {
								//防止s被上个持有锁的改了
								drRWLock.Unlock()
								//重新获取回读锁
								drRWLock.RLock()
								goto restart
							}
							distinctResults[i] = oneResult
							hasSubString = true
							drRWLock.Unlock()
							break
						}
					}
					if !hasSubString {
						drRWLock.RUnlock()
						drRWLock.Lock()
						distinctResults = append(distinctResults, oneResult)
						drRWLock.Unlock()
					}
				}
			}
		}(keySeq)

		return true
	})
	ctx.wg.Wait()
	ProgressBar(0, 0, PROGRESS_FINISH)

	fmt.Println("连接完成，得到结果", len(distinctResults), "条，检查结果正确性...")
	ProgressBar(0, len(distinctResults), PROGRESS_BEGIN)
	doneCnt = 0
	// 验证结果
	for i := 0; i < len(distinctResults); i++ {
		ctx.wg.Add(1)
		go func(index int) {
			defer ctx.wg.Done()
			defer func() {
				atomic.AddInt64(&doneCnt, 1)
				ProgressBar(int(doneCnt), len(distinctResults), PROGRESS_UPDATE)
			}()
			for j := 0; j < len(sequences); j++ {
				if !strings.Contains(*distinctResults[index], sequences[j]) {
					// match != -1可以让这个routine在检查到已经有结果的时候中途退出
					distinctResults[index] = nil
					return
				}
			}
		}(i)
	}
	ctx.wg.Wait()
	ProgressBar(0, 0, PROGRESS_FINISH)

	// 删除不符合的
	retCnt := 0
	shortestLength := math.MaxInt64
	shortestIndex := -1

	for i := 0; i < len(distinctResults); i++ {
		if distinctResults[i] == nil {
			continue
		}
		distinctResults[retCnt] = distinctResults[i]
		if len(*distinctResults[i]) < shortestLength {
			shortestLength = len(*distinctResults[i])
			shortestIndex = retCnt
		}
		retCnt++
	}
	distinctResults = distinctResults[:retCnt]
	fmt.Println("得到符合结果", retCnt, "条")

	if len(distinctResults) == 0 {
		fmt.Println("错误，没有结果")
		return new(string)
	}
	return distinctResults[shortestIndex]
}
