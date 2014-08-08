package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	script, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("error : %v\n", err)
	}

	var result interface{}
	RunV8(string(script), &result)
	fmt.Printf("Result -> %T: %#+v\n", result, result)

}
