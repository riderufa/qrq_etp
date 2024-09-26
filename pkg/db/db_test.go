package db

import (
	"reflect"
	"testing"
)

func TestDB_PreSearch(t *testing.T) {
	db := New()
	ps := PreSearch{
		EtpID:    "EtpID",
		Article:  "Article",
		Brand:    "Brand",
		PartName: "PartName",
	}

	ps.ID = db.NewPreSearch(ps)
	preSearches := db.PreSearches()
	if !reflect.DeepEqual(preSearches[0], ps) {
		t.Errorf("не найден созданный заказ")
	}
}
