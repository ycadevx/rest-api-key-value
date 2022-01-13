package store

import (
	"encoding/json"
	"github/deliveryhero/source/log"
	. "github/deliveryhero/source/models"

	"io/ioutil"

	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	jsonFileName = "data.json"
)

func CheckData() {
	files, err := ioutil.ReadDir(EXPORT_FILE_PATH)
	if err != nil {
		log.Info.Println(err)
		return
	}

	for _, file := range files {
		path := filepath.Join(EXPORT_FILE_PATH, file.Name())
		f, err := os.ReadFile(path)
		if err != nil {
			log.Fatal.Println(err)
		}

		err = json.Unmarshal(f, &InMemoryDB)
		if err != nil {
			log.Fatal.Println(err)
		}
	}
}

func RemoveFileInDirectory() {
	files, err := ioutil.ReadDir(EXPORT_FILE_PATH)
	if err != nil {
		log.Info.Println(err)
		return
	}

	for _, file := range files {
		path := filepath.Join(EXPORT_FILE_PATH, file.Name())
		err := os.Remove(path)
		if err != nil {
			log.Error.Println(err)
		}
	}
}

func SaveData(ticker *time.Ticker, quit chan struct{}) {
	for {
		select {
		case <-ticker.C:
			if _, err := os.Stat(EXPORT_FILE_PATH); os.IsNotExist(err) {
				err := os.Mkdir(EXPORT_FILE_PATH, os.ModePerm)
				if err != nil {
					log.Info.Println(err)
				}
			} else {
				RemoveFileInDirectory()
			}

			dataBytes, err := json.Marshal(InMemoryDB)
			if err != nil {
				log.Fatal.Println(err)
			}

			fileName := filepath.Join(EXPORT_FILE_PATH, jsonFileName)
			err = ioutil.WriteFile(fileName, dataBytes, 0777)

			if err != nil {
				log.Fatal.Println(err)
			} else {
				log.Info.Printf("Data control and recording done -  >: %s", fileName)
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func TimeControl() {
	duration, err := strconv.Atoi(RECORD_TIME)
	if err != nil {
		duration = 2
	}
	ticker := time.NewTicker(time.Duration(duration) * time.Second)
	quit := make(chan struct{})
	go SaveData(ticker, quit)
}
