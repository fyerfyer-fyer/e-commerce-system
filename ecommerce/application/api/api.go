package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/config"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/handler"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"

	"github.com/rs/cors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.RestConf.Timeout = c.Timeout.Http

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 设置限流器
	if c.Middleware.Limit.EnableLocalLimit {
		limiter := limit.NewPeriodLimit(
			int(c.Middleware.Limit.Rate),
			int(c.Middleware.Limit.Burst),
			ctx.Redis,
			"local_limit",
		)
		server.Use(func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				if code, err := limiter.Take(r.URL.Path); err != nil || code == limit.OverQuota {
					http.Error(w, "too many requests", http.StatusTooManyRequests)
					return
				}
				next(w, r)
			}
		})
	}

	// 设置CORS中间件
	if c.CorsConfig.Enable {
		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins:   []string{c.CorsConfig.AllowOrigins},
			AllowedMethods:   []string{c.CorsConfig.AllowMethods},
			AllowedHeaders:   []string{c.CorsConfig.AllowHeaders},
			ExposedHeaders:   []string{c.CorsConfig.ExposeHeaders},
			AllowCredentials: c.CorsConfig.AllowCredentials,
		})
		server.Use(func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				corsMiddleware.Handler(next).ServeHTTP(w, r)
			}
		})
	}

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
