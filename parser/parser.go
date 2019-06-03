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
