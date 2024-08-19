package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math/rand"
	"time"
)

type Apple struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	color  rl.Color
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.InitWindow(screenWidth, screenHeight, "FlappyApples")

	rl.SetTargetFPS(120)
	bird_down := rl.LoadImage("assets/bird-down.png")
	bird_up := rl.LoadImage("assets/bird-up.png")
	texture := rl.LoadTextureFromImage(bird_up)
	rand.Seed(time.Now().UnixNano())
	var apple_loc int = rand.Intn(450-2+1) - 22
	Apples := []Apple{}
	current_apple := Apple{screenWidth, int32(apple_loc), 25, 25, rl.Red}
	Apples = append(Apples, current_apple)

	var x_coords int32 = screenWidth/2 - texture.Width/2
	var y_coords int32 = screenHeight/2 - texture.Height/2 - 40
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawTexture(texture, x_coords, y_coords, rl.White)
		rl.ClearBackground(rl.RayWhite)
		if rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(bird_up)
			y_coords -= 5
		} else {
			texture = rl.LoadTextureFromImage(bird_down)
			y_coords += 5
		}
		for io, current_apple := range Apples {
			rl.DrawRectangle(current_apple.posX, current_apple.posY, current_apple.width, current_apple.height, current_apple.color)
			Apples[io].posX = Apples[io].posX - 5
			if current_apple.posX < 0 {
				Apples[io].posX = 800
				Apples[io].posY = int32(rand.Intn(450-2+1) - 2)

			}

		}

		rl.EndDrawing()
		time.Sleep(50000000)
	}
	rl.UnloadTexture(texture)
	rl.CloseWindow()
}
