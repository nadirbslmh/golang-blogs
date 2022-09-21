package controller

import (
	"encoding/json"
	"golang-blogs/model"
	"golang-blogs/service"
	"io"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	blogService := service.New()

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	var input model.BlogInput = model.BlogInput{}

	json.Unmarshal(reqBody, &input)

	blogService.Create(input)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("success"))
}
