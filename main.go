package main

import (
	"net/http"
	"runtime"
	"time"
	"github.com/zhengjianwen/hr_service/model"
	"golang.org/x/sync/errgroup"
	"github.com/zhengjianwen/hr_service/conf"
	"github.com/gin-gonic/gin"

	"github.com/zhengjianwen/hr_service/views"
	"github.com/zhengjianwen/utils/log"
)

var (
	g errgroup.Group
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

}

func kingadmin() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	views.ConfRouter(r)
	return r
}

func main()  {
	model.InitSQL()

	server01 := &http.Server{
		Addr:         conf.Listen,
		Handler:      kingadmin(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		log.Debugf("[sys] KingAdmin: %s",conf.Listen)
		return server01.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
