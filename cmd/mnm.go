package main

import (
	"encoding/json"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/ski7777/MusicNotesManager/internal/config"
	"github.com/ski7777/MusicNotesManager/internal/database"
	"github.com/ski7777/MusicNotesManager/internal/filestore"
	"github.com/ski7777/MusicNotesManager/internal/mnm"
	"github.com/ski7777/MusicNotesManager/internal/web"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	parser := argparse.NewParser("mnm", "Music Notes Manager")
	configpath := parser.String("c", "config", &argparse.Options{
		Required: false,
		Help:     "path to configuration file",
		Default:  "mnm.config.json",
	})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	conf := &config.Config{}
	conffile, err := os.Open(*configpath)
	if err != nil {
		log.Println(err)
	}
	rawconf, err := ioutil.ReadAll(conffile)
	if err != nil {
		log.Println(err)
	}
	if err = json.Unmarshal(rawconf, conf); err != nil {
		log.Fatal(err)
	}
	db, dbts, err := database.GetDB(conf.Database.Driver, conf.Database.DSN)
	if err != nil {
		log.Println(err)
	}
	fs, err := filestore.NewFileStore(conf.FileStore)
	if err != nil {
		log.Println(err)
	}
	m := mnm.NewMusicNotesManager(db, dbts, fs)
	w, err := web.NewWeb(m)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := w.Start(conf.WebAddr); err != nil {
			log.Fatal(err)
		}
	}()
	if err = m.Run(); err != nil {
		log.Fatal(err)
	}
}
