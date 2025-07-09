package package_http

import (
	"fmt"

	slog "github.com/salamsites/package-log"

	"net/http"
	"strconv"
)

type Middleware struct {
	logger  *slog.Logger
	jwtKey  string
	limiter *RateLimiter
}

type appBaseHandler func(w http.ResponseWriter, r *http.Request) Response
type appAuthHandler func(w http.ResponseWriter, r *http.Request, claims AuthClaims) Response

type AuthClaims struct {
	Id       int64
	DeviceId int64
}

func NewMiddleware(logger *slog.Logger, jwtKey string, limiter *RateLimiter) *Middleware {
	return &Middleware{
		logger:  logger,
		jwtKey:  jwtKey,
		limiter: limiter,
	}
}

func (middleware *Middleware) Base(h appBaseHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		result := h(w, r)

		w.WriteHeader(result.GetStatusCode())
		w.Write(result.Marshal())
	}
}

func (middleware *Middleware) Auth(h appAuthHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("authorization")

		claims, err := TokenClaims(token, middleware.jwtKey)

		if err != nil {

			if err.Error() == "Token is expired" {
				w.WriteHeader(http.StatusNotAcceptable)
			}

			middleware.logger.Error("shttp error jwt: ", err)

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authClaims := AuthClaims{}

		idStr := fmt.Sprint(claims["user_id"])
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			fmt.Println("2 err: ", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		deviceIdStr := fmt.Sprint(claims["device_id"])
		deviceId, err := strconv.ParseInt(deviceIdStr, 10, 64)
		if err != nil {
			fmt.Println("2 err: ", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authClaims.Id = id
		authClaims.DeviceId = deviceId

		result := h(w, r, authClaims)

		w.WriteHeader(result.GetStatusCode())
		w.Write(result.Marshal())
	}

}

func (middleware *Middleware) PAuth(h appAuthHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("authorization")

		claims, err := TokenClaims(token, middleware.jwtKey)

		if err != nil {

			if err.Error() == "Token is expired" {
				w.WriteHeader(http.StatusNotAcceptable)
			}

			middleware.logger.Error("shttp error jwt: ", err)

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authClaims := AuthClaims{}
		fmt.Println(claims["id"])
		idStr := fmt.Sprint(claims["id"])
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			fmt.Println("2 err: ", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authClaims.Id = id

		result := h(w, r, authClaims)

		w.WriteHeader(result.GetStatusCode())
		w.Write(result.Marshal())
	}

}
