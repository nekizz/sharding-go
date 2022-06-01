package model

import (
	"shrading/connection"
	"shrading/helper"
	"strconv"
	"time"
)

type TKB struct {
	ID          uint
	HashKey     uint
	MaMonHoc    string
	TenMon      string
	Lop         string
	KhoaNganh   string
	Nganh       string
	Nhom        string
	ToHop       string
	ToTH        string
	Thu         string
	Kip         string
	SySo        string
	Phong       string
	Nha         string
	HinhThucThi string
	MaGV        string
	TenGV       string
	GhiChu      string
	NgayBD      time.Time
	NgayKT      time.Time
	Khoa        string
	BoMon       string
	SoTC        string
	TSTiet      string //phan bo kthuc
	LT          string
	BT          string
	BTL         string
	THTN        string
	TuHoc       string
}

func ListAll(limit int, offset int) ([]TKB, int64, error) {
	var count int64
	var listTKB []TKB

	listQuery := connection.DB.Model(&TKB{}).Select("*").Limit(limit).Offset(offset)
	countQuery := connection.DB.Model(&TKB{}).Select("*").Count(&count)

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}
	if err := listQuery.Find(&listTKB).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return listTKB, count, nil
}

func GetOne(id int) (*TKB, error) {
	var tkb *TKB

	query := connection.DB.Model(&tkb).Where("id = ?", id).Find(&tkb)
	if err := query.Error; err != nil {
		return nil, err
	}

	return tkb, nil
}

func CreatOneTKB(tkb *TKB) (*TKB, error) {
	query := connection.DB.Model(&TKB{}).Create(tkb)
	if err := query.Error; nil != err {
		return nil, err
	}

	return tkb, nil
}

func CreateManyTKB(list []*TKB) ([]*TKB, error) {
	query := connection.DB.Model(&TKB{}).Create(list)
	if err := query.Error; nil != err {
		return nil, err
	}

	return list, nil
}

func SyncTKBToElasticSearch() error {
	listTKB, _, err := ListAll(2000, 0)
	if err != nil {
		return err
	}
	if len(listTKB) > 0 {
		for _, idx := range listTKB {
			helper.InsertToElasticLivechat(idx, "tkb", strconv.Itoa(int(idx.ID)), "_doc")
		}
	}
	return nil
}
