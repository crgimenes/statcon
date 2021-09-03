package main

import (
	"log"
	"os"

	"github.com/crgimenes/statcon/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	if cfg.Output != "-" {
		cfg.OutputDescriptor, err = os.OpenFile(cfg.Output, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		defer cfg.OutputDescriptor.Close()
	}
	out := cfg.OutputDescriptor

	_, err = out.Write([]byte("test output descriptor"))
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

}
