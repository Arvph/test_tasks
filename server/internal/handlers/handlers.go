package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/arvph/test_tasks/internal/modules"
	"github.com/arvph/test_tasks/internal/services"

	"github.com/gorilla/mux"
)

// POST /tasks - создание новой задачи.
// GET /tasks - получение списка всех задач.
// GET /tasks/{id} - получение задачи по ID.
// PUT /tasks/{id} - обновление задачи по ID.
// DELETE /tasks/{id} - удаление задачи по ID.

// Handler представляет структуру для работы с функциями UseCase слоя.
type Handler struct {
	srv *services.Services
}

// InitHandler создает и возвращает новый экземпляр Handler.
func InitHandler(SRV *services.Services) *Handler {
	return &Handler{
		srv: SRV,
	}
}

// Create godoc
// @Summary Create a new task
// @Description Create a new task with the provided information
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body modules.Task required "Task info"
// @Success 201 {object} map[string]interface{} "Task created successfully"
// @Failure 400 {string} string "Error parsing JSON request body"
// @Failure 500 {string} string "Error creating new record"
// @Router /tasks [post]

// Create - endpoint для добавления новой записи в репозиторий.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Create endpoint is visited")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("Error closing request body: %v\n", err)
		}
	}()

	var task modules.Task

	if err := json.Unmarshal(body, &task); err != nil {
		log.Printf("Error parsing JSON request body: %v\n", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if err := h.srv.Create(r.Context(), task); err != nil {
		log.Printf("Error creating new record: %v\n", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"message": "Task created successfully", "task": task}); err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}

// GetAll - endpoint для получения всех записей из репозитория.
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAll endpoint is visited")

	userid := r.Header.Get("user_id")
	if userid == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	page := r.Header.Get("page")
	size := r.Header.Get("size")

	tasks, err := h.srv.GetAll(r.Context(), userid, page, size)
	if err != nil {
		log.Printf("Error getting all tasks: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		log.Println("unable to encode JSON for sending")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Отправка успешного ответа
	w.WriteHeader(http.StatusOK)
}

// GetByID godoc
// @Summary Get task by ID
// @Description get task by its ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param user_id header string false "User ID" default(1)
// @Success 200 {object} Task "success"
// @Failure 400 {object} ErrorResponse "bad request"
// @Failure 404 {object} ErrorResponse "not found"
// @Failure 500 {object} ErrorResponse "internal error"
// @Router /tasks/{id} [get]

// GetByID - endpoint для получения записи из репозитория по ID и UserID.
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	log.Println("GetByID endpoint is visited")

	userid := r.Header.Get("user_id")
	if userid == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"] // Получение ID из параметров пути

	task, err := h.srv.GetByID(r.Context(), userid, id)
	if err != nil {
		log.Printf("Error getting task by ID: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		log.Println("unable to encode JSON for sending")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// UpdateByID - endpoint для обновления записи из репозитория по ID и UserID.
func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateByID endpoint is visited")

	vars := mux.Vars(r)
	id := vars["id"] // Получение ID из параметров пути

	// Чтение тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var task modules.Task

	// Десериализация JSON в структуру
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Error parsing JSON request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask, err := h.srv.UpdateByID(r.Context(), id, task)
	if err != nil {
		log.Printf("Error updating task by ID: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
		log.Println("unable to encode JSON for sending")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// DeleteByID - endpoint для удаления записи из репозитория по ID и UserID.
func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteByID endpoint is visited")

	userid := r.Header.Get("user_id")
	if userid == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.srv.DeleteByID(r.Context(), userid, id); err != nil {
		log.Printf("Error deleting task by ID: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
