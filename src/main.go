package main

import (
	"GoMars/src/controller"
	"GoMars/src/router"
	"GoMars/src/service"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main()  {
	fmt.Println("Let's go to Mars")
	app := iris.New()
	mvc.Configure(app.Party(router.Car), configureCarMVC)
	_ = app.Run(iris.Addr(":8088"), iris.WithPostMaxMemory(32<<20))
}

func configureCarMVC(app *mvc.Application)  {
	carService := service.NewCarService()
	app.Register(carService)
	app.Handle(new(controller.CarController))
}
