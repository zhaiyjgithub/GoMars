package service

import (
	"GoMars/src/dao"
	"GoMars/src/database"
	"GoMars/src/model"
)

type CarService interface {
	LaunchNow(name string) (error, *model.Car, bool)
	RotateLeft(name string) error
	RotateRight(name string) error
	MoveForward(name string) error
	MoveBack(name string) error
	DetectAround(name string) error
	GetCoverage(name string) float64
	GetCurrentPosition(name string) (error, *model.Car)
}

func NewCarService() CarService {
	return &carService{dao: dao.NewCarDao(database.SetupRedisEngine())}
}

type carService struct {
	dao *dao.CarDao
}

func (s *carService) LaunchNow(name string) (error, *model.Car, bool)  {
	return s.dao.LaunchNow(name)
}

func (s *carService) RotateLeft(name string) error  {
	return s.dao.RotateLeft(name)
}

func (s *carService) RotateRight(name string) error  {
	return s.dao.RotateRight(name)
}

func (s *carService) MoveForward(name string) error  {
	return s.dao.MoveForward(name)
}

func (s *carService) MoveBack(name string) error  {
	return s.dao.MoveBack(name)
}

func (s *carService) DetectAround(name string) error  {
	return s.dao.DetectAround(name)
}

func (s *carService) GetCoverage(name string) float64  {
	return s.dao.GetCoverage(name)
}

func (s *carService)GetCurrentPosition(name string) (error, *model.Car)  {
	return s.dao.GetCurrentPosition(name)
}