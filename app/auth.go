package app

import "net/http"

// JwtAuthentication middleware
var JwtAuthentication = func(next http.Handler) {

	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {

		// Endpoints that does not require auth
		publicRoutes := []string{"/api/user/new", "/api/user/login"}

		// current request path
		requestPath := request.URL.Path

		// Verify authentication
		// On succesful authentication send request
		for _, route := range publicRoutes {
			if route == requestPath {
				next.next.ServeHTTP(write, request)
				return
			}
		}
	})

}
