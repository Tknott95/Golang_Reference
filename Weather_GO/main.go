package main

import (
  "flag"
  "fmt"
  "net/http"
)

var zipcode int
var geocodeUrl = `https://maps.googleapis.com/maps/api/geocode/json?&address=${zipcode}`;

func init() {
  flag.IntVar(&zipcode ,"z", 80203, "Zipcode flag to store into var")
  flag.Parse()
}

func main() {
   fmt.Println("Initiliazing Weather via:", zipcode )

   req, err := http.NewRequest(http.MethodGet, geocodeUrl,, nil)

}
