package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
)

type Save struct {
	WorldName       string     `json:"worldName"`
	SaveName        string     `json:"saveName"`
	TimeElapsed     int        `json:"timeElapsed"`
	WorldSeed       int        `json:"worldSeed"`
	WorldType       int        `json:"worldType"`
	WorldDifficulty int        `json:"worldDifficulty"`
	Player          PlayerData `json:"player"`
	World           []Rooms    `json:"world"`
}

type PlayerData struct {
	PlayerName string    `json:"playerName"`
	Level      int       `json:"level"`
	Exp        int       `json:"exp"`
	Class      int       `json:"class"`
	Statpoints int       `json:"statpoints"`
	Stats      Stats     `json:"stats"`
	Inventory  Inventory `json:"inventory"`
}

type Stats struct {
	Health       int `json:"health"`
	Speed        int `json:"speed"`
	Defense      int `json:"defense"`
	Strength     int `json:"strength"`
	Stamina      int `json:"stamina"`
	Magic        int `json:"magic"`
	Charisma     int `json:"charisma"`
	Intelligence int `json:"intelligence"`
}

func Test() {
	var SaveTest Save
	SaveTest.SaveName = "The Developers Nightmare"
	SaveTest.WorldName = "A Poorly Hand Written Data Set"
	SaveTest.WorldSeed = 42069420
	SaveTest.WorldDifficulty = 3
	SaveTest.WorldType = 2
	SaveTest.TimeElapsed = 6172849
	SaveTest.Player = PlayerData{
		PlayerName: "The MilkMan",
		Level:      1,
		Class:      3,
		Exp:        19,
		Statpoints: 0,
		Stats: Stats{
			Health:       10,
			Defense:      10,
			Charisma:     10,
			Magic:        10,
			Intelligence: 10,
			Strength:     10,
			Stamina:      10,
			Speed:        10,
		},
		Inventory: Inventory{
			Size:       10,
			Helm:       Item{},
			Chestplate: Item{},
			Leggings:   Item{},
			Boots:      Item{},
			Hand:       Item{},
			Slots:      []Item{},
		},
	}
	SaveTest.World = []Rooms{
		{
			RoomID:   1,
			RoomName: "The GameTesters Dream",
			RoomY:    50,
			RoomX:    25,
			Obstacles: []Obstacles{
				{
					Type: 1,
					XY: Coords{
						3,
						5,
					},
				},
				{
					Type: 2,
					XY: Coords{
						19,
						12,
					},
				},
				{
					Type: 3,
					XY: Coords{
						41,
						50,
					},
				},
				{
					Type: 4,
					XY: Coords{
						43,
						35,
					},
				},
			},
			Enemies: []Enemies{
				{
					Type: 9,
					XY: Coords{
						X: 4,
						Y: 19,
					},
				},
				{
					Type: 12,
					XY: Coords{
						X: 2,
						Y: 20,
					},
				},
				{
					Type: 3,
					XY: Coords{
						X: 8,
						Y: 16,
					},
				},
			},
			Chests: []Chests{
				{
					Type: 1,
					XY: Coords{
						X: 6,
						Y: 49,
					},
				},
				{
					Type: 2,
					XY: Coords{
						X: 20,
						Y: 30,
					},
				},
			},
			Doors: []Doors{
				{
					Type: 9,
					XY: Coords{
						X: 4,
						Y: 19,
					},
					Endpoint: 2,
				},
			},
		},
	}

	SaveTestJson, err := json.Marshal(SaveTest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(SaveTestJson))
}

func ProcedualSetTest(seed int64) {
	var table []Rooms
	var ID int

	src := rand.NewSource(seed)
	rnd := rand.New(src)

	for i := 0; i < 5; i++ {
		var new Rooms
		new.RoomX = rnd.Intn(50-10) + 10
		mod := rnd.Intn(1000000000000000-100000000000000) + 100000000000000
		src.Seed(int64(mod))
		new.RoomY = rnd.Intn(50-10) + 10
		new.RoomID = ID
		ID++
		size := new.RoomY * new.RoomX
		for i := 0; i < int(math.Round(float64(size/5))); i++ {
			var ob Obstacles
			mod := rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ob.Type = rand.Intn(3) + 1
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			mk := rand.Intn(new.RoomX)
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ml := rand.Intn(new.RoomY)

			ob.XY = Coords{
				X: mk,
				Y: ml,
			}
			new.Obstacles = append(new.Obstacles, ob)
		}
		for i := 0; i < int(math.Round(float64(size/5))); i++ {
			var ob Enemies
			mod := rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ob.Type = rand.Intn(3) + 1
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			mk := rand.Intn(new.RoomX)
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ml := rand.Intn(new.RoomY)

			ob.XY = Coords{
				X: mk,
				Y: ml,
			}
			new.Enemies = append(new.Enemies, ob)
		}
		for i := 0; i < int(math.Round(float64(size/5))); i++ {
			var ob Chests
			mod := rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ob.Type = rand.Intn(3) + 1
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			mk := rand.Intn(new.RoomX)
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ml := rand.Intn(new.RoomY)

			ob.XY = Coords{
				X: mk,
				Y: ml,
			}
			new.Chests = append(new.Chests, ob)
		}

		mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
		src.Seed(int64(mod))

		var ex Doors
		ex.Type = 1
		ex.XY.X = new.RoomX - 1
		ex.XY.Y = new.RoomY - 1
		ex.Endpoint = ID
		new.Doors = append(new.Doors, ex)

		if rand.Intn(200) == 5 {
			ex.Type = 2
			ex.XY.X = new.RoomX - 1
			ex.XY.Y = 1
			new.Doors = append(new.Doors, ex)
		}
		table = append(table, new)

	}
	tmp, _ := json.Marshal(table)
	fmt.Printf(string(tmp))
}
