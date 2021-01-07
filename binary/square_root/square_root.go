package square_root

import "fmt"

func SquareRoot(n int64) int64 {
	if n == 0 ||n==1{
		return n
	}
	var left,right = int64(1),n

	for left <= right{
		mid :=(right+left)/2
		fmt.Println(left, right, mid)
		if mid * mid >n{
			right = mid - 1
		}else {
			left = mid +1
		}
	}

	return right
}