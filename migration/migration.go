package migration

import (
	"gorm.io/gorm"
	"shrading/migration/versions"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:      "20220512093400",
			Migrate: versions.Version20220512093400,
		},
	})

	return m.Migrate()
}
