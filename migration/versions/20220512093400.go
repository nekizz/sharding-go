package versions

import (
	"gorm.io/gorm"
	"time"
)

func Version20220512093400(tx *gorm.DB) error {
	type TKB struct {
		ID          uint   `gorm:"AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
		HashKey     uint   `gorm:"AUTO_INCREMENT;NOT NULL"`
		MaMonHoc    string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		TenMon      string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Lop         string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		KhoaNganh   string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Nganh       string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Nhom        string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		ToHop       string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		ToTH        string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		SoLop       string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Thu         string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Kip         string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		SySo        string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Phong       string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Nha         string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		HinhThucThi string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		MaGV        string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		TenGV       string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		GhiChu      string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		NgayBD      time.Time
		NgayKT      time.Time
		Khoa        string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		BoMon       string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		SoTC        string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		TSTiet      string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		LT          string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		BT          string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		BTL         string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		THTN        string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		TuHoc       string `gorm:"CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
	}

	return tx.AutoMigrate(&TKB{})
}
