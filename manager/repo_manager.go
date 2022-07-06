package manager

import "go_wmb_gin_refactor/repo"

type RepositoryManager interface {
	MenuRepo() repo.MenuRepository
	MenuPriceRepo() repo.PriceRepository
	TableRepo() repo.TableRepository
}

type repositoryManager struct {
	infra Infra
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
