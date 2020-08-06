package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/nayonacademy/golang-oauth2/api/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	googleOauthConfig *oauth2.Config
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

func (server *Server) handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (server *Server) handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	//content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	state := r.FormValue("state")
	code := r.FormValue("code")
	if state != oauthStateString {
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var info models.UserData

	contents, err := ioutil.ReadAll(response.Body)

	json.Unmarshal(contents, &info)

	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	var user models.User

	if err := server.DB.First(&user, "user_id = ?", info.Id).Error; gorm.IsRecordNotFoundError(err) {
		// record not found
		server.DB.Create(&models.User{
			Token:    token.AccessToken,
			Email:    info.Email,
			Picture:  info.Picture,
			UserID:   info.Id,
		})
	}

	testValue := server.DB.First(&user, "user_id = ?", info.Id).Error
	if testValue != nil{
		fmt.Println("already")
		fmt.Fprintf(w, "%s Already Registered", info.Email)
		server.DB.Model(&user).Where("user_id = ?", info.Id).Update("token", token.AccessToken)
	}else{
		fmt.Println("new")
		fmt.Fprintf(w, "Welcome %s , Congrasulation for New Registration.", info.Email)
	}

}
