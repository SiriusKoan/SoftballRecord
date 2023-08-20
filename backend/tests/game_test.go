package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"softball_record/config"
	"softball_record/db"
	"softball_record/models"
	"time"
)

type GameTestSuite struct {
	suite.Suite
	col   *mongo.Collection
	g1_id string
	g2_id string
    p1_id string
}

func (suite *GameTestSuite) SetupSuite() {
	// setup
	config.Init("testing")
	db.Init()
	suite.col = db.GetGameCollection()
}

func (suite *GameTestSuite) SetupTest() {
	// setup
	g1 := models.Game{HomeTeam: "home1", AwayTeam: "away1"}
	g2 := models.Game{HomeTeam: "home2", AwayTeam: "away2"}
    p1 := models.HittingPlayer{Name: "player1"}
	res, _ := suite.col.InsertOne(context.Background(), g1)
	suite.g1_id = res.InsertedID.(primitive.ObjectID).Hex()
	res, _ = suite.col.InsertOne(context.Background(), g2)
	suite.g2_id = res.InsertedID.(primitive.ObjectID).Hex()
    res, _ = suite.col.InsertOne(context.Background(), p1)
    suite.p1_id = res.InsertedID.(primitive.ObjectID).Hex()
}

func (suite *GameTestSuite) TearDownTest() {
	// teardown
	suite.col.DeleteMany(context.Background(), bson.M{})
}

func (suite *GameTestSuite) TestGetGame() {
	assert := assert.New(suite.T())
	game, _ := models.GetGame(suite.g1_id)
	assert.Equal(game.HomeTeam, "home1")
}

func (suite *GameTestSuite) TestGetAllGames() {
	assert := assert.New(suite.T())
	games, _ := models.GetAllGames()
	assert.Equal(len(games), 2)
}

func (suite *GameTestSuite) TestCreateGame() {
	assert := assert.New(suite.T())
	id, _ := models.CreateGame("home3", "away3", time.Now(), "stadium1")
	assert.NotEqual(id, "")
}

func (suite *GameTestSuite) TestAddRecord() {
	assert := assert.New(suite.T())
	record := models.Record{Team: "test", Player: "player", Runs: 3}
	res := models.AddRecord(suite.g1_id, record)
	assert.Equal(res, true)
	var game models.Game
	suite.col.FindOne(context.Background(), bson.M{"_id": suite.g1_id}).Decode(&game)
	assert.Equal(game.Records[0].Team, "test")
	assert.Equal(game.Records[0].Player, "player")
	assert.Equal(game.Records[0].Runs, 3)
}

// update records
func (suite *GameTestSuite) TestUpdateRecords() {
	assert := assert.New(suite.T())
	r1 := models.Record{Team: "home", Player: "player1"}
	r2 := models.Record{Team: "away", Player: "player2"}
	res := models.UpdateRecords(suite.g1_id, []models.Record{r1, r2})
	assert.Equal(res, true)
	var game models.Game
	suite.col.FindOne(context.Background(), bson.M{"_id": suite.g1_id}).Decode(&game)
	assert.Equal(len(game.Records), 2)
	assert.Equal(game.Records[0].Team, "home")
	assert.Equal(game.Records[0].Player, "player1")
	assert.Equal(game.Records[1].Team, "away")
	assert.Equal(game.Records[1].Player, "player2")
}

func (suite *GameTestSuite) TestSaveGame() {
    assert := assert.New(suite.T())
    record := models.Record{Team: "test", Player: suite.p1_id, Result: "2B", RBI: 1}
    models.AddRecord(suite.g1_id, record)
    models.SaveGame(suite.g1_id)
    var game models.Game
    suite.col.FindOne(context.Background(), bson.M{"_id": suite.g1_id}).Decode(&game)
    assert.Equal(game.Records[0].Player, suite.p1_id)
    var player models.HittingPlayer
    suite.col.FindOne(context.Background(), bson.M{"_id": suite.p1_id}).Decode(&player)
    assert.Equal(player.Doubles, 1)
    assert.Equal(player.RBI, 1)
}

