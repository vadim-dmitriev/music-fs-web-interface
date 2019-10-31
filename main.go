package main

import (
	"github.com/vadim-dmitriev/music-fs-web-interface/common"
	"github.com/vadim-dmitriev/music-fs-web-interface/dir"
)

// service это главный агрегатор
type service struct {
	cfg    *common.Config
	logger *common.Logger
	root   *dir.Node
}

func main() {
	var service = new(service)
	var err error

	if service.cfg, err = common.NewConfig(); err != nil {
		panic(err)
	}

	if service.logger, err = common.NewLogger(service.cfg); err != nil {
		panic(err)
	}

	if service.root, err = dir.NewTree(service.cfg.MusicDir); err != nil {
		panic(err)
	}

}
