package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	var s, sep string

	// practice 1.1
	if len(os.Args) != 0 {
		fmt.Println(os.Args[0]) // file
	}

	//
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	for index, value := range os.Args {
		fmt.Println(index, value)
	}

	// go run ./T1-osArgs.go heelo go
	fmt.Println(strings.Join(os.Args[1:], " ")) // hello go

	/*
		result
	   ~/go/TheGoProgrammingLanguage/T1/ [master+*] go run ./osArgs.go first second third ...
	  1. /var/folders/7q/6dx0j03j0c37sww07vkl3nb00000gn/T/go-build917602424/b001/exe/osArgs
	  2. first second third ../..
	  3.  go run ./T1-osArgs.go heelo go -> return index, value
	    0 /var/folders/7q/6dx0j03j0c37sww07vkl3nb00000gn/T/go-build067001999/b001/exe/T1-osArgs
		1 hello
		2 go
	*/
}
