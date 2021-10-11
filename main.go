package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	installPrometheus(r)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "It works! Try /now form more fun!")
	})

	r.GET("/now", func(c *gin.Context) {
		now := time.Now()
		c.String(200, "%d", now.Unix())
	})

	r.NoRoute(notFound)
	r.NoMethod(notFound)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func notFound(c *gin.Context) {
	c.String(404, "not found")
}

func installPrometheus(r *gin.Engine) {
	p := ginprometheus.NewPrometheus("gin")
	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.FullPath()
		return url
	}
	p.Use(r)
}
