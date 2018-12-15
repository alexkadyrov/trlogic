package main

import (
	"context"
	"fmt"
	"github.com/alexkadyrov/trlogic/api/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	router := gin.Default()
	router.POST("/photo", handlers.GetPhoto)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		s.ListenAndServe()
	}()

	<-stop

	fmt.Println("Stopping...")

	s.Shutdown(context.Background())

	fmt.Println("Stopped")
}
