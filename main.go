package main

import (
	"awesomeProject/list"
	"awesomeProject/slice"
	"fmt"
)

func main() {
	l := list.NewList()
	l.Add(15)
	l.Add(9)
	l.Add(-2)
	l.Add(35)
	l.Add(7)
	l.Add(1)
	l.Delete(3)
	fmt.Println("Sort List")
	l.SortIncLink()
	l.SortDecLink()
	l.Print()

	s := slice.Slice{}
	s.Insert(15)
	s.Insert(9)
	s.Insert(-2)
	s.Insert(35)
	s.Insert(7)
	s.Insert(1)
	s.Print()
	fmt.Println("Sort List")
	s.SortDec()
	s.SortInc()
	s.Print()
}
