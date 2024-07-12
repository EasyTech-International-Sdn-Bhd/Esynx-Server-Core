package customer

import (
	"github.com/easytech-international-sdn-bhd/core/entities"
	iterator "github.com/ledongthuc/goterators"
	"xorm.io/xorm"
)

type CmsCustomerBranchRepository struct {
	db *xorm.Engine
}

func NewCmsCustomerBranchRepository(db *xorm.Engine) *CmsCustomerBranchRepository {
	return &CmsCustomerBranchRepository{
		db: db,
	}
}

func (r *CmsCustomerBranchRepository) Get(branchCode string) (*entities.CmsCustomerBranch, error) {
	var branch entities.CmsCustomerBranch
	has, err := r.db.Where("branch_code = ?", branchCode).Get(&branch)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &branch, nil
}

func (r *CmsCustomerBranchRepository) GetByCustomerCode(custCode string) ([]*entities.CmsCustomerBranch, error) {
	var branches []*entities.CmsCustomerBranch
	err := r.db.Where("cust_code = ? AND branch_active = ?", custCode, 1).Find(&branches)
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func (r *CmsCustomerBranchRepository) GetByAgentId(agentId int64) ([]*entities.CmsCustomerBranch, error) {
	branches, err := r.GetAllStatusByAgentId(agentId)
	if err != nil {
		return nil, err
	}
	if len(branches) == 0 {
		return iterator.Filter(branches, func(item *entities.CmsCustomerBranch) bool {
			return item.BranchActive == 1
		}), nil
	}
	return nil, nil
}

func (r *CmsCustomerBranchRepository) GetAllStatusByAgentId(agentId int64) ([]*entities.CmsCustomerBranch, error) {
	var branches []*entities.CmsCustomerBranch
	err := r.db.Where("agent_id = ?", agentId).Find(&branches)
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func (r *CmsCustomerBranchRepository) InsertMany(records []*entities.CmsCustomerBranch) error {
	_, err := r.db.Insert(iterator.Map(records, func(item *entities.CmsCustomerBranch) *entities.CmsCustomerBranch {
		item.Validate()
		item.ToUpdate()
		return item
	}))
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCustomerBranchRepository) Update(record *entities.CmsCustomerBranch) error {
	_, err := r.db.Where("branch_code = ?", record.BranchCode).Update(record)
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCustomerBranchRepository) UpdateMany(records []*entities.CmsCustomerBranch) error {
	session := r.db.NewSession()
	defer session.Close()
	err := session.Begin()
	if err != nil {
		return err
	}
	var sessionErr error
	rollback := false
	for _, branch := range records {
		branch.Validate()
		_, err = session.Where("branch_code = ?", branch.BranchCode).Update(branch)
		if err != nil {
			rollback = true
			sessionErr = err
			break
		}
	}
	if rollback {
		err := session.Rollback()
		if err != nil {
			return err
		}
		return sessionErr
	}
	err = session.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *CmsCustomerBranchRepository) Delete(record *entities.CmsCustomerBranch) error {
	record.BranchActive = 0
	record.ToUpdate()
	return r.Update(record)
}

func (r *CmsCustomerBranchRepository) DeleteMany(records []*entities.CmsCustomerBranch) error {
	for _, record := range records {
		record.BranchActive = 0
		record.ToUpdate()
	}
	return r.UpdateMany(records)
}
