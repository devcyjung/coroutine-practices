package main

import "fmt"

func main() {
	filePermission := 0666 // octal
	// [1] means "use 1st argument",
	// %#o %#x %#X mean "print 0, 0x, 0X prefix"
	fmt.Printf("filePermission:%d %[1]o %#[1]o\n", filePermission)
	hex := int64(0xdeadbeef)
	fmt.Printf("hex:%d %[1]x %#[1]x %#[1]X\n", hex)
}
