package service

import (
	"fmt"
	"server/app/model/dictionaries"
	"testing"

	"github.com/gogf/gf/frame/g"
)

func TestFindDictionary(t *testing.T) {
	var err error
	var dictionary *dictionaries.DictionaryHasManyDetails
	dictionary = (*dictionaries.DictionaryHasManyDetails)(nil)
	db := g.DB("default").Table("dictionaries").Safe()
	detailDb := g.DB("default").Table("dictionary_details").Safe()
	err = db.Where(g.Map{"id": 5}).Or(g.Map{"`type`": ""}).Struct(&dictionary)
	err = detailDb.Structs(&dictionary.DictionaryDetails, "dictionary_id", dictionary.Id)
	fmt.Println(err)
}
