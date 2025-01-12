package main

import (
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Attribute struct {
	Sex   int
	Age   int
	Orgs  map[string]string
	Tags  []string
	Admin bool
	Role  string
}

type UserWithJSON struct {
	gorm.Model
	Name       string
	Attributes datatypes.JSONType[Attribute]
}

func main() {
	db, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&UserWithJSON{})

	var user = UserWithJSON{
		Name: "hello",
		Attributes: datatypes.NewJSONType(Attribute{
			Age:  18,
			Sex:  1,
			Orgs: map[string]string{"orga": "orga"},
			Tags: []string{"tag1", "tag2", "tag3"},
		}),
	}

	// Create
	db.Create(&user)

	// First
	var result UserWithJSON
	db.First(&result, user.ID)

	fmt.Println(result.Attributes.Data().Age)

	// Update
  jsonMap := UserWithJSON{
		Attributes: datatypes.NewJSONType(Attribute{
			Age:  18,
			Sex:  1,
			Orgs: map[string]string{"orga": "orga"},
			Tags: []string{"tag1", "tag2", "tag3"},
		}),
	}

	db.Model(&user).Updates(jsonMap)

	fmt.Println(user.Attributes.Data().Age)

}
