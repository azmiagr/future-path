package mariadb

import (
	"future-path/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
		&entity.FAQ{},
		&entity.Berita{},
	); err != nil {
		return err
	}
	return nil
}
