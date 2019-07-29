package concat

import (
	"fmt"
	"sync"
)

const PROGRESS_BEGIN = 0
const PROGRESS_UPDATE = 1
const PROGRESS_FINISH = 2

var lastPercent = 0
var barLock = sync.Mutex{}

func ProgressBar(current, total int, op int) {
	barLock.Lock()
	defer barLock.Unlock()
	barWidth := 50
	// 初始化
	if op == PROGRESS_BEGIN {
		lastPercent = 0
		fmt.Print("[-                                                 ] 0 %")
		return
	}
	if current > total {
		fmt.Print("\r[==================================================]100%")
		return
	}
	if op == PROGRESS_FINISH {
		lastPercent = 100
		fmt.Println("\r[==================================================]100%")
		return
	}
	// 计算百分比，并在百分比不变的情况下退出
	percent := int(float64(current) / float64(total) * 100)
	if percent <= lastPercent {
		return
	}
	lastPercent = percent

	// 输出新的进度条
	fmt.Print("\r[")
	buf := make([]byte, 50)
	doubleBar := percent / (100 / barWidth)
	for i := 0; i < doubleBar; i++ {
		buf[i] = '='
	}
	if doubleBar < barWidth {
		buf[doubleBar] = '-'
		for i := doubleBar + 1; i < barWidth; i++ {
			buf[i] = ' '
		}
	}
	fmt.Print(string(buf))
	fmt.Print("]", percent, "%")
}
