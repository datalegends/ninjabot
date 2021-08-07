package series

import "fmt"

// MinMax TODO use talib min max..
func (s Series) MinMax() (float64, float64) {

	min := s[0]
	max := s[0]

	for i := range s {
		if s[i] > max {
			max = s[i]
		}
		if s[i] < min {
			min = s[i]
		}
	}
	return min, max
}

// Distribution return periodic n buckets of within min max
// values in given series in order to analyze the value range.
// [min, m, m, m, max] minumum range.
// [n, n, n, n, n] value counts in ranges
func (s Series) Distribution(bucketCount int) ([]float64, []float64) {
	buckets := make([]float64, bucketCount)
	counts := make([]float64, bucketCount)

	// return empty if using wrong ranges
	if bucketCount > len(s) {
		return []float64{}, []float64{}
	}

	min, max := s.MinMax()

	delta := (max - min) / float64(bucketCount)

	if delta <= 0 {
		return []float64{}, []float64{}
	}

	// divide delta in bucket ranges
	for step := range buckets {
		buckets[step] = min + float64(step)*delta
	}

	// check in which buckets each value fits.
	for _, v := range s {
		for i, minlimit := range buckets {
			// deal with last bucket egde case.
			upperlimit := max
			if len(buckets)-2 >= i {
				upperlimit = buckets[i+1]
			}
			// check if value is within bucket range
			if v >= minlimit && v <= upperlimit {
				counts[i]++
			}
		}
	}
	return buckets, counts
}

func (s Series) ShowDistribution(bucketCount int) {

	buckets, counts := s.Distribution(bucketCount)

	fmt.Println("value distribution:")

	for i, b := range buckets {
		fmt.Printf("%.5f,   %.5f \n", b, counts[i])
	}
}
