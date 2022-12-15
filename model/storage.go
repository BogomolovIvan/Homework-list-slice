package model

type Storage interface {
	Add(data int64) (index int64)
	Delete(index int64) (ok bool)
	Print()
	SortIncrease()
	SortDecrease()
}
