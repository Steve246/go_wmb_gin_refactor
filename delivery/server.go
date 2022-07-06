package delivery

import (
	"go_wmb_gin_refactor/config"
	"go_wmb_gin_refactor/delivery/controller"
	"go_wmb_gin_refactor/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	router := gin.Default()
	appConfig := config.NewConfig()

	infra := manager.NewInfra(appConfig)

	repoManager := manager.NewRepositoryManager(infra)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	host := appConfig.Url

	return &appServer{
		useCaseManager: useCaseManager,
		engine:         router,
		host:           host,
	}
}

func (a *appServer) iniController() {

	controller.NewMenuController(
		a.engine,
		a.useCaseManager.CreateMenuUsecase(),
		a.useCaseManager.DeleteMenuUsecase(),
		a.useCaseManager.UpdateMenuUsecase(),

		a.useCaseManager.CreateMenuPriceUsecase(),
		a.useCaseManager.DeleteMenuPriceUsecase(),
		a.useCaseManager.UpdateMenuPriceUsecase(),

		a.useCaseManager.CreateTableUsecase(),
		a.useCaseManager.DeleteTableUsecase(),
		a.useCaseManager.UpdateTableUsecase(),

		a.useCaseManager.CreateTransUsecase(),
		a.useCaseManager.DeleteTransUsecase(),
		a.useCaseManager.UpdateTransUsecase(),

		a.useCaseManager.CrudDiscountUsecase(),
	)

}

func (a *appServer) Run() {
	a.iniController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
