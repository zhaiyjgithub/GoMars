package controller

import (
	"GoMars/src/response"
	"GoMars/src/router"
	"GoMars/src/service"
	"GoMars/src/utils"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
	"strings"
)

type CarController struct {
	Ctx iris.Context
	CarService service.CarService
}

func (c *CarController) BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle(http.MethodPost, router.LaunchNow, "LaunchNow")
	b.Handle(http.MethodPost, router.SendCommand, "SendCommand")
	b.Handle(http.MethodPost, router.GetCoverage, "GetCoverage")
	b.Handle(http.MethodPost, router.GetCurrentPosition, "GetCurrentPosition")
}

func (c *CarController)LaunchNow()  {
	type Param struct {
		Name string
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	err, car, isNew := c.CarService.LaunchNow(p.Name)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), car)
	}else {
		msg := ""
		if isNew {
			msg = fmt.Sprintf("Launch new car[%s] success", p.Name)
		}else {
			msg = fmt.Sprintf("You have re-control this car[%s].", p.Name)
		}
		response.Success(c.Ctx, msg, car)
	}
}

func (c *CarController)SendCommand()  {
	type Param struct {
		Name string
		Command string
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	upCommand := strings.ToUpper(p.Command)
	for _, cmd := range upCommand {
		if cmd == 'F' { // 向北前进一步
			fmt.Println("Command: move forward")
			err = c.CarService.MoveForward(p.Name)
		}else if cmd == 'B' {
			fmt.Println("Command: move back")
			err = c.CarService.MoveBack(p.Name)
		}else if cmd == 'L' {
			fmt.Println("Command: rotate left")
			err = c.CarService.RotateLeft(p.Name)
		}else if cmd == 'R' {
			fmt.Println("Command: rotate right")
			err = c.CarService.RotateRight(p.Name)
		}else if cmd == 'H' {
			fmt.Println("Command: detect around")
			err = c.CarService.DetectAround(p.Name)
		}

		if err != nil {
			response.Fail(c.Ctx, response.Error, err.Error(), nil)
			fmt.Println("Execute cmd failed, terminated!!!")
			return
		}else {
			err, car := c.CarService.GetCurrentPosition(p.Name)
			fmt.Println(err, "Each command new position: ",car)
		}
	}

	err, car := c.CarService.GetCurrentPosition(p.Name)
	fmt.Println(err, " Current position: ",car)
	response.Success(c.Ctx, "", car)
}

func (c *CarController) GetCoverage() {
	type Param struct {
		Name string
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	rate := c.CarService.GetCoverage(p.Name)
	fmt.Println("current rate:", rate)
	response.Success(c.Ctx, response.Successful, rate)
}

func (c *CarController) GetCurrentPosition()  {
	type Param struct {
		Name string
	}

	var p Param
	err := utils.ValidateParam(c.Ctx, &p)
	if err != nil {
		return
	}

	err, car := c.CarService.GetCurrentPosition(p.Name)
	if err != nil {
		response.Fail(c.Ctx, response.Error, err.Error(), nil)
	}else {
		response.Success(c.Ctx, response.Successful, car)
	}
}