package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/livething/helper"
	"github.com/yaaaaashiki/livething/model"
	"github.com/yaaaaashiki/livething/usecase"
)

const (
	exist = "1"
)

type SetCurrentObjectStatusController struct {
	setCurrentObjectStatusUseCase *usecase.SetCurrentObjectStatusUseCase
}

type InputObjectField struct {
	Value string `binding:"required" json:"value"`
	Name  string `binding:"required" json:"name"`
}

func NewSetCurrentObjectStatusController(setCurrentObjectStatusUseCase *usecase.SetCurrentObjectStatusUseCase) *SetCurrentObjectStatusController {
	return &SetCurrentObjectStatusController{
		setCurrentObjectStatusUseCase: setCurrentObjectStatusUseCase,
	}
}

func CheckStatus(c *gin.Context) bool {
	object := &InputObjectField{}

	if err := c.MustBindWith(object, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return false
	}

	if object.Value == exist {
		return true
	}
	return false
}

func SetStatus(c *gin.Context) {
	object := &model.Object{}
	object.Status = CheckStatus(c)

	//fmt.Println(object.Status) // for debug. TODO remove this code
}

func (s *SetCurrentObjectStatusController) Execute(c *gin.Context) {
	SetStatus(c)
}
