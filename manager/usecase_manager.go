package manager

import "go_wmb_gin_refactor/usecase"

type UseCaseManager interface {
	//Menu
	CreateMenuUsecase() usecase.CreateMenuUsecase
	UpdateMenuUsecase() usecase.UpdateMenuUsecase
	DeleteMenuUsecase() usecase.DeleteMenuUsecase

	//Menu Price
	CreateMenuPriceUsecase() usecase.CreateMenuPriceUsecase
	UpdateMenuPriceUsecase() usecase.UpdateMenuPriceUsecase
	DeleteMenuPriceUsecase() usecase.DeleteMenuPriceUsecase

	//Table
	CreateTableUsecase() usecase.CreateTableUsecase
	UpdateTableUsecase() usecase.UpdateTableUsecase
	DeleteTableUsecase() usecase.DeleteTableUsecase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

//Table

func (u *useCaseManager) DeleteTableUsecase() usecase.DeleteTableUsecase {

	return usecase.NewDeleteTableUsecase(u.repoManager.TableRepo())
}

func (u *useCaseManager) UpdateTableUsecase() usecase.UpdateTableUsecase {

	return usecase.NewUpdateTableUsecase(u.repoManager.TableRepo())
}

func (u *useCaseManager) CreateTableUsecase() usecase.CreateTableUsecase {
	return usecase.NewCreateTableUseCase(u.repoManager.TableRepo())
}

//Menu Price
func (u *useCaseManager) DeleteMenuPriceUsecase() usecase.DeleteMenuPriceUsecase {
	return usecase.NewDeleteMenuPriceUsecase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) UpdateMenuPriceUsecase() usecase.UpdateMenuPriceUsecase {
	return usecase.NewUpdateMenuPriceUsecase(u.repoManager.MenuPriceRepo())
}

func (u *useCaseManager) CreateMenuPriceUsecase() usecase.CreateMenuPriceUsecase {
	return usecase.NewCreateMenuPriceUseCase(u.repoManager.MenuPriceRepo())
}

//Menu

func (u *useCaseManager) DeleteMenuUsecase() usecase.DeleteMenuUsecase {
	return usecase.NewDeleteMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) UpdateMenuUsecase() usecase.UpdateMenuUsecase {
	return usecase.NewUpdateMenuUseCase(u.repoManager.MenuRepo())
}

func (u *useCaseManager) CreateMenuUsecase() usecase.CreateMenuUsecase {
	return usecase.NewCreateMenuUseCase(u.repoManager.MenuRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
