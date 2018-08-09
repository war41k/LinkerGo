package main

import "C"

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

var count int
var mtx sync.Mutex

type Linker string

var Link Linker

func (L Linker) Add(a, b int) int {
	return a + b
}

func (L Linker) Cosine(x float64) float64 {
	return math.Cos(x)
}

func (L Linker) Sort(vals []int) {
	sort.Ints(vals)
}

func (L Linker) Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

func main() {}
