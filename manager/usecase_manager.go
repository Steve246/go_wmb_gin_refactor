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

	//trans
	CreateTransUsecase() usecase.CreateTransUsecase
	UpdateTransUsecase() usecase.UpdateTransUsecase
	DeleteTransUsecase() usecase.DeleteTransUsecase

	//discount
	CrudDiscountUsecase() usecase.CrudDiscount

	//customer

	CrudCustomerUsecase() usecase.CrudCustomer

	//memberActivation
	ActivateMemberUsecase() usecase.MemberActiveUsecase

	//cekMeha

	ValidateTableUsecase() usecase.CekTableUsecase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

//validate table

func (u *useCaseManager) ValidateTableUsecase() usecase.CekTableUsecase {

	return usecase.NewCekTableUsecase(u.repoManager.TableRepo())
}

//activate member usecase
func (u *useCaseManager) ActivateMemberUsecase() usecase.MemberActiveUsecase {

	return usecase.NewMemberActiveUsecase(u.repoManager.CustomerRepo())
}

//customer

func (u *useCaseManager) CrudCustomerUsecase() usecase.CrudCustomer {

	return usecase.NewCrudCustomer(u.repoManager.CustomerRepo())
}

//discount

func (u *useCaseManager) CrudDiscountUsecase() usecase.CrudDiscount {

	return usecase.NewCrudDiscount(u.repoManager.DiscountRepo())
}

//trans

func (u *useCaseManager) DeleteTransUsecase() usecase.DeleteTransUsecase {

	return usecase.NewDeleteTransUsecase(u.repoManager.TransRepo())
}

func (u *useCaseManager) UpdateTransUsecase() usecase.UpdateTransUsecase {

	return usecase.NewUpdateTransUsecase(u.repoManager.TransRepo())
}

func (u *useCaseManager) CreateTransUsecase() usecase.CreateTransUsecase {
	return usecase.NewCreateTransUsecase(u.repoManager.TransRepo())
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
