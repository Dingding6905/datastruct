package main

import (
	"container/list"
	"fmt"
	"log"
	"time"
)

type Tree struct {
	data       int
	leftChild  *Tree
	rightChild *Tree
}

// t 传入的根节点，key 要查找的数据，find 找到后指向的节点，parent 当前节点(无论成功失败)的父节点,
// 如果返回 true 就看 find 指针，如果 false 就看 parent
func SearchBST(t *Tree, key int, find *Tree, parent **Tree) bool {
	log.Println("searchBst")
	if t == nil {
		log.Println("SearchBST t is nil")
		return false
	}

	if key == t.data {
		find = t
		return true
	} else if key < t.data {
		*parent = t
		log.Println("=============")
		// 此时函数已经返回了，相当于返回了一个新函数，和原来的已经没关系了。返回新函数，重新传入新函数值
		return SearchBST(t.leftChild, key, find, parent)
	} else {
		*parent = t
		log.Println("=============xxxxxx")
		return SearchBST(t.rightChild, key, find, parent)
	}
}

func InsertBST(t *Tree, key int) bool {
	var p *Tree = nil

	log.Println("insertBst")
	if t == nil {
		log.Println("InsertBST t is nil")
		return false
	}

	if true == SearchBST(t, key, nil, &p) {
		log.Println("InsertBST key is exist")
		return true
	}

	log.Println("----------", p)
	n := new(Tree)
	n.data = key
	n.leftChild = nil
	n.rightChild = nil
	if key < p.data {
		log.Println("insert left,key=", key, "t.data=", p.data)
		p.leftChild = n
	} else {
		log.Println("insert right,key=", key, "t.data=", p.data)
		p.rightChild = n
	}

	return true
}

func (t *Tree) PrintBST() {
	log.Println("print")

	l := list.New()
	l.PushBack(t)
	len := l.Len()
	iCount := 0
	iSum := len
	iEnd := 0
	for {
		iSum2 := iSum
		iSum = 0
		fmt.Printf("\n")
		for t1 := l.Front(); iCount < iSum2; {

			t2 := t1.Value.(*Tree)
			fmt.Printf("%d ", t2.data)

			if t2.leftChild != nil {
				iSum++
				l.PushBack(t2.leftChild)
			}
			if t2.rightChild != nil {
				iSum++
				l.PushBack(t2.rightChild)
			}
			t3 := t1.Next()
			if t3 != nil {
				l.MoveBefore(t3, t1)
			}
			l.Remove(t1)
			t1 = t3
			iCount++
			if t2.leftChild == nil && t2.rightChild == nil {
				iEnd = 1
			}
		}
		iCount = 0
		time.Sleep(time.Second * 2)
		if iEnd == 1 {
			break
		}
	}
}

func main() {
	var t Tree
	t.data = 5
	t.leftChild = nil
	t.rightChild = nil
	InsertBST(&t, 6)
	InsertBST(&t, 7)
	InsertBST(&t, 10)
	InsertBST(&t, 8)
	InsertBST(&t, 11)
	t.PrintBST()
}
