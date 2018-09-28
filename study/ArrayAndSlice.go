package study

import "fmt"

func array() {
	// 创建数组
	var array [10]int
	// 创建切片
	s := array[:]
	fmt.Printf("array类型：%T\n", array)
	fmt.Printf("slice类型：%T\n", s)

	s = s[1:8]
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))

	// 我们无法这样子增加切片的长度
	s = s[:]
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))

	// 使用append使切片边长，然而底层数组不变，因为len<cap
	s = append(s, 0)
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))

	// 或者这样，切片是可以把后面的设置成超过len的
	s = s[:len(s)+1]
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))

	// 那么可不可以设置len的元素呢？是不可以的，会报错
	//s = s[:len(s) - 1]
	//s[len(s)] = 1

	// s = s[-1:0]是不可以的，因为必须是非负数
	// 前面的元素，如果没有底层的array，就永远没法访问了

	// 再这样就会出错了，因为超过了cap
	//s = s[:len(s) + 1]
	//fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))

	// 这时候需要用到append自动变长cap
	s = append(s, 0, 0)
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))
	// 看到这里增长了一倍的长度。

	// append函数将元素放到切片的末尾后位置，就是len(s) + 1等后面的位置。这个也是最安全的方法了。
	array[0] = 0
	s = array[0:1]
	s = append(s, 1, 2)
	fmt.Println(s)
	fmt.Printf("len(s):%d cap(s):%d\n", len(s), cap(s))
}
