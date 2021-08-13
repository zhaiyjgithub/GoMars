package dao

import (
	"GoMars/src/constant"
	"GoMars/src/database"
	"GoMars/src/model"
	"fmt"
	"github.com/prashantv/gostub"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

var dao = NewCarDao(database.SetupRedisEngine())

func TestCarDao_LaunchNewCar(t *testing.T) {
	convey.Convey("Test launch new car", t, func() {
		name := "Tom2"

		convey.Convey("car name must not be empty", func() {
			convey.So(name, convey.ShouldNotBeEmpty)

			err, car, isNew := dao.LaunchNewCar(name)
			convey.Convey("err must be nil", func() {
				convey.So(err, convey.ShouldBeNil)
			})

			convey.Convey("car must not be nil", func() {
				convey.So(car, convey.ShouldNotBeNil)
			})

			nameExpected := fmt.Sprintf("car.Name must be %s", name)
			convey.Convey(nameExpected, func() {
				convey.So(car.Name, convey.ShouldEqual, name)
			})

			convey.Convey("isNew must be true", func() {
				convey.So(isNew, convey.ShouldBeTrue)
			})
		})
	})
}

func TestCarDao_MoveForward(t *testing.T) {
	convey.Convey("Test move forward", t, func() {
		name:= "Jack-car2"
		convey.Convey("car name must not be empty", func() {
			convey.So(name, convey.ShouldNotBeEmpty)
			_, car, _ := dao.LaunchNewCar(name)

			mockCar := &model.Car{
				Name:      name,
				Position:  model.Position{X: 15, Y: 15},
				Direction: constant.North,
			}
			carStub := gostub.Stub(&car, mockCar)
			defer carStub.Reset()
			fmt.Println("mock car", mockCar.Name, mockCar.X, mockCar.Y, mockCar.Direction)

			_ = dao.updateCurrentPosition(name, mockCar)
			err := dao.MoveForward(name)
			convey.Convey("Move forward one step, err must be nil", func() {
				convey.So(err, convey.ShouldBeNil)

				err, car = dao.GetCurrentPosition(name)
				carxExpected := 14
				convey.Convey("After move forward on step, get the newest position", func() {
					convey.So(car.Y, convey.ShouldEqual, carxExpected)
				})
			})
		})
	})
}