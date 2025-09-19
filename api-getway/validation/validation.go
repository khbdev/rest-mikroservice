package validation

import (
	"errors"
	"net/url"
	"strings"
)



type CreateUserValidation struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateProductValidation struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
}
func ValidationUser(req CreateUserValidation) error {
	if strings.TrimSpace(req.Name) == "" {
		return  errors.New("name is required")
	}

	if !strings.Contains(req.Email, "@") {
		return  errors.New("Invalid email")
	}


if strings.TrimSpace(req.Password) == "" {
    return errors.New("password is required")
}

if len(req.Password) < 6 {
    return errors.New("password must be at least 6 characters")
}
return  nil
}


func ValidationProduct(req CreateProductValidation) error {

	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	
	if req.Price <= 0 {
		return errors.New("price must be greater than 0")
	}


	if strings.TrimSpace(req.Image) == "" {
		return errors.New("image is required")
	}

	
	parsedURL, err := url.ParseRequestURI(req.Image)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return errors.New("image must be a valid URL (http/https)")
	}

	return nil
}