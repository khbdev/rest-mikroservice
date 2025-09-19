package consul

import (
	"fmt"
)

var serviceMap = map[string]string{
	"userservice": "http://localhost:8088",
	"productservice": "http://localhost:8087",
}

func GetServiceAddress(serviceName string) (string, error) {
	addr, ok := serviceMap[serviceName]
	if !ok {
		return "", fmt.Errorf("service %s not found", serviceName)
	}
	return addr, nil
}
