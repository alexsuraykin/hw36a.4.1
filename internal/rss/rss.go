package rss

import (
	"github.com/mmcdole/gofeed"
)

type Post struct {
	ID      string `json:"ID,omitempty"`
	Title   string `json:"Title,omitempty"`
	Content string `json:"Content,omitempty"`
	PubTime int64  `json:"PubTime,omitempty"`
	Link    string `json:"Link,omitempty"`
}

func GetRSS(url string) ([]Post, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	var posts []Post
	if len(feed.Items) == 0 {
		return posts, nil
	}

	for _, item := range feed.Items {
		posts = append(posts, Post{
			ID:      item.GUID,
			Title:   item.Title,
			Content: item.Description,
			PubTime: item.PublishedParsed.Unix(),
			Link:    item.Link,
		})
	}
	return posts, nil
}
