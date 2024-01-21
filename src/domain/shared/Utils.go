package domain

func IsInRange(num, min, max int) bool {
	return num >= min && num <= max
}

func HumidityIsHigh(num, max int) bool {
	return num > max
}
