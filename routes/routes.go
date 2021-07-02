package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"timetracker/db"
	"timetracker/github"
	"timetracker/helpers"

	"github.com/gorilla/mux"
)

// Initialises the HTTP router and pipeline, then listen and serves application
func ListenAndServe() {
	router := configureRouter()
	port := ":" + os.Getenv("PORT")
	log.Printf("Listening on http://localhost%s/", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// Configures the Mux Router for serving the front end and the API
func configureRouter() *mux.Router {
	router := mux.NewRouter()

	// Add any additional middleware
	router.Use(PanicHandler)
	router.Use(CheckAuthHandler)

	// Configure any API routes
	router.HandleFunc("/api/github/url", getGitHubLoginUrl).Methods(http.MethodGet)
	router.HandleFunc("/api/github/login", getGitHubAccessToken).Methods(http.MethodPost)
	router.HandleFunc("/api/github/search", searchRepos).Methods(http.MethodGet)
	router.HandleFunc("/api/github/repo/{owner}/{repo}/branches", getBranches).Methods(http.MethodGet)
	router.HandleFunc("/api/github/repo/{owner}/{repo}/commits", getCommits).Methods(http.MethodGet)

	router.HandleFunc("/api/user", getUser).Methods(http.MethodGet)
	router.HandleFunc("/api/auth/isAuthenticated", isAuthenticated).Methods(http.MethodGet)
	router.HandleFunc("/api/time", timeRouteHandler).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/api/time/{id}", timeRouteHandler).Methods(http.MethodPatch, http.MethodDelete)
	router.HandleFunc("/api/time/tags", getTags).Methods(http.MethodGet)

	// Configure the static file serving for the SPA
	// This must be configured after API routes to stop any /api/
	// requests being redirected to our SPA
	spa := SpaHandler{staticPath: "./ClientApp/dist/", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	return router
}

// Reads the body of the http request to a byte array, handles any errors that may occur
func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleError(err)

	return body
}

// Correctly formats and encodes an API response
func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(call)
}

// Parses the token from the cookie and gets the id of currently logged in user
func getUserId(r *http.Request) uint {
	token, err := parseTokenFromCookie(r)
	helpers.HandleError(err)
	ct, err := github.CheckToken(token)
	helpers.HandleError(err)
	return db.GetUserId(*ct.User.Login)
}
