package main

import (
	"fmt"
	"github.com/daiching/pon"
)

func main() {
	err := pon.Start("8080")
	if err != nil {
		fmt.Println(err)
	}
}
