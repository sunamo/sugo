package sugo

import "fmt"

func exc(e error) {
	if e != nil {
		fmt.Println((e))
	}
}
