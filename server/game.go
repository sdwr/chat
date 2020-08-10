package main


var (
	startingRoom = []string{"a small wooden shack", 
	"You are in a room made of wood. There are boarded-up windows and a door leading west", 
	"A table is in the middle of the room. It has a rusty sword and a lantern on it."}
	blankRooms = []Room{
		{"a sparkling cavern",
		"You are surrounded with amethyst walls that sparkle and shimmer. Light is coming from a floating orb in the center of the cave.", ""},
		{"some palm trees",
		"You see a stand of palm trees with a hammock strung between them. What's that doing here?", ""},
		{"a rusted-out Pontiac Aztek",
		"This car has seen better days. The hubcaps are missing.", ""},
		{"a completely uninteresting field",
		"This space's only remarkable feature is just how unremarkable it is. Boring", ""}}
	roomContents = []string{
		"a vein of copper ore running along a seam in the wall.",
		"a dozen copies of Pet Sounds on compact disc",
		"blue and green cats-eye marbles",
		"a white feather"}
)


type Level struct {
	Rooms [][]Room
}

type Room struct {
	Teaser string `json:"Teaser"`
	Description string `json:Description"`
	Contents string `json:Contents"`
}

type PlayerState struct {
	Coords [2]int
}

func createLevel() Level {
	rooms := [][]Room{}
	for y := 0; y < 5; y++ {
		roomRow := [5]Room{}
		for x := 0; x < 5; x++ {
			roomRow[x] = createRoom()
		}
	}
	return Level{rooms}
}

func createRoom() Room {
	room := blankRooms[Random(len(blankRooms))]
	if(Random(3) == 0) {
		room.Contents = roomContents[Random(len(roomContents))]
	}
	return room
}


