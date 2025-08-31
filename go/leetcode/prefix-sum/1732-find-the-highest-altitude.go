package prefixsum

func LargestAltitude(gain []int) int {
	currentAltitude := 0
	maxAltitude := 0

	for _, g := range gain {
		currentAltitude += g

		maxAltitude = max(maxAltitude, currentAltitude)
	}

	return maxAltitude
}
