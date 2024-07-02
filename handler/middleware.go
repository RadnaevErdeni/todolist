package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "UserId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		errorzResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		errorzResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	if len(headerParts[1]) == 0 {
		errorzResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		errorzResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		errorzResponse(c, http.StatusInternalServerError, "user id is not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		errorzResponse(c, http.StatusInternalServerError, "user id is not found")
		return 0, errors.New("user id not found")
	}
	return idInt, nil
}
