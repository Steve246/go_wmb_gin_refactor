package manager

import "go_wmb_gin_refactor/repo"

type RepositoryManager interface {
	MenuRepo() repo.MenuRepository
	MenuPriceRepo() repo.PriceRepository
}

type repositoryManager struct {
	infra Infra
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
