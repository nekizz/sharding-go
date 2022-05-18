package shard

import (
	"github.com/go-pg/sharding"
	"time"
)

const CreatTableTKB = `CREATE TABLE ?shard.tkbs (id bigint DEFAULT ?shard.next_id(), ma_mon_hoc text, 
ten_mon text, lop text, khoa_nganh text, nganh text, nhom text, to_hop text, to_th text, thu text, kip text, sy_so text, phong text, nha text, 
hinh_thuc_thi text, ma_gv text, ten_gv text, ghi_chu text, ngay_bd timestamp with time zone, ngay_kt timestamp with time zone, khoa text, bo_mon text, so_tc text
, ts_tiet text, lt text, bt text, btl text, thtn text, tu_hoc text)`

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

func CreateTKB(cluster *sharding.Cluster, tkb *TKB) error {
	return cluster.Shard(int64(tkb.ID)).Insert(tkb)
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
