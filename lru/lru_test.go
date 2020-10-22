package main

import (
	"fmt"
	"testing"
)

func TestNewLru(t *testing.T) {
	x := NewLru(	10)
	//x.Del()
	//fmt.Println(x.Start())
	//fmt.Println(x.End())
	x.Put(NewNode("wo","1"))
	//fmt.Println(x.Start())
	fmt.Println(x.End())
	x.Put(NewNode("ai","2"))
	//fmt.Println(x.Start())
	fmt.Println(x.End())
	//
	//x.Del()
	//fmt.Println(x.Start())
	//fmt.Println(x.End())
	//x.Del()
	//x.Del()
	//x.Del()
	//x.Del()
	fmt.Println(x.Cache())

	//fmt.Println(x.Start())
	//fmt.Println(x.End())
	x.Put(NewNode("ni","3"))

	fmt.Println("llllll", x.Len())
	//x.Del()
	fmt.Println(x.Len())
	fmt.Println(x.End())
	x.Put(NewNode("ya","4"))
	fmt.Println(x.End())

	fmt.Println("")
	fmt.Println(x.Start())


	fmt.Println(x.Val("ya"))
	fmt.Println(x.Val("ni"))
	fmt.Println(x.Val("ai"))
	fmt.Println(x.Val("ni"))
	fmt.Println(x.Start())
	fmt.Println(x.Val("wo"))
	//fmt.Println(x.Val(""))
	fmt.Println(x.Len())
	fmt.Println(x.MaxLen())
	fmt.Println(x.Cache())
	fmt.Println(x.Start())

}
