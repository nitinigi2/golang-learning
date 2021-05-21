package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := os.Args[1]
	convertToMeter(arg)
}

func convertToMeter(arg string) {
	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		errMessage := arg + " is not a number"
		errorObj := errors.New(errMessage)
		log.Fatal(errorObj)
	} else {
		meters := feet * 0.3048
		fmt.Printf("%g feet is %g meters.\n", feet, meters)
	}
}
