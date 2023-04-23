package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"test-go/internal/config"
	"test-go/internal/user"
	"test-go/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("router created")
	router := httprouter.New()

	cfg := config.GetConfig()

	//mongoDBClient, err := mongodb.NewClient(
	//	context.Background(),
	//	cfg.MongoDB.Host,
	//	cfg.MongoDB.Port,
	//	cfg.MongoDB.Username,
	//	cfg.MongoDB.Password,
	//	cfg.MongoDB.Database,
	//	cfg.MongoDB.AuthDB,
	//)
	//if err != nil {
	//	panic(err)
	//}

	//user1 := user.User{
	//	ID:           "",
	//	Email:        "text@mail.ru",
	//	Username:     "text",
	//	PasswordHash: "bgejrbgergv",
	//}
	//
	//storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Collection, logger)
	//user1ID, err := storage.Create(context.Background(), user1)
	//if err != nil {
	//	panic(err)
	//}
	//logger.Info(user1ID)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("soket path is: %s", socketPath)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
	} else {
		logger.Infof("listen tcp on %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("server listening")
	log.Fatal(server.Serve(listener))
}
