package main

import "fmt"

type Number int

// 1
func (n Number) isEven() bool {
	if n%2 == 0 {
		return true
	}
	return false
}

// 2
type Speaker interface {
	speak()
}
type Person struct {
	age int
}

func (p Person) speak() {
	fmt.Println("A person speaks")
}

func describe(s Speaker) {
	s.speak()
}

// 3
type Streamer interface {
	play()
	pause()
	stop()
}
type MusicPlayer struct {
	model string
}

func (m MusicPlayer) play() {
	fmt.Println("playing the music")
}
func (m MusicPlayer) pause() {
	fmt.Println("pausing the music")
}
func (m MusicPlayer) stop() {
	fmt.Println("stop the music")
}

func main() {
	num := Number(23)
	fmt.Println(num.isEven())

	person1 := Person{
		age: 45,
	}
	describe(person1)

	myMusicPlayer := MusicPlayer{
		model: "PowerDVD",
	}
	myMusicPlayer.play()
	myMusicPlayer.stop()
}
