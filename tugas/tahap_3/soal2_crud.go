package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
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

	testNum := 3
	init_items := []Item{}
	init_item_details := []ItemDetail{}
	ids := []string{"27126432-cfd2-4ed2-9fc4-8b5ac1820621", "892138c8-2a14-483c-af54-3ff43ef14cc1", "0ce24a3a-b9ab-4336-a15f-f9342c5d1c9e"}
	names := []string{"desk", "cupboard", "fridge"}
	statuses := []string{"active", "active", "inactive"}
	for i := 0; i < testNum; i++ {
		init_items = append(init_items, Item{Id: ids[i], Name: names[i], Status: statuses[i], Amount: rand.Intn(100)})
		init_item_details = append(init_item_details, ItemDetail{Item_id: ids[i], Name: names[i] + " long description"})
		ids = append(ids, ids[i])
	}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&init_items)
	if result.Error != nil {
		panic(result.Error)
	}
	result = db.Create(&init_item_details)
	if result.Error != nil {
		panic(result.Error)
	}

	var items []Item
	var itemDetails []ItemDetail
	fmt.Println("=== querying ===")
	db.Find(&items)
	db.Find(&itemDetails)
	fmt.Println("Item data:")
	PrintItem(items)
	fmt.Println("\nItemDetail data:")
	PrintItemDetail(itemDetails)

	fmt.Println("\n\n= query with filter")
	db.Find(&items, "status <> ?", "inactive")
	fmt.Println("\nItem data:")
	PrintItem(items)
	fmt.Println("\nItemDetail data:")
	PrintItemDetail(itemDetails)

	fmt.Println("\n=== updating ===")
	randNum := rand.Intn(100)
	result = db.Model(&Item{}).Where("name = ?", "cupboard").Update("amount", randNum)
	if result.Error != nil {
		panic("Update failed")
	}
	db.Find(&items)
	PrintItem(items)

	fmt.Println("\n=== deleting ===")
	db.Find(&items)
	db.Find(&itemDetails)
	result = db.Where("item_id = ?", ids[0]).Delete(&itemDetails)
	if result.Error != nil {
		panic("Delete failed")
	}
	fmt.Println("Deleted rows: ", result.RowsAffected)
}