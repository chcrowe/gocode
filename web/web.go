package main

import (
  "fmt"
  "net/http"
  "time"
  "os"
  "os/signal"
  "runtime"
  "path/filepath"
)

func main() {

  // memStats := &runtime.MemStats{}
  // runtime.ReadMemStats(memStats)
  // fmt.Printf("Memory: %+v\n", memStats)

  fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())

  done := make(chan bool, 1) //this is the main service channel

  go signalCtrlC_Handler(done) //register the CTRL+C handler

  go startListening() //start the HTTP server
  
  <-done //wait for the CTRL+C interupt
}

func startListening(){
  http.HandleFunc("/time", timeHandler)
  http.HandleFunc("/urlecho", urlechoHandler)
  
  dir, f := filepath.Split(os.Args[0])
  fmt.Printf("%q started [:8080] in folder %q\n", f, dir)
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
  fmt.Fprintf(w, "{Method=%v,\nURL.Path=%v,\nRemoteAddr=%v,\nUserAgent=%v,\nTime=%v} ", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Now().UTC())
 }

func timeHandler(w http.ResponseWriter, r *http.Request) {  
  //fmt.Fprintf(w, "{\"now\":\"%v\",\"utc\":\"%v\"}\n", time.Now(), time.Now().UTC())
  fmt.Fprintf(w, "{\"now\":\"%v\"}\n", time.Now())
  fmt.Fprintf(w, "{\"utc\":\"%v\"}", time.Now().UTC())
}
