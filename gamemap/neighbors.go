package gamemap

import (
	"context"

	"github.com/smxlong/mud/ent"
)

// Opposite returns the opposite of the given direction
func Opposite(direction string) string {
	switch direction {
	case "north":
		return "south"
	case "south":
		return "north"
	case "east":
		return "west"
	case "west":
		return "east"
	case "up":
		return "down"
	case "down":
		return "up"
	}
	panic("invalid direction")
}

// Neighbor is returned by the Neighbors method.
type Neighbor struct {
	Room      *ent.Room
	Door      *ent.Door
	Direction string
}

// Neighbors returns the neighbors of a room.
func Neighbors(ctx context.Context, room *ent.Room) ([]*Neighbor, error) {
	neighbors := make([]*Neighbor, 0)
	// 1. Doors for which this is the from room
	doors, err := room.QueryDoors().WithTo().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, door := range doors {
		neighbors = append(neighbors, &Neighbor{
			Room:      door.Edges.To,
			Door:      door,
			Direction: door.Direction.String(),
		})
	}
	// 2. Doors for which this is the to room
	doorsIn, err := room.QueryDoorsIn().WithFrom().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, door := range doorsIn {
		neighbors = append(neighbors, &Neighbor{
			Room:      door.Edges.From,
			Door:      door,
			Direction: Opposite(door.Direction.String()),
		})
	}
	return neighbors, nil
}
