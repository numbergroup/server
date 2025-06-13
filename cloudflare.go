package server

import (
	"github.com/gin-gonic/gin"
)

// CloudflareOriginIP gets the origin IP from the Cloudflare header, or falls back to the client IP if the header is not set.
func CloudflareOriginIP(c *gin.Context) string {
	ip := c.GetHeader("CF-Connecting-IP")
	if len(ip) == 0 { // fallback to client ip
		ip = c.ClientIP()
	}
	return ip
}

// CloudflareIPCountry gets the country code from the Cloudflare header "CF-IPCountry",
// or returns "", false if the header is not set.
// The country code is a two-letter ISO 3166-1 alpha-2 code.
func CloudflareIPCountry(c *gin.Context) (string, bool) {
	country := c.GetHeader("CF-IPCountry")
	if len(country) == 0 { // fallback to empty string if not set
		return "", false
	}
	return country, true
}
