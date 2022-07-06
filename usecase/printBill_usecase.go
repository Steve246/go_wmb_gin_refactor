package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type PrintBillUsecase interface {
	CetakBill(billId string) (model.Bill_Detail, error)
}

type printBillUsecase struct {
	billDetail  repo.BillDetailRepository
	bill        repo.BillRepository
	tableDetail repo.TableRepository
}

func (c *printBillUsecase) CetakBill(billId string) (model.Bill_Detail, error) {
	resultStrukPembelian, err1 := c.billDetail.FindById(billId)

	if err1 != nil {
		return resultStrukPembelian, err1
	}

	// balikin nilai model table berdasarkan id

	selectTableId, err2 := c.bill.FindByIdBill(billId) //balikin bill yang ada table id yang sama dengan input

	if err2 != nil {
		return resultStrukPembelian, err2
	}

	tabl, err := c.tableDetail.FindByIdTable(selectTableId.Id) //pake id ini buat cari model table

	err3 := c.tableDetail.UpdateStatus(&tabl, map[string]interface{}{"is_available": 0})

	//update nilainya jadi is_available = 0

	if err3 != nil {
		return resultStrukPembelian, err
	}

	return resultStrukPembelian, nil

}

func NewPrintBillUsecase(repo repo.BillDetailRepository) printBillUsecase {
	return printBillUsecase{
		billDetail: repo,
	}
}
