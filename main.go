package main

import (
	"fmt"

	myutils "github.com/shuxiaow/greeting_go/greetings/utils"
)

func main() {
	fmt.Printf("version: %s\n", myutils.GetVersion())
}
