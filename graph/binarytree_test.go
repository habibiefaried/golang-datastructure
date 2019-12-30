package main 

import (
	"testing"
	"math/rand"
)


func BenchmarkInsertInOrder(b *testing.B){
	root := Node{body: 15}
	i := 0
	var t int
	for i < b.N {
		t = rand.Intn(b.N * 5)
		root.insert(&Node{body: t})
		i++
	}
}