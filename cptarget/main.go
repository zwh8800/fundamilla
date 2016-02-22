package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"os/signal"

	"github.com/howeyc/fsnotify"
)

type Config struct {
	BaseFolder string            `json:"baseFolder"`
	WatchList  map[string]string `json:"watchList"`
}

func readConfig(filename string) (*Config, error) {
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func handleFile(srcFilename string, destFilename string) {
	log.Println("Open file ", srcFilename, "to read")
	srcFile, err := os.Open(srcFilename)
	if err != nil {
		log.Println("error occurs:", err)
		return
	}
	defer srcFile.Close()
	log.Println("Open file ", destFilename, "to write")
	destFile, err := os.OpenFile(destFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("error occurs:", err)
		return
	}
	defer destFile.Close()
	n, err := io.Copy(destFile, srcFile)
	if err != nil {
		log.Println("error occurs:", err)
		return
	}
	log.Println(n, "bytes copyed")
}

func main() {
	configFile := flag.String("config", "watchlist.json", "specify a watchlist")
	flag.Parse()

	config, err := readConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	if err := watcher.Watch(config.BaseFolder); err != nil {
		log.Fatal(err)
	}

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Kill, os.Interrupt)

	log.Println("start watch")

	for {
		select {
		case ev := <-watcher.Event:
			log.Println(ev.String())
			_, srcName := path.Split(ev.Name)
			destFilename, ok := config.WatchList[srcName]
			if ok {
				if !path.IsAbs(destFilename) {
					destFilename = path.Join(config.BaseFolder, destFilename)
				}
				handleFile(ev.Name, destFilename)
			}
		case err := <-watcher.Error:
			log.Println("error occurs:", err)
		case sig := <-sigChannel:
			log.Println("signal caught:", sig)
			log.Println("graceful exit")
			return
		}
	}
}
