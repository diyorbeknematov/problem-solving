package arraysstrings

func CanPlaceFlowers(flowerbed []int, n int) bool {
	res := 0

	for i := 0; i < len(flowerbed); i++ {
		if i == 0 && i == len(flowerbed)-1 && flowerbed[i] == 0 {
			res++
			flowerbed[i] = 1
		}
		if (i == 0 && i < len(flowerbed)-1) && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			res++
			flowerbed[i] = 1
		}
		if (i > 0 && i < len(flowerbed)-1) && flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			res++
			flowerbed[i] = 1
		}
		if (i > 0 && i == len(flowerbed)-1) && flowerbed[i] == 0 && flowerbed[i-1] == 0 {
			res++
			flowerbed[i] = 1
		}
	}

	return res >= n
}
