package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 处理边界情况
	if nums1 == nil || len(nums1) == 0 {
		return float64(nums2[(len(nums2)-1)/2]+nums2[len(nums2)/2]) / 2
	}
	if nums2 == nil || len(nums2) == 0 {
		return float64(nums1[(len(nums1)-1)/2]+nums1[len(nums1)/2]) / 2
	}

	// 确保m小于等于n
	if len(nums2) < len(nums1) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}
	m := len(nums1)
	n := len(nums2)
	halfLen := (m + n + 1) / 2
	iMin := 0
	iMax := m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else {
			// 找到了i
			// 找左边最大与右边最小
			var leftMax int
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					leftMax = nums1[i-1]
				} else {
					leftMax = nums2[j-1]
				}
			}
			if (m+n)&1 == 1 {
				// 奇数则直接返回
				return float64(leftMax)
			}

			var rightMin int
			if i == m {
				rightMin = nums2[j]
			} else if j == n {
				rightMin = nums1[i]
			} else {
				if nums1[i] < nums2[j] {
					rightMin = nums1[i]
				} else {
					rightMin = nums2[j]
				}
			}
			return float64(leftMax+rightMin) / 2
		}
	}
	return 0
}
