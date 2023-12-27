package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

var (
	mtrx      = [][]string{}
	gameOver  = false
	shapeList = []*Shape{}
	nxtMtx    = [][]string{}
	score     = 0
	highscore = 0
)

func main() {
	for i := 0; i < 23; i++ {
		tmp := []string{}
		for j := 0; j < 10; j++ {
			tmp = append(tmp, ".")
		}
		mtrx = append(mtrx, tmp)
	}
	for i := 0; i < 3; i++ {
		tmp := []string{}
		for j := 0; j < 3; j++ {
			tmp = append(tmp, ".")
		}
		nxtMtx = append(nxtMtx, tmp)
	}
	file, err := os.ReadFile("score.txt")
	if err == nil {
		v, _ := strconv.Atoi(string(file))
		highscore = v
	}
	shapeList = getList()
	crnt := Object{}
	crnt.OverWrite(RandomShape())
	next := RandomShape()
	UpdateNxt(next)
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	Out()
	go func() {
		for !gameOver {
			char, key, err := keyboard.GetKey()
			if err != nil {
				gameOver = true
				return
			}
			if key == keyboard.KeyEsc {
				gameOver = true
				return
			}
			if char == 's' {
				crnt.MoveDown()
			}
			if char == 'a' {
				crnt.MoveLeft()
			}
			if char == 'd' {
				crnt.MoveRight()
			}
			if char == 'w' {
				crnt.Rotate()
			}
			if crnt.End() {
				crnt.Change()
				Check(crnt.shape)
				if gameOver {
					return
				}
				crnt.OverWrite(next)
				next = RandomShape()
				UpdateNxt(next)
			}
			Out()
		}
	}()
	for range time.Tick(time.Millisecond * 500) {
		crnt.MoveDown()
		Out()
		if crnt.End() {
			crnt.Change()
			Check(crnt.shape)
			crnt.OverWrite(next)
			next = RandomShape()
			UpdateNxt(next)
		}
		if gameOver {
			break
		}
	}
	if score > highscore {
		err := os.WriteFile("score.txt", []byte(fmt.Sprint(score)), 0755)
		if err != nil {
			panic(err)
		}
	}
}

func Out() {
	Clear()
	for i := 3; i < 23; i++ {
		if i == 6 {
			fmt.Println(mtrx[i], "        highscore:", highscore)
		} else if i == 7 {
			fmt.Println(mtrx[i], "        score:", score)
		} else if i == 9 {
			fmt.Println(mtrx[i], "        nextShape:")
		} else if i >= 10 && i < 13 {
			fmt.Println(mtrx[i], "       ", nxtMtx[i-10])
		} else {
			fmt.Println(mtrx[i])
		}
	}
}

func RandomShape() *Shape {
	return shapeList[rand.Intn(len(shapeList)-1)]
}

func CheckRow(y int) bool {
	for x := 0; x < len(mtrx[y]); x++ {
		if mtrx[y][x] == "." {
			return false
		}
	}
	return true
}

func DeleteRow(y int) {
	for ; y >= 0; y-- {
		for x := 0; x < len(mtrx[y]); x++ {
			mtrx[y+1][x] = mtrx[y][x]
		}
	}
}

func Check(a [][]int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	arr := Filter(a)
	for i := range arr {
		if arr[i] <= 2 {
			gameOver = true
			return
		}
		if CheckRow(arr[i]) {
			score += 100
			DeleteRow(arr[i] - 1)
		}
	}
}

func Filter(a [][]int) []int {
	arr := []int{}
	for i := range a {
		if len(arr) > 0 && arr[len(arr)-1] == a[i][0] {
			continue
		}
		arr = append(arr, a[i][0])
	}
	return arr
}

func UpdateNxt(s *Shape) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			nxtMtx[i][j] = "."
		}
	}
	for j := range s.body {
		tmp := strings.Split(s.body[j], ",")
		y, _ := strconv.Atoi(tmp[0])
		x, _ := strconv.Atoi(tmp[1])
		nxtMtx[y][x-3] = "1"
	}
}
