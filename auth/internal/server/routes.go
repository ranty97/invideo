package server

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
	"html/template"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc("/auth/{provider}/callback", s.getAuthCallback).Methods("GET")
	r.HandleFunc("/logout/{provider}", s.logout).Methods("GET")
	r.HandleFunc("/auth/{provider}", s.getUser).Methods("GET")

	return r
}

func (s *Server) getAuthCallback(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)["provider"]
	log.Println(p)

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		_, err := fmt.Fprintln(w, err)
		if err != nil {
			return
		}
	}

	t, _ := template.New("foo").Parse(userTemplate)
	err = t.Execute(w, user)
	if err != nil {
		return
	}

	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	err := gothic.Logout(w, r)
	if err != nil {
		return
	}
	w.Header().Set("Location", "/")
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		err := t.Execute(w, gothUser)
		if err != nil {
			return
		}
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
