package controllers

import (
	"github.com/askaroe/geminiApiGolang/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) GetUserSummary(g *gin.Context) {
	var body struct {
		ID     int `json:"id"`
		Height int `json:"height"`
		Weight int `json:"weight"`
		Age    int `json:"age"`
	}

	var err error
	var resp string

	err = g.Bind(&body)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := &model.User{
		ID:     body.ID,
		Height: body.Height,
		Weight: body.Weight,
		Age:    body.Age,
	}

	resp, err = c.s.SummarizeUserInfo(u)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"response": resp,
	})
	return
}
