package handlers

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	var innerHtml = `
	<html>
		<body>
			<a href="/auth/google/login" style="text-transform:none">
                            <div class="left">
                                <img width="30px" alt="Google &quot;G&quot; Logo"
                                     src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/512px-Google_%22G%22_Logo.svg.png"/>
                            </div>
                            Login with Google
                        </a>
		</body>
	</html>
	`
	fmt.Fprintf(w, innerHtml)
}

func (server *Server) handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (server *Server) handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "Content: %s\n", content)
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}