package main

import (
	"bufio"
	"bytes"
	"ekyu.moe/leb128"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path"
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
	fileInfos, e := ioutil.ReadDir(replayDir)
	logError(e)

	replayName := fileInfos[0].Name()
	replayBytes, e := ioutil.ReadFile(path.Join(replayDir, replayName))
	fmt.Println(replayName)
	logError(e)

	// Create reader
	buffer := bytes.NewReader(replayBytes)
	reader := bufio.NewReader(buffer)

	mode, e := reader.ReadByte()
	logError(e)

	version, e := readInteger(reader)
	logError(e)

	beatmapHash, e := readString(reader)
	logError(e)
	playerName, e := readString(reader)
	logError(e)
	replayHash, e := readString(reader)
	logError(e)
	amount300s, e := readShort(reader)
	logError(e)
	amount100s, e := readShort(reader)
	logError(e)
	amount50s, e := readShort(reader)
	logError(e)
	amountSpecial300s, e := readShort(reader)
	logError(e)
	amountSpecial100s, e := readShort(reader)
	logError(e)
	misses, e := readShort(reader)
	logError(e)
	score, e := readInteger(reader)
	logError(e)
	combo, e := readShort(reader)
	logError(e)
	isPerfect, e := readBoolean(reader)
	logError(e)
	mods, e := readInteger(reader)
	logError(e)
	healthGraph, e := readString(reader)
	logError(e)
	timestamp, e := readLong(reader)
	logError(e)

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

	replayJson, e := json.Marshal(replay)
	logError(e)

	fmt.Println(string(replayJson))
}

func readLong(buffer io.Reader) (uint64, error) {
	integer := uint64(0)
	e := binary.Read(buffer, binary.LittleEndian, &integer)
	return integer, e
}

func readInteger(buffer io.Reader) (uint32, error) {
	integer := uint32(0)
	e := binary.Read(buffer, binary.LittleEndian, &integer)
	return integer, e
}
func readShort(buffer io.Reader) (uint16, error) {
	short := uint16(0)
	e := binary.Read(buffer, binary.LittleEndian, &short)
	return short, e
}

func readBoolean(buffer *bufio.Reader) (bool, error) {
	next, e := buffer.ReadByte()
	if e != nil {
		return false, e
	}
	return next == 0x1, nil
}

// Get the value of a variable length integer
func readUleb(reader *bufio.Reader) (uint64, error) {
	next, e := reader.Peek(10)
	if e != nil {
		return 0, e
	}
	value, length := leb128.DecodeUleb128(next)
	_, e = reader.Discard(int(length))
	if e != nil {
		return 0, e
	}
	return value, nil
}

// Get the string value from bufio.Reader
func readString(reader *bufio.Reader) (string, error) {
	b, e := reader.ReadByte()
	if e != nil {
		return "", e
	}
	if b == 0x00 {
		return "", nil
	}
	if b != 0x0b {
		return "", errors.New("could not find string")
	}
	length, e := readUleb(reader)
	if e != nil {
		return "", e
	}
	valueArray := make([]byte, length)
	_, e = reader.Read(valueArray)
	if e != nil {
		return "", e
	}
	return string(valueArray), nil
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
