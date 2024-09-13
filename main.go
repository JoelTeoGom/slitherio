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
	game.GameWidth = 1000
	game.GameHeight = 500
	game.PlayerSize = 50
}

type Snake struct {
	Length   int
	Position []rl.Rectangle
}

var player *Player

type Player struct {
	ID          int
	SnakePlayer Snake
	Camera      rl.Camera2D
	Direction   int
}

var playerCount = 0

func newPlayer() {
	player = new(Player)
	player.SnakePlayer.Length = 1
	player.ID = playerCount + 1
	player.Direction = 1
}

func main() {

	newGame()
	newPlayer()

	rl.InitWindow(game.GameWidth, game.GameHeight, "classic game: snake")
	rl.SetTargetFPS(60)

	if !game.GameOver {
		for !rl.WindowShouldClose() {

			if rl.IsKeyDown(rl.KeyA) {
				player.Direction += -1
			} //left
			if rl.IsKeyDown(rl.KeyD) {
				player.Direction += 1
			} //right
			if rl.IsKeyDown(rl.KeyW) {
				player.Direction += 2
			} //up
			if rl.IsKeyDown(rl.KeyS) {
				player.Direction += -2
			} //down

			drawSnake()

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

func drawSnake() {

	for i := 0; i < player.SnakePlayer.Length; i++ {

	}
}
