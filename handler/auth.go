package handler

import (
	todo "TODO"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	err := c.BindJSON(&input)
	if err != nil {
		errorzResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if id == 0 {
		errorzResponse(c, http.StatusInternalServerError, "username already exists")
		return
	} else {
		c.JSON(http.StatusCreated, map[string]interface{}{
			"id": id,
		})
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var input todo.User
	err := c.BindJSON(&input)
	if err != nil {
		errorzResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, "user not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
