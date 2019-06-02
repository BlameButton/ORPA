package parser

import (
	"encoding/json"
	"fmt"
)

type (
	// OSU! mode (Standard, Taiko, Catch the Beat, Mania)
	Mode byte

	Common struct {
		Mode              Mode   `json:"mode"`
		GameVersion       uint32 `json:"game_version"`
		BeatMapHash       string `json:"beat_map_hash"`
		PlayerName        string `json:"player_name"`
		ReplayHash        string `json:"replay_hash"`
		Amount300s        uint16 `json:"amount_300_s"`
		Amount100s        uint16 `json:"amount_100_s"`
		Amount50s         uint16 `json:"amount_50_s"`
		AmountSpecial300s uint16 `json:"amount_special_300_s"`
		AmountSpecial100s uint16 `json:"amount_special_100_s"`
		AmountMisses      uint16 `json:"amount_misses"`
		TotalScore        uint32 `json:"total_score"`
		LongestCombo      uint16 `json:"longest_combo"`
		IsPerfect         bool   `json:"is_perfect"`
		Timestamp         uint64 `json:"timestamp"`
		ScoreID           uint64 `json:"score_id"`
	}

	Key string

	Action struct {
		Millis  int64   `json:"millis"`
		X       float32 `json:"x"`
		Y       float32 `json:"y"`
		Pressed []Key   `json:"pressed"`
	}

	Mod struct {
		Name       string  `json:"name"`
		Multiplier float32 `json:"multiplier"`
		bitOffset  uint32
	}

	HealthGraphPoint struct {
		Millis uint64  `json:"millis"`
		Life   float32 `json:"life"`
	}

	// export replay structure
	ReplayData struct {
		Common
		Mods        []Mod              `json:"mods"`
		HealthGraph []HealthGraphPoint `json:"health_graph"`
		ReplayData  []Action           `json:"replay_data"`
	}

	// rawReplay data structure
	rawReplay struct {
		Common
		Mods        uint32 `json:"mods"`
		HealthGraph string `json:"health_graph"`
		ReplayData  string `json:"replay_data"`
	}
)

func (m Mod) String() string {
	return fmt.Sprintf("%s (x%.2f)", m.Name, m.Multiplier)
}

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
		fallthrough
	case ModeStandard:
		return ModeStandardName // osu!standard
	case ModeTaiko:
		return ModeTaikoName // osu!taiko
	case ModeCTB:
		return ModeCTBName // osu!catch
	case ModeMania:
		return ModeManiaName // osu!mania
	}
}

func (m Mode) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}
