package main

import (
	"alfredd/dump/demo"
	"alfredd/dump/demo2"
	"alfredd/dump/demo2/demo3"
	"fmt"
)

func main() {
	fmt.Println(demo.Echo("Hello World!"))
	fmt.Println(demo2.Echo2("Hello World! "))
	fmt.Println(demo3.Echo3("Hello World! "))
}
