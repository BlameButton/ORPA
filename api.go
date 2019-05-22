package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type (
	BeatmapResponse struct {
		BeatmapsetID     string `json:"beatmapset_id"`
		BeatmapID        string `json:"beatmap_id"`
		Approved         string `json:"approved"`
		TotalLength      string `json:"total_length"`
		HitLength        string `json:"hit_length"`
		Version          string `json:"version"`
		FileMd5          string `json:"file_md5"`
		DiffSize         string `json:"diff_size"`
		DiffOverall      string `json:"diff_overall"`
		DiffApproach     string `json:"diff_approach"`
		DiffDrain        string `json:"diff_drain"`
		Mode             string `json:"mode"`
		SubmitDate       string `json:"submit_date"`
		ApprovedDate     string `json:"approved_date"`
		LastUpdate       string `json:"last_update"`
		Artist           string `json:"artist"`
		Title            string `json:"title"`
		Creator          string `json:"creator"`
		CreatorID        string `json:"creator_id"`
		Bpm              string `json:"bpm"`
		Source           string `json:"source"`
		Tags             string `json:"tags"`
		GenreID          string `json:"genre_id"`
		LanguageID       string `json:"language_id"`
		FavouriteCount   string `json:"favourite_count"`
		Rating           string `json:"rating"`
		Playcount        string `json:"playcount"`
		Passcount        string `json:"passcount"`
		MaxCombo         string `json:"max_combo"`
		DiffAim          string `json:"diff_aim"`
		DiffSpeed        string `json:"diff_speed"`
		Difficultyrating string `json:"difficultyrating"`
	}
)

func (b BeatmapResponse) String() string {
	bytes, e := json.Marshal(b)
	if e != nil {
		return ""
	}
	return string(bytes)
}

func GetBeatmap(hash string) []BeatmapResponse {
	url := getApiBaseUrl("get_beatmaps", getApiKey())
	url = url + "&h=" + hash
	response, e := getRequest(url)
	if e != nil {
		log.Fatal(e)
	}
	responseBytes, e := getResponseBody(response)
	if e != nil {
		log.Fatal(e)
	}
	beatmapResponse := make([]BeatmapResponse, 0)
	e = json.Unmarshal(responseBytes, &beatmapResponse)
	if e != nil {
		log.Fatal(e)
	}
	return beatmapResponse
}

func getResponseBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func getRequest(url string) (*http.Response, error) {
	return http.Get(url)
}

// Get the base URL for an endpoint, i.e. "get_beatmaps"
func getApiBaseUrl(endpoint string, token string) string {
	return fmt.Sprintf("https://osu.ppy.sh/api/%s?k=%s", endpoint, token)
}

func getApiKey() string {
	return os.Getenv("OSU_API_TOKEN")
}
