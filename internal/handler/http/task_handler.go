package handler

import (
	"TZ-API/internal/usecase"
	"TZ-API/pkg/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type TaskHandler struct {
	taskUseCase *usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	task, err := h.taskUseCase.CreateTask(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create task", err)
		return
	}

	response.Success(w, http.StatusCreated, "Task created successfully", task)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := uuid.Parse(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid task ID", err)
		return
	}

	task, err := h.taskUseCase.GetTask(r.Context(), taskID)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Task not found", err)
		return
	}

	response.Success(w, http.StatusOK, "Task retrieved successfully", task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := uuid.Parse(vars["id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid task ID", err)
		return
	}

	err = h.taskUseCase.DeleteTask(r.Context(), taskID)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Task not found", err)
		return
	}

	response.Success(w, http.StatusOK, "Task deleted successfully", nil)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskUseCase.ListTasks(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve tasks", err)
		return
	}

	response.Success(w, http.StatusOK, "Tasks retrieved successfully", tasks)
}
