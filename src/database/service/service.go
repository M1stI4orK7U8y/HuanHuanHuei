package service

import (
	"gitlab.com/packtumi9722/huanhuanhuei/src/database/dao"
)

// Service service for db operation
type Service struct{}

// Close close service
func (*Service) Close() {
	dao.Instance().Close()
}
