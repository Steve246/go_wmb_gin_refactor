package manager

import "go_wmb_gin_refactor/repo"

type RepositoryManager interface {
	MenuRepo() repo.MenuRepository
	MenuPriceRepo() repo.PriceRepository
	TableRepo() repo.TableRepository
	TransRepo() repo.TransTypeRepository
	DiscountRepo() repo.DiscountRepository
	CustomerRepo() repo.CustRepository
	tBillRepo() repo.BillRepository
	tBillDetailRepo() repo.BillDetailRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) tBillDetailRepo() repo.BillDetailRepository {
	return repo.NewBillDetailRepository(r.infra.SqlDb())
}

func (r *repositoryManager) tBillRepo() repo.BillRepository {
	return repo.NewBillRepository(r.infra.SqlDb())
}

func (r *repositoryManager) CustomerRepo() repo.CustRepository {
	return repo.NewCustRepository(r.infra.SqlDb())
}

func (r *repositoryManager) DiscountRepo() repo.DiscountRepository {
	return repo.NewDiscountRepository(r.infra.SqlDb())
}

func (r *repositoryManager) TransRepo() repo.TransTypeRepository {
	return repo.NewTransTypeRepository(r.infra.SqlDb())
}

func (r *repositoryManager) TableRepo() repo.TableRepository {
	return repo.NewTableRepository(r.infra.SqlDb())
}

func (r *repositoryManager) MenuPriceRepo() repo.PriceRepository {
	return repo.NewPriceRepository(r.infra.SqlDb())
}

func (r *repositoryManager) MenuRepo() repo.MenuRepository {
	return repo.NewMenuRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
