package analysis_test

import (
	"math"
	"testing"

	"github.com/sciack/simulation/analysis"
)

func TestAverage(t *testing.T) {
	samples := []int{1, 2, 3}
	result := analysis.ComputeStats(&samples)
	if result.Average != 2 {
		t.Errorf("Expecting 2 got %f", result.Average)
	}
}

func TestVariance(t *testing.T) {
	samples := []int{1, 2, 3}
	result := analysis.ComputeStats(&samples)
	if math.Abs(result.Variance-0.6666666) > 0.0001 {
		t.Errorf("Expecting 1] got %f", result.Variance)
	}
}
