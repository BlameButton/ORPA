package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

// OSU! mode (Standard, Taiko, Catch the Beat, Mania)
type Mode byte

const (
	// Hex values for all modes
	ModeStandard Mode = 0x0
	ModeTaiko    Mode = 0x1
	ModeCTB      Mode = 0x2
	ModeMania    Mode = 0x3
	// Names for all modes
	ModeStandardName = "osu!standard"
	ModeTaikoName    = "osu!taiko"
	ModeCTBName      = "osu!catch"
	ModeManiaName    = "osu!mania"
)

func (m Mode) String() string {
	switch m {
	default:
	case ModeStandard:
		return ModeStandardName
	case ModeTaiko:
		return ModeTaikoName
	case ModeCTB:
		return ModeCTBName
	case ModeMania:
		return ModeManiaName
	}
	return ModeStandardName
}

func (m Mode) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// Replay data structure
type Replay struct {
	Mode              Mode
	GameVersion       uint32
	BeatMapHash       string
	PlayerName        string
	ReplayHash        string
	Amount300s        uint16
	Amount100s        uint16
	Amount50s         uint16
	AmountSpecial300s uint16
	AmountSpecial100s uint16
	AmountMisses      uint16
	TotalScore        uint32
	LongestCombo      uint16
	IsPerfect         bool
	Mods              uint32
	HealthGraph       string
	Timestamp         uint64
	ReplayData        []byte
	ScoreID           uint64
}

func main() {
	const replayDir = "replays"
	replays, e := ioutil.ReadDir(replayDir)
	LogError(e)

	fmt.Printf("Found %d replays\n\n", len(replays))
	for _, replayFileInfo := range replays {
		replayFile := replayFileInfo.Name()
		if strings.HasPrefix(replayFile, ".") {
			continue
		}
		replayPath := path.Join(replayDir, replayFile)
		replay := getReplay(replayPath)
		beatmap := GetBeatmap(replay.BeatMapHash)
		fmt.Printf("%s\n", beatmap[0].Artist)
	}
}

func getReplay(replayFile string) Replay {
	replayBytes, e := ioutil.ReadFile(replayFile)
	LogError(e)

	// Create reader
	buffer := bytes.NewReader(replayBytes)
	reader := bufio.NewReader(buffer)

	mode, e := reader.ReadByte()
	LogError(e)

	version, e := ReadInteger(reader)
	LogError(e)

	beatmapHash, e := ReadString(reader)
	LogError(e)
	playerName, e := ReadString(reader)
	LogError(e)
	replayHash, e := ReadString(reader)
	LogError(e)
	amount300s, e := ReadShort(reader)
	LogError(e)
	amount100s, e := ReadShort(reader)
	LogError(e)
	amount50s, e := ReadShort(reader)
	LogError(e)
	amountSpecial300s, e := ReadShort(reader)
	LogError(e)
	amountSpecial100s, e := ReadShort(reader)
	LogError(e)
	misses, e := ReadShort(reader)
	LogError(e)
	score, e := ReadInteger(reader)
	LogError(e)
	combo, e := ReadShort(reader)
	LogError(e)
	isPerfect, e := ReadBoolean(reader)
	LogError(e)
	mods, e := ReadInteger(reader)
	LogError(e)
	healthGraph, e := ReadString(reader)
	LogError(e)
	timestamp, e := ReadLong(reader)
	LogError(e)

	replay := Replay{
		Mode:              Mode(mode),
		GameVersion:       version,
		BeatMapHash:       beatmapHash,
		PlayerName:        playerName,
		ReplayHash:        replayHash,
		Amount300s:        amount300s,
		Amount100s:        amount100s,
		Amount50s:         amount50s,
		AmountSpecial300s: amountSpecial300s,
		AmountSpecial100s: amountSpecial100s,
		AmountMisses:      misses,
		TotalScore:        score,
		LongestCombo:      combo,
		IsPerfect:         isPerfect,
		Mods:              mods,
		HealthGraph:       healthGraph,
		Timestamp:         timestamp,
	}

	LogError(e)

	return replay
}
