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

	//menu controller

	ucCreateMenu usecase.CreateMenuUsecase
	ucDeleteMenu usecase.DeleteMenuUsecase
	ucUpdateMenu usecase.UpdateMenuUsecase

	//price controller

	ucCreateMenuPrice usecase.CreateMenuPriceUsecase
	ucDeleteMenuPrice usecase.DeleteMenuPriceUsecase
	ucUpdateMenuPrice usecase.UpdateMenuPriceUsecase

	//table controller

	ucCreateTable usecase.CreateTableUsecase
	ucDeleteTable usecase.DeleteTableUsecase
	ucUpdateTable usecase.UpdateTableUsecase

	api.BaseApi
}

//UPDATE

func (m *MenuController) updateTable(c *gin.Context) {
	var table *model.Table

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucUpdateTable.UpdateTable(table, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Update is Success",
			"message": table,
		})

	}

}

func (m *MenuController) updateMenuPrice(c *gin.Context) {
	var menuPrice *model.Menu_Price

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&menuPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucUpdateMenuPrice.UpdateMenuPrice(menuPrice, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Update is Success",
			"message": menuPrice,
		})

	}

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

//DELETE

func (m *MenuController) deleteTable(c *gin.Context) {
	var table *model.Table

	if err := c.BindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucDeleteTable.DeleteTable(table)
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

func (m *MenuController) deletePriceMenu(c *gin.Context) {
	var menuPrice *model.Menu_Price

	if err := c.BindJSON(&menuPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucDeleteMenuPrice.DeletePrice(menuPrice)
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

//CREATE

func (m *MenuController) createTable(c *gin.Context) {
	var newTable *model.Table

	if err := c.BindJSON(&newTable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCreateTable.CreateTable(newTable)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating menu",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newTable,
		})

	}
}

func (m *MenuController) createNewMenuPrice(c *gin.Context) {
	var newMenuPrice *model.Menu_Price

	if err := c.BindJSON(&newMenuPrice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCreateMenuPrice.CreateMenuPrice(newMenuPrice)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating menu",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newMenuPrice,
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

func NewMenuController(router *gin.Engine,
	ucCreateMenu usecase.CreateMenuUsecase,
	ucDeleteMenu usecase.DeleteMenuUsecase,
	ucUpdateMenu usecase.UpdateMenuUsecase,

	ucCreateMenuPrice usecase.CreateMenuPriceUsecase, ucDeleteMenuPrice usecase.DeleteMenuPriceUsecase, ucUpdateMenuPrice usecase.UpdateMenuPriceUsecase,

	ucCreateTable usecase.CreateTableUsecase,
	ucDeleteTable usecase.DeleteTableUsecase,
	ucUpdateTable usecase.UpdateTableUsecase) *MenuController {

	controller := MenuController{
		router:       router,
		ucCreateMenu: ucCreateMenu,
		ucDeleteMenu: ucDeleteMenu,
		ucUpdateMenu: ucUpdateMenu,

		ucCreateMenuPrice: ucCreateMenuPrice,
		ucDeleteMenuPrice: ucDeleteMenuPrice,
		ucUpdateMenuPrice: ucUpdateMenuPrice,

		ucCreateTable: ucCreateTable,
		ucDeleteTable: ucDeleteTable,
		ucUpdateTable: ucUpdateTable,
	}

	//                      SOAL 1 - CRUD MENU

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

	//                    SOAL 2 - CRUD MENU PRICE

	//masukin data menu price

	router.POST("/price", controller.createNewMenuPrice)
	// http://localhost:8888/price

	// soft delete data menu price
	router.DELETE("/priceDelete", controller.deletePriceMenu)
	// http://localhost:8888/priceDelete

	//update data menu price

	router.PATCH("/priceUpdate/:id", controller.updateMenuPrice)
	// http://localhost:8888/priceUpdate/C0002

	//                    SOAL 3 - CRUD MENU PRICE

	//masukin data table

	router.POST("/table", controller.createTable)
	// http://localhost:8888/table

	// soft delete data tabel
	router.DELETE("/tableDelete", controller.deleteTable)
	// http://localhost:8888/priceDelete

	//update data tabel

	router.PATCH("/tableUpdate/:id", controller.updateTable)
	// http://localhost:8888/tableUpdate/:id

	return &controller

}
