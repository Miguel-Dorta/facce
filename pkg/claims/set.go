package claims

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
)

func Set(c *auth.Client, uid string, newClaims map[string]interface{}) error {
	claims, err := getClaims(c, uid)
	if err != nil {
		return err
	}

	if claims == nil {
		claims = newClaims
	} else {
		for k, v := range newClaims {
			claims[k] = v
		}
	}

	if err := c.SetCustomUserClaims(context.Background(), uid, claims); err != nil {
		return fmt.Errorf("error setting new claims: %w", err)
	}
	return nil
}
