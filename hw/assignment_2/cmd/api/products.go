package main

import (
	"fmt"
	"kaspi.nurgalym.net/internal/data"
	"net/http"
	"time"
)

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new product")
}

func (app *application) showProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	product := data.Product{
		ID:         id,
		CreatedAt:  time.Now(),
		Title:      "Iphone 15",
		Price:      500000,
		Categories: []string{"technic", "phone", "iphone"},
		Version:    1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		//app.logger.Println(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	//fmt.Fprintf(w, "show the details of product %d\n", id)
}
