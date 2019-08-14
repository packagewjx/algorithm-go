package leetcode

type stack struct {
	arr []int
}

func (s *stack) push(num int) {
	s.arr = append(s.arr, num)
}

func (s *stack) pop() int {
	temp := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return temp
}

func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *stack) peek() int {
	return s.arr[len(s.arr)-1]
}

func validateStackSequences(pushed []int, popped []int) bool {
	st := &stack{arr: make([]int, 0, len(pushed)/2)}
	popIndex := 0
	for i := 0; i < len(pushed); i++ {
		st.push(pushed[i])
		for ; popIndex < len(popped); popIndex++ {
			if st.isEmpty() || st.peek() != popped[popIndex] {
				break
			}
			st.pop()
		}
	}
	return popIndex == len(popped)
}
