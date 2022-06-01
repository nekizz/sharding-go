package shard

import (
	"github.com/go-pg/sharding"
	"time"
)

type TKB struct {
	tableName string `sql:"?shard.tkbs"`

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
	SoLop       string
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

type ABC struct {
	name string
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

func (t *TKB) CreateTKB(cluster *sharding.Cluster, tkb *TKB) error {
	return cluster.Shard(int64(tkb.ID)).Insert(tkb)
}

func (t *TKB) UpdateTKB(cluster *sharding.Cluster, tkb *TKB) error {
	return cluster.Shard(int64(tkb.ID)).Update(tkb)
}

func (t *TKB) DeleteTKB(cluster *sharding.Cluster, tkb *TKB) error {
	err := cluster.Shard(int64(tkb.ID)).Delete(&tkb)
	return err
}

func (t *TKB) GetTKB(cluster *sharding.Cluster, id int64) (*TKB, error) {
	var tkb TKB
	err := cluster.Shard(id).Model(&tkb).Where("id = ?", id).Select()
	return &tkb, err
}

func (t *TKB) ListAllTKB(cluster *sharding.Cluster, id int64) ([]TKB, error) {
	var tkbs []TKB
	err := cluster.Shard(id).Model(&tkbs).Where("id = ?", id).Select()
	return tkbs, err
}

//Register Subject Activity

func NewRS() *RegisterSubject {
	return &RegisterSubject{}
}

func (r *RegisterSubject) CreateRS(cluster *sharding.Cluster, rs *RegisterSubject) error {
	return cluster.Shard(int64(rs.ID)).Insert(rs)
}

func (r *RegisterSubject) DeleteRS(cluster *sharding.Cluster, rs *RegisterSubject) error {
	err := cluster.Shard(int64(rs.ID)).Delete(rs)
	return err
}

func (r *RegisterSubject) GetRS(cluster *sharding.Cluster, id int64) (*RegisterSubject, error) {
	var rs RegisterSubject
	err := cluster.Shard(id).Model(&rs).Where("id = ?", id).Select()
	return &rs, err
}

func (r *RegisterSubject) ListAllRS(cluster *sharding.Cluster, id int64) ([]RegisterSubject, error) {
	var rs []RegisterSubject
	err := cluster.Shard(id).Model(&rs).Where("id = ?", id).Select()
	return rs, err
}
