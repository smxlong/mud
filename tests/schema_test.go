package tests

// This file contains tests of the schema itself. It uses an in-memory SQLite
// database to test the schema and its constraints.

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/smxlong/mud/ent"
	"github.com/smxlong/mud/ent/playerrole"
	"github.com/smxlong/mud/gamemap"
	"github.com/smxlong/mud/password"
	"github.com/stretchr/testify/require"
)

const TEST_EMAIL = "test@example.com"

// newEntClient returns a new ent.Client for testing.
func newEntClient(t *testing.T) *ent.Client {
	t.Helper()
	cli, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	require.NoError(t, err)
	require.NoError(t, cli.Schema.Create(context.Background()))
	return cli
}

func Test_Simple_World(t *testing.T) {
	cli := newEntClient(t)
	defer cli.Close()
	ctx := context.Background()

	// Create a new room.
	room, err := cli.Room.Create().
		SetName("test room").
		Save(ctx)
	require.NoError(t, err)

	// Create a second room
	room2, err := cli.Room.Create().
		SetName("test room 2").
		Save(ctx)
	require.NoError(t, err)

	// Create a door
	door, err := cli.Door.Create().
		SetName("test door").
		SetDirection("east").
		SetFrom(room).
		SetTo(room2).
		Save(ctx)
	require.NoError(t, err)

	// Check that the door is in the first room
	doors, err := room.QueryDoors().All(ctx)
	require.NoError(t, err)
	require.Len(t, doors, 1)
	require.Equal(t, doors[0].ID, door.ID)

	// Check that the door is in the second room
	doors, err = room2.QueryDoorsIn().All(ctx)
	require.NoError(t, err)
	require.Len(t, doors, 1)
	require.Equal(t, doors[0].ID, door.ID)

	// Get all the neighbors of room
	neighbors, err := gamemap.Neighbors(ctx, room)
	require.NoError(t, err)
	require.Len(t, neighbors, 1)
	require.Equal(t, neighbors[0].Room.ID, room2.ID)
	require.Equal(t, neighbors[0].Door.ID, door.ID)
	require.Equal(t, neighbors[0].Direction, "east")

	// Get all the neighbors of room2
	neighbors, err = gamemap.Neighbors(ctx, room2)
	require.NoError(t, err)
	require.Len(t, neighbors, 1)
	require.Equal(t, neighbors[0].Room.ID, room.ID)
	require.Equal(t, neighbors[0].Door.ID, door.ID)
	require.Equal(t, neighbors[0].Direction, "west")
}

func Test_Tetrahedron_Of_Rooms(t *testing.T) {
	cli := newEntClient(t)
	defer cli.Close()
	ctx := context.Background()

	// Create a tetrahedron of rooms. A is the apex node, B, C, D are the base nodes.
	a, err := cli.Room.Create().
		SetName("A").
		Save(ctx)
	require.NoError(t, err)

	b, err := cli.Room.Create().
		SetName("B").
		Save(ctx)
	require.NoError(t, err)

	c, err := cli.Room.Create().
		SetName("C").
		Save(ctx)
	require.NoError(t, err)

	d, err := cli.Room.Create().
		SetName("D").
		Save(ctx)
	require.NoError(t, err)

	// A-B, A-C, A-D, B-C, C-D, D-B
	_, err = cli.Door.Create().
		SetName("A-B").
		SetDirection("north").
		SetFrom(a).
		SetTo(b).
		Save(ctx)
	require.NoError(t, err)

	_, err = cli.Door.Create().
		SetName("A-C").
		SetDirection("east").
		SetFrom(a).
		SetTo(c).
		Save(ctx)
	require.NoError(t, err)

	_, err = cli.Door.Create().
		SetName("A-D").
		SetDirection("south").
		SetFrom(a).
		SetTo(d).
		Save(ctx)
	require.NoError(t, err)

	_, err = cli.Door.Create().
		SetName("B-C").
		SetDirection("west").
		SetFrom(b).
		SetTo(c).
		Save(ctx)
	require.NoError(t, err)

	_, err = cli.Door.Create().
		SetName("C-D").
		SetDirection("up").
		SetFrom(c).
		SetTo(d).
		Save(ctx)
	require.NoError(t, err)

	_, err = cli.Door.Create().
		SetName("D-B").
		SetDirection("down").
		SetFrom(d).
		SetTo(b).
		Save(ctx)
	require.NoError(t, err)

	// Check that the tetrahedron is correct
	neighbors, err := gamemap.Neighbors(ctx, a)
	require.NoError(t, err)
	require.Len(t, neighbors, 3)
	require.True(t, hasNeighbor(neighbors, b, "north"))
	require.True(t, hasNeighbor(neighbors, c, "east"))
	require.True(t, hasNeighbor(neighbors, d, "south"))

	neighbors, err = gamemap.Neighbors(ctx, b)
	require.NoError(t, err)
	require.Len(t, neighbors, 3)
	require.True(t, hasNeighbor(neighbors, a, "south"))
	require.True(t, hasNeighbor(neighbors, c, "west"))
	require.True(t, hasNeighbor(neighbors, d, "up"))

	neighbors, err = gamemap.Neighbors(ctx, c)
	require.NoError(t, err)
	require.Len(t, neighbors, 3)
	require.True(t, hasNeighbor(neighbors, a, "west"))
	require.True(t, hasNeighbor(neighbors, b, "east"))
	require.True(t, hasNeighbor(neighbors, d, "up"))

	neighbors, err = gamemap.Neighbors(ctx, d)
	require.NoError(t, err)
	require.Len(t, neighbors, 3)
	require.True(t, hasNeighbor(neighbors, a, "north"))
	require.True(t, hasNeighbor(neighbors, b, "down"))
	require.True(t, hasNeighbor(neighbors, c, "down"))
}

func hasNeighbor(neighbors []*gamemap.Neighbor, room *ent.Room, direction string) bool {
	for _, n := range neighbors {
		if n.Room.ID == room.ID && n.Direction == direction {
			return true
		}
	}
	return false
}

func Test_Player_In_Room(t *testing.T) {
	cli := newEntClient(t)
	defer cli.Close()
	ctx := context.Background()

	// Create a room
	room, err := cli.Room.Create().
		SetName("test room").
		Save(ctx)
	require.NoError(t, err)

	// Create a player
	player, err := cli.Player.Create().
		SetName("test player").
		SetPassword(password.Hash("password")).
		SetEmail(TEST_EMAIL).
		SetRoom(room).
		Save(ctx)
	require.NoError(t, err)

	// Check that the player is in the room
	players, err := room.QueryPlayers().WithRoom().All(ctx)
	require.NoError(t, err)
	require.Len(t, players, 1)
	require.Equal(t, players[0].ID, player.ID)
	require.Equal(t, players[0].Edges.Room.ID, room.ID)
}

func Test_Player_Assign_Role(t *testing.T) {
	cli := newEntClient(t)
	defer cli.Close()
	ctx := context.Background()

	// Create a player
	player, err := cli.Player.Create().
		SetName("test player").
		SetPassword(password.Hash("password")).
		SetEmail(TEST_EMAIL).
		Save(ctx)
	require.NoError(t, err)

	// Create a role
	role, err := cli.PlayerRole.Create().
		SetName("test role").
		Save(ctx)
	require.NoError(t, err)

	// Assign the role to the player
	_, err = player.Update().
		AddPlayerRoles(role).
		Save(ctx)
	require.NoError(t, err)

	// Check that the player has the role
	roles, err := player.QueryPlayerRoles().All(ctx)
	require.NoError(t, err)
	require.Len(t, roles, 1)
	require.Equal(t, roles[0].ID, role.ID)
	// use Where
	hasRole, err := player.QueryPlayerRoles().Where(playerrole.ID(role.ID)).Exist(ctx)
	require.NoError(t, err)
	require.True(t, hasRole)
}
