package app

import (
	utils "rest-api-go/utils"
	"net/http"
)

// JwtAuthentication middleware
var JwtAuthentication = func(next http.Handler) {

	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {

		// Endpoints that does not require auth
		publicRoutes := []string{"/api/user/new", "/api/user/login"}

		// current request path
		requestPath := request.URL.Path

		// Verify authentication on routes
		// On succesful authentication send request
		for _, route := range publicRoutes {
			if route == requestPath {
				next.next.ServeHTTP(write, request)
				return
			}
		}

		response := make(map[string]interface{})

		// Get token from header
		tokenHeader := request.Header.Get("Authorization")

		// When token is missing return error code 403 unauthorized
		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			write.WriteHeader(http.StatusForbidden)
			write.Header().add("Content-Type", "application/json")
			utils.response(write, response)
			return
		}

		// Reformat Bearer token
		splitted := string.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			reponse utils.Message(false, "Invalid or Malformed auth token")
			write.WriteHeader(http.StatusForbidden)
			write.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		// Get Token
		tokenPart := splitted[1]
		tk := &modles.Token{}

		token, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})
		
		// Malformed token
		// Return 403
		if err != nil {
		 respone = utils.Messages(false, "Malformed authentication token")	
		 write.WriteHeader(http.StatusForbidden)
		 write.Header().Add("Content-Type", "application/json")
		 utils.Responde(write, response)
		 return
		}

		// Invalid token 
		if !token.Valid {
			response = utils.Message(false, "Token is not valid")
			write.WriteHeader(http.StatusForbidden)
			write.Header().Add("Content-Type", "application/json")
			utils.Responde(write, response)
			return
		} 

		// Valid Token 

		// Log username
		fmt.Sprintf("User %", tk.Username)
		ctx := context.WithValue(response.Context, "user", tk.UserId)
		response = response.WithContext(ctx)
		next.ServeHttp(write, response)

	})

}
