package lib

import "testing"

func TestOps(t *testing.T) {

	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager faild!")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed!not empty!")
	}
	music1 := &Music{"1", "name", "songgl", "https://music.io/12", "mp3"}
	mm.Add(music1)
	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed!")
	}
	m := mm.Find(music1.Name)
	if m == nil {
		t.Error("MusicManager.Find failed!")
	}
	if m.Id != music1.Id || m.Name != music1.Name || m.Source != music1.Source || m.Author != music1.Author || m.Type != music1.Type {
		t.Error("MusicManager.Find failed!Found item mismatch")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed!", err)
	}
	m = mm.Remove(m.Name)
	if m == nil || mm.Len() != 0 {
		t.Error("MustcManager.Remove() failed")
	}

}
