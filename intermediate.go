package main

import (
	"fmt"
	"log"
	"net/http"
)

var actionPage = `
<!DOCTYPE html>
<html>
<head>
    <title>Intermediate Server</title>
</head>
<body>
    <p>Performing action, please wait...</p>
    <script type="text/javascript">
        setTimeout(function() {
            window.location.href = "/perform-action?redirect_uri=%s";
        }, 5000);
    </script>
</body>
</html>
`

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/perform-action", performActionHandler)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect_uri")

	// Simulate a user performing an action
	fmt.Fprintf(w, actionPage, redirectURI)
}

func performActionHandler(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect_uri")

	// Simulate a successful Facebook login and set the user ID
	userID := "exampleUserId" // In a real application, this would come from Facebook
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    userID,
		HttpOnly: true,
		Path:     "/",
	})

	redirectURL := fmt.Sprintf("%s?user_id=%s", redirectURI, userID)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
