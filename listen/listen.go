package main

import (
  "fmt"
  "net/http"
  "time"
  "os"
  "os/signal"
)

func main() {
  
  done := make(chan bool, 1) //this is the main service channel

  go signalCtrlC_Handler(done) //register the CTRL+C handler

  go startListening() //start the HTTP server
  
  <-done //wait for the CTRL+C interupt
}

func startListening(){
  http.HandleFunc("/time", timeHandler)
  http.HandleFunc("/urlecho", urlechoHandler)
  
  fmt.Println("listening [:8080]...")
  http.ListenAndServe(":8080", nil)
}

func signalCtrlC_Handler(done chan bool){

  sig := make(chan os.Signal, 1) //signal channel to capture CTRL+C
  signal.Notify(sig, os.Interrupt)
  <-sig

  fmt.Print(" shutting down ")
  for j := 0; j < 5; j++ {
    time.Sleep(time.Second * 1)
    fmt.Print(".")
  }
  fmt.Println(" [done]")

  done <- true
}

func urlechoHandler(w http.ResponseWriter, r *http.Request) {  
  fmt.Fprintf(w, "{%v: %v: %v: %v} ", r.Method, r.URL.Path, r.RemoteAddr, r.TLS)
 }

func timeHandler(w http.ResponseWriter, r *http.Request) {  
  //fmt.Fprintf(w, "{\"now\":\"%v\",\"utc\":\"%v\"}\n", time.Now(), time.Now().UTC())
  fmt.Fprintf(w, "{\"now\":\"%v\"}\n", time.Now())
  fmt.Fprintf(w, "{\"utc\":\"%v\"}", time.Now().UTC())
}
