package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/cstack/helper"
)

type Illumination struct {
	Value int `binding:"required" json:"value"`
}

func Execute(c *gin.Context) {
	//rawIlluminationValue, err := strconv.Atoi(c.Param("illumination"))
	//IlluminationValue := uint(rawIlluminationValue)

	if err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	val := &IlluminationValue{}

	if err := c.MustBindWith(val, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

}
