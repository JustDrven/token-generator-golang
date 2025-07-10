package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"petrnemecek.dev/token-generator/network"
	"petrnemecek.dev/token-generator/utils"
)

func Make(pattern string, handler func(http.ResponseWriter, *http.Request))  {
	http.HandleFunc(pattern, handler)
}

func Init()  {
	Make("/api/token/generate", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var length = 10
		value, err := strconv.Atoi(r.Header.Get("token-length"))
		if (err == nil) {
			length = value
		}
		
		json := json.NewEncoder(w)

		token, err := utils.GenerateToken(length)

		if err != nil {
			json.Encode(&network.Error{
				Error: true,
				Message: err.Error(),
			})
			return
		}


		json.Encode(&network.Token{
			Token: token,
		})

		
	})
}