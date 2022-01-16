package main

import "testing"

func TestGetFruits(testHandler *testing.T) {
	fruits := getFruits()
	if len(fruits) != 10 {
		testHandler.Errorf("expected fruits len 10, but got %d", len(fruits))
	}
}
