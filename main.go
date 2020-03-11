package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	helpers "github.com/josenoriegaa/exam-unit-2/helpers"

	structs "github.com/josenoriegaa/exam-unit-2/structs"
)

func main() {
	var data [5]structs.Student
	reader := bufio.NewReader(os.Stdin)

	option := -1
	userError := ""
	for {
		helpers.Clear()
		fmt.Println("1) Capture students")
		fmt.Println("2) Save students")
		fmt.Println("3) Print stored students")
		fmt.Println("4) Exit")
		if len(userError) > 0 {
			fmt.Println(userError)
		}
		fmt.Println("Select an option from the menu by enter its related number")
		fmt.Print("> ")
		optStr, _ := reader.ReadString('\n')
		optStr = strings.ToLower(optStr)
		optStr = strings.TrimSpace(optStr)

		// Validate entry
		optIsValid, _ := regexp.MatchString("[0-9]", optStr)
		if !optIsValid {
			userError = "Please, enter a valid option"
			continue
		}
		option, _ = strconv.Atoi(optStr)
		switch option {
		case 1:
			helpers.Clear()
			for i := 0; i < 5; i++ {
				student := structs.Student{}
				student.Capture()
				student.UUID = strconv.Itoa(i)
				data[i] = student
			}
			break
		case 2:
			for i := 0; i < 5; i++ {
				data[i].Save()
			}
			break
		case 3:
			for i := 0; i < 5; i++ {
				idStr := strconv.Itoa(i)
				data[i].UUID = idStr
				data[i].Read()
				data[i].ToString()
			}
			break
		}
		if option == 4 {
			break
		}
		fmt.Println("Press any key to continue")
		reader.ReadString('\n')
	}
}
