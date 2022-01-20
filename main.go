package main

import (
	"fmt"
	"net/http"

	"gin-example.com/v0/pkg/setting"
	"gin-example.com/v0/routers"
)

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// r.Run("0.0.0.0:80") 中包含 s.ListenAndServe()
	s.ListenAndServe()
}
