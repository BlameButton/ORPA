package parser

import (
	"bufio"
	"bytes"
	"io/ioutil"
)

func GetFileReplay(file string) (*ReplayData, error) {
	reader, e := getReplayReader(file)
	if e != nil {
		return nil, e
	}
	// Make a new replay and read it data into the replay
	replay, e := parseReplay(reader)
	if e != nil {
		return nil, e
	}
	return replay, nil
}

func getReplayReader(file string) (*bufio.Reader, error) {
	replayBytes, e := ioutil.ReadFile(file)
	if e != nil {
		return nil, e
	}
	buffer := bytes.NewReader(replayBytes)
	reader := bufio.NewReader(buffer)
	return reader, e
}

func parseReplay(reader *bufio.Reader) (*ReplayData, error) {
	raw := &rawReplay{}
	mode, e := reader.ReadByte()
	if e != nil {
		return nil, e
	}
	raw.Mode = Mode(mode)
	version, e := ReadInteger(reader)
	if e != nil {
		return nil, e
	}
	raw.GameVersion = version
	beatmapHash, e := ReadString(reader)
	if e != nil {
		return nil, e
	}
	raw.BeatMapHash = beatmapHash
	playerName, e := ReadString(reader)
	if e != nil {
		return nil, e
	}
	raw.PlayerName = playerName
	replayHash, e := ReadString(reader)
	if e != nil {
		return nil, e
	}
	raw.ReplayHash = replayHash
	amount300s, e := ReadShort(reader)
	if e != nil {
		return nil, e
	}
	raw.Amount300s = amount300s
	amount100s, e := ReadShort(reader)
	if e != nil {
		return nil, e
	}
	raw.Amount100s = amount100s
	amount50s, e := ReadShort(reader)
	if e != nil {
		return nil, e
	}
	raw.Amount50s = amount50s
	amountSpecial300s, e := ReadShort(reader)
	if e != nil {
		return nil, e
	}
	raw.AmountSpecial300s = amountSpecial300s
	amountSpecial100s, e := ReadShort(reader)
	if e != nil {
		return nil, e
	}
	raw.AmountSpecial100s = amountSpecial100s
	misses, e := ReadShort(reader)
	if e != nil {
		return nil, e
	}
	raw.AmountMisses = misses
	score, e := ReadInteger(reader)
	if e != nil {
		return nil, e
	}
	raw.TotalScore = score
	combo, e := ReadShort(reader)
	if e != nil {
		return nil, e
	}
	raw.LongestCombo = combo
	isPerfect, e := ReadBoolean(reader)
	if e != nil {
		return nil, e
	}
	raw.IsPerfect = isPerfect
	mods, e := ReadInteger(reader)
	if e != nil {
		return nil, e
	}
	raw.Mods = mods
	healthGraph, e := ReadString(reader)
	if e != nil {
		return nil, e
	}
	raw.HealthGraph = healthGraph
	timestamp, e := ReadLong(reader)
	if e != nil {
		return nil, e
	}
	raw.Timestamp = timestamp
	dataLength, e := ReadInteger(reader)
	if e != nil {
		return nil, e
	}
	data, e := ReadLZMA(reader, dataLength)
	if e != nil {
		return nil, e
	}
	raw.ReplayData = data
	scoreId, e := ReadLong(reader)
	if e != nil {
		return nil, e
	}
	raw.ScoreID = scoreId
	// Now format the data into a proper replay to be exported
	replay := parseRawReplay(raw)
	return replay, nil
}

var availableMods = []Mod{
	{Name: "NoFail", Multiplier: 0.5, bitOffset: 1},
	{Name: "Easy", Multiplier: 0.5, bitOffset: 2},
	{Name: "TouchDevice", Multiplier: -1, bitOffset: 4},
	{Name: "Hidden", Multiplier: 1.06, bitOffset: 8},
	{Name: "Hard Rock", Multiplier: 1.06, bitOffset: 16},
	{Name: "Sudden Death", Multiplier: 1.0, bitOffset: 32},
	{Name: "Double Time", Multiplier: 1.12, bitOffset: 64},
	{Name: "Relax", Multiplier: 0.0, bitOffset: 128},
	{Name: "Half Time", Multiplier: 0.30, bitOffset: 256},
	{Name: "Nightcore", Multiplier: 1.12, bitOffset: 512},
	{Name: "Flashlight", Multiplier: 1.12, bitOffset: 1024},
	{Name: "Auto", Multiplier: 1.0, bitOffset: 2048},
	{Name: "Spun Out", Multiplier: 0.90, bitOffset: 4096},
	{Name: "Auto Pilot", Multiplier: 0.0, bitOffset: 8192},
	{Name: "Perfect", Multiplier: 1.0, bitOffset: 16384},
}

func parseRawReplay(replay *rawReplay) *ReplayData {
	mods := make([]Mod, 0)
	for _, mod := range availableMods {
		if replay.Mods&mod.bitOffset == mod.bitOffset {
			mods = append(mods, mod)
		}
	}
	return &ReplayData{
		Common:      replay.Common,
		Mods:        mods,
		ReplayData:  []Action{},
		HealthGraph: []HealthGraphPoint{},
	}
}
