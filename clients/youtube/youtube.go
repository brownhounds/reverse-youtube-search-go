package youtube

import (
	"encoding/json"
	"go-reverse-youtube-search/env"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	BASE_URL = "https://www.googleapis.com/youtube/v3/search"
)

type YoutubeApiResponse struct {
	Items []struct {
		Snippet struct {
			ChannelTitle string `json:"channelTitle"`
		} `json:"snippet"`
	} `json:"items"`
	PageInfo struct {
		TotalResults int `json:"totalResults"`
	} `json:"pageInfo"`
}

type SearchResponse struct {
	ContentCreators string `json:"contentCreators"`
	SearchTerm      string `json:"searchTerm"`
	TotalResults    int    `json:"totalResults"`
}

func Search(search string) (error, *SearchResponse) {
	parsedURL, err := url.Parse(BASE_URL)
	if err != nil {
		log.Println(err.Error())
		return err, nil
	}

	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("maxResults", "50")
	params.Add("q", search)
	params.Add("type", "video")
	params.Add("key", os.Getenv(env.YOUTUBE_API_KEY))

	parsedURL.RawQuery = params.Encode()

	r, err := http.Get(parsedURL.String())
	if err != nil {
		log.Println(err.Error())
		return err, nil
	}

	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return err, nil
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return err, nil
	}

	var apiResponse YoutubeApiResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Println(err.Error())
		return err, nil
	}

	return nil, &SearchResponse{
		SearchTerm:      search,
		TotalResults:    apiResponse.PageInfo.TotalResults,
		ContentCreators: getUniqueContentCreators(&apiResponse),
	}
}

func getUniqueContentCreators(data *YoutubeApiResponse) string {
	var creators []string

	for _, item := range data.Items {
		title := item.Snippet.ChannelTitle

		if !sliceContainsValue(creators, title) {
			creators = append(creators, title)
		}
	}

	return strings.Join(creators, ", ")
}

func sliceContainsValue(s []string, value string) bool {
	for _, a := range s {
		if a == value {
			return true
		}
	}
	return false
}
