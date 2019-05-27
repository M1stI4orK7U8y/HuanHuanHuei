package service

import "gitlab.com/packtumi9722/huanhuanhuei/src/database/dao"

// DeletePending delete pending
func (*Service) DeletePending(id string) error {
	return dao.Instance().DeletePending(id)
}
