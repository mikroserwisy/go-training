package main

import "testing"

type testPair struct {
	values []float64
	result float64
}

var testData = []testPair{
	{[]float64{1.0,-2.0}, 3.0},
	{[]float64{1.0,2.0}, 3.0}}

func TestHi(t *testing.T)  {
	result := hi("Jan")
	if result != "Hi Jan!" {
		t.Error("Expected Hi Jan, got", result)
	}
}

func TestAbsAdd(t *testing.T)  {
	for _, pair := range testData {
		result := absAdd(pair.values[0], pair.values[1])
		if result != pair.result {
			t.Error("For ", pair.values, "expected", pair.result)
		}
	}
}