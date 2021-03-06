package avltree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestAVLTree(t *testing.T) {
	n := 1000
	arr := rand.Perm(n)
	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	if !tree.IsEmpty() {
		t.Error("IsEmpty() != true error")
	}

	var m = make(map[int]int)
	for k, v := range arr {
		m[v] = k
		tree.Put(v, k)
	}

	testFunc := func() {
		ids := rand.Perm(2 * len(m))
		for _, v := range ids {
			_, ok := m[v]
			if tree.Contains(v) != ok {
				t.Error("Contains() error")
			}
			if (tree.Get(v) != nil) != ok {
				t.Error("Get() error")
			}
		}

		if tree.IsEmpty() != (len(m) == 0) {
			t.Error("IsEmpty() error")
		}
		if tree.GetSize() != len(m) {
			t.Error("GetSize() error")
		}
		if !tree.isBST() {
			t.Error("isBST() error")
		}
		if !tree.isBalanced() {
			t.Error("isBalanced() error")
		}
		for k, v := range m {
			if tree.Get(k) != v {
				t.Error("Get(v) error")
			}
		}
	}

	testFunc()

	for k, v := range arr {
		tree.Put(k, v)
		m[k] = v
	}
	testFunc()

	for len(m) > 0 {
		for k := range m {
			tree.Remove(k)
			delete(m, k)
			testFunc()
		}
	}
}

func TestNode_String(t *testing.T) {
	bst := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := rand.Perm(10)
	for i := 0; i < len(arr); i++ {
		bst.Put(arr[i], nil)
	}
	fmt.Println(bst)
}

func TestAVLTree_Put(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		for i := 0; i < len(arr); i++ {
			tree.Put(arr[i], nil)
			if !tree.isBalanced() {
				t.Fatal("balance error", arr)
			}
			if !tree.Contains(arr[i]) {
				t.Fatal("contains error", arr)
			}
			if tree.GetSize() != i+1 {
				t.Fatal("size error", arr)
			}
		}
	}
}

func TestAVLTree_Remove(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		arr2 := make([]int, n)
		copy(arr2, arr)
		for i := 0; i < len(arr)/2; i++ {
			tree.Put(arr[i], nil)
		}

		delArr := []int{}

		for i := 0; i < n; i++ {
			j := rand.Intn(len(arr2))
			key := arr2[j]
			arr2 = append(arr2[:j], arr2[j+1:]...)

			delArr = append(delArr, key)

			size := tree.GetSize()
			has := tree.Contains(key)

			tree.Remove(key)
			if !tree.isBalanced() {
				t.Fatal("balance error", arr[:len(arr)/2], delArr)
			}
			if tree.Contains(key) {
				t.Fatal("contains error")
			}
			if (has && tree.GetSize() != size-1) || (!has && tree.GetSize() != size) {
				t.Fatal("size error")
			}
		}
	}
}
