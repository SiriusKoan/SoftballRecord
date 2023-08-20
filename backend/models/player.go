package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"softball_record/db"
)

type PlayerMeta struct {
	Name string `bson:"name" json:"name"`
}

type HittingPlayer struct {
	Name    string `bson:"name" json:"name"`
	Games   int    `bson:"games" json:"games"`
	AB      int    `bson:"ab" json:"ab"`
	Runs    int    `bson:"runs" json:"runs"`
	Hits    int    `bson:"hits" json:"hits"`
	Doubles int    `bson:"doubles" json:"doubles"`
	Triples int    `bson:"triples" json:"triples"`
	HR      int    `bson:"hr" json:"hr"`
	SF      int    `bson:"sf" json:"sf"`
	RBI     int    `bson:"rbi" json:"rbi"`
	BB      int    `bson:"bb" json:"bb"`
	SO      int    `bson:"so" json:"so"`
}

type PitchingPlayer struct {
	Name   string  `bson:"name" json:"name"`
	Games  int     `bson:"games" json:"games"`
	Wins   int     `bson:"wins" json:"wins"`
	Losses int     `bson:"losses" json:"losses"`
	IP     float64 `bson:"ip" json:"ip"`
	Hits   int     `bson:"hits" json:"hits"`
	Runs   int     `bson:"runs" json:"runs"`
	ER     int     `bson:"er" json:"er"`
	HR     int     `bson:"hr" json:"hr"`
	BB     int     `bson:"bb" json:"bb"`
	SO     int     `bson:"so" json:"so"`
}

// hitting player methods
func (h HittingPlayer) GetAVG() float64 {
	return float64(h.Hits) / float64(h.AB)
}

func (h HittingPlayer) GetOBP() float64 {
	return float64(h.Hits+h.BB) / float64(h.AB+h.BB+h.SF)
}

func (h HittingPlayer) GetSLG() float64 {
	return float64(h.Hits+h.Doubles+2*h.Triples+3*h.HR) / float64(h.AB)
}

func (h HittingPlayer) GetOPS() float64 {
	return h.GetOBP() + h.GetSLG()
}

func GetAllHittingPlayers() ([]HittingPlayer, error) {
	col := db.GetHittingPlayerCollection()
	var players []HittingPlayer
	cur, err := col.Find(context.Background(), bson.M{})
	if err != nil {
		return players, err
	}
	for cur.Next(context.Background()) {
		var player HittingPlayer
		err := cur.Decode(&player)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return players, nil
			}
			return players, err
		}
		players = append(players, player)
	}
	return players, nil
}

func GetHittingPlayer(id string) (HittingPlayer, error) {
    col := db.GetHittingPlayerCollection()
    var player HittingPlayer
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return player, err
    }
    err = col.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&player)
    if err != nil {
        return player, err
    }
    return player, nil
}

func GetHittingPlayerByName(name string) (HittingPlayer, error) {
	col := db.GetHittingPlayerCollection()
	var player HittingPlayer
	err := col.FindOne(context.Background(), bson.M{"name": name}).Decode(&player)
	if err != nil {
		return player, err
	}
	return player, nil
}

func CreateHittingPlayer(name string) (string, error) {
	h := HittingPlayer{
		Name:    name,
		Games:   0,
		AB:      0,
		Runs:    0,
		Hits:    0,
		Doubles: 0,
		Triples: 0,
		HR:      0,
		SF:      0,
		RBI:     0,
		BB:      0,
		SO:      0,
	}
	col := db.GetHittingPlayerCollection()
	res, err := col.InsertOne(context.Background(), h)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func UpdateHittingPlayer(data bson.M) (HittingPlayer, error) {
	col := db.GetHittingPlayerCollection()
	var player HittingPlayer
	err := col.FindOneAndUpdate(context.Background(), bson.M{"name": data["name"]}, bson.M{"$set": data}).Decode(&player)
	if err != nil {
		return player, err
	} else {
		return player, nil
	}
}

// pitching player methods
func (p PitchingPlayer) GetWHIP() float64 {
	return float64(p.BB+p.Hits) / float64(p.IP)
}

func (p PitchingPlayer) GetERA() float64 {
	return float64(p.ER) / (float64(p.IP) * 9)
}

func (p PitchingPlayer) GetAVG() float64 {
	return float64(p.Hits) / float64(p.IP)
}

func GetAllPitchingPlayers() ([]PitchingPlayer, error) {
	col := db.GetPitchingPlayerCollection()
	var players []PitchingPlayer
	cur, err := col.Find(context.Background(), bson.M{})
	if err != nil {
		return players, err
	}
	for cur.Next(context.Background()) {
		var player PitchingPlayer
		err := cur.Decode(&player)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return players, nil
			}
			return players, err
		}
		players = append(players, player)
	}
	return players, nil
}

func GetPitchingPlayer(id string) (PitchingPlayer, error) {
    col := db.GetPitchingPlayerCollection()
    var player PitchingPlayer
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return player, err
    }
    err = col.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&player)
    if err != nil {
        return player, err
    }
    return player, nil
}

func GetPitchingPlayerByName(name string) (PitchingPlayer, error) {
	col := db.GetPitchingPlayerCollection()
	var player PitchingPlayer
	err := col.FindOne(context.Background(), bson.M{"name": name}).Decode(&player)
	if err != nil {
		return player, err
	}
	return player, nil
}

func CreatePitchingPlayer(name string) (string, error) {
	p := PitchingPlayer{
		Name:   name,
		Games:  0,
		Wins:   0,
		Losses: 0,
		IP:     0,
		Hits:   0,
		Runs:   0,
		ER:     0,
		HR:     0,
		BB:     0,
		SO:     0,
	}
	col := db.GetPitchingPlayerCollection()
	res, err := col.InsertOne(context.Background(), p)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func UpdatePitchingPlayer(data bson.M) (PitchingPlayer, error) {
	col := db.GetPitchingPlayerCollection()
	var player PitchingPlayer
	err := col.FindOneAndUpdate(context.Background(), bson.M{"name": data["name"]}, bson.M{"$set": data}).Decode(&player)
	if err != nil {
		return player, err
	} else {
		return player, nil
	}
}
