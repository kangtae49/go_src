package main

//https://lawlessguy.wordpress.com/2013/07/17/simulating-try-catch-in-go-golang/

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			(func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Println("defer ij", i, j)
					}
				}()

				fmt.Println(i, j)
				panic("PANIC")

			})()

		}
	}
}

/*
type Pos struct {
	name string
	x    int
	y    int
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	pos := g()
	fmt.Println(pos)
}

func g() *Pos {
	pos := Pos{"aa", 1, 2}
	defer func() {
		pos.x = 10
	}()

	return &pos
}

*/
