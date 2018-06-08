package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main()  {
	counts := make(map[string]int)
	//input := bufio.NewScanner(os.Stdin)

	//files := os.Args[1:]

	//if len(files) == 0 {
	//	countLines(os.Stdin, counts)
	//} else {
	//	for _, arg := range files {
	//		f, err := os.Open(arg)
	//		if err != nil {
	//			fmt.Fprintf(os.Stderr, "dup1:%v\n", err)
	//			continue
	//		}
	//	}
	//}

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup1: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}


	//for input.Scan() {
	//	counts[input.Text()]++
	//}
	//
	//for line, n := range counts {
	//	if n > 1 {
	//		fmt.Printf("%d\t%s\n", n, line)
	//	}
	//}
}