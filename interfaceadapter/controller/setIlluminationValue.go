package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/livething/helper"
)

const (
	exist = "1"
)

type InputObjectField struct {
	Value  string `binding:"required" json:"value"`
	Name   string `binding:"required" json:"name"`
	Status bool
}

func setCurrentObjectStatus(c *gin.Context) {

	illumination := &InputObjectField{}

	if err := c.MustBindWith(illumination, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if illumination.Value == exist {
		illumination.Status = true
	} else {
		illumination.Status = false
	}

	fmt.Println(illumination.Value)
}
