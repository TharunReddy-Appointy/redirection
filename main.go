package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Main Server</title>
</head>
<body>
    <a href="/start-login">Book Now</a>
</body>
</html>
`))

var callbackTpl = template.Must(template.New("callback").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Callback</title>
</head>
<body>
    <p>User is authenticated, user_id: {{.UserID}}</p>
</body>
</html>
`))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/start-login", startLoginHandler)
	http.HandleFunc("/callback", callbackHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tpl.Execute(w, nil)
}

func startLoginHandler(w http.ResponseWriter, r *http.Request) {
	redirectURL := "http://localhost:3001/login?redirect_uri=http://localhost:3000/callback"
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID != "" {
		// Set the cookie in the response
		http.SetCookie(w, &http.Cookie{
			Name:     "user_id",
			Value:    userID,
			HttpOnly: true,
			Path:     "/",
		})
		w.Header().Set("Content-Type", "text/html")
		callbackTpl.Execute(w, map[string]string{"UserID": userID})
	} else {
		http.Error(w, "User ID not found", http.StatusBadRequest)
	}
}
