package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Kartik30R/studentsGo.git/internal/config"
	"github.com/Kartik30R/studentsGo.git/internal/http/handler/students"
	"github.com/Kartik30R/studentsGo.git/storage/storage/sqllite"
)

func main() {
	fmt.Println("Welcome to student API")

	// Load config
	cfg := config.MustLoad()
	//database setup

	storage, err:= sqllite.New(cfg)
	if err!=nil{
		log.Fatal(err)
	}

	slog.Info("storage initialized",slog.String("env",cfg.Env),slog.String("version","1.0.0"))

	// Setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /api/students", students.New(storage))

	// Setup and start server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Println("Server starting on", cfg.HTTPServer.Addr)
	
	done:= make(chan os.Signal, 1)


	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func(){

		if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	}()

<-done

slog.Info("shutting down the server")
ctx,cancle:=context.WithTimeout(context.Background(),5* time.Second)
defer cancle()
if err:= server.Shutdown(ctx); err!= nil {
	slog.Error("failed to shutdown serve", slog.String("error", err.Error()))
}
slog.Info("server shutdown success")

}
