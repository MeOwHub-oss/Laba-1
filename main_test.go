package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
)

 func TestHandler(t *testing.T){
   handler := http.HandlerFunc(mainHandler)
    assert.HTTPStatusCode(t,handler,"GET","/",nil,http.StatusOK)
  }
  func TestHandlerGetById(t *testing.T){
    handler := http.HandlerFunc(mainHandler)
   assert.HTTPStatusCode(t,handler,"GET","/1",nil,http.StatusOK)
  }
  func TestHandlerDeletById(t *testing.T){
    handler := http.HandlerFunc(mainHandler)
    assert.HTTPStatusCode(t,handler,"DELETE","/1",nil,http.StatusOK)
 }
  func TestHandlerPut(t *testing.T){
    handler := http.HandlerFunc(mainHandler)
    assert.HTTPStatusCode(t,handler,"PUT","/1",nil,http.StatusInternalServerError)
  }
  func TestHandlerPost(t *testing.T){
    handler := http.HandlerFunc(mainHandler)
    assert.HTTPStatusCode(t,handler,"POST","/1?id=1",nil,http.StatusBadRequest)
  }
