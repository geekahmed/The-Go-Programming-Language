/*
By: Ahmed Moustafa (a.k.a geekahmed)
Email: geekahmed1@gmail.com
linkedIn: https://www.linkedin.com/in/geekahmed
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	
	for i, args := range os.Args[1:]{
		fmt.Println(i, args)
	}

}
