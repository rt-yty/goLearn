package main

import (
	"fmt"
	"goLearn/internal/pkg/storage"
)

func main() {
	s := storage.NewStorage()

	s.Set("1", "123")
	res11 := s.Get("1")
	res21 := s.GetKind("1")

	s.Set("2", "abc")
	res12 := s.Get("2")
	res22 := s.GetKind("2")

	fmt.Println(*res11, res21)
	fmt.Println(*res12, res22)
}