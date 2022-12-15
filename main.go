package main

import (
	"awesomeProject/list"
	"awesomeProject/slice"
	"fmt"
	//"awesomeProject/model"
)

func main() {
	l := list.NewList()
	l.Add(15)
	l.Add(9)
	l.Add(-2)
	l.Add(35)
	l.Add(7)
	l.Add(1)
	//model.Delete(l, 3)
	//model.Print(l)
	fmt.Println("Sort List")
	//model.SortInc(l)
	//l.SortIncLink()
	//model.SortDec(l)
	//l.SortDecLink()
	//model.Print(l)

	s := slice.Slice{}
	l.Add(15)
	l.Add(9)
	l.Add(-2)
	l.Add(35)
	l.Add(7)
	l.Add(1)
	l.Print()
	fmt.Println("Sort List")
	//model.SortInc(&s)
	//model.SortDec(&s)
	s.Print()
}
