package basicauth

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber"
)

// Config defines the config for BasicAuth middleware
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(*fiber.Ctx) bool
	// Users defines the allowed credentials
	// Required. Default: map[string]string{}
	Users map[string]string
	// Realm is a string to define realm attribute of BasicAuth.
	// the realm identifies the system to authenticate against
	// and can be used by clients to save credentials
	// Optional. Default: "Restricted".
	Realm string
	// Authorizer defines a function you can pass
	// to check the credentials however you want.
	// It will be called with a username and password
	// and is expected to return true or false to indicate
	// that the credentials were approved or not.
	// Optional. Default: nil.
	Authorizer func(*fiber.Ctx, string, string) bool
	// Unauthorized defines the response body for unauthorized responses.
	// Optional. Default: nil
	Unauthorized func(*fiber.Ctx)
}

func New(config ...Config) func(*fiber.Ctx) {
	// Init config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if cfg.Users == nil {
		cfg.Users = map[string]string{}
	}
	if cfg.Realm == "" {
		cfg.Realm = "Restricted"
	}
	if cfg.Authorizer == nil {
		cfg.Authorizer = func(c *fiber.Ctx, user, pass string) bool {
			if user == "" || pass == "" {
				return false
			}
			return cfg.Users[user] == pass
		}
	}
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) {
			c.Set(fiber.HeaderWWWAuthenticate, "basic realm="+cfg.Realm)
			c.SendStatus(401)
		}
	}
	// Return middleware handler
	return func(c *fiber.Ctx) {
		// Filter request to skip middleware
		if cfg.Filter != nil && cfg.Filter(c) {
			c.Next()
			return
		}
		// Get authorization header
		auth := c.Get(fiber.HeaderAuthorization)
		// Check if header is valid
		if len(auth) > 6 && strings.ToLower(auth[:5]) == "basic" {
			// Try to decode
			if raw, err := base64.StdEncoding.DecodeString(auth[6:]); err == nil {
				// Convert to string
				cred := string(raw)
				// Find semicolumn
				for i := 0; i < len(cred); i++ {
					if cred[i] == ':' {
						// Split into user & pass
						user := cred[:i]
						pass := cred[i+1:]
						// If exist & match in Users, we let him pass
						if cfg.Authorizer(c, user, pass) {
							c.Next()
							return
						}
					}
				}
			}
		}
		// Authentication failed
		cfg.Unauthorized(c)
	}
}
