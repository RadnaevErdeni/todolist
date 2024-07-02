package handler

import (
	todo "TODO"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createStr(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorzResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	var input todo.TodoStr
	if err := c.BindJSON(&input); err != nil {
		errorzResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	id, err := h.services.TodoStr.Create(userId, listId, input)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllStr(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorzResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	str, err := h.services.TodoStr.GetAll(userId, listId)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, str)
}

func (h *Handler) getStrById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	strId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorzResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	str, err := h.services.TodoStr.GetById(userId, strId)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, str)
}

func (h *Handler) updateStr(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorzResponse(c, http.StatusBadRequest, "invalid str id param")
		return
	}
	var input todo.UpdateStrInput
	if err := c.BindJSON(&input); err != nil {
		errorzResponse(c, http.StatusBadRequest, "invalid str id param")
		return
	}
	if err := h.services.UpdateStr(userId, id, input); err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
	})
}

func (h *Handler) deleteStr(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	strId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorzResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}
	err = h.services.TodoStr.Delete(userId, strId)
	if err != nil {
		errorzResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
	})
}
