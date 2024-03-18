package app

import (
	"log"
	"vktest/internal/bootstrap"
	"vktest/internal/controller"
	"vktest/internal/repository"
	"vktest/internal/service"
)

func Run() error {

	db, err := repository.ConnectDb()
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	store := repository.NewCinematheque(db)

	srv := service.NewService(store)

	cnt := controller.NewController(srv)

	serv := bootstrap.NewServer(cnt)

	serv.InitRoutes()

	serv.Run()

	return nil
}
