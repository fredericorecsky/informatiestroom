package auth

import ( 
    "fmt"
    "net/http"
    "time"

    "github.com/satori/go.uuid"
)

func Logged ( w http.ResponseWriter, r *http.Request ) error {
    c, err := r.Cookie( "session_token" );

    if c != nil {
        fmt.Println ( "not nill" )
    }

    fmt.Println( err )

    if  err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader( http.StatusUnauthorized )
            return err
        }

        w.WriteHeader( http.StatusBadRequest )
        return err
    }

    sessionToken := c.Value
    _ = sessionToken

    fmt.Println( sessionToken )

    return err
}

func Login ( w http.ResponseWriter, r *http.Request) error {

    // If the user is who he said so

    sessionToken, err := uuid.NewV4()

    // Save the token 

    http.SetCookie( w, &http.Cookie{
        Name: "session_token",
        Value: sessionToken.String(),
        Path: "/",
        Expires: time.Now().Add( 120 * time.Second ),
    })

    return err

}

func loginHandler( w http.ResponseWriter, r *http.Request ) {
    err := Login( w, r )

    fmt.Println( err )

    fmt.Fprintf( w, "%s", "Logged" )

    return;
}

func init() {
    http.HandleFunc( "/login/", loginHandler )
}

