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

	students := new(Student)

	students.Next = nil

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
			Next: students.Next,
		}

		tempStu.age, _ = strconv.Atoi(td[1])

		students.Next = tempStu
	}

	println("Our list of students\n")

	/* while here is a next student keep itterating over unitl we hit the end (When next student is nil) */

	for s := students.Next; s != nil; s = s.Next {
		println(s.name, s.age, s.ssn)
	}

}
