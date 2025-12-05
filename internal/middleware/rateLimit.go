package middleware

import "github.com/juniorAkp/backend-boilerplate/internal/server"

type RateLimitMiddleware struct {
	server *server.Server
}

func NewRateLimitMiddleware(s *server.Server) *RateLimitMiddleware {
	return &RateLimitMiddleware{server: s}
}

func (r *RateLimitMiddleware) RecordRateLimitHit(endpoint string) {

	if r.server.LoggerService != nil && r.server.LoggerService.GetApplication() != nil {
		app := r.server.LoggerService.GetApplication()
		app.RecordCustomEvent("RateLimitHit", map[string]any{
			"endpoint": endpoint,
		})
	}

}
