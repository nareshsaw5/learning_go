package main

import (
	"fmt"

	"github.com/0101binarybard/mdm/service"
	"github.com/0101binarybard/mdm/utils"
)

/**
Import declaration          Local name of Sin

import   "lib/math"         math.Sin
import M "lib/math"         M.Sin
import . "lib/math"         Sin

the dot(.) import is not recommented
*/

func main() {
	device := service.GetDeviceName("dummy")
	fmt.Println(device)
	newName := utils.ChangeName("Naresh")
	fmt.Println(newName)
}
