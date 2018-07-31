package page

import (
    "fmt"
    "html/template"
    "net/http"

    "github.com/fredericorecsky/informatiestroom/auth"
    "github.com/fredericorecsky/informatiestroom/documents/page"
//    "github.com/fredericorecsky/informatiestroom/httperror"
)

// TODO lib
func errorHandler( w http.ResponseWriter, r *http.Request, status int ) {
    w.WriteHeader( status )
    if status == http.StatusNotFound {
        fmt.Fprint( w, "404 not found" )
    }
    if status == http.StatusUnauthorized {
        fmt.Fprint( w, "401 not authorized" )
    }
}

func editHandler( w http.ResponseWriter, r *http.Request ) {
    if auth.Logged( w, r )  != nil {
        errorHandler( w, r, http.StatusUnauthorized )
        return
    }

    address := r.URL.Path[len("/edit/"):]

    fmt.Println( "Stay clean" )

    if r.Method == "POST" {
        errorHandler( w, r, http.StatusNotFound )
        return
    }else{
        page := &Document.Page{ Filename: address }
        err := page.Load()

        fmt.Println( address )
        fmt.Println( err )
        if err != nil {
            errorHandler( w, r, http.StatusNotFound )
            return
        }
        t , err := template.ParseFiles("./Templates/page/edit.t")
        fmt.Println( err )
        t.Execute( w, page )
    }
}

func init() {
    http.HandleFunc( "/edit/", editHandler )
}
