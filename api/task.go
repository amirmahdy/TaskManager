package api

import (
	"database/sql"
	"net/http"
	db "taskmanager/db/model"
	"taskmanager/token"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Priority    int    `json:"priority"`
	CreatedAt   string `json:"created_at"`
}
type createTaskRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DueDate     time.Time `json:"due_date" binding:"required"`
	Priority    int32     `json:"priority" binding:"required"`
}

// @Summary Task
// @Schemes
// @Description	get task
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} Task
// @Router /task [get]
func (server *Server) getTask(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tasks, err := server.store.GetTasks(ctx, user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

// @Summary Task
// @Schemes
// @Description task create
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param createTaskRequest	body createTaskRequest true "Create Task Request"
// @Success 200 {object} Task
// @Router /task [post]
func (server *Server) createTask(ctx *gin.Context) {
	var req createTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	task, err := server.store.CreateTask(ctx, db.CreateTaskParams{
		Username:    user.Username,
		Title:       sql.NullString{String: req.Title, Valid: true},
		Description: sql.NullString{String: req.Description, Valid: true},
		DueDate:     sql.NullTime{Time: req.DueDate, Valid: true},
		Priority:    sql.NullInt32{Int32: req.Priority, Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}
