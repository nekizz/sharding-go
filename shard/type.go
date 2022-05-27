package shard

import (
	"github.com/go-pg/sharding"
	"time"
)

type TKB struct {
	tableName string `sql:"?shard.tkbs"`

	ID          uint
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
	SoChoConLai uint
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

type RegisterSubject struct {
	tableName string `sql:"?shard.register_subject"`

	ID       uint
	MaSV     string
	MaMonHoc string
	NhomLop  string
}

//TKB Activity

func NewTKB() *TKB {
	return &TKB{}
}

func CreateTKB(cluster *sharding.Cluster, tkb *TKB) error {
	return cluster.Shard(int64(tkb.ID)).Insert(tkb)
}

func UpdateTKB(cluster *sharding.Cluster, tkb *TKB) error {
	return cluster.Shard(int64(tkb.ID)).Update(tkb)
}

func DeleteTKB(cluster *sharding.Cluster, tkb *TKB) error {
	err := cluster.Shard(int64(tkb.ID)).Delete(&tkb)
	return err
}

func GetTKB(cluster *sharding.Cluster, id int64) (*TKB, error) {
	var tkb TKB
	err := cluster.Shard(id).Model(&tkb).Where("id = ?", id).Select()
	return &tkb, err
}

func ListAllTKB(cluster *sharding.Cluster, id int64) ([]TKB, error) {
	var tkbs []TKB
	err := cluster.Shard(id).Model(&tkbs).Where("id = ?", id).Select()
	return tkbs, err
}

//Register Subject Activity

func NewRS() *RegisterSubject {
	return &RegisterSubject{}
}

func CreateRS(cluster *sharding.Cluster, rs *RegisterSubject) error {
	return cluster.Shard(int64(rs.ID)).Insert(rs)
}

func DeleteRS(cluster *sharding.Cluster, rs *RegisterSubject) error {
	err := cluster.Shard(int64(rs.ID)).Delete(rs)
	return err
}

func GetRS(cluster *sharding.Cluster, id int64) (*RegisterSubject, error) {
	var rs RegisterSubject
	err := cluster.Shard(id).Model(&rs).Where("id = ?", id).Select()
	return &rs, err
}

func ListAllRS(cluster *sharding.Cluster, id int64) ([]RegisterSubject, error) {
	var rs []RegisterSubject
	err := cluster.Shard(id).Model(&rs).Where("id = ?", id).Select()
	return rs, err
}
