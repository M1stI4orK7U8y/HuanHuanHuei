package service

import (
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/dao"
)

// GetRecord get record
func (*Service) GetRecord(id string) ([]byte, error) {
	return dao.Instance().GetRecord(id)
}

// GetPendings get all pending items
func (*Service) GetPendings() (map[string][]byte, error) {
	return dao.Instance().GetPendings()
}
