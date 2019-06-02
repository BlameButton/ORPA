package main

import (
	"fmt"
	"github.com/blamebutton/orpa/api"
	"github.com/blamebutton/orpa/parser"
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const ReplayDir = "replays"

func main() {
	// Get a list of all replay files and print the amount that will be processed
	replays := getReplayFiles(ReplayDir)
	fmt.Printf("Processing %d replays...\n\n", len(replays))
	// Set up the table writer
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Beatmap", "Difficulty", "Score", "Mods"})
	table.SetAutoWrapText(false)
	// Loop over all the replays
	for k, v := range replays {
		filepath := path.Join(ReplayDir, v)
		replay, e := parser.GetFileReplay(filepath)
		if e != nil {
			continue
		}
		beatmaps, e := api.GetBeatmap(getApiKey(), replay.BeatMapHash)
		if e != nil || len(beatmaps) < 1 {
			continue
		}
		beatmap := beatmaps[0]
		// Append data to table writer
		mods := "None"
		if len(replay.Mods) > 0 {
			mods = strings.Trim(fmt.Sprintf("%+v", replay.Mods), "[]")
		}
		table.Append([]string{
			fmt.Sprintf("%d", k),
			fmt.Sprintf("%s - %s", beatmap.Artist, beatmap.Title),
			beatmap.Version,
			humanize.Comma(int64(replay.TotalScore)),
			mods,
		})
	}
	// Render the table to console
	table.Render()
	// Print how many replays were succesfully processed.
	fmt.Printf("\nSuccesfully processed %d replays.\n", table.NumLines())
}

// Get the list of files in a directory
func getReplayFiles(directory string) []string {
	replays, e := ioutil.ReadDir(directory)
	if e != nil {
		log.Fatal(e)
	}
	files := make([]string, 0)
	for _, v := range replays {
		name := v.Name()
		if !strings.HasSuffix(name, ".osr") {
			continue
		}
		if strings.HasPrefix(name, ".") {
			continue
		}
		files = append(files, v.Name())
	}
	return files
}

func getApiKey() string {
	return os.Getenv("OSU_API_TOKEN")
}
