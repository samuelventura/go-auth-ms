package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func rest(args Args) (func(), error) {
	dao := args.Get("dao").(Dao)
	endpoint := args.Get("endpoint").(string)
	gin.SetMode(gin.ReleaseMode) //remove debug warning
	router := gin.New()          //remove default logger
	router.Use(gin.Recovery())   //looks important
	rapi := router.Group("/api")
	header := func(r *http.Request, name string) (string, error) {
		key := http.CanonicalHeaderKey(name)
		values := r.Header[key]
		if len(values) != 1 {
			return "", fmt.Errorf("invalid header %s %v", key, values)
		}
		return values[0], nil
	}
	rapi.POST("/access", func(c *gin.Context) {
		app, err := header(c.Request, "Auth-App")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		dev, err := header(c.Request, "Auth-Dev")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		email, err := header(c.Request, "Auth-Email")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		_, err = dao.GetApp(app)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = dao.DisableCodes(app, email)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		dro := CodeDro{}
		dro.Code = codegen()
		dro.Created = time.Now()
		dro.Expires = time.Now().Add(5 * time.Minute)
		dro.App = app
		dro.Dev = dev
		dro.Email = email
		dao.AddCode(dro)
		c.JSON(200, gin.H{"code": dro.Code})
	})
	rapi.GET("/app/:name", func(c *gin.Context) {
		name := c.Param("name")
		dro, err := dao.GetApp(name)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": dro.Name})
	})
	rapi.GET("/app", func(c *gin.Context) {
		dros, err := dao.GetApps()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		names := make([]string, 0, len(*dros))
		for _, dro := range *dros {
			names = append(names, dro.Name)
		}
		c.JSON(200, gin.H{"names": names})
	})
	rapi.POST("/app/:name", func(c *gin.Context) {
		name := c.Param("name")
		err := dao.AddApp(name)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": name})
	})
	rapi.DELETE("/app/:name", func(c *gin.Context) {
		name := c.Param("name")
		dro, err := dao.GetApp(name)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = dao.DelApp(name)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": dro.Name})
	})
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, err
	}
	port := listen.Addr().(*net.TCPAddr).Port
	log.Println("port", port)
	server := &http.Server{
		Addr:    endpoint,
		Handler: router,
	}
	exit := make(chan interface{})
	go func() {
		err = server.Serve(listen)
		if err != nil {
			log.Println(err)
		}
		close(exit)
	}()
	closer := func() {
		ctx := context.Background()
		server.Shutdown(ctx)
		<-exit
	}
	return closer, nil
}
