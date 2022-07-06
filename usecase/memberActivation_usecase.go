package usecase

import (
	"go_wmb_gin_refactor/model"
	"go_wmb_gin_refactor/repo"
)

type MemberActiveUsecase interface {
	ActivateMember(name string, noHp string) (model.Customer, error)
}

type memberRepo struct {
	custRepo repo.CustRepository
}

func (ma memberRepo) ActivateMember(name string, noHp string) (model.Customer, error) {

	//kasus customer udh ada di database, tapi mau aktifin status

	//cari customer dari nama dan no hp
	cust, err := ma.custRepo.FindFirstBy(map[string]interface{}{
		"customer_name":   name,
		"mobile_phone_no": noHp,
	})

	if err != nil {
		return model.Customer{}, err
	}

	err = ma.custRepo.UpdateForActivate(&cust, map[string]interface{}{"is_member": 1})

	if err != nil {
		return model.Customer{}, err
	}

	return cust, nil
}

func NewMemberActiveUsecase(repo repo.CustRepository) MemberActiveUsecase {
	return memberRepo{
		custRepo: repo,
	}

}
