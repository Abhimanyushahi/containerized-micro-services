package main

import (
    "fmt"
    "log"
    "net/http"
    )
    
func main(){
    http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World, From Abhimanyu Shahi")
    })

    
    log.Fatal(http.ListenAndServe(":81",nil))
}
