package handler

import (
	"apiGetWay/validation"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)


func runValidation(c *gin.Context, serviceName, path string) bool {
	if serviceName == "userservice" && path == "/users" && c.Request.Method == http.MethodPost {
		var req validation.CreateUserValidation

	
		body, _ := ioutil.ReadAll(c.Request.Body)
		_ = json.Unmarshal(body, &req)


		if err := validation.ValidationUser(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return false
		}

	
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	if serviceName == "productservice" && path == "/products" && c.Request.Method == http.MethodPost {
    var req validation.CreateProductValidation
    body, _ := ioutil.ReadAll(c.Request.Body)
    _ = json.Unmarshal(body, &req)

    if err := validation.ValidationProduct(req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return false
    }

    c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
}

	return true
}
