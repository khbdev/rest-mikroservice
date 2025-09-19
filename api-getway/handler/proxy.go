package handler

import (
	"apiGetWay/consul"
	"apiGetWay/response"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyToService(c *gin.Context) {
	serviceName := c.Param("service")
	path := c.Param("path")

	// 🔹 Validation check
	if ok := runValidation(c, serviceName, path); !ok {
		return
	}

	// 🔹 Service discovery
	addr, err := consul.GetServiceAddress(serviceName)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}

	if addr[len(addr)-1] == '/' && len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	url := addr + path

	// 🔹 Forward request
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, nil, err)
		return
	}

	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		response.JSON(c, http.StatusBadGateway, nil, err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// 🔹 Try decode JSON
	var parsed interface{}
	if err := json.Unmarshal(body, &parsed); err != nil {
		// Agar JSON emas bo‘lsa, raw textni chiqaramiz
		parsed = string(body)
	}

	// 🔹 Normalize response
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		response.JSON(c, resp.StatusCode, parsed, nil)
	} else {
		response.JSON(c, resp.StatusCode, nil, err)
	}
}
