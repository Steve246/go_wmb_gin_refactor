package manager

import "go_wmb_gin_refactor/usecase"

type UseCaseManager interface {
	CreateMenuUsecase() usecase.CreateMenuUsecase
	UpdateMenuUsecase() usecase.UpdateMenuUsecase
	DeleteMenuUsecase() usecase.DeleteMenuUsecase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

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
