package scale

import (
	"fmt"
	"testing"
)

func TestFindSymbools(t *testing.T) {
	// pathd := "M100 100 C220 98, 220 104, 150 150M150 150 C80 196, 80 202, 200 200Z"
	pathd := "M14.412 5.652c-1.28 0-2.283-.977-2.283-2.224 0-1.205a.736.736 0 0 0-.71-.76z"
	rs := []rune(pathd)
	loc := findSymbools(rs)
	t.Log(loc)
	logD(pathd, loc)
}

func TestConv(t *testing.T) {
	// pathd := "M100 100 C220 98, 220 104, 150 150M150 150 C80 196, 80 202, 200 200Z"
	// pathd := "M14.412 5.652c-1.28 0-2.283-.977-2.283-2.224 0-1.205zm0-4.306a.736.736 0 0 0-.71-.76z"
	pathd := "M0 0zm0-4.306a.736.736 0 0 0-.71-.76z"
	rs := []rune(pathd)
	loc := findSymbools(rs)
	fmt.Println(loc)
	logD(pathd, loc)
	ps := conv(rs, loc)

	t.Log(Path2D(ps))

	t.Log(Scale(pathd, 10))
}

func logD(pathd string, loc []int) {
	fmt.Println(string(pathd[loc[len(loc)-1]]))
	fmt.Println(pathd)
	for i := 1; i < len(loc); i++ {
		fmt.Printf("%s|", pathd[loc[i-1]+1:loc[i]])
	}
	fmt.Printf("%s", pathd[loc[len(loc)-1]+1:])
	fmt.Println()
}
