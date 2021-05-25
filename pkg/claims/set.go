package claims

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"strings"
)

func Set(c *auth.Client, uid string, newClaims map[string]interface{}) error {
	claims, err := getClaims(c, uid)
	if err != nil {
		return err
	}
	if claims == nil {
		claims = make(map[string]interface{}, 10)
	}

	for k, v := range newClaims {
		if err := setPath(claims, k, v); err != nil {
			return err
		}
	}

	if err := c.SetCustomUserClaims(context.Background(), uid, claims); err != nil {
		return fmt.Errorf("error setting new claims: %w", err)
	}
	return nil
}

func setPath(claims map[string]interface{}, path string, value interface{}) error {
	pathParts := strings.Split(path, ".")
	for _, part := range pathParts[:len(pathParts) - 1] {
		child, exists := claims[part]
		if !exists {
			child = make(map[string]interface{}, 10)
		}
		subMap, ok := child.(map[string]interface{})
		if !ok {
			return fmt.Errorf("error setting value in path '%s': field '%s' is not an object", path, part)
		}
		claims = subMap
	}

	claims[pathParts[len(pathParts)-1]] = value
	return nil
}
