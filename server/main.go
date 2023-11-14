package main

import (
	"fmt"
	"gin_admin/pkg/setting"
	"gin_admin/routers"
	_ "gin_admin/routers"
	"net/http"
)

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routers.Router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
