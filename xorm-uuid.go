package xorm_uuid

import (
	"github.com/google/uuid"
)

type XormUUID uuid.UUID

func (u *XormUUID) FromDB(data []byte) error {
	a, err := uuid.FromBytes(data)
	if err != nil {
		return err
	}
	*u = XormUUID(a)
	return nil
}

func (u *XormUUID) ToDB() ([]byte, error) {
	return u[:], nil
}
