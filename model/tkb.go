package model

import "time"

type TKB struct {
	MaMonHoc    string
	TenMon      string
	Lop         string
	KhoaNganh   int
	Nganh       string
	Nhom        int
	ToHop       int
	Thu         int
	Kip         int
	SySo        int
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
	SoTC        int
	PhanBoCT    PhanBoCT
}

type PhanBoCT struct {
	TSTiet int
	LT     int
	BT     int
	BTL    int
	THTN   int
	TuHoc  int
}
