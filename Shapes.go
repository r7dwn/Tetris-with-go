package main

import (
	"strconv"
	"strings"
)

type Shape struct {
	body   []string
	rotate [][]string
}

type Object struct {
	shape [][]int
	idx   int
	dir   [][][]int
}

func (i *Object) Update() {
	for j := range i.shape {
		y, x := i.shape[j][0], i.shape[j][1]
		mtrx[y][x] = "1"
	}
}

func (i *Object) OverWrite(s *Shape) {
	i.idx = 0
	shp := [][]int{}
	for j := range s.body {
		tmp := strings.Split(s.body[j], ",")
		y, _ := strconv.Atoi(tmp[0])
		x, _ := strconv.Atoi(tmp[1])
		shp = append(shp, []int{y, x})
	}
	rote := [][][]int{}
	for x := range s.rotate {
		rte := [][]int{}
		for j := range s.rotate[x] {
			tmp := strings.Split(s.rotate[x][j], ",")
			y, _ := strconv.Atoi(tmp[0])
			x, _ := strconv.Atoi(tmp[1])
			rte = append(rte, []int{y, x})
		}
		rote = append(rote, rte)
	}
	i.shape = shp
	i.dir = rote
	i.Update()
}

func (i *Object) Rotate() {
	if len(i.dir) == 0 {
		return
	}
	if i.idx == len(i.dir)-1 {
		i.idx = 0
	} else {
		i.idx++
	}
	if i.CheckBondery() {
		return
	}
	for j := range i.shape {
		mtrx[i.shape[j][0]][i.shape[j][1]] = "."
		i.shape[j][0] += i.dir[i.idx][j][0]
		i.shape[j][1] += i.dir[i.idx][j][1]
	}
	i.Update()
}

func (i *Object) CheckBondery() bool {
	xx, yy := 0, 0
	for j := range i.shape {
		x := i.shape[j][1] + i.dir[i.idx][j][1]
		y := i.shape[j][0] + i.dir[i.idx][j][0]
		if x < 0 {
			xx = 1
			break
		} else if x >= 10 {
			xx = -1
			break
		} else if mtrx[i.shape[j][0]][x] == "0" {
			i.idx--
			return true
		}
		if y < 0 {
			yy = 1
			break
		} else if y >= 20 {
			yy = -1
			break
		} else if mtrx[y][i.shape[j][1]] == "0" {
			i.idx--
			return true
		}
	}
	if xx != 0 {
		for j := range i.shape {
			i.shape[j][1] += xx
		}
	}
	if yy != 0 {
		for j := range i.shape {
			i.shape[j][0] += yy
		}
	}
	return false
}

func (i *Object) MoveLeft() {
	for j := range i.shape {
		if i.shape[j][1] == 0 {
			return
		}
		if mtrx[i.shape[j][0]][i.shape[j][1]-1] == "0" {
			return
		}
	}
	for j := range i.shape {
		y, x := i.shape[j][0], i.shape[j][1]
		mtrx[y][x] = "."
		i.shape[j][1] -= 1
	}
	i.Update()
}

func (i *Object) MoveRight() {
	for j := range i.shape {
		if i.shape[j][1] == 9 {
			return
		}
		if mtrx[i.shape[j][0]][i.shape[j][1]+1] == "0" {
			return
		}
	}
	for j := range i.shape {
		y, x := i.shape[j][0], i.shape[j][1]
		mtrx[y][x] = "."
		i.shape[j][1] += 1
	}
	i.Update()
}

func (i *Object) MoveDown() {
	for j := range i.shape {
		if i.shape[j][0] == 22 {
			return
		}
		if mtrx[i.shape[j][0]+1][i.shape[j][1]] == "0" {
			return
		}
	}
	for j := range i.shape {
		y, x := i.shape[j][0], i.shape[j][1]
		mtrx[y][x] = "."
		i.shape[j][0] += 1
	}
	i.Update()
}

func (i *Object) End() bool {
	for j := range i.shape {
		if i.shape[j][0] == 22 {
			return true
		}
		if mtrx[i.shape[j][0]+1][i.shape[j][1]] == "0" {
			return true
		}
	}
	return false
}

func (i *Object) Change() {
	for j := range i.shape {
		y, x := i.shape[j][0], i.shape[j][1]
		mtrx[y][x] = "0"
	}
}
