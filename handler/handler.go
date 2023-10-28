package handler

import (
	"bufio"
	"fmt"
	"os"

	"github.com/amirthapa27/precize-test/database"
	"github.com/amirthapa27/precize-test/helpers"
	"github.com/amirthapa27/precize-test/models"

	"github.com/olekukonko/tablewriter"
)

func InsertData(scanner *bufio.Scanner) {
	// Enter and check the details
	name := helpers.ReadField(scanner, "Name", false)
	address := helpers.ReadField(scanner, "Address", false)
	city := helpers.ReadField(scanner, "City", false)
	country := helpers.ReadField(scanner, "Country", false)
	pincode := helpers.ReadField(scanner, "Pincode", true)
	score := helpers.ReadField(scanner, "Score", true)
	passed := helpers.CalculateIfPassed(score.(int))
	// store all the data to create a new entry
	newStudent := models.SAT_Results{
		Name:    name.(string),
		Address: address.(string),
		City:    city.(string),
		Country: country.(string),
		Pincode: pincode.(int),
		Score:   score.(int),
		Passed:  passed,
	}
	// check if the student already exists
	var student []models.SAT_Results
	database.DB.Where("name=?", name).Find(&student)
	if len(student) > 0 {
		fmt.Println("")
		fmt.Println("***Name already exists***")
		fmt.Println("")
		return
	}
	// create a new student
	result := database.DB.Create(&newStudent)
	// handle error
	if helpers.CheckError(result.Error) {
		return
	}
	fmt.Println("")
	fmt.Println("---Data inserted successfully---")
	fmt.Println("")
}

func ViewAllData() {
	// define a varible of the model type
	var SAT_Results []models.SAT_Results
	// fetch the data from database and store in the variable
	result := database.DB.Find(&SAT_Results)
	// handle error
	if helpers.CheckError(result.Error) {
		return
	}
	// create a table view using tablewritter package
	table := tablewriter.NewWriter(os.Stdout)
	// define headers
	table.SetHeader([]string{"ID", "Name", "Address", "City", "Country", "Pincode", "Score", "Passed"})
	// store all the data in the table
	for _, student := range SAT_Results {
		dataRow := []string{
			fmt.Sprint(student.ID),
			student.Name,
			student.Address,
			student.City,
			student.Country,
			fmt.Sprint(student.Pincode),
			fmt.Sprint(student.Score),
			student.Passed,
		}
		table.Append(dataRow)
	}
	// print the table
	fmt.Println("")
	table.Render()
	fmt.Println("")
}

func GetRank(scanner *bufio.Scanner) {
	name := helpers.ReadField(scanner, "Name", false)
	var rank int64
	var student models.SAT_Results
	// check if the student exists
	result := database.DB.Where("Name = ?", name).First(&student)
	if helpers.CheckUserDoseNotExists(result.Error) {
		return
	}
	// query for getting the rank of the student
	result = database.DB.Model(&models.SAT_Results{}).Where("score > ?", student.Score).Count(&rank)
	if helpers.CheckError(result.Error) {
		return
	}
	// increment the rank
	rank++
	fmt.Println("")
	fmt.Printf("---%s ranks on number %d---\n", name, rank)
	fmt.Println("")
}

func UpdateScore(scanner *bufio.Scanner) {
	name := helpers.ReadField(scanner, "Name", false)
	newScore := helpers.ReadField(scanner, "Score", true)
	passed := helpers.CalculateIfPassed(newScore.(int))
	var student models.SAT_Results
	result := database.DB.Where("Name = ?", name).First(&student)
	if helpers.CheckUserDoseNotExists(result.Error) {
		return
	}
	student.Score = newScore.(int)
	student.Passed = passed
	// update the new score
	result = database.DB.Save(&student)
	if helpers.CheckError(result.Error) {
		return
	}
	fmt.Println("")
	fmt.Println("---Score updated---")
	fmt.Println("")
}

func DeleteOne(scanner *bufio.Scanner) {
	name := helpers.ReadField(scanner, "Name", false)
	var student models.SAT_Results
	result := database.DB.Where("Name = ?", name).First(&student)
	if helpers.CheckUserDoseNotExists(result.Error) {
		return
	}
	// delete the SAT_Results record
	result = database.DB.Delete(&student)
	if helpers.CheckError(result.Error) {
		return
	}
	fmt.Println("")
	fmt.Println("---Deleted Successfully---")
	fmt.Println("")
}
