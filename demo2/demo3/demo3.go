package demo3

import (
	"alfredd/dump/demo"
	"alfredd/dump/demo2"
)

func Echo3(s string) string {
	return demo2.Echo2(s) + demo.Echo(s)
}
