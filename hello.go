package main

import "C"

func main() {

}
func Random() int {
	return int(C.random())
}
func Seed(i int) {
	C.srandom(C.uint(i))
}
