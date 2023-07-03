package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"softball_record/config"
	"softball_record/db"
	"softball_record/models"
	"testing"
)

// HittingPlayerTestSuite
type HittingPlayerTestSuite struct {
	suite.Suite
}

func (suite *HittingPlayerTestSuite) SetupTest() {
	// setup
	config.Init("testing")
	db.Init()
	h1 := models.HittingPlayer{Name: "test"}
	h2 := models.HittingPlayer{Name: "test2"}
	col := db.GetHittingPlayerCollection()
	col.InsertOne(context.Background(), h1)
	col.InsertOne(context.Background(), h2)
}

func (suite *HittingPlayerTestSuite) TearDownTest() {
	// teardown
	col := db.GetHittingPlayerCollection()
	col.DeleteMany(context.Background(), bson.M{})
}

func (suite *HittingPlayerTestSuite) TestGetHittingPlayerByName() {
	assert := assert.New(suite.T())
	player := models.GetHittingPlayerByName("test")
	assert.Equal(player.Name, "test")
}

func (suite *HittingPlayerTestSuite) TestGetAllHittingPlayers() {
	assert := assert.New(suite.T())
	players := models.GetAllHittingPlayers()
	assert.Equal(len(players), 2)
}

func (suite *HittingPlayerTestSuite) TestCreateHittingPlayer() {
	assert := assert.New(suite.T())
	id := models.CreateHittingPlayer("test3")
	assert.NotEqual(id, "")
}

func (suite *HittingPlayerTestSuite) TestUpdateHittingPlayer() {
	assert := assert.New(suite.T())
	player := models.GetHittingPlayerByName("test")
	player.AB = 1
	models.UpdateHittingPlayer(player)
	player = models.GetHittingPlayerByName("test")
	assert.Equal(player.AB, 1)
}

func TestHittingPlayerTestSuite(t *testing.T) {
	suite.Run(t, new(HittingPlayerTestSuite))
}

// PitchingPlayerTestSuite
type PitchingPlayerTestSuite struct {
	suite.Suite
}

func (suite *PitchingPlayerTestSuite) SetupTest() {
	// setup
	config.Init("testing")
	db.Init()
	p1 := models.PitchingPlayer{Name: "test"}
	p2 := models.PitchingPlayer{Name: "test2"}
	col := db.GetPitchingPlayerCollection()
	col.InsertOne(context.Background(), p1)
	col.InsertOne(context.Background(), p2)
}

func (suite *PitchingPlayerTestSuite) TearDownTest() {
	// teardown
	col := db.GetPitchingPlayerCollection()
	col.DeleteMany(context.Background(), bson.M{})
}

func (suite *PitchingPlayerTestSuite) TestGetPitchingPlayerByName() {
	assert := assert.New(suite.T())
	player := models.GetPitchingPlayerByName("test")
	assert.Equal(player.Name, "test")
}

func (suite *PitchingPlayerTestSuite) TestGetAllPitchingPlayers() {
	assert := assert.New(suite.T())
	players := models.GetAllPitchingPlayers()
	assert.Equal(len(players), 2)
}

func (suite *PitchingPlayerTestSuite) TestCreatePitchingPlayer() {
	assert := assert.New(suite.T())
	id := models.CreatePitchingPlayer("test3")
	assert.NotEqual(id, "")
}

func (suite *PitchingPlayerTestSuite) TestUpdatePitchingPlayer() {
	assert := assert.New(suite.T())
	player := models.GetPitchingPlayerByName("test")
	player.IP = 1.0
	models.UpdatePitchingPlayer(player)
	player = models.GetPitchingPlayerByName("test")
	assert.Equal(player.IP, 1.0)
}

func TestPitchingPlayerTestSuite(t *testing.T) {
	suite.Run(t, new(PitchingPlayerTestSuite))
}
