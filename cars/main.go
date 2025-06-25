package cars

const (
	MINUTES_IN_HOUR = 60
	SCOPE_COST      = 95000
	UNIT_COST       = 10000
)

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {

	checkValidityOfSuccessRate(successRate)

	productionPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	productionPerMinute := int(float64(productionPerHour) / MINUTES_IN_HOUR)

	return productionPerMinute
}

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

	checkValidityOfSuccessRate(successRate)

	return float64(productionRate) * (successRate / 100)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	scopesCount := carsCount / 10
	unitCount := carsCount % 10

	totalCost := uint(scopesCount*SCOPE_COST + unitCount*UNIT_COST)

	return totalCost
}

func checkValidityOfSuccessRate(successRate float64) {

	if successRate < 0 || successRate > 100 {
		panic("There is not valid success rate")
	}
}

// For another solution by clamping Success Rate
func ClampSuccessRate(successRate float64) float64 {

	switch {
	case successRate < 0:
		return 0
	case successRate > 100:
		return 100
	default:
		return successRate
	}
}
