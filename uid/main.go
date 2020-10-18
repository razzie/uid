package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/razzie/uid"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: uid <bits>")
		os.Exit(0)
	}
	bits, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(uid.NewGenerator(bits).UID())
}
