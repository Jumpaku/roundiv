package div_test

import (
	"fmt"
	"testing"

	"github.com/Jumpaku/go-div"
)

func TestDivFloor_Simple(t *testing.T) {
	testCases := []testcaseDiv{
		{
			-10, -10,
			1,
		}, // 1.000000
		{
			-10, -9,
			1,
		}, // 1.111111
		{
			-10, -8,
			1,
		}, // 1.250000
		{
			-10, -7,
			1,
		}, // 1.428571
		{
			-10, -6,
			1,
		}, // 1.666667
		{
			-10, -5,
			2,
		}, // 2.000000
		{
			-10, -4,
			2,
		}, // 2.500000
		{
			-10, -3,
			3,
		}, // 3.333333
		{
			-10, -2,
			5,
		}, // 5.000000
		{
			-10, -1,
			10,
		}, // 10.000000
		{
			-10, 1,
			-10,
		}, // -10.000000
		{
			-10, 2,
			-5,
		}, // -5.000000
		{
			-10, 3,
			-4,
		}, // -3.333333
		{
			-10, 4,
			-3,
		}, // -2.500000
		{
			-10, 5,
			-2,
		}, // -2.000000
		{
			-10, 6,
			-2,
		}, // -1.666667
		{
			-10, 7,
			-2,
		}, // -1.428571
		{
			-10, 8,
			-2,
		}, // -1.250000
		{
			-10, 9,
			-2,
		}, // -1.111111
		{
			-10, 10,
			-1,
		}, // -1.000000
		{
			-9, -10,
			0,
		}, // 0.900000
		{
			-9, -9,
			1,
		}, // 1.000000
		{
			-9, -8,
			1,
		}, // 1.125000
		{
			-9, -7,
			1,
		}, // 1.285714
		{
			-9, -6,
			1,
		}, // 1.500000
		{
			-9, -5,
			1,
		}, // 1.800000
		{
			-9, -4,
			2,
		}, // 2.250000
		{
			-9, -3,
			3,
		}, // 3.000000
		{
			-9, -2,
			4,
		}, // 4.500000
		{
			-9, -1,
			9,
		}, // 9.000000
		{
			-9, 1,
			-9,
		}, // -9.000000
		{
			-9, 2,
			-5,
		}, // -4.500000
		{
			-9, 3,
			-3,
		}, // -3.000000
		{
			-9, 4,
			-3,
		}, // -2.250000
		{
			-9, 5,
			-2,
		}, // -1.800000
		{
			-9, 6,
			-2,
		}, // -1.500000
		{
			-9, 7,
			-2,
		}, // -1.285714
		{
			-9, 8,
			-2,
		}, // -1.125000
		{
			-9, 9,
			-1,
		}, // -1.000000
		{
			-9, 10,
			-1,
		}, // -0.900000
		{
			-8, -10,
			0,
		}, // 0.800000
		{
			-8, -9,
			0,
		}, // 0.888889
		{
			-8, -8,
			1,
		}, // 1.000000
		{
			-8, -7,
			1,
		}, // 1.142857
		{
			-8, -6,
			1,
		}, // 1.333333
		{
			-8, -5,
			1,
		}, // 1.600000
		{
			-8, -4,
			2,
		}, // 2.000000
		{
			-8, -3,
			2,
		}, // 2.666667
		{
			-8, -2,
			4,
		}, // 4.000000
		{
			-8, -1,
			8,
		}, // 8.000000
		{
			-8, 1,
			-8,
		}, // -8.000000
		{
			-8, 2,
			-4,
		}, // -4.000000
		{
			-8, 3,
			-3,
		}, // -2.666667
		{
			-8, 4,
			-2,
		}, // -2.000000
		{
			-8, 5,
			-2,
		}, // -1.600000
		{
			-8, 6,
			-2,
		}, // -1.333333
		{
			-8, 7,
			-2,
		}, // -1.142857
		{
			-8, 8,
			-1,
		}, // -1.000000
		{
			-8, 9,
			-1,
		}, // -0.888889
		{
			-8, 10,
			-1,
		}, // -0.800000
		{
			-7, -10,
			0,
		}, // 0.700000
		{
			-7, -9,
			0,
		}, // 0.777778
		{
			-7, -8,
			0,
		}, // 0.875000
		{
			-7, -7,
			1,
		}, // 1.000000
		{
			-7, -6,
			1,
		}, // 1.166667
		{
			-7, -5,
			1,
		}, // 1.400000
		{
			-7, -4,
			1,
		}, // 1.750000
		{
			-7, -3,
			2,
		}, // 2.333333
		{
			-7, -2,
			3,
		}, // 3.500000
		{
			-7, -1,
			7,
		}, // 7.000000
		{
			-7, 1,
			-7,
		}, // -7.000000
		{
			-7, 2,
			-4,
		}, // -3.500000
		{
			-7, 3,
			-3,
		}, // -2.333333
		{
			-7, 4,
			-2,
		}, // -1.750000
		{
			-7, 5,
			-2,
		}, // -1.400000
		{
			-7, 6,
			-2,
		}, // -1.166667
		{
			-7, 7,
			-1,
		}, // -1.000000
		{
			-7, 8,
			-1,
		}, // -0.875000
		{
			-7, 9,
			-1,
		}, // -0.777778
		{
			-7, 10,
			-1,
		}, // -0.700000
		{
			-6, -10,
			0,
		}, // 0.600000
		{
			-6, -9,
			0,
		}, // 0.666667
		{
			-6, -8,
			0,
		}, // 0.750000
		{
			-6, -7,
			0,
		}, // 0.857143
		{
			-6, -6,
			1,
		}, // 1.000000
		{
			-6, -5,
			1,
		}, // 1.200000
		{
			-6, -4,
			1,
		}, // 1.500000
		{
			-6, -3,
			2,
		}, // 2.000000
		{
			-6, -2,
			3,
		}, // 3.000000
		{
			-6, -1,
			6,
		}, // 6.000000
		{
			-6, 1,
			-6,
		}, // -6.000000
		{
			-6, 2,
			-3,
		}, // -3.000000
		{
			-6, 3,
			-2,
		}, // -2.000000
		{
			-6, 4,
			-2,
		}, // -1.500000
		{
			-6, 5,
			-2,
		}, // -1.200000
		{
			-6, 6,
			-1,
		}, // -1.000000
		{
			-6, 7,
			-1,
		}, // -0.857143
		{
			-6, 8,
			-1,
		}, // -0.750000
		{
			-6, 9,
			-1,
		}, // -0.666667
		{
			-6, 10,
			-1,
		}, // -0.600000
		{
			-5, -10,
			0,
		}, // 0.500000
		{
			-5, -9,
			0,
		}, // 0.555556
		{
			-5, -8,
			0,
		}, // 0.625000
		{
			-5, -7,
			0,
		}, // 0.714286
		{
			-5, -6,
			0,
		}, // 0.833333
		{
			-5, -5,
			1,
		}, // 1.000000
		{
			-5, -4,
			1,
		}, // 1.250000
		{
			-5, -3,
			1,
		}, // 1.666667
		{
			-5, -2,
			2,
		}, // 2.500000
		{
			-5, -1,
			5,
		}, // 5.000000
		{
			-5, 1,
			-5,
		}, // -5.000000
		{
			-5, 2,
			-3,
		}, // -2.500000
		{
			-5, 3,
			-2,
		}, // -1.666667
		{
			-5, 4,
			-2,
		}, // -1.250000
		{
			-5, 5,
			-1,
		}, // -1.000000
		{
			-5, 6,
			-1,
		}, // -0.833333
		{
			-5, 7,
			-1,
		}, // -0.714286
		{
			-5, 8,
			-1,
		}, // -0.625000
		{
			-5, 9,
			-1,
		}, // -0.555556
		{
			-5, 10,
			-1,
		}, // -0.500000
		{
			-4, -10,
			0,
		}, // 0.400000
		{
			-4, -9,
			0,
		}, // 0.444444
		{
			-4, -8,
			0,
		}, // 0.500000
		{
			-4, -7,
			0,
		}, // 0.571429
		{
			-4, -6,
			0,
		}, // 0.666667
		{
			-4, -5,
			0,
		}, // 0.800000
		{
			-4, -4,
			1,
		}, // 1.000000
		{
			-4, -3,
			1,
		}, // 1.333333
		{
			-4, -2,
			2,
		}, // 2.000000
		{
			-4, -1,
			4,
		}, // 4.000000
		{
			-4, 1,
			-4,
		}, // -4.000000
		{
			-4, 2,
			-2,
		}, // -2.000000
		{
			-4, 3,
			-2,
		}, // -1.333333
		{
			-4, 4,
			-1,
		}, // -1.000000
		{
			-4, 5,
			-1,
		}, // -0.800000
		{
			-4, 6,
			-1,
		}, // -0.666667
		{
			-4, 7,
			-1,
		}, // -0.571429
		{
			-4, 8,
			-1,
		}, // -0.500000
		{
			-4, 9,
			-1,
		}, // -0.444444
		{
			-4, 10,
			-1,
		}, // -0.400000
		{
			-3, -10,
			0,
		}, // 0.300000
		{
			-3, -9,
			0,
		}, // 0.333333
		{
			-3, -8,
			0,
		}, // 0.375000
		{
			-3, -7,
			0,
		}, // 0.428571
		{
			-3, -6,
			0,
		}, // 0.500000
		{
			-3, -5,
			0,
		}, // 0.600000
		{
			-3, -4,
			0,
		}, // 0.750000
		{
			-3, -3,
			1,
		}, // 1.000000
		{
			-3, -2,
			1,
		}, // 1.500000
		{
			-3, -1,
			3,
		}, // 3.000000
		{
			-3, 1,
			-3,
		}, // -3.000000
		{
			-3, 2,
			-2,
		}, // -1.500000
		{
			-3, 3,
			-1,
		}, // -1.000000
		{
			-3, 4,
			-1,
		}, // -0.750000
		{
			-3, 5,
			-1,
		}, // -0.600000
		{
			-3, 6,
			-1,
		}, // -0.500000
		{
			-3, 7,
			-1,
		}, // -0.428571
		{
			-3, 8,
			-1,
		}, // -0.375000
		{
			-3, 9,
			-1,
		}, // -0.333333
		{
			-3, 10,
			-1,
		}, // -0.300000
		{
			-2, -10,
			0,
		}, // 0.200000
		{
			-2, -9,
			0,
		}, // 0.222222
		{
			-2, -8,
			0,
		}, // 0.250000
		{
			-2, -7,
			0,
		}, // 0.285714
		{
			-2, -6,
			0,
		}, // 0.333333
		{
			-2, -5,
			0,
		}, // 0.400000
		{
			-2, -4,
			0,
		}, // 0.500000
		{
			-2, -3,
			0,
		}, // 0.666667
		{
			-2, -2,
			1,
		}, // 1.000000
		{
			-2, -1,
			2,
		}, // 2.000000
		{
			-2, 1,
			-2,
		}, // -2.000000
		{
			-2, 2,
			-1,
		}, // -1.000000
		{
			-2, 3,
			-1,
		}, // -0.666667
		{
			-2, 4,
			-1,
		}, // -0.500000
		{
			-2, 5,
			-1,
		}, // -0.400000
		{
			-2, 6,
			-1,
		}, // -0.333333
		{
			-2, 7,
			-1,
		}, // -0.285714
		{
			-2, 8,
			-1,
		}, // -0.250000
		{
			-2, 9,
			-1,
		}, // -0.222222
		{
			-2, 10,
			-1,
		}, // -0.200000
		{
			-1, -10,
			0,
		}, // 0.100000
		{
			-1, -9,
			0,
		}, // 0.111111
		{
			-1, -8,
			0,
		}, // 0.125000
		{
			-1, -7,
			0,
		}, // 0.142857
		{
			-1, -6,
			0,
		}, // 0.166667
		{
			-1, -5,
			0,
		}, // 0.200000
		{
			-1, -4,
			0,
		}, // 0.250000
		{
			-1, -3,
			0,
		}, // 0.333333
		{
			-1, -2,
			0,
		}, // 0.500000
		{
			-1, -1,
			1,
		}, // 1.000000
		{
			-1, 1,
			-1,
		}, // -1.000000
		{
			-1, 2,
			-1,
		}, // -0.500000
		{
			-1, 3,
			-1,
		}, // -0.333333
		{
			-1, 4,
			-1,
		}, // -0.250000
		{
			-1, 5,
			-1,
		}, // -0.200000
		{
			-1, 6,
			-1,
		}, // -0.166667
		{
			-1, 7,
			-1,
		}, // -0.142857
		{
			-1, 8,
			-1,
		}, // -0.125000
		{
			-1, 9,
			-1,
		}, // -0.111111
		{
			-1, 10,
			-1,
		}, // -0.100000
		{
			0, -10,
			0,
		}, // -0.000000
		{
			0, -9,
			0,
		}, // -0.000000
		{
			0, -8,
			0,
		}, // -0.000000
		{
			0, -7,
			0,
		}, // -0.000000
		{
			0, -6,
			0,
		}, // -0.000000
		{
			0, -5,
			0,
		}, // -0.000000
		{
			0, -4,
			0,
		}, // -0.000000
		{
			0, -3,
			0,
		}, // -0.000000
		{
			0, -2,
			0,
		}, // -0.000000
		{
			0, -1,
			0,
		}, // -0.000000
		{
			0, 1,
			0,
		}, // 0.000000
		{
			0, 2,
			0,
		}, // 0.000000
		{
			0, 3,
			0,
		}, // 0.000000
		{
			0, 4,
			0,
		}, // 0.000000
		{
			0, 5,
			0,
		}, // 0.000000
		{
			0, 6,
			0,
		}, // 0.000000
		{
			0, 7,
			0,
		}, // 0.000000
		{
			0, 8,
			0,
		}, // 0.000000
		{
			0, 9,
			0,
		}, // 0.000000
		{
			0, 10,
			0,
		}, // 0.000000
		{
			1, -10,
			-1,
		}, // -0.100000
		{
			1, -9,
			-1,
		}, // -0.111111
		{
			1, -8,
			-1,
		}, // -0.125000
		{
			1, -7,
			-1,
		}, // -0.142857
		{
			1, -6,
			-1,
		}, // -0.166667
		{
			1, -5,
			-1,
		}, // -0.200000
		{
			1, -4,
			-1,
		}, // -0.250000
		{
			1, -3,
			-1,
		}, // -0.333333
		{
			1, -2,
			-1,
		}, // -0.500000
		{
			1, -1,
			-1,
		}, // -1.000000
		{
			1, 1,
			1,
		}, // 1.000000
		{
			1, 2,
			0,
		}, // 0.500000
		{
			1, 3,
			0,
		}, // 0.333333
		{
			1, 4,
			0,
		}, // 0.250000
		{
			1, 5,
			0,
		}, // 0.200000
		{
			1, 6,
			0,
		}, // 0.166667
		{
			1, 7,
			0,
		}, // 0.142857
		{
			1, 8,
			0,
		}, // 0.125000
		{
			1, 9,
			0,
		}, // 0.111111
		{
			1, 10,
			0,
		}, // 0.100000
		{
			2, -10,
			-1,
		}, // -0.200000
		{
			2, -9,
			-1,
		}, // -0.222222
		{
			2, -8,
			-1,
		}, // -0.250000
		{
			2, -7,
			-1,
		}, // -0.285714
		{
			2, -6,
			-1,
		}, // -0.333333
		{
			2, -5,
			-1,
		}, // -0.400000
		{
			2, -4,
			-1,
		}, // -0.500000
		{
			2, -3,
			-1,
		}, // -0.666667
		{
			2, -2,
			-1,
		}, // -1.000000
		{
			2, -1,
			-2,
		}, // -2.000000
		{
			2, 1,
			2,
		}, // 2.000000
		{
			2, 2,
			1,
		}, // 1.000000
		{
			2, 3,
			0,
		}, // 0.666667
		{
			2, 4,
			0,
		}, // 0.500000
		{
			2, 5,
			0,
		}, // 0.400000
		{
			2, 6,
			0,
		}, // 0.333333
		{
			2, 7,
			0,
		}, // 0.285714
		{
			2, 8,
			0,
		}, // 0.250000
		{
			2, 9,
			0,
		}, // 0.222222
		{
			2, 10,
			0,
		}, // 0.200000
		{
			3, -10,
			-1,
		}, // -0.300000
		{
			3, -9,
			-1,
		}, // -0.333333
		{
			3, -8,
			-1,
		}, // -0.375000
		{
			3, -7,
			-1,
		}, // -0.428571
		{
			3, -6,
			-1,
		}, // -0.500000
		{
			3, -5,
			-1,
		}, // -0.600000
		{
			3, -4,
			-1,
		}, // -0.750000
		{
			3, -3,
			-1,
		}, // -1.000000
		{
			3, -2,
			-2,
		}, // -1.500000
		{
			3, -1,
			-3,
		}, // -3.000000
		{
			3, 1,
			3,
		}, // 3.000000
		{
			3, 2,
			1,
		}, // 1.500000
		{
			3, 3,
			1,
		}, // 1.000000
		{
			3, 4,
			0,
		}, // 0.750000
		{
			3, 5,
			0,
		}, // 0.600000
		{
			3, 6,
			0,
		}, // 0.500000
		{
			3, 7,
			0,
		}, // 0.428571
		{
			3, 8,
			0,
		}, // 0.375000
		{
			3, 9,
			0,
		}, // 0.333333
		{
			3, 10,
			0,
		}, // 0.300000
		{
			4, -10,
			-1,
		}, // -0.400000
		{
			4, -9,
			-1,
		}, // -0.444444
		{
			4, -8,
			-1,
		}, // -0.500000
		{
			4, -7,
			-1,
		}, // -0.571429
		{
			4, -6,
			-1,
		}, // -0.666667
		{
			4, -5,
			-1,
		}, // -0.800000
		{
			4, -4,
			-1,
		}, // -1.000000
		{
			4, -3,
			-2,
		}, // -1.333333
		{
			4, -2,
			-2,
		}, // -2.000000
		{
			4, -1,
			-4,
		}, // -4.000000
		{
			4, 1,
			4,
		}, // 4.000000
		{
			4, 2,
			2,
		}, // 2.000000
		{
			4, 3,
			1,
		}, // 1.333333
		{
			4, 4,
			1,
		}, // 1.000000
		{
			4, 5,
			0,
		}, // 0.800000
		{
			4, 6,
			0,
		}, // 0.666667
		{
			4, 7,
			0,
		}, // 0.571429
		{
			4, 8,
			0,
		}, // 0.500000
		{
			4, 9,
			0,
		}, // 0.444444
		{
			4, 10,
			0,
		}, // 0.400000
		{
			5, -10,
			-1,
		}, // -0.500000
		{
			5, -9,
			-1,
		}, // -0.555556
		{
			5, -8,
			-1,
		}, // -0.625000
		{
			5, -7,
			-1,
		}, // -0.714286
		{
			5, -6,
			-1,
		}, // -0.833333
		{
			5, -5,
			-1,
		}, // -1.000000
		{
			5, -4,
			-2,
		}, // -1.250000
		{
			5, -3,
			-2,
		}, // -1.666667
		{
			5, -2,
			-3,
		}, // -2.500000
		{
			5, -1,
			-5,
		}, // -5.000000
		{
			5, 1,
			5,
		}, // 5.000000
		{
			5, 2,
			2,
		}, // 2.500000
		{
			5, 3,
			1,
		}, // 1.666667
		{
			5, 4,
			1,
		}, // 1.250000
		{
			5, 5,
			1,
		}, // 1.000000
		{
			5, 6,
			0,
		}, // 0.833333
		{
			5, 7,
			0,
		}, // 0.714286
		{
			5, 8,
			0,
		}, // 0.625000
		{
			5, 9,
			0,
		}, // 0.555556
		{
			5, 10,
			0,
		}, // 0.500000
		{
			6, -10,
			-1,
		}, // -0.600000
		{
			6, -9,
			-1,
		}, // -0.666667
		{
			6, -8,
			-1,
		}, // -0.750000
		{
			6, -7,
			-1,
		}, // -0.857143
		{
			6, -6,
			-1,
		}, // -1.000000
		{
			6, -5,
			-2,
		}, // -1.200000
		{
			6, -4,
			-2,
		}, // -1.500000
		{
			6, -3,
			-2,
		}, // -2.000000
		{
			6, -2,
			-3,
		}, // -3.000000
		{
			6, -1,
			-6,
		}, // -6.000000
		{
			6, 1,
			6,
		}, // 6.000000
		{
			6, 2,
			3,
		}, // 3.000000
		{
			6, 3,
			2,
		}, // 2.000000
		{
			6, 4,
			1,
		}, // 1.500000
		{
			6, 5,
			1,
		}, // 1.200000
		{
			6, 6,
			1,
		}, // 1.000000
		{
			6, 7,
			0,
		}, // 0.857143
		{
			6, 8,
			0,
		}, // 0.750000
		{
			6, 9,
			0,
		}, // 0.666667
		{
			6, 10,
			0,
		}, // 0.600000
		{
			7, -10,
			-1,
		}, // -0.700000
		{
			7, -9,
			-1,
		}, // -0.777778
		{
			7, -8,
			-1,
		}, // -0.875000
		{
			7, -7,
			-1,
		}, // -1.000000
		{
			7, -6,
			-2,
		}, // -1.166667
		{
			7, -5,
			-2,
		}, // -1.400000
		{
			7, -4,
			-2,
		}, // -1.750000
		{
			7, -3,
			-3,
		}, // -2.333333
		{
			7, -2,
			-4,
		}, // -3.500000
		{
			7, -1,
			-7,
		}, // -7.000000
		{
			7, 1,
			7,
		}, // 7.000000
		{
			7, 2,
			3,
		}, // 3.500000
		{
			7, 3,
			2,
		}, // 2.333333
		{
			7, 4,
			1,
		}, // 1.750000
		{
			7, 5,
			1,
		}, // 1.400000
		{
			7, 6,
			1,
		}, // 1.166667
		{
			7, 7,
			1,
		}, // 1.000000
		{
			7, 8,
			0,
		}, // 0.875000
		{
			7, 9,
			0,
		}, // 0.777778
		{
			7, 10,
			0,
		}, // 0.700000
		{
			8, -10,
			-1,
		}, // -0.800000
		{
			8, -9,
			-1,
		}, // -0.888889
		{
			8, -8,
			-1,
		}, // -1.000000
		{
			8, -7,
			-2,
		}, // -1.142857
		{
			8, -6,
			-2,
		}, // -1.333333
		{
			8, -5,
			-2,
		}, // -1.600000
		{
			8, -4,
			-2,
		}, // -2.000000
		{
			8, -3,
			-3,
		}, // -2.666667
		{
			8, -2,
			-4,
		}, // -4.000000
		{
			8, -1,
			-8,
		}, // -8.000000
		{
			8, 1,
			8,
		}, // 8.000000
		{
			8, 2,
			4,
		}, // 4.000000
		{
			8, 3,
			2,
		}, // 2.666667
		{
			8, 4,
			2,
		}, // 2.000000
		{
			8, 5,
			1,
		}, // 1.600000
		{
			8, 6,
			1,
		}, // 1.333333
		{
			8, 7,
			1,
		}, // 1.142857
		{
			8, 8,
			1,
		}, // 1.000000
		{
			8, 9,
			0,
		}, // 0.888889
		{
			8, 10,
			0,
		}, // 0.800000
		{
			9, -10,
			-1,
		}, // -0.900000
		{
			9, -9,
			-1,
		}, // -1.000000
		{
			9, -8,
			-2,
		}, // -1.125000
		{
			9, -7,
			-2,
		}, // -1.285714
		{
			9, -6,
			-2,
		}, // -1.500000
		{
			9, -5,
			-2,
		}, // -1.800000
		{
			9, -4,
			-3,
		}, // -2.250000
		{
			9, -3,
			-3,
		}, // -3.000000
		{
			9, -2,
			-5,
		}, // -4.500000
		{
			9, -1,
			-9,
		}, // -9.000000
		{
			9, 1,
			9,
		}, // 9.000000
		{
			9, 2,
			4,
		}, // 4.500000
		{
			9, 3,
			3,
		}, // 3.000000
		{
			9, 4,
			2,
		}, // 2.250000
		{
			9, 5,
			1,
		}, // 1.800000
		{
			9, 6,
			1,
		}, // 1.500000
		{
			9, 7,
			1,
		}, // 1.285714
		{
			9, 8,
			1,
		}, // 1.125000
		{
			9, 9,
			1,
		}, // 1.000000
		{
			9, 10,
			0,
		}, // 0.900000
		{
			10, -10,
			-1,
		}, // -1.000000
		{
			10, -9,
			-2,
		}, // -1.111111
		{
			10, -8,
			-2,
		}, // -1.250000
		{
			10, -7,
			-2,
		}, // -1.428571
		{
			10, -6,
			-2,
		}, // -1.666667
		{
			10, -5,
			-2,
		}, // -2.000000
		{
			10, -4,
			-3,
		}, // -2.500000
		{
			10, -3,
			-4,
		}, // -3.333333
		{
			10, -2,
			-5,
		}, // -5.000000
		{
			10, -1,
			-10,
		}, // -10.000000
		{
			10, 1,
			10,
		}, // 10.000000
		{
			10, 2,
			5,
		}, // 5.000000
		{
			10, 3,
			3,
		}, // 3.333333
		{
			10, 4,
			2,
		}, // 2.500000
		{
			10, 5,
			2,
		}, // 2.000000
		{
			10, 6,
			1,
		}, // 1.666667
		{
			10, 7,
			1,
		}, // 1.428571
		{
			10, 8,
			1,
		}, // 1.250000
		{
			10, 9,
			1,
		}, // 1.111111
		{
			10, 10,
			1,
		}, // 1.000000

	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`(%3d/%3d)`, tc.a, tc.b), func(t *testing.T) {
			doTestDiv(t, div.DivFloor, tc)
		})
	}
}

func TestDivFloor_Small(t *testing.T) {
	testCases := []testcaseDiv{
		{
			-2, -1234567890123456789,
			0,
		},
		{
			-2, 1234567890123456789,
			-1,
		},
		{
			-1, -1234567890123456789,
			0,
		},
		{
			-1, 1234567890123456789,
			-1,
		},
		{
			0, -1234567890123456789,
			0,
		},
		{
			0, 1234567890123456789,
			0,
		},
		{
			1, -1234567890123456789,
			-1,
		},
		{
			1, 1234567890123456789,
			0,
		},
		{
			2, -1234567890123456789,
			-1,
		},
		{
			2, 1234567890123456789,
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`(%3d/%20d)`, tc.a, tc.b), func(t *testing.T) {
			doTestDiv(t, div.DivFloor, tc)
		})
	}
}

func TestDivFloor_Half(t *testing.T) {
	testCases := []testcaseDiv{
		{
			-25000000000000001, -10000000000000000,
			2,
		},
		{
			-25000000000000001, 10000000000000000,
			-3,
		},
		{
			-25000000000000000, -10000000000000000,
			2,
		},
		{
			-25000000000000000, 10000000000000000,
			-3,
		},
		{
			-24999999999999999, -10000000000000000,
			2,
		},
		{
			-24999999999999999, 10000000000000000,
			-3,
		},
		{
			25000000000000001, -10000000000000000,
			-3,
		},
		{
			25000000000000001, 10000000000000000,
			2,
		},
		{
			25000000000000000, -10000000000000000,
			-3,
		},
		{
			25000000000000000, 10000000000000000,
			2,
		},
		{
			24999999999999999, -10000000000000000,
			-3,
		},
		{
			24999999999999999, 10000000000000000,
			2,
		},
		{
			-35000000000000001, -10000000000000000,
			3,
		},
		{
			-35000000000000001, 10000000000000000,
			-4,
		},
		{
			-35000000000000000, -10000000000000000,
			3,
		},
		{
			-35000000000000000, 10000000000000000,
			-4,
		},
		{
			-34999999999999999, -10000000000000000,
			3,
		},
		{
			-34999999999999999, 10000000000000000,
			-4,
		},
		{
			35000000000000001, -10000000000000000,
			-4,
		},
		{
			35000000000000001, 10000000000000000,
			3,
		},
		{
			35000000000000000, -10000000000000000,
			-4,
		},
		{
			35000000000000000, 10000000000000000,
			3,
		},
		{
			34999999999999999, -10000000000000000,
			-4,
		},
		{
			34999999999999999, 10000000000000000,
			3,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`(%20d/%20d)`, tc.a, tc.b), func(t *testing.T) {
			doTestDiv(t, div.DivFloor, tc)
		})
	}
}

func TestDivFloor_Close(t *testing.T) {
	testCases := []testcaseDiv{
		{
			-20000000000000001, -20000000000000000,
			1,
		},
		{
			-20000000000000001, 20000000000000000,
			-2,
		},
		{
			-20000000000000000, -20000000000000000,
			1,
		},
		{
			-20000000000000000, 20000000000000000,
			-1,
		},
		{
			-19999999999999999, -20000000000000000,
			0,
		},
		{
			-19999999999999999, 20000000000000000,
			-1,
		},
		{
			20000000000000001, -20000000000000000,
			-2,
		},
		{
			20000000000000001, 20000000000000000,
			1,
		},
		{
			20000000000000000, -20000000000000000,
			-1,
		},
		{
			20000000000000000, 20000000000000000,
			1,
		},
		{
			19999999999999999, -20000000000000000,
			-1,
		},
		{
			19999999999999999, 20000000000000000,
			0,
		},
		{
			-30000000000000001, -30000000000000000,
			1,
		},
		{
			-30000000000000001, 30000000000000000,
			-2,
		},
		{
			-30000000000000000, -30000000000000000,
			1,
		},
		{
			-30000000000000000, 30000000000000000,
			-1,
		},
		{
			-29999999999999999, -30000000000000000,
			0,
		},
		{
			-29999999999999999, 30000000000000000,
			-1,
		},
		{
			30000000000000001, -30000000000000000,
			-2,
		},
		{
			30000000000000001, 30000000000000000,
			1,
		},
		{
			30000000000000000, -30000000000000000,
			-1,
		},
		{
			30000000000000000, 30000000000000000,
			1,
		},
		{
			29999999999999999, -30000000000000000,
			-1,
		},
		{
			29999999999999999, 30000000000000000,
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(`(%20d/%20d)`, tc.a, tc.b), func(t *testing.T) {
			doTestDiv(t, div.DivFloor, tc)
		})
	}
}
