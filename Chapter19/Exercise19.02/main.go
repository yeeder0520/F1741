package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Name    string
	Age     int
	Balance float64
	Member  bool
}

func main() {
	u1 := User{
		Name:    "Tracy",
		Age:     51,
		Balance: 98.43,
		Member:  true,
	}
	fmt.Println(u1)

	fmt.Println("Size/offset:")
	fmt.Println("Name   ", unsafe.Sizeof(u1.Name), unsafe.Offsetof(u1.Name))
	fmt.Println("Age    ", unsafe.Sizeof(u1.Age), unsafe.Offsetof(u1.Age))
	fmt.Println("Balance", unsafe.Sizeof(u1.Balance), unsafe.Offsetof(u1.Balance))
	fmt.Println("Member ", unsafe.Sizeof(u1.Member), unsafe.Offsetof(u1.Member))
	fmt.Println("u1 align:", unsafe.Alignof(u1))
	fmt.Println("u1 size :", unsafe.Sizeof(u1))

	balance := (*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(&u1)) + unsafe.Sizeof(u1.Name) + unsafe.Sizeof(u1.Age)))
	*balance += 10000

	member := (*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(&u1)) + unsafe.Offsetof(u1.Member)))
	*member = false

	fmt.Println(u1)
}
