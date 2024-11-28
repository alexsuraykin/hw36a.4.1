package main

import (
	"log"
	"net/http"
	"time"

	"hw36a.4.1/internal/api"
	"hw36a.4.1/internal/conf"
	"hw36a.4.1/internal/postgres"
	"hw36a.4.1/internal/rss"
)

const (
	pathRSSConfig = "./cmd/server/config.json"
	pathBDConfig  = "./cmd/server/BD.json"
)

type server struct {
	api *api.API
}

func main() {
	// Uploading RSS configuration
	rssConf, err := conf.NewRSS(pathRSSConfig)
	if err != nil {
		log.Fatal("ошибка загрузки конфига RSS:", err)
	}

	// Uploading DB configuration
	bdConf, err := conf.NewBD(pathBDConfig)
	if err != nil {
		log.Fatal("ошибка загрузки конфига BD:", err)
	}

	// DB initialization
	data, err := postgres.New(bdConf)
	if err != nil {
		log.Fatal("ошибка инициализации БД:,", err)
	}

	t := time.Minute * time.Duration(rssConf.RequestPeriod)

	for _, url := range rssConf.UrlsRSS {
		pipe := cache(readRSS(url, t))
		go reporter(pipe, data)
	}

	var srv server

	// Creating API handlers
	srv.api = api.New(*data)

	// Run server
	_ = http.ListenAndServe(":80", srv.api.Router())

}

// Reading news
func readRSS(url string, timer time.Duration) chan rss.Post {
	out := make(chan rss.Post)

	go func() {
		defer close(out)
		for {
			arPosts, err := rss.GetRSS(url)
			if err != nil {
				log.Println(err)
				return
			}
			for _, post := range arPosts {
				out <- post
			}
			time.Sleep(timer)
		}
	}()
	return out
}

// One attempt to handle news
func cache(input <-chan rss.Post) chan rss.Post {
	output := make(chan rss.Post)
	go func() {
		defer close(output)

		cacheMap := make(map[string]bool)

		for {
			select {
			case value, ok := <-input:
				if !ok {
					return
				}
				if cacheMap[value.ID] {
					continue
				} else {
					cacheMap[value.ID] = true
					output <- value
				}
			}
		}
	}()
	return output
}

func reporter(input <-chan rss.Post, dataBase *postgres.Store) {

	go func() {
		for value := range input {
			// Writing news in DB
			err := dataBase.AddPost(value)
			if err != nil {
				log.Println("ошибка добавления новости в БД:", err)
			}
			log.Printf("Добавлена новость id: '%s' title: '%s'\n", value.ID, value.Title)
		}
	}()
}
