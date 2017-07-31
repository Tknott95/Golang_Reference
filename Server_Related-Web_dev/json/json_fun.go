package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
														<html lang="en">
																<head>
																		<meta charset="UTF-8">
																				<title>FOO</title>
																						</head>
																								<body>
																										You are at foo
																												</body>
																														</html>`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "Floyd",
		Lname: "Mayweather",
		Items: []string{"Glass Chin", "Old Age", "Death to his career by Mcgregor"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "Connor",
		Lname: "Mcgregor",
		Items: []string{"Fist", "Kick", "Will take over your village."},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
