package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

type Item struct {
	Id		string
	Name	string
	Status	string
	Amount	int
}

type ItemDetail struct {
	Item_id	string
	Name	string
}

// override Go behavior that pluralize struct names when generating/targeting a table.
func (Item) TableName() string {
	return "item"
}

// override Go behavior that pluralize struct names when generating/targeting a table.
func (ItemDetail) TableName() string {
	return "item_detail"
}


func main() {
	dsn := "host=localhost user=sugi dbname=hsi_sandbox port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.Transaction(func(tx *gorm.DB) error {
		testId := "0ce24a3a-b9ab-4336-a15f-f9342c5d1c9e"
		if err := tx.Create(ItemDetail{Item_id: testId, Name: "desk another long long description"}).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.Model(&Item{}).Where("id = ?", testId).Update("amount", 999).Update("status", "inactive").Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
}