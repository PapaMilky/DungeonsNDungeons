package main

import (
	utils "changeme/Utils"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	xrand "golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) TestingFunction() {
	log.Fatal("FUCK")
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Shutdown() {
	os.Exit(0)
}

func (a *App) CreateNewWorld(worldName, worldSeed, worldType, worldDiff, worldSize, playerName, playerClass string) string { //TODO Implement Infinite-Floor Mode For Generation
	var table []utils.Rooms
	var ID int
	seed, _ := strconv.Atoi(worldSeed)

	var world utils.Save
	world.WorldDifficulty, _ = strconv.Atoi(worldDiff)
	world.WorldName = worldName
	world.WorldType, _ = strconv.Atoi(worldType)
	world.WorldSeed = seed

	src := rand.NewSource(int64(seed))
	rnd := rand.New(src)
	var size int

	switch worldSize {
	case "1":
		size = 20
	case "2":
		size = 40
	case "3":
		size = 50
	case "4":
		size = 100
	case "5":
		size = 150
	case "6":
		size = 1000
	}

	source := xrand.NewSource(uint64(time.Now().UnixNano()))
	wDiff, _ := strconv.Atoi(worldDiff)

	for i := 0; i < size; i++ {
		beta := distuv.Beta{
			Alpha: float64(wDiff) * (float64(ID+1) / 50) * 1.5,
			Beta:  float64(5 - wDiff),
			Src:   source,
		}
		chestsType := distuv.Beta{
			Alpha: 0.001,
			Beta:  float64(3/2 + wDiff*(1/2)),
			Src:   source,
		}
		var buildRooms utils.Rooms
		buildRooms.RoomX = rnd.Intn(50-10) + 10
		mod := rnd.Intn(1000000000000000-100000000000000) + 100000000000000
		src.Seed(int64(mod))
		buildRooms.RoomY = rnd.Intn(50-10) + 10
		buildRooms.RoomID = ID
		size := buildRooms.RoomY * buildRooms.RoomX
		for i := 0; i < int(math.Round(float64(size/5))); i++ {
			var ob utils.Obstacles
			mod := rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ob.Type = rand.Intn(3) + 1
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			mk := rand.Intn(buildRooms.RoomX)
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ml := rand.Intn(buildRooms.RoomY)

			ob.XY = utils.Coords{
				X: mk,
				Y: ml,
			}
			buildRooms.Obstacles = append(buildRooms.Obstacles, ob)
		}
		for i := 0; i < int(math.Round(float64(size/(10+(ID/10))))); i++ {
			var ob utils.Enemies
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			mk := rand.Intn(buildRooms.RoomX)
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ml := rand.Intn(buildRooms.RoomY)

			ob.XY = utils.Coords{
				X: mk,
				Y: ml,
			}
			ob.Type = int(math.Round(beta.Rand() * 100))

			buildRooms.Enemies = append(buildRooms.Enemies, ob)
		}
		mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
		src.Seed(int64(mod))
		for i := 0; i < rnd.Intn(2)+1; i++ {
			var ob utils.Chests

			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			mk := rand.Intn(buildRooms.RoomX)
			mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
			src.Seed(int64(mod))
			ml := rand.Intn(buildRooms.RoomY)

			ob.XY = utils.Coords{
				X: mk,
				Y: ml,
			}

			ob.Type = int(math.Round(chestsType.Rand() * 10))

			buildRooms.Chests = append(buildRooms.Chests, ob)
		}
		ID++
		mod = rnd.Intn(1000000000000000-100000000000000) + 100000000000000
		src.Seed(int64(mod))

		var ex utils.Doors
		ex.Type = 1
		ex.XY.X = buildRooms.RoomX - 1
		ex.XY.Y = buildRooms.RoomY - 1
		ex.Endpoint = ID
		buildRooms.Doors = append(buildRooms.Doors, ex)

		if rand.Intn(200) == 5 {
			ex.Type = 2
			ex.XY.X = buildRooms.RoomX - 1
			ex.XY.Y = 1
			buildRooms.Doors = append(buildRooms.Doors, ex)
		}
		table = append(table, buildRooms)

	}

	world.World = table
	world.SaveName = worldName

	tmp, _ := json.Marshal(world)
	return string(tmp)
}

func (a *App) SaveGame(save string) { //TODO Implement Encoding
	d1 := []byte(save)
	var temp utils.Save
	err := json.Unmarshal(d1, &temp)
	if err != nil {
		return
	}
	path := utils.GetWorldDir()
	d2 := base64.StdEncoding.EncodeToString(d1)
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s.dat", path, temp.WorldName), []byte(d2), 0666)
	if err != nil {
		panic("Failed To Save!")
	}
}

func (a *App) ListWorlds() string {
	path := utils.GetWorldDir()
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var list string
	for _, file := range files {
		name := strings.ReplaceAll(file.Name(), ".dat", "")
		list = list + fmt.Sprintf(`<div>
        <label for="loadableWorldName">%s</label>
        <button id="%s" class="MainButtons" onclick="loadWorld('%s')">%s</button>
</div>`, name, name, file.Name(), name)
	}
	list = list + `<div>
      <button class="MainButtons" onclick="back()">Back</button>
    </div>
<div class="Alert" id="failAlert">
  <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
  Fail
</div>`
	return list
}

func (a *App) LoadGameFromString(nameString string) (string, error) {
	data, err := os.ReadFile(utils.GetWorldDir() + "/" + nameString)
	if err != nil {
		return "", err
	}
	decodeString, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return "", err
	}
	return string(decodeString), nil
}

func (a *App) ExistingCheck(name string) error {
	_, err := os.ReadFile(utils.GetWorldDir() + "/" + name + ".dat")
	if err != nil {
		return nil
	}

	return errors.New("failed")
}
