package mocks

import (
	"time"

	"greenlight.karronqiu.github.com/internal/data"
)

var MockMovie = &data.Movie{
	ID:        1,
	Title:     "Tiger",
	Genres:    []string{"action"},
	Year:      2002,
	Runtime:   100,
	Version:   1,
	CreatedAt: time.Date(2022, time.December, 31, 20, 00, 00, 0, time.Now().UTC().Location()),
}

type MockMovieModel struct {
}

func (m MockMovieModel) Insert(movie *data.Movie) error {
	movie.ID = MockMovie.ID
	return nil
}

func (m MockMovieModel) Get(id int64) (*data.Movie, error) {
	return MockMovie, nil
}

func (m MockMovieModel) Update(movie *data.Movie) error {
	return nil
}

func (m MockMovieModel) Delete(id int64) error {
	return nil
}

func (m MockMovieModel) GetAll(title string, genres []string, filters data.Filters) ([]*data.Movie, data.Metadata, error) {
	return []*data.Movie{MockMovie}, data.Metadata{}, nil
}
