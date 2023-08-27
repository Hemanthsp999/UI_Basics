package main

import (
	"fmt"
	"net/http"
)




func main(){
    // starting web server
    http.ListenAndServe(":8080",nil)
    fmt.Println("Starting server at port 8989")
}
