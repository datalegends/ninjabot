package series

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSeries_Last(t *testing.T) {
	series := Series([]float64{1, 2, 3, 4, 5})
	require.Equal(t, 5.0, series.Last(0))
	require.Equal(t, 3.0, series.Last(2))
}

func TestSeries_LastValues(t *testing.T) {
	t.Run("with value", func(t *testing.T) {
		series := Series([]float64{1, 2, 3, 4, 5})
		require.Equal(t, []float64{4, 5}, series.LastValues(2))
	})

	t.Run("empty", func(t *testing.T) {
		series := Series([]float64{})
		require.Empty(t, series.LastValues(2))
	})
}

func TestSeries_Distribution(t *testing.T) {

	t.Run("with 1-6", func(t *testing.T) {
		series := Series([]float64{1, 2, 3, 4, 5, 6})
		b, c := series.Distribution(2)
		require.Equal(t, []float64{1, 3.5}, b)
		require.Equal(t, []float64{3, 3}, c)
	})

	t.Run("with -2 6", func(t *testing.T) {
		series := Series([]float64{-2, -1, 0, 1, 2, 3, 4, 5, 6})
		b, c := series.Distribution(5)
		buckets := []float64{
			-2,
			-0.3999999999999999,
			1.2000000000000002,
			2.8000000000000007,
			4.4,
		}
		require.Equal(t, buckets, b)
		require.Equal(t, []float64{2, 2, 1, 2, 2}, c)
	})

	t.Run("with -2 6", func(t *testing.T) {
		series := Series([]float64{-2, -1, 0, 1, 2, 3, 4, 5, 6})
		b, c := series.Distribution(3)
		require.Equal(t, []float64{-2, 0.6666666666666665, 3.333333333333333}, b)
		require.Equal(t, []float64{3, 3, 3}, c)
	})

	t.Run("with bad values", func(t *testing.T) {
		series := Series([]float64{-2, -1, 0, 1, 2, 3, 4, 5, 6})
		b, c := series.Distribution(40)
		require.Empty(t, b)
		require.Empty(t, c)
	})

}
