package postgres

import (
	"reflect"
	"testing"
	"time"

	"hw36a.4.1/internal/conf"
	"hw36a.4.1/internal/rss"
)

var confBD = conf.BDConfig{
	Name:     "rss",
	Port:     5432,
	Table:    "posts",
	User:     "sandbox",
	Password: "sandbox",
}

var Serv *Store

func TestNew(t *testing.T) {
	var err error
	Serv, err = New(&confBD)
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
	if Serv == nil {
		t.Errorf("New() got = nil, want not nil")
	}
}

func TestStore_AddPost(t *testing.T) {
	type args struct {
		p rss.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Add Good Post",
			args: args{
				p: rss.Post{
					ID:      "id001",
					Title:   "Test Title",
					Content: "Test Content",
					PubTime: time.Now().Unix(),
					Link:    "https://test.ru/id001",
				},
			},
			wantErr: false,
		},
		{
			name: "Add Good Post 2",
			args: args{
				p: rss.Post{
					ID:      "id002",
					Title:   "Test Title",
					Content: "Test Content",
					PubTime: time.Now().Unix(),
					Link:    "https://test.ru/id002",
				},
			},
			wantErr: false,
		},
		{
			name: "Add Bad Post",
			args: args{
				p: rss.Post{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Serv.AddPost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_Posts(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []rss.Post
		wantErr bool
	}{
		{
			name: "Get 1 Post",
			args: args{
				n: 1,
			},
			want: []rss.Post{{
				ID:      "id001",
				Title:   "Test Title",
				Content: "Test Content",
				PubTime: time.Now().Unix(),
				Link:    "https://test.ru/id001",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Serv.Posts(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Posts() got = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}
