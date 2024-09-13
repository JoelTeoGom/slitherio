package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var game *GameState

type GameState struct {
	GameWidth  int32
	GameHeight int32
	GameOver   bool
	PlayerSize int
}

func newGame() {
	// Inicializar la variable global
	game = new(GameState)
	game.GameOver = false
	game.GameWidth = 800
	game.GameHeight = 450
	game.PlayerSize = 40
}

func main() {

	newGame()
	rl.InitWindow(game.GameWidth, game.GameHeight, "classic game: snake")
	rl.SetTargetFPS(60)

	if !game.GameOver {

		for !rl.WindowShouldClose() {

			rl.BeginDrawing()

			rl.ClearBackground(rl.RayWhite)
			drawTable()

			rl.EndDrawing()
		}
		//we init the map

	}

	defer rl.CloseWindow()

}

// funcion that prints the map
func drawTable() {

	for i := 0; i < int(game.GameWidth)/game.PlayerSize; i++ {
		rl.DrawLineV(rl.Vector2{float32(i) * float32(game.PlayerSize), 0}, rl.Vector2{float32(i) * float32(game.PlayerSize), float32(game.GameHeight)}, rl.Black)
	}

	for i := 0; i < int(game.GameHeight)/game.PlayerSize; i++ {
		rl.DrawLineV(rl.Vector2{0, float32(i) * float32(game.PlayerSize)}, rl.Vector2{float32(game.GameWidth), float32(i) * float32(game.PlayerSize)}, rl.Black)
	}

}
