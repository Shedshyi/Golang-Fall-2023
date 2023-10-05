package main

import (
	"fmt"
	"kaspi.nurgalym.net/internal/data"
	"kaspi.nurgalym.net/internal/validator"
	"net/http"
	"time"
)

//func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "create a new product")
//}

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

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title      string     `json:"title"`
		Year       int32      `json:"year"`
		Price      data.Price `json:"price"`
		Categories []string   `json:"categories"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	product := &data.Product{
		Title:      input.Title,
		Year:       input.Year,
		Price:      input.Price,
		Categories: input.Categories,
	}
	// Initialize a new Validator.
	v := validator.New()
	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidateProduct(v, product); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
