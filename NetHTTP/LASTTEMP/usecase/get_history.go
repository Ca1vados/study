package usecase

import (
	"fmt"
	"temp/entity"
)

func (u *UseCase) GetHistory() ([][]entity.ApiBinanceResponse, error) {
	data, err := u.db.ReadAllData()
	if err != nil {
		return data, fmt.Errorf("ошибка u.db.ReadAllData: %v", err)
	}
	return data, nil
}
