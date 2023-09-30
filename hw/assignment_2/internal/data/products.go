package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type Product struct {
	ID         int64     `json:"id"`                     // Unique integer ID for the product
	CreatedAt  time.Time `json:"-"`                      // Timestamp for when the product is added to our database
	Title      string    `json:"title"`                  // Product title
	Year       int32     `json:"year,omitempty"`         // Product release year
	Price      int32     `json:"price,omitempty,string"` // Product price (in tenge)
	Categories []string  `json:"categories,omitempty"`   // Product category (technic, for home, etc.)
	Version    int32     `json:"version"`                // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}

func (p Product) MarshalJSON() ([]byte, error) {
	// Create a variable holding the custom runtime string, just like before.
	var price string
	if p.Price != 0 {
		price = fmt.Sprintf("%d tenge", p.Price)
	}
	// Define a MovieAlias type which has the underlying type Movie. Due to the way that
	// Go handles type definitions (https://golang.org/ref/spec#Type_definitions) the
	// MovieAlias type will contain all the fields that our Movie struct has but,
	// importantly, none of the methods.
	type ProductAlias Product
	// Embed the MovieAlias type inside the anonymous struct, along with a Runtime field
	// that has the type string and the necessary struct tags. It's important that we
	// embed the MovieAlias type here, rather than the Movie type directly, to avoid
	// inheriting the MarshalJSON() method of the Movie type (which would result in an
	// infinite loop during encoding).
	aux := struct {
		ProductAlias
		Price string `json:"price,omitempty"`
	}{
		ProductAlias: ProductAlias(p),
		Price:        price,
	}
	return json.Marshal(aux)
}
