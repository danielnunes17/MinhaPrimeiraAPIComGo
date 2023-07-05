package main

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/labstack/echo/v4"
)

type Car struct {
	Name  string
	Price float64
}

var cars []Car

func generateCar() {
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
	cars = append(cars, Car{Name: "Fusca", Price: 60})
	cars = append(cars, Car{Name: "Scort", Price: 50})
	cars = append(cars, Car{Name: "Gol", Price: 40})
}

func main() {
	generateCar()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCars)
	e.Logger.Fatal(e.Start(":8080"))

}
func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCars(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); c != nil {
		return err
	}
	cars = append(cars, *car)
	saveCars(*car)
	return c.JSON(200, cars)
}

func saveCars(car Car) error {
	db, err := sql.Open("mssqldb", "Car.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES (@1, @2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}
