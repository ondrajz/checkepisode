package main

import (
	"time"

	"github.com/garfunkel/go-tvdb"
)

type ShowStore interface {
	Update(s *Show) error
	Get(id uint64) (Show, error)
	GetByName(name string) (Show, error)
	Remove(id uint64) error
}

type Show struct {
	ID         uint64
	TvdbID     uint64
	SeriesID   string
	SeriesName string
	ImdbID     string
	Overview   string
	Genre      []string
	Actors     []string
	Status     string
	Network    string
	NetworkID  string
	FirstAired time.Time
	Banner     string
	Fanart     string
	Poster     string
}

func SearchShows(name string) ([]Show, error) {
	list, err := tvdb.SearchSeries(name, 5)
	if err != nil {
		return nil, err
	}
	var shows []Show
	for _, s := range list.Series {
		firstAired, _ := time.Parse("2006-01-02", s.FirstAired)
		shows = append(shows, Show{
			TvdbID:     s.ID,
			SeriesID:   s.SeriesID,
			SeriesName: s.SeriesName,
			ImdbID:     s.ImdbID,
			Overview:   s.Overview,
			Genre:      s.Genre,
			Actors:     s.Actors,
			Status:     s.Status,
			FirstAired: firstAired,
			Network:    s.Network,
			NetworkID:  s.NetworkID,
			Banner:     s.Banner,
			Fanart:     s.Fanart,
			Poster:     s.Poster,
		})
	}
	return shows, nil
}
