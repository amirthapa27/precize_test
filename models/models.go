package models

type SAT_Results struct {
	ID      uint32 `gorm:"primaryKey"`
	Name    string `gorm:"unique"`
	Address string
	City    string
	Country string
	Pincode int
	Score   int
	Passed  string
}
