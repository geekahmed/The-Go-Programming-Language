/*
By: Ahmed Moustafa (a.k.a geekahmed)
Email: geekahmed1@gmail.com
linkedIn: https://www.linkedin.com/in/geekahmed
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// inefficient
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	t := time.Now()
	elapsed := t.Sub(start)

	// Efficient
	start1 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	t1 := time.Now()
	elapsed1 := t1.Sub(start1)

	fmt.Println("Version 1 Execution time: ", elapsed)
	fmt.Println("Version 2 Execution time: ", elapsed1)
}

/*
	==== Results ====
./main aaa aaa kdkl askld dlk
aaa aaa kdkl askld dlk
aaa aaa kdkl askld dlk
Version 1 Execution time:  107.126µs
Version 2 Execution time:  15.232µs
 */
