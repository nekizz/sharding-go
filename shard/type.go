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
	SoChoConLai string
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
}

func CreateTKB(cluster *sharding.Cluster, tkb *TKB) error {
	return cluster.Shard(int64(tkb.ID)).Insert(tkb)
}

func UpdateTKB(cluster *sharding.Cluster, tkb *TKB) error {
	return cluster.Shard(int64(tkb.ID)).Update(tkb)
}

// GetTKB splits shard from user id and fetches tkb from the shard.
func GetTKB(cluster *sharding.Cluster, id int64) (*TKB, error) {
	var tkb TKB
	err := cluster.SplitShard(id).Model(&TKB{}).Where("id = ?", id).Select()
	return &tkb, err
}

// GetTKbs picks shard by account id and fetches tkb from the shard.
func GetTKBs(cluster *sharding.Cluster, id int64) ([]TKB, error) {
	var tkbs []TKB
	err := cluster.Shard(id).Model(&tkbs).Where("id = ?", id).Select()
	return tkbs, err
}

func CreateRS(cluster *sharding.Cluster, rs *RegisterSubject) error {
	return cluster.Shard(int64(rs.ID)).Insert(rs)
}

func GetRS(cluster *sharding.Cluster, id int64) (*RegisterSubject, error) {
	var rs RegisterSubject
	err := cluster.SplitShard(id).Model(&RegisterSubject{}).Where("id = ?", id).Select()
	return &rs, err
}

func GetRSs(cluster *sharding.Cluster, id int64) ([]RegisterSubject, error) {
	var rs []RegisterSubject
	err := cluster.Shard(id).Model(&rs).Where("id = ?", id).Select()
	return rs, err
}
