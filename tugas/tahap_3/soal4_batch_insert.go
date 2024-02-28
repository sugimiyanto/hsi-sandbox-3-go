package main

import (
	"fmt"
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

func PrintItem(items []Item) {
	for _, item := range items {
		fmt.Println(item.Name, item.Status, item.Amount)
	}
}

func PrintItemDetail(itemDetails []ItemDetail) {
	for _, itemDetail := range itemDetails {
		fmt.Println(itemDetail.Item_id, itemDetail.Name)
	}
}


func main() {
	dsn := "host=localhost user=sugi dbname=hsi_sandbox port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	var itemDetailRecords = []*ItemDetail{
		{Item_id: "27126432-cfd2-4ed2-9fc4-8b5ac1820621", Name: "fridge batch insert 1"},
		{Item_id: "0ce24a3a-b9ab-4336-a15f-f9342c5d1c9e", Name: "desk batch insert 1"},
		{Item_id: "27126432-cfd2-4ed2-9fc4-8b5ac1820621", Name: "fridge batch insert 2"},
		{Item_id: "892138c8-2a14-483c-af54-3ff43ef14cc1", Name: "cupboard batch insert 1"},
		{Item_id: "0ce24a3a-b9ab-4336-a15f-f9342c5d1c9e", Name: "desk batch insert 2"},
		{Item_id: "892138c8-2a14-483c-af54-3ff43ef14cc1", Name: "cupboard batch insert 2"},
	}

	batchInsertSize := 2
	db.Transaction(func(tx *gorm.DB) error {
		tx.CreateInBatches(itemDetailRecords, batchInsertSize)

		// return nil will commit the whole transaction
		return nil
	})
}