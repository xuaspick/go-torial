package main

func Task1Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z)
	}
	return z
}
