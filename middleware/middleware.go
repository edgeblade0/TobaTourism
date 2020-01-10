package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	u "github.com/TobaTourism/pkg/common/util"
	"github.com/TobaTourism/pkg/models"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// InitMiddleware intialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT")
		// c.Response().Header().Set("Access-Control-Allow-Headers", "*")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-Custom-Header, Upgrade-Insecure-Requests")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		// Access-Control-Allow-Credentials
		if c.Request().Method == "OPTIONS" {
			// c.Response().Header()
			return c.JSON(http.StatusOK, "ok")
		}

		return next(c)
	}
}

func (m *GoMiddleware) JwtAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {

		notAuth := []string{"/api/restaurant", "/api/culinary"}
		requestPath := c.Request().URL.Path
		log.Println(requestPath)
		for _, value := range notAuth {
			if value == requestPath {
				next(c)
				return nil
			}
		}

		response := make(map[string]interface{})
		tokenHeader := c.Request().Header.Get("Authorization")
		log.Println(tokenHeader)

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = u.Message(false, "Missing auth token")
			c.Response().Header().Set(`X-Cursor`, "header")
			return c.JSON(http.StatusForbidden, response)
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			c.Response().Header().Set(`X-Cursor`, "header")
			return c.JSON(http.StatusForbidden, response)
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = u.Message(false, "Malformed authentication token")
			c.Response().Header().Set(`X-Cursor`, "header")
			return c.JSON(http.StatusForbidden, response)
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid.")
			c.Response().Header().Set(`X-Cursor`, "header")
			return c.JSON(http.StatusForbidden, response)
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.UserId) //Useful for monitoring
		r := c.Request().WithContext(context.WithValue(c.Request().Context(), "user", tk.UserId))
		c.SetRequest(r)
		next(c) //proceed in the middleware chain!

		return nil
	})
}

// var JwtAuthentication = func(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		notAuth := []string{"/api/culinary"} //List of endpoints that doesn't require auth
// 		requestPath := r.URL.Path            //current request path

// 		//check if request does not need authentication, serve the request if it doesn't need it
// 		for _, value := range notAuth {

// 			if value == requestPath {
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 		}

// 		response := make(map[string]interface{})
// 		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

// 		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
// 			response = u.Message(false, "Missing auth token")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Add("Content-Type", "application/json")
// 			u.Respond(w, response)
// 			return
// 		}

// 		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
// 		if len(splitted) != 2 {
// 			response = u.Message(false, "Invalid/Malformed auth token")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Add("Content-Type", "application/json")
// 			u.Respond(w, response)
// 			return
// 		}

// 		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
// 		tk := &models.Token{}

// 		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
// 			return []byte(os.Getenv("token_password")), nil
// 		})

// 		if err != nil { //Malformed token, returns with http code 403 as usual
// 			response = u.Message(false, "Malformed authentication token")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Add("Content-Type", "application/json")
// 			u.Respond(w, response)
// 			return
// 		}

// 		if !token.Valid { //Token is invalid, maybe not signed on this server
// 			response = u.Message(false, "Token is not valid.")
// 			w.WriteHeader(http.StatusForbidden)
// 			w.Header().Add("Content-Type", "application/json")
// 			u.Respond(w, response)
// 			return
// 		}

// 		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
// 		fmt.Sprintf("User %", tk.UserId) //Useful for monitoring
// 		ctx := context.WithValue(r.Context(), "user", tk.UserId)
// 		r = r.WithContext(ctx)
// 		next.ServeHTTP(w, r) //proceed in the middleware chain!
// 	})
// }
