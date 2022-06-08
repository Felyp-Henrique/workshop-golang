package main

// #include <stdio.h>
//
// void helloWorld() {
// 		printf("Hello World com C\n");
// }
//
import "C"

func main() {
	C.helloWorld()
}
