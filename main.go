package main

import (
	"strconv"

	structs "github.com/josenoriegaa/exam-unit-2/structs"
)

func main() {
	for i := 0; i < 5; i++ {
		student := structs.Student{}
		student.Capture()
		student.UUID = strconv.Itoa(i)
		student.Save()
	}

	for i := 0; i < 5; i++ {
		student := structs.Student{}
		student.UUID = strconv.Itoa(i)
		student.Read()
		student.ToString()
	}
}
