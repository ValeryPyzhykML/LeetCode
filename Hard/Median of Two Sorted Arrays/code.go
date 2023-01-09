package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	allLen := len(nums1) + len(nums2)
	odd := allLen % 2
	margin := (allLen - 2 + odd) / 2

	if odd == 0 {
		return findMedianSortedArraysEven(nums1, nums2, margin)
	}

	return findMedianSortedArraysOdd(nums1, nums2, margin)
}

func findMedianSortedArraysEven(nums1 []int, nums2 []int, margin int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[margin]+nums2[margin+1]) / 2
	}
	if len(nums2) == 0 {
		return float64(nums1[margin]+nums1[margin+1]) / 2
	}
	if len(nums1) == 1 && len(nums2) == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}

	margin++
	c1, c2 := 0, 0

	nums1_last := len(nums1) - 1
	nums2_last := len(nums2) - 1

	var arr []int
	var firstBig int

	arraysNorEqual := len(nums1) != len(nums2)
	isBig1 := false

	if nums1[0] > nums2[0] {
		arr = nums2
		firstBig = nums1[0]
		isBig1 = true
	} else {
		arr = nums1
		firstBig = nums2[0]
	}

	i := 0
	for i+1 < len(arr) && arr[i+1] <= firstBig {
		i++
		if i == margin {
			return float64((arr[i-1] + arr[i])) / 2
		}
		if arraysNorEqual && i+1 == len(arr) {
			if isBig1 {
				return findMedianSortedArraysEvenCompute(nums2, nums1)
			} else {
				return findMedianSortedArraysEvenCompute(nums1, nums2)
			}
		}
	}

	if isBig1 {
		c2 = i
	} else {
		c1 = i
	}

	last := firstBig
	prev := arr[i]
	i++

	for i < margin {
		if c1 == nums1_last {
			c2++
			if arraysNorEqual && nums2[c2] >= nums1[c1] {
				return findMedianSortedArraysEvenCompute(nums1, nums2)
			}
			prev = last
			last = nums2[c2]
		} else if c2 == nums2_last {
			c1++
			if arraysNorEqual && nums1[c1] >= nums2[c2] {
				return findMedianSortedArraysEvenCompute(nums2, nums1)
			}
			prev = last
			last = nums1[c1]
		} else if nums1[c1+1] == nums2[c2+1] {
			if nums1[c1] > nums2[c2] {
				c2++
				prev = last
				last = nums2[c2]
			} else {
				c1++
				prev = last
				last = nums1[c1]
			}
		} else if nums1[c1+1] > nums2[c2+1] {
			c2++
			prev = last
			last = nums2[c2]
		} else {
			c1++
			prev = last
			last = nums1[c1]
		}

		i++
	}

	return float64(prev+last) / 2
}

func findMedianSortedArraysEvenCompute(short []int, long []int) float64 {
	last := short[len(short)-1]
	rightOfPair := +(len(long)+len(short))/2 - len(short)

	if rightOfPair-1 < 0 || last > long[rightOfPair-1] {
		return float64(last+long[rightOfPair]) / 2
	}
	return float64(long[rightOfPair-1]+long[rightOfPair]) / 2
}

func findMedianSortedArraysOddCompute(short []int, long []int) float64 {
	rightOfPair := +(len(long)+len(short))/2 - len(short)
	return float64(long[rightOfPair])
}

func findMedianSortedArraysOdd(nums1 []int, nums2 []int, margin int) float64 {
	if len(nums1) == 0 {
		return float64(nums2[margin])
	}
	if len(nums2) == 0 {
		return float64(nums1[margin])
	}

	c1, c2 := 0, 0

	nums1_last := len(nums1) - 1
	nums2_last := len(nums2) - 1
	i := 0
	last := 0

	var arr []int
	var firstBig int

	is_not_equal := len(nums1) != len(nums2)
	isBig1 := false
	if nums1[0] > nums2[0] {
		arr = nums2
		firstBig = nums1[0]
		isBig1 = true
	} else {
		arr = nums1
		firstBig = nums2[0]
	}

	for i+1 < len(arr) && arr[i+1] <= firstBig {
		i++
		if i == margin {
			return float64(arr[i])
		}
		if is_not_equal && i+1 == len(arr) {
			if isBig1 {
				return findMedianSortedArraysOddCompute(nums2, nums1)
			} else {
				return findMedianSortedArraysOddCompute(nums1, nums2)
			}
		}
	}

	if isBig1 {
		c2 = i
	} else {
		c1 = i
	}

	last = firstBig
	i++

	for i < margin {
		if c1 == nums1_last {
			c2++
			if is_not_equal && nums2[c2] >= nums1[c1] {
				return findMedianSortedArraysOddCompute(nums1, nums2)
			}
			last = nums2[c2]
		} else if c2 == nums2_last {
			c1++
			if is_not_equal && nums1[c1] >= nums2[c2] {
				return findMedianSortedArraysOddCompute(nums2, nums1)
			}
			last = nums1[c1]
		} else if nums1[c1+1] == nums2[c2+1] {
			if nums1[c1] > nums2[c2] {
				c2++
				last = nums2[c2]
			} else {
				c1++
				last = nums1[c1]
			}
		} else if nums1[c1+1] > nums2[c2+1] {
			c2++
			last = nums2[c2]
		} else {
			c1++
			last = nums1[c1]
		}
		i++
	}
	return float64(last)
}
