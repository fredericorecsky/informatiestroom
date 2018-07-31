package page 

import (
    "fmt"
    "html/template"
//    "log"
    "net/http"

    "github.com/fredericorecsky/informatiestroom/documents/page"
    "github.com/fredericorecsky/informatiestroom/auth"
)

func errorHandler( w http.ResponseWriter, r *http.Request, status int ) {
    w.WriteHeader( status )
    if status == http.StatusNotFound {
        fmt.Fprint( w, "404 not found" )
    }
    if status == http.StatusUnauthorized {
        fmt.Fprint( w, "401 not authorized" )
    }
}

func pageHandler( w http.ResponseWriter, r *http.Request ) {
    address := r.URL.Path[len("/page/"):]
    
    page := &Document.Page{ Filename: address }
    err := page.Load()

    if err != nil {
        errorHandler(w, r, http.StatusNotFound )
        return
    }

    t , _ := template.ParseFiles("./Templates/page/view.t")
    t.Execute( w, page )
}

func rawHandler( w http.ResponseWriter, r *http.Request ) {
    if auth.Logged( w, r) != nil {
        errorHandler(w, r, http.StatusUnauthorized )
        return
    }
    address := r.URL.Path[len("/raw/"):]

    page := &Document.Page{ Filename: address }
    err := page.Load()

    if err != nil {
        errorHandler(w, r, http.StatusNotFound )
        return
    }

    fmt.Fprintf( w, "%s", page.Body )
}


func init() {
    http.HandleFunc( "/page/", pageHandler )
    http.HandleFunc( "/raw/", rawHandler )
}
