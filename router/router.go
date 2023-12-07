package router

import (
	"account-service/config"
	impl_controller "account-service/controller/implController"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Router() http.Handler {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler)

	register := impl_controller.RegisterCotrollerInit()

	r.Route("/api/v1", func(v1 chi.Router) {
		v1.Route("/public", func(public chi.Router) {
			public.Post("/send_info", register.SendInfoRegister)
			public.Post("/confirm_code", register.ConfirmCodeRegister)
		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.GetAppPort()+"/swagger/doc.json"),
	))

	log.Println("http://localhost:" + config.GetAppPort())
	log.Println("http://localhost:" + config.GetAppPort() + "/swagger/index.html")

	return r
}
