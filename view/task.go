package view

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ussii39/go_rest_api/model"
)

type tasksResponse struct {
	Total int           `json:"total"`
	Tasks []*model.Task `json:"tasks"`
}

type Task struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func RenderTasks(w http.ResponseWriter, tasks []*model.Task) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	outputJson, err := json.Marshal(tasks)
	fmt.Print(string(outputJson))
	s, err := json.Marshal(tasksResponse{Total: len(tasks), Tasks: tasks})
	if err != nil {
		RenderInternalServerError(w, "cant't encode tasks response json")
		return
	}
	fmt.Fprintln(w, string(s))
}

func RenderTask(w http.ResponseWriter, task *model.Task, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	s, err := json.Marshal(task)
	if err != nil {
		RenderInternalServerError(w, "cant't encode task response json")
		return
	}
	fmt.Fprintln(w, string(s))
}
