package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"golang.org/x/image/draw"
)

func scaleFile(srcPath, destPath string, size image.Rectangle) error {
	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer file.Close()

	src, err := png.Decode(file)
	if err != nil {
		return fmt.Errorf("%q - can't decode as PNG - %w", srcPath, err)
	}

	dest := image.NewRGBA(size)
	bounds := src.Bounds()
	draw.NearestNeighbor.Scale(dest, size, src, bounds, draw.Over, nil)

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	if err := png.Encode(out, dest); err != nil {
		out.Close()
		os.Remove(destPath)
		return err
	}

	return nil
}

func scaleDir(srcDir, destDir string, size image.Rectangle) error {
	var mu sync.Mutex
	var errs error

	pool := make(chan bool, runtime.GOMAXPROCS(0))
	var wg sync.WaitGroup

	srcFiles, err := filepath.Glob(filepath.Join(srcDir, "*.png"))
	if err != nil {
		return err
	}

	wg.Add(len(srcFiles))
	for _, src := range srcFiles {
		go func(srcFile string) {
			defer wg.Done()

			pool <- true
			defer func() { <-pool }()

			destFile := path.Join(destDir, filepath.Base(srcFile))
			if err := scaleFile(srcFile, destFile, size); err != nil {
				mu.Lock()
				errs = errors.Join(errs, err)
				mu.Unlock()
			}
		}(src)
	}

	wg.Wait()
	return errs
}

func main() {
	size := image.Rect(0, 0, 400, 400)
	if err := scaleDir("/tmp/src", "/tmp/dest", size); err != nil {
		log.Fatal(err)
	}
}
