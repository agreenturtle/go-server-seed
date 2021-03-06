package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
    code  = ""
    token = ""
)

// Your credentials should be obtained from the Google
// Developer Console (https://console.developers.google.com).
var oauthCfg = &oauth2.Config{
  ClientID:     "CLIENT_ID",
  ClientSecret: "CLIENT_SECRET",
  RedirectURL:  "http://localhost:8000/auth/google/callback",
  Scopes: []string{"profile"},
  Endpoint: google.Endpoint,
}

func GoogleHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
  http.Redirect(w,r, oauthCfg.AuthCodeURL("state"), http.StatusMovedPermanently)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
    code := r.FormValue("code")

    token, err := oauthCfg.Exchange(oauth2.NoContext, code)
    if err != nil {
				log.Println("Failed Login")
        http.Redirect(w, r, "/fail", http.StatusMovedPermanently)
    } else {
      fmt.Println(token)
      http.Redirect(w, r, "/success", http.StatusMovedPermanently)
    }
}

func main() {
    http.HandleFunc("/auth/google/", GoogleHandler)

    http.HandleFunc("/auth/google/callback", GoogleCallbackHandler)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "views/login/index.html")
    })

    http.HandleFunc("/success/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "views/success/index.html")
    })

    http.HandleFunc("/fail/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "views/fail/index.html")
    })

    log.Fatal(http.ListenAndServe(":8000", nil))
}
