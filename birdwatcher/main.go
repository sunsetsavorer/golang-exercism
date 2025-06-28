package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {

	sum := 0

	for _, count := range birdsPerDay {
		sum += count
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	result := 0
	start := (week - 1) * 7
	end := week * 7

	for day := start; day < end; day++ {
		result += birdsPerDay[day]
	}

	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}
