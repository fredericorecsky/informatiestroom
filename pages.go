package page 

import (
    "fmt"
    "html/template"
    "log"
    "net/http"

    "github.com/fredericorecsky/informatiestroom/Documents"
)

func errorHandler( w http.ResponseWriter, r *http.Request, status int ) {
    w.WriteHeader( status )
    if status == http.StatusNotFound {
        fmt.Fprint( w, "404 not found" )
    }
}

func pageHandler( w http.ResponseWriter, r *http.Request ) {
    address := r.URL.Path[len("/page/"):]
    
    // Load page
    page := &Document.Page{ Filename: address }
    err := page.Load()

    if err != nil {
        errorHandler(w, r, http.StatusNotFound )
        return
    }

    t , _ := template.ParseFiles("./Templates/page.t")
    t.Execute( w, page )
}



func init() {
    http.HandleFunc( "/page/", pageHandler )
    //log.Fatal( http.ListenAndServe(":8080", nil ))
}
