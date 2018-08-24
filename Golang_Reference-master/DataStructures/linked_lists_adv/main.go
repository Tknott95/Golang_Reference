package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name string
	age  int
	ssn  string
	Next *Student
}

func main() {

	var students *Student

	students = nil

	if len(os.Args) < 2 {
		println("There is no file")
		return
	}

	filename := os.Args[1]

	data, _ := ioutil.ReadFile(filename)

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		td := strings.Split(string(line), " ")

		tempStu := &Student{
			name: td[0],
			ssn:  td[2],
			Next: students,
		}

		tempStu.age, _ = strconv.Atoi(td[1])

		students = tempStu
	}

	println("Our list of students\n")

	/* while here is a next student keep itterating over unitl we hit the end (When next student is nil) */

	for s := students; s != nil; s = s.Next {
		println(s.name, s.age, s.ssn)
	}

	println("\nOur REVRESED list of students\n")

	/* while here is a next student keep itterating over unitl we hit the end (When next student is nil) */

	students = reverse(students)

	for s := students; s != nil; s = s.Next {
		println(s.name, s.age, s.ssn)
	}

}

func reverse(curr *Student) *Student {
	if curr.Next == nil {
		return curr
	} else {
		newHead := reverse(curr.Next)
		curr.Next.Next = curr
		curr.Next = nil
		return newHead
	}
}
