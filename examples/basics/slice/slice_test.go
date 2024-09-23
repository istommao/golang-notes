package slice

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar := foo[1:4]
	bar[1] = 99
	fmt.Println("foo: ", foo)
	fmt.Println("bar: ", bar)

	// foo:  [0 0 99 42 100]
	// bar:  [0 99 42]
}

func TestSlice2(t *testing.T) {

	a := make([]int, 32)
	b := a[1:16]
	a = append(a, 1)
	a[2] = 42

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	// a:  [0 0 42 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
	// b:  [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func TestSlice3(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB
}

func TestSlice3_2(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex:sepIndex] // Full Slice Expression Limited Capacity
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => BBBBBBBBB
}
