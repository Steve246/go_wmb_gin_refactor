package manager

import "go_wmb_gin_refactor/repo"

type RepositoryManager interface {
	MenuRepo() repo.MenuRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) MenuRepo() repo.MenuRepository {
	return repo.NewMenuRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
