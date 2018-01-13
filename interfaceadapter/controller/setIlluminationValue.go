package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/livething/helper"
)

type Illumination struct {
	Value int `binding:"required" json:"value"`
}

func setCurrentIlluminationValue(c *gin.Context) {

	illumination := &Illumination{}

	if err := c.MustBindWith(illumination, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(illumination.Value)
}
