package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/gcfg.v1"
)

func config() (string, string, string, string, int64) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	cfgpath := dir + "\\godville.cfg"

	file, err := ioutil.ReadFile(cfgpath) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	cfgStr := string(file)

	cfg := struct {
		Gv struct {
			Timeout string
			Godname string
			Apikey  string
			Bot     string
			Userid  int64
		}
	}{}
	err = gcfg.ReadStringInto(&cfg, cfgStr)
	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}

	return cfg.Gv.Timeout, cfg.Gv.Godname, cfg.Gv.Apikey, cfg.Gv.Bot, cfg.Gv.Userid

}
