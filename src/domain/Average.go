package domain

// Média Móvel
func CalculateMovingAverage(readings []float64, window int) []float64 {
	movingAverage := make([]float64, len(readings))
	for i := 0; i < len(readings); i++ {
		start := Max(0, i-window+1)
		sum := 0.0
		count := 0
		for j := start; j <= i; j++ {
			sum += readings[j]
			count++
		}
		movingAverage[i] = sum / float64(count)
	}
	return movingAverage
}

// Média Exponencial
func CalculateExponentialAverage(readings []float64, smoothingFactor float64) []float64 {
	exponentialAverage := make([]float64, len(readings))
	exponentialAverage[0] = readings[0]
	for i := 1; i < len(readings); i++ {
		exponentialAverage[i] = smoothingFactor*readings[i] + (1-smoothingFactor)*exponentialAverage[i-1]
	}
	return exponentialAverage
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
