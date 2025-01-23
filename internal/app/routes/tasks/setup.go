package tasksrouter

import "github.com/go-chi/chi/v5"

func SetupRouter(r *chi.Mux) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", GetTasks)
		r.Post("/", CreateTask)
		r.Put("/{id}", UpdateTask)
		r.Delete("/{id}", DeleteTask)
	})
}