package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testCase/pkg/repository"
)

func (h *Handler) getUserStatus(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := h.services.GetUserById(userID)
	if err != nil {
		if err == repository.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	completedTasks, err := h.services.GetCompletedTasksByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch completed tasks"})
		return
	}

	response := gin.H{
		"id":              user.ID,
		"username":        user.UserName,
		"balance":         user.Balance,
		"referrer_id":     user.ReferrerID,
		"completed_tasks": completedTasks,
	}

	c.JSON(http.StatusOK, response)

}
func (h *Handler) getLeaderboard(c *gin.Context) {
	leaders, err := h.services.User.GetLeaderboard()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leaderboard"})
		return
	}
	var response []map[string]interface{}
	for _, leader := range leaders {
		response = append(response, map[string]interface{}{
			"name":    leader.UserName,
			"balance": leader.Balance,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"leaderboard": response,
	})
}
func (h *Handler) completeTask(c *gin.Context) {
	var input struct {
		TaskID string `json:"task_id" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	err := h.services.Task.CompleteTask(userID, input.TaskID)
	if err != nil {
		if err == repository.ErrTaskNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else if err == repository.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task completed successfully"})
}
func (h *Handler) addReferrer(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var input struct {
		ReferrerID string `json:"referrer_id" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.services.User.AddReferrer(userID, input.ReferrerID)
	if err != nil {
		if err == repository.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else if err == repository.ErrReferrerNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Referrer not found"})
		} else if err == repository.ErrReferrerAlreadySet {
			c.JSON(http.StatusConflict, gin.H{"error": "Referrer already set"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Referrer added successfully"})
}
