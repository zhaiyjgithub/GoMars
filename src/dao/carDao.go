package dao

import (
	"GoMars/src/constant"
	"GoMars/src/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

type CarDao struct {
	engine *redis.Client
}

func NewCarDao(engine *redis.Client) *CarDao {
	return &CarDao{engine: engine}
}

func (d *CarDao) LaunchNow(name string) (error, *model.Car, bool) {
	err, car, _ := d.CheckCarIsExisting(name)
	if err == nil && car != nil {
		return err, car, false
	}
	return d.LaunchNewCar(name)
}

func (d *CarDao) LaunchNewCar(name string) (error, *model.Car, bool) {
	x := rand.Intn(24)
	y := rand.Intn(24)

	car := new(model.Car)
	car.Name = name
	car.X = x
	car.Y = y
	car.Direction = constant.North
	err := d.updateCurrentPosition(name, car)
	if err != nil {
		return err, car, true
	}
	_, err = d.addTrackRecord(name, &model.Position{X: car.X, Y: car.Y})
	return err, car, true
}

func (d *CarDao) CheckCarIsExisting(name string) (error, *model.Car, bool) {
	err, car := d.GetCurrentPosition(name)
	return err, car, true
}

func (d *CarDao) RotateLeft(name string) error{
	err, car := d.GetCurrentPosition(name)
	if err != nil {
		return err
	}

	dir := car.Direction
	if dir == constant.North {
		car.Direction = constant.West
	}else if dir == constant.West {
		car.Direction = constant.South
	}else if dir == constant.South {
		car.Direction = constant.East
	}else if dir == constant.East {
		car.Direction = constant.North
	}
	fmt.Println("rotate left, new position: ", car)
	err = d.updateCurrentPosition(name, car)
	return err
}

func (d *CarDao) RotateRight(name string) error{
	err, car := d.GetCurrentPosition(name)
	if err != nil {
		return err
	}

	dir := car.Direction
	if dir == constant.North {
		car.Direction = constant.East
	}else if dir == constant.East {
		car.Direction = constant.South
	}else if dir == constant.South {
		car.Direction = constant.West
	}else if dir == constant.West {
		car.Direction = constant.North
	}
	fmt.Println("rotate right, new position: ", car)
	err = d.updateCurrentPosition(name, car)
	return err
}

func (d *CarDao) MoveForward(name string) error {
	err, car := d.GetCurrentPosition(name)
	if err != nil {
		return err
	}
	dir := car.Direction
	if dir == constant.North {
		if car.Y == 0 {
			err = errors.New("error: up border now, operation is terminated")
			return err
		}
		car.Y = car.Y - 1
	}else if dir == constant.East {
		if car.X == 24 {
			err = errors.New("error: right border now, operation is terminated")
			return err
		}
		car.X = car.X + 1
	}else if dir == constant.South {
		if car.Y == 24 {
			err = errors.New("error: bottom border now, operation is terminated")
			return err
		}
		car.Y = car.Y + 1
	}else if dir == constant.West {
		if car.X == 0 {
			err = errors.New("error: left border now, operation is terminated")
			return err
		}
		car.X = car.X - 1
	}
	err = d.updateCurrentPosition(name, car)
	if err != nil {
		return err
	}
	_, err = d.addTrackRecord(name, &model.Position{X: car.X, Y: car.Y})
	return err
}

func (d *CarDao) MoveBack(name string) error {
	err, car := d.GetCurrentPosition(name)
	if err != nil {
		return err
	}
	dir := car.Direction
	if dir == constant.North {
		if car.Y == 24 {
			err = errors.New("move back error: bottom border now, operation is terminated")
			return err
		}
		car.Y = car.Y + 1
	}else if dir == constant.East {
		if car.X == 0 {
			err = errors.New("error: right border now, operation is terminated")
			return err
		}
		car.X = car.X - 1
	}else if dir == constant.South {
		if car.Y == 0 {
			err = errors.New("error: bottom border now, operation is terminated")
			return err
		}
		car.Y = car.Y - 1
	}else if dir == constant.West {
		if car.X == 24 {
			err = errors.New("error: left border now, operation is terminated")
			return err
		}
		car.X = car.X + 1
	}
	err = d.updateCurrentPosition(name, car)
	if err != nil {
		return err
	}
	_, err = d.addTrackRecord(name, &model.Position{X: car.X, Y: car.Y})
	return err
}

func (d *CarDao) DetectAround(name string) error {
	err, car := d.GetCurrentPosition(name)
	if err != nil {
		return err
	}

	list := make([]*model.Position, 0)
	p0 := &model.Position{X: car.X, Y: car.Y - 1} //
	if d.isPositionValid(p0) {
		list = append(list, p0)
	}
	p1 := &model.Position{X: car.X + 1, Y: car.Y - 1}
	if d.isPositionValid(p1) {
		list = append(list, p1)
	}
	p2 := &model.Position{X: car.X + 1, Y: car.Y}
	if d.isPositionValid(p2) {
		list = append(list, p2)
	}
	p3 := &model.Position{X: car.X + 1, Y: car.Y + 1}
	if d.isPositionValid(p3) {
		list = append(list, p3)
	}
	p4 := &model.Position{X: car.X, Y: car.Y + 1}
	if d.isPositionValid(p4) {
		list = append(list, p4)
	}
	p5 := &model.Position{X: car.X - 1, Y: car.Y + 1}
	if d.isPositionValid(p5) {
		list = append(list, p5)
	}
	p6 := &model.Position{X: car.X - 1, Y: car.Y}
	if d.isPositionValid(p6) {
		list = append(list, p6)
	}
	p7 := &model.Position{X: car.X - 1, Y: car.Y - 1}
	if d.isPositionValid(p7) {
		list = append(list, p7)
	}

	for _, pos := range list {
		fmt.Println("info: add detected pos: {", pos.X, ",", pos.Y, "}")
		_, err = d.addTrackRecord(name, &model.Position{X: pos.X, Y: pos.Y})
	}

	return nil
}

func (d *CarDao) isPositionValid(pos *model.Position) bool {
	return pos.X >= 0 && pos.X <= 24 && pos.Y >= 0 && pos.Y <= 24
}

func (d *CarDao) GetCoverage(name string) float64 {
	key := name + "-" + "track"
	items, err := d.engine.SMembers(context.Background(), key).Result()
	if err != nil {
		fmt.Println("error: get coverage error ", err)
	}
	fmt.Println("all track records: ", len(items))
	return float64(len(items))/625.0
}

func (d *CarDao) updateCurrentPosition(name string, car *model.Car) error {
	key := name + "-" + "position"
	b, err := json.Marshal(car)
	if err != nil {
		fmt.Println("Error: Parse track record position failed")
		return  err
	}
	fmt.Println("Begin to update current position:", car)
	err = d.engine.Set(context.Background(), key, string(b), time.Hour*20).Err()

	return err
}

func (d *CarDao) GetCurrentPosition(name string) (error, *model.Car) {
	key := name + "-" + "position"
	val, err := d.engine.Get(context.Background(), key).Result()
	if err != nil {
		return err, nil
	}
	var car model.Car
	err = json.Unmarshal([]byte(val), &car)
	return err, &car
}

func (d *CarDao) addTrackRecord(name string, pos *model.Position) (int64, error) {
	key := name + "-" + "track"
	b, err := json.Marshal(pos)
	if err != nil {
		fmt.Println("Error: Parse track record position failed")
		return 0, err
	}
	fmt.Println("Begin to insert new position:", pos)
	code, err := d.engine.SAdd(context.Background(), key, string(b)).Result()

	return code, err
}

