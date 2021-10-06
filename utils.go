package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func codegen() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	var b strings.Builder
	for i := 0; i < 6; i++ {
		d := r.Intn(10)
		b.WriteString(fmt.Sprint(d))
	}
	return b.String()
}

func getenv(name string, defval string) string {
	value := os.Getenv(name)
	if len(strings.TrimSpace(value)) > 0 {
		log.Println(name, value)
		return value
	}
	log.Println(name, defval)
	return defval
}

func withext(ext string) (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(exe)
	base := filepath.Base(exe)
	file := base + "." + ext
	return filepath.Join(dir, file), nil
}
