package sugo

import "fmt"

func Exc(e error) {
	if e != nil {
		fmt.Println((e))
	}
}
