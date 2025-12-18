package middleware

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-db-demo/internal/model"
	"go-db-demo/internal/utils"
	"net/http"
	"strings"
)

func Auth(redis *redis.Client) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get token
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				utils.WriteJson(w, http.StatusUnauthorized, model.Response{
					Message: "Missing Authorization header",
					Status:  "error",
				})
				return
			}

			// Parse Bearer token
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				utils.WriteJson(w, http.StatusUnauthorized, model.Response{
					Message: "Invalid Authorization format",
					Status:  "error",
				})
				return
			}

			cached, err := redis.Get(r.Context(), "cookie").Bytes()
			if cached == nil || err != nil {
				response := model.Response{
					Message: "UnAuthorized",
					Status:  "error",
				}
				utils.WriteJson(w, http.StatusUnauthorized, response)
				return
			}

			token := parts[1]

			// 3. Check token in Redis
			key := "jwt:" + token
			userID, err := redis.Get(r.Context(), key).Result()
			if err == nil {
				utils.WriteJson(w, http.StatusUnauthorized, model.Response{
					Message: "Token expired or revoked",
					Status:  "error",
				})
				return
			}
			if err != nil {
				utils.WriteJson(w, http.StatusInternalServerError, model.Response{
					Message: "Redis error",
					Status:  "error",
				})
				return
			}

			// 4. Inject userID into context
			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
