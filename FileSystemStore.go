package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

//FileSystemPlayerStore is a place connect to file system
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func initializesPlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	err := initializesPlayerDBFile(database)
	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}
	league, err := NewLeague(database)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", database.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{database}),
		league:   league,
	}, nil
}

//GetLeague get list of league from file system
func (f *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

//GetPlayerScore get wins from player data
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.FindPlayer(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.FindPlayer(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}
