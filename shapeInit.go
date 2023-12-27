package main

import (
	"os"
	"os/exec"
	"runtime"
)

func getList() []*Shape {
	arr := []*Shape{}
	// shape I
	tmpI := Shape{}
	tmpI.body = []string{"0,4", "1,4", "2,4"}
	tmpI.rotate = append(tmpI.rotate, []string{"-1,1", "0,0", "1,-1"})
	tmpI.rotate = append(tmpI.rotate, []string{"1,-1", "0,0", "-1,1"})
	arr = append(arr, &tmpI)

	// shapeO
	tmpO := Shape{}
	tmpO.body = []string{"0,4", "0,5", "1,4", "1,5"}
	arr = append(arr, &tmpO)

	// shapeL
	tmpL := Shape{}
	tmpL.body = []string{"0,4", "1,4", "2,4", "2,5"}
	tmpL.rotate = append(tmpL.rotate, []string{"-1,-2", "0,-1", "1,0", "0,1"})
	tmpL.rotate = append(tmpL.rotate, []string{"1,0", "0,1", "-1,2", "-2,1"})
	tmpL.rotate = append(tmpL.rotate, []string{"1,1", "0,0", "-1,-1", "0,-2"})
	tmpL.rotate = append(tmpL.rotate, []string{"-1,1", "0,0", "1,-1", "2,0"})
	arr = append(arr, &tmpL)

	// ShapeJ
	tmpJ := Shape{}
	tmpJ.body = []string{"0,5", "1,5", "2,5", "2,4"}
	tmpJ.rotate = append(tmpJ.rotate, []string{"-1,-1", "0,0", "1,1", "2,0"})
	tmpJ.rotate = append(tmpJ.rotate, []string{"1,-1", "0,0", "-1,1", "0,2"})
	tmpJ.rotate = append(tmpJ.rotate, []string{"1,0", "0,-1", "-1,-2", "-2,-1"})
	tmpJ.rotate = append(tmpJ.rotate, []string{"-1,2", "0,1", "1,0", "0,-1"})
	arr = append(arr, &tmpJ)

	// ShapeN
	tmpN := Shape{}
	tmpN.body = []string{"0,4", "1,4", "1,5", "2,5"}
	tmpN.rotate = append(tmpN.rotate, []string{"-1,0", "0,-1", "1,0", "2,-1"})
	tmpN.rotate = append(tmpN.rotate, []string{"1,0", "0,1", "-1,0", "-2,1"})
	arr = append(arr, &tmpN)

	// ShapeZ
	tmpZ := Shape{}
	tmpZ.body = []string{"0,5", "1,5", "1,4", "2,4"}
	tmpZ.rotate = append(tmpZ.rotate, []string{"0,1", "1,0", "0,-1", "1,-2"})
	tmpZ.rotate = append(tmpZ.rotate, []string{"0,-1", "-1,0", "0,1", "-1,2"})
	arr = append(arr, &tmpZ)

	return arr
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Clear() {
	switch runtime.GOOS {
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}
