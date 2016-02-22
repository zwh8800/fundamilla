package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"path"
	"io"
)

const baseFolder = "/Users/zzz/go/src/code.meican.com/diffusion/MIS/archangel.git/public/mealReports/"

var watchList = map[string]string {
	"/Users/zzz/go/src/code.meican.com/diffusion/MIS/archangel.git/public/mealReports/main.bundle.css": "/Users/zzz/java-dev/meican-web/public/stylesheets/mealreports.spa",
	"/Users/zzz/go/src/code.meican.com/diffusion/MIS/archangel.git/public/mealReports/main.bundle.js": "/Users/zzz/java-dev/meican-web/public/javascripts/mealreports.spa",
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	if err := watcher.Watch(baseFolder); err != nil {
		log.Fatal(err)
	}

	log.Println("start watch")

	for {
		select {
		case ev := <-watcher.Event:
			log.Println("file ", ev.Name, ev.IsModify())
			dest, ok := watchList[ev.Name]
			if ok {
				srcFile, err := os.Open(ev.Name)
				if err != nil {
					log.Println("error occurs: ", err)
				}
				_, filename := path.Split(ev.Name)
				destFilename := path.Join(dest, filename)
				log.Println("Open file ", destFilename, "to write")
				destFile, err := os.OpenFile(destFilename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
				if err != nil {
					log.Println("error occurs: ", err)
				}
				n, err := io.Copy(destFile, srcFile)
				if err != nil {
					log.Println("error occurs: ", err)
				}
				log.Println(n, "bytes copyed")

				srcFile.Close()
				destFile.Close()
			}

		case err := <-watcher.Error:
			log.Println("error occurs: ", err)
		}
	}
}
