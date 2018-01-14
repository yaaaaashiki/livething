package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/labstack/gommon/log"
	"github.com/yaaaaashiki/livething/helper"
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

//Reference this varibale to check object status
type Object struct {
	Status bool
	Name   string
}

func NewSetCurrentObjectStatusController(setCurrentObjectStatusUseCase *usecase.SetCurrentObjectStatusUseCase) *SetCurrentObjectStatusController {
	return &SetCurrentObjectStatusController{
		setCurrentObjectStatusUseCase: setCurrentObjectStatusUseCase,
	}
}

func CheckStatus(c *gin.Context) (bool, error) {
	object := &InputObjectField{}

	if err := c.MustBindWith(object, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return false, err
	}

	if object.Value == exist {
		return true, nil
	}
	return false, nil
}

func SetStatus(c *gin.Context) {
	object = &Object{}
	object.Status, err := CheckStatus(c)
	if err != nil {
		log.Errorf(err.Error())
	}
	fmt.Println(object.Status) // for debug. TODO remove this code
}

func (s *SetCurrentObjectStatusController) Execute(c *gin.Context) {
	SetStatus(c)
}
