package main

import ( 
    "net/http"
    "log"

    _ "github.com/fredericorecsky/informatiestroom/view/page"
)

func main () {
    log.Fatal( http.ListenAndServe(":8080", nil ) )
}
