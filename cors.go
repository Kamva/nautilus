package nautilus

import (
	"strings"

	"github.com/kataras/iris"
)

// CORS is struct containing data of CORS related headers.
type CORS struct {
	AllowedOrigins     []string
	AllowedMethods     []string
	AllowedHeaders     []string
	DisableCredentials bool
}

// Handle is a middleware for adding CORS headers in response.
func (c CORS) Handle(context iris.Context) {
	origins := "*"
	methods := "GET,HEAD,OPTIONS,POST,PUT,PATCH,DELETE"
	headers := "Accept,Authorization,Cache-Control,Content-Type,X-Requested-With"
	allowCredentials := "true"

	if len(c.AllowedOrigins) > 0 {
		origins = strings.Join(c.AllowedOrigins, ",")
	}

	if len(c.AllowedMethods) > 0 {
		methods = strings.Join(c.AllowedMethods, ",")
	}

	if len(c.AllowedHeaders) > 0 {
		headers = strings.Join(c.AllowedHeaders, ",")
	}

	if c.DisableCredentials {
		allowCredentials = "false"
	}

	context.Header("Access-Control-Allow-Origin", origins)
	context.Header("Access-Control-Allow-Methods", methods)
	context.Header("Access-Control-Allow-Headers", headers)
	context.Header("Access-Control-Allow-Credentials", allowCredentials)
	context.Next()
}
