package main

import (
	"log"
	"os"
)

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func CloseFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func First[T, U, V any](val T, _ U, _ V) T {
	return val
}
