package main

import (
	"embed"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"io/fs"
)

var (
	//go:embed media/*
	mediaList embed.FS
)

type SurfTexture struct {
	W int32
	H int32
	T *sdl.Texture
	S *sdl.Surface
}

func GetImage(fileName string) *sdl.RWops {
	file, _ := mediaList.Open("media/" + fileName)
	return GetResource(file)
}

func GetResource(file fs.File) *sdl.RWops {
	stat, _ := file.Stat()
	size := stat.Size()
	buf := make([]byte, size)
	file.Read(buf)
	rwOps, _ := sdl.RWFromMem(buf)
	return rwOps
}

func LoadTexture(fileName string, sdlRenderer *sdl.Renderer) *sdl.Texture {
	return LoadSurfTexture(fileName, sdlRenderer).T
}

func LoadSurfTexture(fileName string, sdlRenderer *sdl.Renderer) *SurfTexture {
	surface, err := img.LoadRW(GetImage(fileName), true)
	if err != nil {
		println(err.Error())
	}
	defer surface.Free()
	txt, err := sdlRenderer.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	return &SurfTexture{T: txt, W: surface.W, H: surface.H}
}
