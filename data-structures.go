package main

import "encoding/json"

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
