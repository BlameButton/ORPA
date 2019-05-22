package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
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
		replay := getReplay(filepath)
		beatmaps := GetBeatmap(replay.BeatMapHash)
		if len(beatmaps) < 1 {
			continue
		}
		beatmap := beatmaps[0]
		// Append data to table writer
		table.Append([]string{
			fmt.Sprintf("%d", k),
			fmt.Sprintf("%s - %s", beatmap.Artist, beatmap.Title),
			beatmap.Version,
			humanize.Comma(int64(replay.TotalScore)),
			fmt.Sprintf("%b", replay.Mods),
		})
	}
	// Render the table to console
	table.Render()
	// Print how many replays were succesfully processed.
	fmt.Printf("\nSuccesfully processed %d replays.\n", table.NumLines())
}

func getReplayFiles(directory string) []string {
	replays, e := ioutil.ReadDir(directory)
	LogError(e)
	files := make([]string, 0)
	for _, v := range replays {
		name := v.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		files = append(files, v.Name())
	}
	return files
}

func getReplay(replayFile string) Replay {
	reader, e := getReplayReader(replayFile)
	LogError(e)
	// Make a new replay and read it data into the replay
	replay := Replay{}
	readReplay(&replay, reader)
	return replay
}

func getReplayReader(replayFile string) (*bufio.Reader, error) {
	replayBytes, e := ioutil.ReadFile(replayFile)
	if e != nil {
		return nil, e
	}
	buffer := bytes.NewReader(replayBytes)
	reader := bufio.NewReader(buffer)
	return reader, e
}

func readReplay(replay *Replay, reader *bufio.Reader) {
	mode, e := reader.ReadByte()
	LogError(e)
	replay.Mode = Mode(mode)

	version, e := ReadInteger(reader)
	LogError(e)
	replay.GameVersion = version

	beatmapHash, e := ReadString(reader)
	LogError(e)
	replay.BeatMapHash = beatmapHash

	playerName, e := ReadString(reader)
	LogError(e)
	replay.PlayerName = playerName

	replayHash, e := ReadString(reader)
	LogError(e)
	replay.ReplayHash = replayHash

	amount300s, e := ReadShort(reader)
	LogError(e)
	replay.Amount300s = amount300s

	amount100s, e := ReadShort(reader)
	LogError(e)
	replay.Amount100s = amount100s

	amount50s, e := ReadShort(reader)
	LogError(e)
	replay.Amount50s = amount50s

	amountSpecial300s, e := ReadShort(reader)
	LogError(e)
	replay.AmountSpecial300s = amountSpecial300s

	amountSpecial100s, e := ReadShort(reader)
	LogError(e)
	replay.AmountSpecial100s = amountSpecial100s

	misses, e := ReadShort(reader)
	LogError(e)
	replay.AmountMisses = misses

	score, e := ReadInteger(reader)
	LogError(e)
	replay.TotalScore = score

	combo, e := ReadShort(reader)
	LogError(e)
	replay.LongestCombo = combo

	isPerfect, e := ReadBoolean(reader)
	LogError(e)
	replay.IsPerfect = isPerfect

	mods, e := ReadInteger(reader)
	LogError(e)
	replay.Mods = mods

	healthGraph, e := ReadString(reader)
	LogError(e)
	replay.HealthGraph = healthGraph

	timestamp, e := ReadLong(reader)
	LogError(e)
	replay.Timestamp = timestamp
}
