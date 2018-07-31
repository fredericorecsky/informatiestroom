package main 

import (
    "fmt"
    "github.com/fredericorecsky/informatiestroom/Documents"
)

func main() {
    page := &Document.Page{ Filename: "im", Title: "Test", Body: []byte("this is a test")}

    fmt.Println( string(page.Body) )

    page.Save()
    fmt.Printf( "It works [%s]\n", page.Json )

    page2 := &Document.Page{ Filename: "im" }
    page2.Load()

    fmt.Printf( "It works [%s] also \n", page2.Body )

}
