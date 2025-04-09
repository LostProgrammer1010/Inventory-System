package post

import "net/http"

func checkMethod(r *http.Request) bool {
	return r.Method != http.MethodPost
}
