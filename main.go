package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type Timeline struct {
	gorm.Model
	Name   string `json:"name" gorm:"column:name"`
	Text   string `json:"text" gorm:"column:text"`
	UserID int    `json:"user_id" gorm:"column:user_id"`
}

func main() {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Timeline{})

	e := echo.New()
	e.HideBanner = true
	e.GET("/timeline", func(c echo.Context) error {
		var timelines []Timeline
		db.Find(&timelines)
		return c.JSON(200, timelines)
	})

	e.POST("/timeline", func(c echo.Context) error {
		// リクエストパラメーターを取得
		timeline := new(Timeline)
		err := c.Bind(timeline)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		// DBにINSERT
		if db.Create(&timeline).Error != nil {
			return c.String(http.StatusInternalServerError, "DB Error")
		}
		// レスポンス
		return c.JSON(http.StatusOK, timeline)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
