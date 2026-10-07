package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdCount := 0
	for _, v := range birdsPerDay {
		birdCount += v
	}
	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	endCount := week * 7
	startCount := endCount - 7
	weekCount := 0
	for i, v := range birdsPerDay {
		if i >= startCount && i <= endCount {
			weekCount += v
		}
	}
	return weekCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	fixedBirdsPerDay := birdsPerDay
	for i, v := range birdsPerDay {
		if i%2 == 0 {
			fixedBirdsPerDay[i] = v + 1
		}
	}
	return fixedBirdsPerDay
}
