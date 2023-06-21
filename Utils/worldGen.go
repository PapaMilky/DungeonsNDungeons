package utils

type Rooms struct {
	RoomID    int         `json:"roomID"`
	RoomX     int         `json:"roomX"`
	RoomY     int         `json:"roomY"`
	RoomName  string      `json:"roomName"`
	Obstacles []Obstacles `json:"obstacles"`
	Enemies   []Enemies   `json:"enemies"`
	Chests    []Chests    `json:"chests"`
	Doors     []Doors     `json:"doors"`
}

type Obstacles struct {
	Type int    `json:"type"`
	XY   Coords `json:"XY"`
}

type Enemies struct {
	Type int    `json:"type"`
	XY   Coords `json:"XY"`
}

type Chests struct {
	Type int    `json:"type"`
	XY   Coords `json:"XY"`
}

type Doors struct {
	Type     int    `json:"type"`
	XY       Coords `json:"XY"`
	Endpoint int    `json:"endpoint"`
}

type Coords struct {
	X int `json:"X"`
	Y int `json:"Y"`
}
