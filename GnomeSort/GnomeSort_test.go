package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestGnomeSort(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int, random.Intn(100))
	for i := range array {
		array[i] = random.Intn(100)
	}
	array2 := make(sort.IntSlice, len(array))
	copy(array2, array)
	GnomeSort(array)
	array2.Sort()
	for i := range array2 {
		if array[i] != array2[i] {
			t.Fail()
		}
	}
}
