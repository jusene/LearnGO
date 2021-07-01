package main

import "fmt"

type MediaPlayer interface {
	Play(audioType string, filename string)
}

type AdvanceMediaPlayer interface {
	PlayVlc(filename string)
	PlayMp4(filename string)
}

type VlcPlayer struct {
}

func (s *VlcPlayer) PlayVlc(filename string) {
	fmt.Println("Playing vlc file. Name: " + filename)
}

func (s *VlcPlayer) PlayMp4(filename string) {

}

type Mp4Player struct {
}

func (s *Mp4Player) PlayVlc(filename string) {
}

func (s *Mp4Player) PlayMp4(filename string) {
	fmt.Println("Playing mp4 file. Name: " + filename)
}

type MediaAdapter struct {
	advanceMediaPlayer AdvanceMediaPlayer
}

func (s *MediaAdapter) MediaAdapter(audioType string) {
	if audioType == "vlc" {
		s.advanceMediaPlayer = new(VlcPlayer)
	} else if audioType == "mp4" {
		s.advanceMediaPlayer = new(Mp4Player)
	}
}

func (s *MediaAdapter) Play(audioType string, filename string) {
	if audioType == "vlc" {
		s.advanceMediaPlayer.PlayVlc(filename)
	} else if audioType == "mp4" {
		s.advanceMediaPlayer.PlayMp4(filename)
	}
}

type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (s *AudioPlayer) Play(audioType string, filename string) {
	if audioType == "mp3" {
		fmt.Println("Playing mp3 file. Name: " + filename)
	} else if audioType == "vlc" || audioType == "mp4" {
		s.mediaAdapter = MediaAdapter{}
		s.mediaAdapter.MediaAdapter(audioType)
		s.mediaAdapter.Play(audioType, filename)
	} else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}
}

func main() {
	audioPlayer := AudioPlayer{}

	audioPlayer.Play("mp3", "beyond.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "away.vlc")
	audioPlayer.Play("avi", "mind.avi")
}
