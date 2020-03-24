package baseAuth

import (
	"net/http"
)

func BasicAuth(r *http.Request,USERNAME string,PASSWORD string) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		//w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		//w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}
