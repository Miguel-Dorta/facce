package claims

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"strings"
)

func Unset(c *auth.Client, uid string, fieldsToRemove []string) error {
	claims, err := getClaims(c, uid)
	if err != nil {
		return err
	}
	if claims == nil {
		return nil
	}

	for _, field := range fieldsToRemove {
		if err := removeClaim(claims, field); err != nil {
			return err
		}
	}

	if err := c.SetCustomUserClaims(context.Background(), uid, claims); err != nil {
		return fmt.Errorf("error setting new claims: %w", err)
	}
	return nil
}

func removeClaim(claims map[string]interface{}, path string) error {
	pathParts := strings.Split(path, ".")
	for _, part := range pathParts[:len(pathParts) - 1] {
		child, exists := claims[part]
		if !exists {
			return nil
		}
		subMap, ok := child.(map[string]interface{})
		if !ok {
			return fmt.Errorf("error removing value in path '%s': field '%s' is not an object", path, part)
		}
		claims = subMap
	}

	delete(claims, pathParts[len(pathParts)-1])
	return nil
}
