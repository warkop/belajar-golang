package main

import "math"

// Kubus is
type Kubus struct {
	Sisi float64
}

// Volume is
func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

// Luas is
func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

// Keliling is
func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}
