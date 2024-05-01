package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request-utility/internal/db"

	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	DB := db.StartConnection()

	SeedDatabase(DB)

	fmt.Println("Seed applied!")
}

func SeedDatabase(DB *gorm.DB) {
	var formCount int64
	DB.Model(&dto.Form{}).Count(&formCount)
	if formCount == 0 {
		DB.Create([]*dto.Form{
			{
				Name:     "Driving License",
				Category: "General",
				Sections: []*dto.Section{
					{
						Name: "Biodata",
						Fields: []*dto.Field{
							{
								Label:      "Name",
								Type:       "String",
								FieldOrder: 1,
							},
							{
								Label:      "Birth Date",
								Type:       "Date",
								FieldOrder: 2,
							},
							{
								Label:      "Gender",
								Type:       "Option",
								FieldOrder: 3,
							},
						},
					},
					{
						Name: "General Knowledge",
						Fields: []*dto.Field{
							{
								Label:      "Technical Skill (1-10)",
								Type:       "Number",
								FieldOrder: 1,
							},
							{
								Label:      "Practical Experience (1-10)",
								Type:       "Number",
								FieldOrder: 2,
							},
							{
								Label:      "Note",
								Type:       "Text",
								FieldOrder: 3,
							},
						},
					},
				},
			},
		})
	}

	var userCount int64
	DB.Model(&dto.User{}).Count(&userCount)
	if userCount == 0 {
		DB.Create([]*dto.User{
			{
				ID:            "ADMIN-001",
				Email:         "admin@email.com",
				Picture:       "https://lh3.googleusercontent.com/a-/ALV-UjUiYmyOrvOAaQsnGfnVDs6QrtTn7sLD5ICee2OcVz32pAi6Pwj1=s96-c",
				VerifiedEmail: true,
				Password:      "",
				Role:          "ADMIN",
			},
			{
				ID:            "USER-001",
				Email:         "user@email.com",
				Picture:       "https://lh3.googleusercontent.com/a-/ALV-UjUiYmyOrvOAaQsnGfnVDs6QrtTn7sLD5ICee2OcVz32pAi6Pwj1=s96-c",
				VerifiedEmail: true,
				Password:      "",
				Role:          "USER",
			},
		})
	}

	var applicationCount int64
	DB.Model(&dto.Application{}).Count(&applicationCount)
	if applicationCount == 0 {
		DB.Create([]*dto.Application{
			{
				Number: "DRV-001",
				UserID: "USER-001",
				FormID: 1,
				Responses: []*dto.Response{
					{
						FieldID: 1,
						Value:   "Basic User",
					},
					{
						FieldID: 2,
						Value:   "2001-28-03",
					},
					{
						FieldID: 3,
						Value:   "Female",
					},
					{
						FieldID: 4,
						Value:   "8",
					},
					{
						FieldID: 5,
						Value:   "7",
					},
					{
						FieldID: 6,
						Value:   "Little bit too aggressive in corner but has very good handling and fast response",
					},
				},
			},
			{
				Number: "DRV-002",
				UserID: "USER-001",
				FormID: 1,
				Responses: []*dto.Response{
					{
						FieldID: 1,
						Value:   "Basic User",
					},
					{
						FieldID: 3,
						Value:   "Female",
					},
					{
						FieldID: 4,
						Value:   "8",
					},
				},
			},
		})
	}

}
