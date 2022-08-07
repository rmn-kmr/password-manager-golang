package server

import (
	"errors"
	"net/http"
	"net/url"
	"passwordmanager/data"
	"text/template"

	"github.com/go-chi/chi/v5"
)

var store = data.Store{}

func (app *Config) addNewPassword(w http.ResponseWriter, r *http.Request) {
	// request payload struct
	var rp struct {
		Websitelink string `json:"websiteLink"`
		UserName    string `json:"userName"`
		Password    string `json:"password"`
	}
	// decode json to rp
	err := app.readJSON(w, r, &rp)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	// validation check for Websitelink
	_, err = url.ParseRequestURI(rp.Websitelink)
	if err != nil {
		app.errorJSON(w, errors.New("invalid websitelink"), http.StatusBadRequest)
		return
	}
	// length validation check for Websitelink UserName Password
	if len(rp.Websitelink) == 0 || len(rp.UserName) == 0 || len(rp.Password) == 0 {
		app.errorJSON(w, errors.New("websitelink or userName or password cannot be blank"), http.StatusBadRequest)
		return
	}
	// append new records
	err = store.Password.AddNewRecord(data.PasswordData{Website: rp.Websitelink, UserName: rp.UserName, Password: rp.Password})
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "new password added",
	}
	app.writeJSON(w, http.StatusCreated, res)
}

func (app *Config) GetPasswordList(w http.ResponseWriter, r *http.Request) {
	// fetch password list
	passwordList, err := store.Password.GetPasswordList()
	if err != nil {
		app.errorJSON(w, errors.New("no Data Found"), http.StatusInternalServerError)
		return
	}
	res := jsonResponse{
		Status:  true,
		Message: "Password list fetched",
		Data:    passwordList,
	}
	app.writeJSON(w, http.StatusOK, res)
}

func (app *Config) GetPlainPassword(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	// decode ciphertext
	decodedStr, err := store.Password.GetPlainPassword(key)
	if err != nil {
		app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}
	res := jsonResponse{
		Status: true,
		Data:   decodedStr,
	}
	app.writeJSON(w, http.StatusOK, res)
}

func (app *Config) RenderUI(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	// http.FileServer(http.Dir("../templates/index.html"))
	// http.ServeFile(w, r, "/templates/index.html")

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
