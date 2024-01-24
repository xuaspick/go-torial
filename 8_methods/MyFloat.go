package main

import "fmt"

type MyFloat float64

func (mf MyFloat) Abs() float64 {
	if mf < 0 {
		return float64(-mf)
	}
	return float64(mf)
}

func (mf MyFloat) M() {
	fmt.Println(mf)
}

func (mf MyFloat) get() string {
	return fmt.Sprintf("%f", mf)
}
