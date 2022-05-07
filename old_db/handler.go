package old_db

import (
	"shrading/connection"
	"shrading/model"
)

func ListAll(limit int, offset int) ([]model.Employee, int64, error) {
	var count int64
	var listEmployee []model.Employee

	listQuery := connection.DB.Model(&model.Employee{}).Select("*").Limit(limit).Offset(offset)
	countQuery := connection.DB.Model(&model.Employee{}).Select("*").Count(&count)

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&listEmployee).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return listEmployee, count, nil
}
