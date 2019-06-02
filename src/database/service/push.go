package service

import (
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/dao"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/model/pending"
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/model/record"
)

// UpdateRecord update record
func (*Service) UpdateRecord(input *record.Record) error {
	return dao.Instance().UpdateRecord(input)
}

// UpdatePending update upending
func (*Service) UpdatePending(input *pending.Pending) error {
	return dao.Instance().UpdatePending(input)
}
