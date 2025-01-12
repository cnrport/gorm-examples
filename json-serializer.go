package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type User struct {
  gorm.Model
	Name        []byte                 `gorm:"serializer:json"`
	Roles       Roles                  `gorm:"serializer:json"`
	Contracts   map[string]interface{} `gorm:"serializer:json"`
	JobInfo     Job                    `gorm:"type:bytes;serializer:gob"`
	CreatedTime int64                  `gorm:"serializer:unixtime;type:time"` // store int as datetime into database
}

type Roles []string

type Job struct {
	Title    string
	Location string
	IsIntern bool
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	createdAt := time.Date(2020, 1, 1, 0, 8, 0, 0, time.UTC)
	data := User{
		Name:        []byte("jinzhu"),
		Roles:       []string{"admin", "owner"},
		Contracts:   map[string]interface{}{"name": "jinzhu", "age": 10},
		CreatedTime: createdAt.Unix(),
		JobInfo: Job{
			Title:    "Developer",
			Location: "NY",
			IsIntern: false,
		},
	}

	db.Create(&data)
	// INSERT INTO `users` (`name`,`roles`,`contracts`,`job_info`,`created_time`) VALUES
	//   ("\"amluemh1\"","[\"admin\",\"owner\"]","{\"age\":10,\"name\":\"jinzhu\"}",<gob binary>,"2020-01-01 00:08:00")

	var result User
	db.First(&result, "id = ?", data.Model.ID)
	// result => User{
	//   Name:        []byte("jinzhu"),
	//   Roles:       []string{"admin", "owner"},
	//   Contracts:   map[string]interface{}{"name": "jinzhu", "age": 10},
	//   CreatedTime: createdAt.Unix(),
	//   JobInfo: Job{
	//     Title:    "Developer",
	//     Location: "NY",
	//     IsIntern: false,
	//   },
	// }

	db.Where(User{Name: []byte("jinzhu")}).Take(&result)
	// SELECT * FROM `users` WHERE `users`.`name` = "\"amluemh1\"

	fmt.Println(result.JobInfo.Title)
  fmt.Println(result.JobInfo)
}
