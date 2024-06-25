package format

import "fmt"

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
