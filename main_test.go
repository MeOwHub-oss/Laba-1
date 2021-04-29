package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
)


func TestHandler(t *testing.T){
  req,err := http.NewRequest("GET","http://localhost:8080/",nil)
  if err != nil{
    t.Fatalf("could not create request: %v",err)
  }
  
  rec := httptest.NewRecorder()
  handler := http.HandlerFunc(mainHandler)
  handler.ServeHTTP(rec,req)
  
  assert.HTTPStatusCode(t,handler,"GET","/",nil,http.StatusOK)

}

func TestHandlerGetOrder(t *testing.T){
  req,err := http.NewRequest("GET","http://localhost:8080/1",nil)
  if err != nil{
    t.Fatalf("could not create request: %v",err)
  }
  
  rec := httptest.NewRecorder()
  handler := http.HandlerFunc(mainHandler)
  handler.ServeHTTP(rec,req)
  
  assert.HTTPStatusCode(t,handler,"GET","/1",nil,http.StatusOK)
}
func TestHandlerDeleteOrder(t *testing.T){
  req,err := http.NewRequest("DELETE","http://localhost:8080/1",nil)
  if err != nil{
    t.Fatalf("could not create request: %v",err)
  }
  
  rec := httptest.NewRecorder()
  handler := http.HandlerFunc(mainHandler)
  handler.ServeHTTP(rec,req)
  
  assert.HTTPStatusCode(t,handler,"DELETE","/1",nil,http.StatusOK)
}
func TestHandlerPutOrder(t *testing.T){
  req,err := http.NewRequest("PUT","http://localhost:8080/1",nil)
  if err != nil{
    t.Fatalf("could not create request: %v",err)
  }
  
  rec := httptest.NewRecorder()
  handler := http.HandlerFunc(mainHandler)
  handler.ServeHTTP(rec,req)
  
  assert.HTTPStatusCode(t,handler,"PUT","/1",nil,http.StatusBadRequest)
}
