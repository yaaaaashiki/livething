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

Objchan := make(chan model.Object, 1)

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

func CheckStatus(c *gin.Context) (bool, string) {
	object := &InputObjectField{}

	if err := c.MustBindWith(object, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return false, ""
	}

	if object.Value == exist {
		return true, object.Name
	}
	return false, ""
}

func SetStatus(c *gin.Context) {
	object := &model.Object{}
	object.Status, object.Name = CheckStatus(c)

	

	Objchan.Status <- object.Status
	Objchan.Name <- object.Name
	/*
		fmt.Printf("object status: ")
		fmt.Println(object.Status)
		fmt.Println(object.Name)
	*/
	//fmt.Println(object.Status) // for debug. TODO remove this code
}

func (s *SetCurrentObjectStatusController) Execute(c *gin.Context) {
	SetStatus(c)
}
