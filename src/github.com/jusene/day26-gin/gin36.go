package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	g errgroup.Group
)

func router1() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Welcome server 1",
		})
	})
	return e
}

func router2() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "Welcome server 2",
		})
	})
	return e
}

func main() {
	server1 := &http.Server{
		Addr:         ":8080",
		Handler:      router1(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server2 := &http.Server{
		Addr:         ":8081",
		Handler:      router2(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server1.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		err := server2.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
