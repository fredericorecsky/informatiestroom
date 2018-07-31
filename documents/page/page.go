package Document 

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
)

type Page struct{
    Filename string
    Title string
    Body []byte
    Json []byte
}

func ( p *Page) String() string {
    return fmt.Sprintf( "%s", p.Body )
}

func ( p *Page ) Error() string {
    return fmt.Sprintf( "Error %s", p.Body )
}

func (p *Page) Save() error {
    b, _ := json.Marshal( p )

    p.Json = []byte(b)

    return ioutil.WriteFile( p.Filename, p.Json, 0600 )
}

func (p *Page) Load() error {
    body, err := ioutil.ReadFile( p.Filename )

    if err != nil {
        return err;
    }

    err = json.Unmarshal( body, &p )

    fmt.Println( p.Body )

    return err 

}
