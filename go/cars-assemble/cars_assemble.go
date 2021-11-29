package cars

import "fmt"

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	switch speed {
	case 0:
		return 0
	case 1, 2, 3, 4:
		return 1
	case 5, 6, 7, 8:
		return 0.9
	case 9, 10:
		return 0.77
	default:
		fmt.Println("The success rate could not be determined")
		return 1
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	productionRatePerHour := SuccessRate(speed) * (float64(speed * 221))
	return productionRatePerHour
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	productionRatePerMinute := ((float64((speed * 221)) * SuccessRate(speed)) / float64(60))
	return int(productionRatePerMinute)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	hourlyProductionRate := (float64((speed * 221)) * SuccessRate(speed))
	if hourlyProductionRate >= limit {
		hourlyProductionRate = limit
		return hourlyProductionRate
	}
	return hourlyProductionRate
}
