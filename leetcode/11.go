package leetcode

/*
双端指针法。每次移动较短的那一边。解释，假设左边为p，右边为q，p较短。若移动较长的一边，假设height[q-1]大于height[p]，面积是减少的因为p还是短。
而假设height[q-1]小于height[p]的话，那么面积就建的更小了。因此，其实可以看出，我们无需再考虑p位置不变，q往左移动的情况了，因为面积只会变小。
而如果是移动较短的一边，由于可能会有更长的边出现，因此面积会有进一步提升的空间。
如果两边高度一样，则移动到下一个边更长的一边，保证能够有提升。如果下一边都一样长，那选择哪一边都一样，因为再继续的话，只会得到更小的面积，而不会再变大。
*/
func maxArea(height []int) int {
	p := 0
	q := len(height) - 1
	result := 0
	for p < q {
		area := 0
		if height[p] < height[q] {
			area = height[p] * (q - p)
			p++
		} else if height[q] < height[p] {
			area = height[q] * (q - p)
			q--
		} else {
			// 两端相等时，移向下一个更大的那个
			area = height[q] * (q - p)
			if height[p+1] > height[q-1] {
				p++
			} else {
				q--
			}
		}
		if area > result {
			result = area
		}
	}
	return result
}
