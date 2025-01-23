package tasksrouter

import (
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetTasks"))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateTask"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateTask"))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteTask"))
}