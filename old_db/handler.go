package old_db

import (
	"shrading/connection"
	"shrading/model"
)

func ListAll(limit int, offset int) ([]model.TKB, int64, error) {
	var count int64
	var listTKB []model.TKB

	listQuery := connection.DB.Model(&model.TKB{}).Select("*").Limit(limit).Offset(offset)
	countQuery := connection.DB.Model(&model.TKB{}).Select("*").Count(&count)

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&listTKB).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return listTKB, count, nil
}

func CreatOne(tkb *model.TKB) (*model.TKB, error) {
	query := connection.DB.Model(&model.TKB{}).Create(tkb)

	if err := query.Error; nil != err {
		return nil, err
	}

	return tkb, nil
}
