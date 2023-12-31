package router

import (
	"account-service/config"
	impl_controller "account-service/controller/implController"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Router() http.Handler {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler)

	registerController := impl_controller.RegisterCotrollerInit()
	loginController := impl_controller.LoginControllerInit()
	projectController := impl_controller.ProjectControllerInit()
	accountController := impl_controller.AccountControllerInit()
	taskController := impl_controller.TaskControllerInit()

	r.Route("/api/v1", func(v1 chi.Router) {
		v1.Route("/public", func(public chi.Router) {
			public.Post("/send_info", registerController.SendInfoRegister)
			public.Post("/confirm_code", registerController.ConfirmCodeRegister)

			public.Post("/login", loginController.Login)
		})

		v1.Route("/protected", func(protected chi.Router) {
			protected.Use(jwtauth.Verifier(config.GetJWT()))
			protected.Use(jwtauth.Authenticator(config.GetJWT()))

			protected.Post("/login_token", loginController.LoginToken)

			protected.Route("/project", func(project chi.Router) {
				project.Get("/creater_id", projectController.GetProjectByCreaterId)
				project.Get("/joined", projectController.GetProjectJoined)
				project.Get("/creater_id_detail/{id}", projectController.GetProjectCreaterById)
				project.Get("/joined_detail/{id}", projectController.GetProjectJoinedById)
				project.Post("/create", projectController.CreateProject)
				project.Delete("/delete", projectController.DeleteProject)
			})

			protected.Route("/account", func(account chi.Router) {
				account.Get("/get_account_project/{id}", accountController.GetUserProject)
				account.Post("/get_account", accountController.GetAccount)
				account.Post("/add_to_project", accountController.AddUserToProject)
			})

			protected.Route("/task", func(task chi.Router) {
				task.Get("/{id}", taskController.GetTaskOfProject)
				task.Post("/create", taskController.CreateTask)
				task.Put("/update", taskController.UpdateStatusTask)
				task.Delete("/delete", taskController.DeleteTask)

				task.Get("/all_user_task", taskController.GetAllUserOfTask)
				task.Post("/add_user", taskController.AddUserToTask)
				task.Delete("/remove_user", taskController.RemoveUserToTask)
			})

		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+config.GetAppPort()+"/swagger/doc.json"),
	))

	log.Println("http://localhost:" + config.GetAppPort())
	log.Println("http://localhost:" + config.GetAppPort() + "/swagger/index.html")

	return r
}
