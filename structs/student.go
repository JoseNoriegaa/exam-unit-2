package structs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	helpers "github.com/josenoriegaa/exam-unit-2/helpers"
)

// Student struct
type Student struct {
	UUID      string
	FirstName string
	LastName  string
	Notes     [2]uint16
}

// DataPath returns the absolute path related to the student
func (s *Student) DataPath() (path string) {
	exPath, _ := filepath.Abs("./")
	path = exPath + "/data/" + s.UUID
	return
}

// Exist returns true if the student is already stored into de data directory
func (s *Student) Exist() (result bool) {
	result = false
	_, err := os.Stat(s.DataPath())
	if err == nil {
		result = true
	}
	return
}

// Save saves the student information into a file
func (s *Student) Save() (result bool) {
	result = true
	var file *os.File
	var err error
	file, err = os.Create(s.DataPath())
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data := "firstName:" + s.FirstName + "\n"
	data = data + "lastName:" + s.LastName + "\n"
	data = data + "note[0]:" + strconv.Itoa(int(s.Notes[0])) + "\n"
	data = data + "note[1]:" + strconv.Itoa(int(s.Notes[1])) + "\n"
	buf := bufio.NewWriter(file)
	_, err = buf.WriteString(data)
	if err != nil {
		result = false
		panic(err)
	}
	buf.Flush()
	return
}

// Read searchs the student related data file and parses the stored data to the struct
func (s *Student) Read() (result bool) {
	if s.UUID == "" {
		panic(errors.New("Error: UUID must be setted up"))
	} else {
		if s.Exist() {
			file, _ := os.Open(s.DataPath())
			defer file.Close()
			reader := bufio.NewReader(file)
			for {
				data, _ := reader.ReadString('\n')
				data = strings.TrimSuffix(data, "\n")
				set := strings.Split(data, ":")
				switch set[0] {
				case "firstName":
					s.FirstName = set[1]
					break
				case "lastName":
					s.LastName = set[1]
					break
				case "note[0]":
					val, _ := strconv.Atoi(set[1])
					s.Notes[0] = uint16(val)
					break
				case "note[1]":
					val, _ := strconv.Atoi(set[1])
					s.Notes[1] = uint16(val)
					break
				}
				if len(data) <= 0 {
					break
				}
			}
			result = true
		} else {
			result = false
		}
	}
	return
}

// Capture shows a form to enter the student information
func (s *Student) Capture() {
	var userError string
	var firstName string
	var lastName string
	var note1 string
	var note2 string

	reader := bufio.NewReader(os.Stdin)
	for len(firstName) == 0 || len(lastName) == 0 || len(note1) == 0 || len(note2) == 0 {
		helpers.Clear()
		// Header
		fmt.Println(strings.Repeat("=", 15))
		fmt.Println("Student Form")
		fmt.Println(strings.Repeat("=", 15))
		// User error
		if len(userError) > 0 {
			fmt.Println(userError)
		}
		// Inputs
		if len(firstName) <= 0 {
			fmt.Println("Enter the first name")
			fmt.Print("> ")
			firstName, _ = reader.ReadString('\n')
			firstName = strings.ToLower(firstName)
			firstName = strings.TrimSpace(strings.Title(firstName))
		}
		// Validate entry
		firstNameIsValid, _ := regexp.MatchString("[a-zA-Z]", firstName)
		if !firstNameIsValid {
			firstName = ""
			userError = "Please, enter a valid name"
			continue
		}

		if len(lastName) <= 0 {
			fmt.Println("Enter the last name")
			fmt.Print("> ")
			lastName, _ = reader.ReadString('\n')
			lastName = strings.ToLower(lastName)
			lastName = strings.TrimSpace(strings.Title(lastName))
		}
		// Validate entry
		lastNameIsValid, _ := regexp.MatchString("[a-zA-Z]", lastName)
		if !lastNameIsValid {
			lastName = ""
			userError = "Please, enter a valid last name"
			continue
		}

		if len(note1) <= 0 {
			fmt.Println("Enter the first note")
			fmt.Print("> ")
			note1, _ = reader.ReadString('\n')
			note1 = strings.ToLower(note1)
			note1 = strings.TrimSpace(note1)
		}
		// Validate entry
		note1IsValid, _ := regexp.MatchString("[0-9]", note1)
		if !note1IsValid {
			note1 = ""
			userError = "Please, enter a valid note"
			continue
		}

		if len(note2) <= 0 {
			fmt.Println("Enter the second note")
			fmt.Print("> ")
			note2, _ = reader.ReadString('\n')
			note2 = strings.ToLower(note2)
			note2 = strings.TrimSpace(note2)
		}
		// Validate entry
		note2IsValid, _ := regexp.MatchString("[0-9]", note2)
		if !note2IsValid {
			note2 = ""
			userError = "Please, enter a valid note"
		}
	}

	s.FirstName = firstName
	s.LastName = lastName
	note1Parsed, _ := strconv.Atoi(note1)
	s.Notes[0] = uint16(note1Parsed)
	note2Parsed, _ := strconv.Atoi(note2)
	s.Notes[1] = uint16(note2Parsed)
}

// ToString shows the student information
func (s *Student) ToString() {
	fmt.Println(strings.Repeat("=", 15))
	fmt.Printf("First Name: %s", s.FirstName)
	fmt.Println("")
	fmt.Printf("Last Name: %s", s.LastName)
	fmt.Println("")
	fmt.Printf("Note 1: %d", s.Notes[0])
	fmt.Println("")
	fmt.Printf("Note 2: %d", s.Notes[1])
	fmt.Println("")
	fmt.Println(strings.Repeat("=", 15))
}
