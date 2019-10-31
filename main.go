package main

import (
	"github.com/vadim-dmitriev/music-fs-web-interface/common"
	"github.com/vadim-dmitriev/music-fs-web-interface/dir"
)

func main() {
	cfg, err := common.NewConfig()
	if err != nil {
		panic(err)
	}

	_, err = dir.NewTree(cfg.MusicDir)
	if err != nil {
		panic(err)
	}

}
