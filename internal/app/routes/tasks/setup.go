package tasksrouter

import (
	"github.com/go-chi/chi/v5"

	"github.com/yigit433/go-rest-template/internal/app/repository"
	"github.com/yigit433/go-rest-template/internal/database"
)

func SetupRouter(r *chi.Mux) {
    taskRepository := repository.NewTaskRepository(database.Mongo.TasksCollection)
    
    taskHandler := NewTaskHandler(taskRepository)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", taskHandler.GetTasks)
		r.Post("/", taskHandler.CreateTask)
		r.Put("/{id}", taskHandler.UpdateTask)
		r.Delete("/{id}", taskHandler.DeleteTask)
	})
}