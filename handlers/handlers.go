package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ReleaseList struct {
	gorm.Model
	Year int    `json:"year"`
	Name string `json:"name"`
	City string `json:"city"`
}

type ErrorMassage struct {
	Massage  string `json:"massage,omitempty"`
	NORecord int    `json:"effected records,omitempty"`
	Err      error  `json:"error,omitempty"`
	//Data     ReleaseList `json:"data,omitempty"`
	ID string `json:"id,omitempty"`
}

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("gta.db"), &gorm.Config{})
	if err != nil {
		log.Panicf("Faild to connect database: %v", err)
	}
}

// GetData is GET methods handler
func GetData(c echo.Context) error {
	var fullList []ReleaseList
	result := db.Find(&fullList)
	if result.Error != nil {
		return result.Error
	}
	return c.JSON(http.StatusOK, fullList)
}

// InsertData is Post methods handler it creates on or more records.
// data struct must be between two pair of square brackets
func InsertData(c echo.Context) error {
	db.AutoMigrate(&ReleaseList{})

	var newRealiseList []ReleaseList
	if err := c.Bind(&newRealiseList); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorMassage{Massage: "Unable to bind request body"})
	}
	insertedData := db.Create(&newRealiseList)

	if insertedData.Error != nil {
		log.Printf("Unable to insert the data: %v", insertedData.Error)

		return c.JSON(http.StatusBadRequest, ErrorMassage{Massage: "Unable to insert the data"})
	}
	return c.JSON(http.StatusCreated, ErrorMassage{Massage: "Data inserted", NORecord: int(insertedData.RowsAffected)})
}

func UpdateData(c echo.Context) error {
	var data ReleaseList
	id := c.Param("id")
	result := db.First(&data, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, ErrorMassage{Massage: "Unable to find a record with this id"})
	}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorMassage{Massage: "Unable to bind request payload.", Err: err})
	}
	saveResult := db.Save(&data)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, ErrorMassage{Massage: "Unable to save.", Err: result.Error})
	}
	return c.JSON(http.StatusOK, ErrorMassage{Massage: "Updated", ID: id, NORecord: int(saveResult.RowsAffected)})
}

// findOne finds a record with its id
func findOne(id string, c echo.Context) (ReleaseList, int) {
	var data ReleaseList
	result := db.First(&data, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return data, http.StatusNotFound
	}
	if result.Error != nil {
		return data, http.StatusInternalServerError
	}
	return data, 0
}

// DeleteRecord is handler for DELETE method
// it will soft delete a record with given id
func DeleteRecord(c echo.Context) error {
	id := c.Param("id")
	_, errNum := findOne(id, c)
	if errNum != 0 {
		if errNum == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, ErrorMassage{Massage: "Unable to find a record with this id", ID: id})
		}
		if errNum == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, ErrorMassage{Massage: "Error while finding record"})
		}
	}
	result := db.Delete(&ReleaseList{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorMassage{Massage: "Unable to delete the record with this id", ID: id})
	}
	return c.JSON(http.StatusOK, ErrorMassage{Massage: "Record deleted", ID: id})
}

func GetSingle(c echo.Context) error {
	id := c.Param("id")
	data, errNum := findOne(id, c)
	if errNum != 0 {
		if errNum == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, ErrorMassage{Massage: "Unable to find a record with this id", ID: id})
		}
		if errNum == http.StatusInternalServerError {
			return c.JSON(http.StatusInternalServerError, ErrorMassage{Massage: "Error while finding record"})
		}
	}
	return c.JSON(http.StatusOK, data)
}
