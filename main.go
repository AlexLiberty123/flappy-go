package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var window_width int = 600
	var window_height int = 450

	rl.InitWindow(int32(window_width), int32(window_height), "Flappy Go")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var character_y float32 = 225
	var velocity float32 = 0

	var isJump bool = false

	var tubes_x float32 = 500
	var tubes_speed float32 = 1

	var down_tube_y = 300
	var down_tube_height = 150
	var top_tube_height = 125

	var score int = 0
	var string_score string = strconv.Itoa(score)
	var isScore bool = true

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		//Падение
		if isJump {
			character_y += velocity
			velocity -= -9.81 / 75
		}

		rl.ClearBackground(rl.White)

		character := rl.Rectangle{X: 150, Y: float32(character_y), Width: 75, Height: 50}
		top_tube := rl.Rectangle{X: tubes_x, Y: 0, Width: 100, Height: float32(top_tube_height)}
		down_tube := rl.Rectangle{X: tubes_x, Y: float32(down_tube_y), Width: 100, Height: float32(down_tube_height)}
		score_line := rl.Rectangle{X: tubes_x + 100, Y: 0, Width: 5, Height: float32(window_height)}

		rl.DrawRectangle(int32(score_line.X), 0, int32(score_line.Width), int32(score_line.Height), rl.White)
		rl.DrawRectangle(int32(tubes_x), top_tube.ToInt32().Y, top_tube.ToInt32().Width, int32(top_tube_height), rl.DarkGray)
		rl.DrawRectangle(int32(tubes_x), int32(down_tube_y), down_tube.ToInt32().Width, int32(down_tube_height), rl.DarkGray)
		rl.DrawRectangle(int32(character.X), int32(character_y), int32(character.Width), int32(character.Height), rl.LightGray)

		rl.DrawText(string_score, 9, 9, 50, rl.Black)

		//Изменение координат труб
		if isJump {
			tubes_x -= tubes_speed
		}
		if tubes_x <= 0-top_tube.Width {
			tubes_x = float32(window_width)

			top_tube_height = rand.Intn(100) + 100
			down_tube_height = window_height - top_tube_height - 175
			down_tube_y = window_height - down_tube_height

			isScore = true
		}

		//Cтолкновение труб с героем
		if rl.CheckCollisionRecs(character, top_tube) || rl.CheckCollisionRecs(character, down_tube) {
			isJump = false
			character_y = 225
			tubes_speed = 1
			tubes_x = 500
			score = 0
			string_score = strconv.Itoa(score)
		}

		//Прыжок
		if rl.IsKeyPressed(rl.KeySpace) {
			character_y -= 40
			velocity = 0
			isJump = true
			tubes_speed += 0.05
		}

		//Столкновение с землёй
		if character_y >= float32(window_height) {
			isJump = false
			character_y = 225
			tubes_speed = 1
			tubes_x = 500
			score = 0
			string_score = strconv.Itoa(score)
		}

		//Начисление очков
		if rl.CheckCollisionRecs(character, score_line) && isScore {
			score += 1
			string_score = strconv.Itoa(score)
			isScore = false
			fmt.Println(score)
		}

		rl.EndDrawing()
	}
}
