package mariadb

import (
	"future-path/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Role{},
		&entity.FAQ{},
		&entity.Berita{},
		&entity.Admin{},
		&entity.User{},
	); err != nil {
		return err
	}
	return nil
}
