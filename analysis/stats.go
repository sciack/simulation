package analysis

import "math"

type Confidence struct {
	Min float64
	Max float64
}

type Stats struct {
	Average      float64
	Variance     float64
	StdDev       float64
	Confidence95 Confidence
}

func ComputeStats(samples *[]int) Stats {
	average := avg(samples)
	variance := variance(samples, average)
	stdDev := math.Sqrt(variance)
	return Stats{
		Average:  average,
		Variance: variance,
		StdDev:   stdDev,
		Confidence95: Confidence{
			Min: average - (1.96 * stdDev),
			Max: average + (1.96 * stdDev),
		},
	}
}

func avg(samples *[]int) float64 {
	sum := 0
	for _, value := range *samples {
		sum += value
	}
	return float64(sum) / float64(len(*samples))
}

func variance(samples *[]int, avg float64) float64 {
	sum := 0.0
	for _, value := range *samples {
		sum += math.Pow((float64(value) - avg), 2)
	}
	return sum / float64(len(*samples))
}
