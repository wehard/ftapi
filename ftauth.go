package ftauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

type CodeRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
	State        string `json:"state"`
}

type ClientCredentials struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	CreatedAt   int    `json:"created_at"`
}

var oauthConfig *oauth2.Config
var oauthStateString = "kakkorvarflygerhem"

func Init() {
	endpoint := oauth2.Endpoint{AuthURL: "https://api.intra.42.fr/oauth/authorize"}
	oauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("HM_CLIENT_ID"),
		ClientSecret: os.Getenv("HM_CLIENT_SECRET"),
		Scopes:       []string{},
		Endpoint:     endpoint,
	}
	fmt.Println("Open http://localhost:8080 to continue.")
}

func RequestAuth() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var html = `<html>
<body>
	<a href="/login">42 Auth login</a>
</body>
</html>`
	fmt.Fprintf(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	token := getAccessToken(r.FormValue("state"), r.FormValue("code"))
	fmt.Println("token: ", token)
}

func getAccessToken(state string, code string) string {
	if state != oauthStateString {
		fmt.Println("invalid oauth state")
	}
	codeRequest := CodeRequest{
		GrantType:    "authorization_code",
		ClientID:     os.Getenv("HM_CLIENT_ID"),
		ClientSecret: os.Getenv("HM_CLIENT_SECRET"),
		Code:         code,
		RedirectURI:  oauthConfig.RedirectURL,
		State:        oauthStateString,
	}
	jsonBytes, _ := json.Marshal(codeRequest)
	response, err := http.Post("https://api.intra.42.fr/oauth/token", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Println("failed getting user token!")
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("could not read response body!")
	}
	var clientCredentials ClientCredentials
	err = json.Unmarshal(contents, &clientCredentials)
	if err != nil {
		fmt.Println("could not unmarshal body content!")
	}
	return clientCredentials.AccessToken
}
