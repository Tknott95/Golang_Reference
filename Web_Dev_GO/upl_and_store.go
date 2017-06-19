package main

import (
 "io"
 "fmt"
 "io/ioutil"
 "net/http"
)

func main() {
  http.HandleFunc("/", uplIndex)
  http.ListenAndServe(":8080", nil)
}

func uplIndex(w http.ResponseWriter, req *http.Request) {
  var s string
  fmt.Println(req.Method)
  if req.Method == http.MethodPost {

   // open file
   f, h, err := req.FormFile("upl")
   if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }
  
   defer f.Close()

   fmt.Println("\file:", f, "\nheader:", h, "\nerr", err)  
  
  // read file
  bs, err := ioutil.ReadAll(f)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  
  s = string(bs)

  // store file on server
  dst, err := os.Create(filepath.Join("./files/", h.Filename))
  
  if err != nil {
   http.Error(w, err.Error(), http.StatusInternalServerError)
   return
  }
 
  defer dst.Close()

  _, err = dst.Write(bs)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

}

 w.Header().Set("Content-Type", "text/html; charset=utf-8")
    io.WriteString(w, `
    <form method="POST" enctype="multipart/form-data">
    <input type="file" name="upl">
    <input type="submit">
    </form>
    <br>`+s)
  
}
