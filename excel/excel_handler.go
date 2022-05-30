package excel

import (
	"fmt"
	"shrading/connection"
	"shrading/helper"
	"shrading/model"
	"strconv"
	"time"
)

func ReadDataFromExcel() error {
	var check = true
	var tkb *model.TKB
	var listTKB []*model.TKB

	conn, err := connection.Excelconn()
	if err != nil {
		fmt.Println(err)
	}

	if check {
		for i := 13; i <= 1525; i++ {
			tt, _ := conn.GetCellValue("DKGD", "A"+strconv.Itoa(i))
			maMonHoc, _ := conn.GetCellValue("DKGD", "B"+strconv.Itoa(i))
			tenMon, _ := conn.GetCellValue("DKGD", "C"+strconv.Itoa(i))
			lop, _ := conn.GetCellValue("DKGD", "D"+strconv.Itoa(i))
			khoaNganh, _ := conn.GetCellValue("DKGD", "E"+strconv.Itoa(i))
			nganh, _ := conn.GetCellValue("DKGD", "F"+strconv.Itoa(i))
			nhom, _ := conn.GetCellValue("DKGD", "G"+strconv.Itoa(i))
			toHop, _ := conn.GetCellValue("DKGD", "H"+strconv.Itoa(i))
			thu, _ := conn.GetCellValue("DKGD", "K"+strconv.Itoa(i))
			kip, _ := conn.GetCellValue("DKGD", "L"+strconv.Itoa(i))
			sySo, _ := conn.GetCellValue("DKGD", "M"+strconv.Itoa(i))
			phong, _ := conn.GetCellValue("DKGD", "N"+strconv.Itoa(i))
			nha, _ := conn.GetCellValue("DKGD", "O"+strconv.Itoa(i))
			hinhThucThi, _ := conn.GetCellValue("DKGD", "P"+strconv.Itoa(i))
			maGV, _ := conn.GetCellValue("DKGD", "Q"+strconv.Itoa(i))
			tenGV, _ := conn.GetCellValue("DKGD", "R"+strconv.Itoa(i))
			ghiChu, _ := conn.GetCellValue("DKGD", "S"+strconv.Itoa(i))
			//NgayBD, _ := conn.GetCellValue("DKGD", "A"+a)
			//NgayKT, _ := conn.GetCellValue("DKGD", "A"+a)
			khoa, _ := conn.GetCellValue("DKGD", "AL"+strconv.Itoa(i))
			boMon, _ := conn.GetCellValue("DKGD", "AM"+strconv.Itoa(i))
			soTC, _ := conn.GetCellValue("DKGD", "AN"+strconv.Itoa(i))
			tsTiet, _ := conn.GetCellValue("DKGD", "AO"+strconv.Itoa(i))
			lt, _ := conn.GetCellValue("DKGD", "AP"+strconv.Itoa(i))
			bt, _ := conn.GetCellValue("DKGD", "AQ"+strconv.Itoa(i))
			btl, _ := conn.GetCellValue("DKGD", "AR"+strconv.Itoa(i))
			thtn, _ := conn.GetCellValue("DKGD", "AS"+strconv.Itoa(i))
			tuHoc, _ := conn.GetCellValue("DKGD", "AT"+strconv.Itoa(i))

			tkb = &model.TKB{
				ID:          uint(helper.HashToInt(maMonHoc + nhom + tt)),
				HashKey:     uint(helper.StringToInt(tt)),
				MaMonHoc:    maMonHoc,
				TenMon:      tenMon,
				Lop:         lop,
				KhoaNganh:   khoaNganh,
				Nganh:       nganh,
				Nhom:        nhom,
				ToHop:       toHop,
				ToTH:        "",
				Thu:         thu,
				Kip:         kip,
				SySo:        sySo,
				Phong:       phong,
				Nha:         nha,
				HinhThucThi: hinhThucThi,
				MaGV:        maGV,
				TenGV:       tenGV,
				GhiChu:      ghiChu,
				NgayBD:      time.Time{},
				NgayKT:      time.Time{},
				Khoa:        khoa,
				BoMon:       boMon,
				SoTC:        soTC,
				TSTiet:      tsTiet,
				LT:          lt,
				BT:          bt,
				BTL:         btl,
				THTN:        thtn,
				TuHoc:       tuHoc,
			}

			listTKB = append(listTKB, tkb)

			//_, err := model.CreatOneTKB(tkb)
			//if err != nil {
			//	fmt.Println(err)
			//}
		}

		_, err := model.CreateManyTKB(listTKB)
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
