package main

import (
	"io"
	"log"
	"os"

	"github.com/howeyc/fsnotify"
)

const baseFolder = "/Users/zzz/go/src/code.meican.com/diffusion/MIS/archangel.git/public/mealReports/"

var watchList = map[string]string{
	"/Users/zzz/go/src/code.meican.com/diffusion/MIS/archangel.git/public/mealReports/main.bundle.css": "/Users/zzz/java-dev/meican-web/public/stylesheets/mealreports.spa/corp_meal_reports.spa.css",
	"/Users/zzz/go/src/code.meican.com/diffusion/MIS/archangel.git/public/mealReports/main.bundle.js":  "/Users/zzz/java-dev/meican-web/public/javascripts/mealreports.spa/corp_meal_reports.spa.js",
}

func handleFile(srcFilename string, destFilename string) {
	srcFile, err := os.Open(srcFilename)
	if err != nil {
		log.Println("error occurs: ", err)
	}
	defer srcFile.Close()
	log.Println("Open file ", destFilename, "to write")
	destFile, err := os.OpenFile(destFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println("error occurs: ", err)
	}
	defer destFile.Close()
	n, err := io.Copy(destFile, srcFile)
	if err != nil {
		log.Println("error occurs: ", err)
	}
	log.Println(n, "bytes copyed")
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
			destFilename, ok := watchList[ev.Name]
			if ok {
				handleFile(ev.Name, destFilename)
			}
		case err := <-watcher.Error:
			log.Println("error occurs: ", err)
		}
	}
}
