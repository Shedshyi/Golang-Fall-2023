package main

import (
	"errors"
	"fmt"
	"kaspi.nurgalym.net/internal/data"
	"kaspi.nurgalym.net/internal/validator"
	"net/http"
)

//func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "create a new product")
//}

func (app *application) showProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	product, err := app.models.Products.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	//product := data.Product{
	//	ID:         id,
	//	CreatedAt:  time.Now(),
	//	Title:      "Iphone 15",
	//	Price:      500000,
	//	Categories: []string{"technic", "phone", "iphone"},
	//	Version:    1,
	//}
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

	if data.ValidateProduct(v, product); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Products.Insert(product)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/products/%d", product.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"product": product}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	product, err := app.models.Products.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title      *string     `json:"title"`
		Year       *int32      `json:"year"`
		Price      *data.Price `json:"price"`
		Categories []string    `json:"categories"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//product.Title = *input.Title
	//product.Year = *input.Year
	//product.Price = *input.Price
	//product.Categories = input.Categories

	if input.Title != nil {
		product.Title = *input.Title
	}
	if input.Year != nil {
		product.Year = *input.Year
	}
	if input.Price != nil {
		product.Price = *input.Price
	}
	if input.Categories != nil {
		product.Categories = input.Categories // Note that we don't need to dereference a slice.
	}

	v := validator.New()
	if data.ValidateProduct(v, product); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Products.Update(product)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Products.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Return a 200 OK status code along with a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "product successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listProductsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title      string
		Categories []string
		data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Categories = app.readCSV(qs, "categories", []string{})
	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "price", "-id", "-title", "-year", "-price"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	products, metadata, err := app.models.Products.GetAll(input.Title, input.Categories, input.Filters)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"products": products, "metadata": metadata}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
