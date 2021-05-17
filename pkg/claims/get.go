package claims

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"strings"
)

func Get(c *auth.Client, uid string, path string) (string, error) {
	claims, err := getClaims(c, uid)
	if err != nil {
		return "", err
	}

	var result interface{}
	if path == "" {
		result = claims
	} else {
		result, err = getPath(claims, path)
		if err != nil {
			return "", err
		}
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formating JSON type (%+v): %w", result, err)
	}
	return string(data), nil
}

func getClaims(c *auth.Client, uid string) (map[string]interface{}, error) {
	usr, err := c.GetUser(context.Background(), uid)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return usr.CustomClaims, nil
}

func getPath(m map[string]interface{}, path string) (interface{}, error) {
	pathParts := strings.Split(path, ".")
	var result interface{} = m
	for _, part := range pathParts {
		resultMap, ok := result.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("path not found: %s", path)
		}
		result = resultMap[part]
	}
	return result, nil
}
