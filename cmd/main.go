package main

import (
	"github.com/BakiOztel/practice-1/pkg/controller"
	"github.com/BakiOztel/practice-1/pkg/model"
)

func main() {

	model.DB()
	controller.ServerStart()
}
