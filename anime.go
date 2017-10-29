package kitsu

import (
	"encoding/json"
	"fmt"

	"github.com/Jeffail/gabs"
)

// Anime struct with all the anime data from kitsu
type Anime struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Attributes struct {
		TBA               string   `json:"tba"`
		AbbreviatedTitles []string `json:"abbreviatedTitles"`
		AverageRating     string   `json:"averageRating"`
		Status            string   `json:"status"`
		AgeRating         string   `json:"ageRating"`
		Subtype           string   `json:"subtype"`
		CanonicalTitle    string   `json:"canonicalTitle"`
		EpisodeLength     int      `json:"episodeLength"`
		CoverImage        struct {
			Original string `json:"original"`
			Tiny     string `json:"tiny"`
			Small    string `json:"small"`
			Large    string `json:"large"`
		} `json:"coverImage"`
		Slug   string `json:"slug"`
		Titles struct {
			EnJp string `json:"en_jp"`
			JaJp string `json:"ja_jp"`
		} `json:"titles"`
		AgeRatingGuide string `json:"ageRatingGuide"`
		StartDate      string `json:"startDate"`
		EpisodeCount   int    `json:"episodeCount"`
		FavoritesCount int    `json:"favoritesCount"`
		NSFW           bool   `json:"nsfw"`
		EndDate        string `json:"endDate"`
		RatingRank     int    `json:"ratingRank"`
		PosterImage    struct {
			Tiny     string `json:"tiny"`
			Small    string `json:"small"`
			Medium   string `json:"medium"`
			Large    string `json:"large"`
			Original string `json:"original"`
		} `json:"posterImage"`
		Synopsis       string `json:"synopsis"`
		ShowType       string `json:"showType"`
		UserCount      int    `json:"userCount"`
		PopularityRank int    `json:"popularityRank"`
	} `json:"attributes"`
	YoutubeVideoID string `json:"youtubeVideoId"`
}

// SearchAnime search for an anime on kitsu.io
// query being the anime to search for
// offset being the page offset
func SearchAnime(query string, offset int) (*Anime, error) {
	uri := fmt.Sprintf("%s/anime?filter[text]=%s&page[offset]=%v", baseURL, query, offset)

	parJSON, e := gabs.ParseJSON(get(uri))
	if e != nil {
		return nil, e
	}

	anime := parJSON.Path("data").Data().([]interface{})

	ani := new(Anime)

	resJSON, er := json.Marshal(anime[0]) // Right now I'm doing anime[0] because I have no idea how to handle it when it would return more than 1 result.
	if er != nil {
		return nil, er
	}
	err := json.Unmarshal(resJSON, &ani)
	if err != nil {
		return nil, err
	}
	return ani, nil
}