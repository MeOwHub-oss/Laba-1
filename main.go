package main 

import (
  "log"
  "net/http"
)

func main(){
  http.HandleFunc("/",mainHandler)

  log.Fatal(http.ListenAndServe(":8080",nil))


}
