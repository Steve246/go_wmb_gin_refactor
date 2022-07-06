package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type CekTableUsecase interface {
	RunCekMeja(namaMeja string, statusDuduk int) (model.Table, error)
}

type cekTableUsecase struct {
	table repo.TableRepository
}

func (tr cekTableUsecase) RunCekMeja(namaMeja string, statusDuduk int) (model.Table, error) {

	//cari tabel dari tabel_id (find first)

	result, err := tr.table.FindFirstByTable(map[string]interface{}{
		"table_description": namaMeja,
		"is_available":      statusDuduk,
	})

	if err != nil {
		return result, err //kalau ada yang duduk, balikin error
	}

	return result, nil //kalau gak ad yang duduk balikin datanya

}

func NewCekTableUsecase(repo repo.TableRepository) cekTableUsecase {
	return cekTableUsecase{
		table: repo,
	}

}
