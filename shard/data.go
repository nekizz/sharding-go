package shard

import (
	"shrading/helper"
	"shrading/model"
	"time"
)

func changeDataType(tkb model.TKB) *TKB {
	return &TKB{
		ID:          tkb.ID,
		HashKey:     tkb.HashKey,
		MaMonHoc:    tkb.MaMonHoc,
		TenMon:      tkb.TenMon,
		Lop:         tkb.Lop,
		KhoaNganh:   tkb.KhoaNganh,
		Nganh:       tkb.Nganh,
		Nhom:        tkb.Nhom,
		ToHop:       tkb.ToHop,
		ToTH:        tkb.ToTH,
		Thu:         tkb.Thu,
		Kip:         tkb.Kip,
		SoChoConLai: uint(helper.StringToInt(tkb.SySo)),
		SySo:        tkb.SySo,
		Phong:       tkb.Phong,
		Nha:         tkb.Nha,
		HinhThucThi: tkb.HinhThucThi,
		MaGV:        tkb.MaGV,
		TenGV:       tkb.TenGV,
		GhiChu:      tkb.GhiChu,
		NgayBD:      time.Time{},
		NgayKT:      time.Time{},
		Khoa:        tkb.Khoa,
		BoMon:       tkb.BoMon,
		SoTC:        tkb.SoTC,
		TSTiet:      tkb.TSTiet,
		LT:          tkb.LT,
		BT:          tkb.BT,
		BTL:         tkb.BTL,
		THTN:        tkb.THTN,
		TuHoc:       tkb.TuHoc,
	}
}
