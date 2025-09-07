package stack

func AsteroidCollision(asteroids []int) []int {
	ans := make([]int, 0, len(asteroids)) 

	for _, a := range asteroids {
		alive := true

		for alive && a < 0 && len(ans) > 0 && ans[len(ans)-1] > 0 {
			top := ans[len(ans)-1]
			if abs(top) < abs(a) {
				ans = ans[:len(ans)-1]
				continue 
			} else if abs(top) == abs(a) {
				ans = ans[:len(ans)-1]
				alive = false
			} else {
				alive = false
			}
		}

		if alive {
			ans = append(ans, a)
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		x = x * (-1)
	}

	return x
}