package orm

import (
	"log"

	"github.com/biFebriansyah/gobackend/src/database/orm/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrasi database",
	RunE:  dbMigrate,
}

var migUp bool
var migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "0001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.CartItem{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&models.User{}, &models.Product{}, &models.Cart{}, &models.CartItem{})
			},
		},
	})

	if migUp {
		if err := m.Migrate(); err != nil {
			return err
		}
		log.Println("Migration up done")
		return nil
	}

	if migDown {
		if err := m.RollbackLast(); err != nil {
			return err
		}
		log.Println("Rollback database done")
		return nil
	}

	log.Println("init schema database done")
	return nil
}
