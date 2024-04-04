package models

import "github.com/gin-gonic/gin"

func NewError(c *gin.Context, status int, err error) {
	er := HttpError{
		Error: err.Error(),
	}
	c.JSON(status, er)
}

type HttpError struct {
	Error string `json:"error" example:"..."`
}
