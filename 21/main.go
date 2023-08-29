package main

import "fmt"

type Player interface {
	PlayMusic()
}

// Sony - структура, которая может проигрывать музыку в формате mp3
type Sony struct {
}

func (s *Sony) PlayMusic() {
	fmt.Println("Sony: Проигрывание музыки в mp3 формате")
}

// Apple - структура, которая не может проигрывать музыку в формате mp3, но может в aac
type Apple struct {
}

func (a *Apple) PlayAAC() {
	fmt.Println("Apple: Проигрывание музыки в aac формате")
}

// AppleToMP3Adapter - адаптер, чтобы проигрывать музыку в формате mp3 на Apple
type AppleToMP3Adapter struct {
	apple *Apple
}

func (adapter *AppleToMP3Adapter) PlayMusic() { // Использование структуры Adapter чтобы удовлетворять интерфесу Player
	fmt.Println("Подключение адптера к Apple")
	adapter.apple.PlayAAC()
}

func PlayMP3(player Player) { // Метод вызывающий интерфейс Playmusic в соотв-ии с переданной структурой
	player.PlayMusic()
}

func main() {
	sony := &Sony{}
	apple := &Apple{}
	adapter := &AppleToMP3Adapter{apple: apple}

	PlayMP3(sony)
	PlayMP3(adapter)
}
