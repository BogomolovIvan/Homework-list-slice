package list

import "fmt"

type node struct {
	data     int64
	nextNode *node
}

type List struct {
	len       int64
	firstNode *node
}

func NewList() (l *List) {
	l = &List{
		len:       0,
		firstNode: nil,
	}
	return
}

func (l *List) Add(data int64) (index int64) {
	if l.firstNode == nil {
		n := &node{}
		n.data = data
		l.firstNode = n
		l.len++
		return 0
	}
	cn := l.firstNode
	for {
		if cn.nextNode == nil {
			break
		}
		cn = cn.nextNode
	}
	n := &node{}
	n.data = data
	cn.nextNode = n
	l.len++
	return l.len - 1
}

func (l *List) Delete(index int64) (ok bool) {
	if l.firstNode == nil {
		return false
	}
	if index < 0 {
		return false
	}
	if index == 0 {
		l.firstNode = l.firstNode.nextNode
		l.len--
		return true
	}
	if l.len-1 >= index {
		cn := l.firstNode
		for i := int64(0); i < index-1; i++ {
			cn = cn.nextNode
		}
		previousNode := cn
		currentNode := cn.nextNode
		previousNode.nextNode = currentNode.nextNode
		currentNode.nextNode = nil
		l.len--
	}
	return true
}

func (l *List) Print() {
	fmt.Println("Len", l.len)
	if l.firstNode == nil {
		fmt.Println("Пустой список")
		return
	}
	cn := l.firstNode
	for {
		if cn.nextNode == nil {
			fmt.Println(cn.data)
			break
		}
		fmt.Println(cn.data)
		cn = cn.nextNode
	}
}

func (l *List) SortIncrease() {
	if l.firstNode == nil {
		fmt.Println("Пустой список")
		return
	}

	cn := l.firstNode
	max := cn
	for i := int64(0); i < l.len; i++ {
		for j := int64(0); j < l.len-i-1; j++ {
			if cn.data > max.data {
				max = cn
			}
			cn = cn.nextNode
		}
		cn.data, max.data = max.data, cn.data
		cn = l.firstNode
		max = cn
	}
	cn.nextNode.data, max.data = max.data, cn.nextNode.data
}

func (l *List) SortDecrease() {
	if l.firstNode == nil {
		fmt.Println("Пустой список")
		return
	}

	cn := l.firstNode
	max := cn
	for i := int64(0); i < l.len; i++ {
		for j := int64(0); j < l.len-i-1; j++ {
			if cn.data < max.data {
				max = cn
			}
			cn = cn.nextNode
		}
		cn.data, max.data = max.data, cn.data
		cn = l.firstNode
		max = cn
	}
	cn.nextNode.data, max.data = max.data, cn.nextNode.data
}

func (l *List) SortDecLink() {

}

func (l *List) SortIncLink() {

}
