package object

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yaaaaashiki/cstack/helper"
)

const (
	exist = "1"
)

type InputObjectField struct {
	Value string `binding:"required" json:"value"`
	Name  string `binding:"required" json:"name"`
}

//Reference this varibale to check object status
var Status bool

func CheckStatus() bool {
	object := &InputObjectField{}

	if err := c.MustBindWith(object, binding.JSON); err != nil {
		helper.ResponseErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}

	if object.Value == exist {
		return true
	}
	return false
}

func SetStatus(c *gin.Context) {
	Status = CheckStatus()
	fmt.Println(Status) // for debug. TODO remove this code
}
