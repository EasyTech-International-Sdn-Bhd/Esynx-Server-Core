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

func (r *CmsCustomerBranchRepository) Resolve(branchCode string) (*entities.CmsCustomerBranch, error) {
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
	err := r.db.Where("cust_code = ?", custCode).Where("branch_active = ?", 1).Find(&branches)
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

func (r *CmsCustomerBranchRepository) InsertBatch(records []*entities.CmsCustomerBranch) error {
	_, err := r.db.Insert(records)
	if err != nil {
		return err
	}
	return nil
}
