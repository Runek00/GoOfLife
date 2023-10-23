package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const sizex = 10
const sizey = 10

var arr = [sizex][sizey]bool{}
var arr2 = [sizex][sizey]bool{}

func main() {
	fmt.Println(arr)
	for i := 0; i < sizex; i++ {
		for j := 0; j < sizey; j++ {
			arr[i][j] = rand.Int()%2 == 0
		}
	}

	s := 0
	for s < 10 {
		for i := 0; i < sizex; i++ {
			var line string
			for j := 0; j < sizey; j++ {
				if arr[i][j] {
					line = line + "#"
				} else {
					line = line + " "
				}
			}
			fmt.Println(line)
		}
		step()
		s++
	}
}

func step() {
	wg := sync.WaitGroup{}
	for i := 0; i < sizex; i++ {
		for j := 0; j < sizey; j++ {
			wg.Add(1)
			go calculatePixel(i, j, &wg)
		}
	}
	wg.Wait()
	arr = arr2
}

func calculatePixel(i int, j int, wg *sync.WaitGroup) {
	sum := 0
	defer wg.Done()
	for x := max(i-1, 0); x < min(sizex-1, i+2); x++ {
		for y := max(j-1, 0); y < min(sizey-1, j+2); y++ {
			if x == i && y == j {
				continue
			}
			if arr[x][y] {
				sum++
			}
		}
	}
	if sum < 2 || sum > 3 {
		arr2[i][j] = false
	} else if sum == 3 {
		arr2[i][j] = true
	} else {
		arr2[i][j] = arr[i][j]
	}
}
