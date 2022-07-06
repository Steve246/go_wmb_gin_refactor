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
	ucDeleteMenu usecase.DeleteMenuUsecase
	ucUpdateMenu usecase.UpdateMenuUsecase

	api.BaseApi
}

func (m *MenuController) updateMenu(c *gin.Context) {
	var menu *model.Menu

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucUpdateMenu.UpdateMenu(menu, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Update is Success",
			"message": menu,
		})

	}

}

func (m *MenuController) deleteMenu(c *gin.Context) {
	var menu *model.Menu

	if err := c.BindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucDeleteMenu.DeleteMenu(menu)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "Delete Is Success",
			// "message": newMenu,
		})

	}

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
				"message": "Error when creating menu",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newMenu,
		})

	}
}

func NewMenuController(router *gin.Engine, ucCreateMenu usecase.CreateMenuUsecase, ucDeleteMenu usecase.DeleteMenuUsecase, ucUpdateMenu usecase.UpdateMenuUsecase) *MenuController {
	controller := MenuController{
		router:       router,
		ucCreateMenu: ucCreateMenu,
		ucDeleteMenu: ucDeleteMenu,
		ucUpdateMenu: ucUpdateMenu,
	}

	//masukin data menu

	router.POST("/product", controller.createNewMenu)
	// http://localhost:8888/product

	//masukin data dibawah
	// {
	//   "menuId": "C0001",
	//   "menuText": "MakananEnajk"
	// }

	// soft delete data menu
	router.DELETE("/productDelete", controller.deleteMenu)
	// http://localhost:8888/productDelete

	//masukin data dibawah
	// {
	//   "menuId": "C0001",
	//   "menuText": "MakananEnajk"
	// }

	//update data menu

	router.PATCH("/productUpdate/:id", controller.updateMenu)
	// http://localhost:8888/productUpdate/C0002

	//masukin data dibawah
	// {
	//
	//   "menuText": "Gorengan Mie Sedap"
	// }

	return &controller

}
