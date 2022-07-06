package controller

import (
	"go_wmb_gin_refactor/delivery/api"
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Nambain usecase

type MenuController struct {
	router *gin.Engine

	ucCreateMenu usecase.CreateMenuUsecase
	// ucUpdateMenu usecase.UpdateMenuUsecase
	// ucDeleteMenu usecase.DeleteMenuUsecase

	api.BaseApi
}

func (m *MenuController) createNewMenu(c *gin.Context) {
	var newMenu *model.Menu

	if err := c.BindJSON(&newMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCreateMenu.CreateMenu(newMenu)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating Product",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newMenu,
		})

	}
}

func NewMenuController(router *gin.Engine, ucCreateMenu usecase.CreateMenuUsecase) *MenuController {
	controller := MenuController{
		router:       router,
		ucCreateMenu: ucCreateMenu,
	}

	//masukin post atau get disini

	router.POST("/product", controller.createNewMenu)
	// http://localhost:8888/product

	//masukin data dibawah
	// {
	//   "menuId": "C0001",
	//   "menuText": "MakananEnajk"
	// }

	return &controller

}
