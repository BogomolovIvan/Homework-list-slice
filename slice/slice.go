package slice

import (
	"fmt"
	"sort"
)

type Slice struct {
	sl []int64
}

func (s *Slice) Insert(data int64) (index int64) {
	s.sl = append(s.sl, data)
	return int64(len(s.sl))
}

func (s *Slice) Remove(index int64) (ok bool) {
	if int64(len(s.sl))-1 < index || index < 0 {
		return false
	}

	slRight := s.sl[index+1:]
	slLeft := s.sl[:index]

	s.sl = append(slLeft, slRight...)

	return true
}

func (s *Slice) Print() {
	for _, el := range s.sl {
		fmt.Printf("%d, ", el)
	}
	fmt.Println()
}

func (s *Slice) SortInc() {
	sort.Slice(s.sl, func(i, j int) bool {
		return s.sl[i] < s.sl[j]
	})
}

func (s *Slice) SortDec() {
	sort.Slice(s.sl, func(i, j int) bool {
		return s.sl[i] > s.sl[j]
	})
}
