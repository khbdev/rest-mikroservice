package handler

import (
	"apiGetWay/consul"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)


func ProxyToService(c *gin.Context) {
	serviceName := c.Param("service")
	path := c.Param("path")

	addr, err := consul.GetServiceAddress(serviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if addr[len(addr)-1] == '/' && len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	url := addr + path

	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// headers ko‘chirish
	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
