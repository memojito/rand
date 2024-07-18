package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/memojito/igapi/types"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func ParseJSON(r *http.Request, v any) error {
	if r.Body == nil {
		return types.APIError{
			Status: 403,
			Msg:    "missing body",
		}
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}

func PrepareWhereINString(count int) string {
	var paramrefs string
	for i := 0; i < count; i++ {
		paramrefs += `$` + strconv.Itoa(i+1) + `,`
	}
	paramrefs = paramrefs[:len(paramrefs)-1]
	return paramrefs
}
