package mongodb

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
)

// 参考: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
// mongo操作对象:
// bson.M: map[string]interface{}
// bson.A: []interface{}
// bson.E: struct{Key string, Value interface{}}
// bson.D: []bson.E

const (
	logFormat     = "[mongo] [info]  %s [CMD] %s\n"
	logTimeFormat = "2006/01/02 15:04:05.999999"
)

var (
	mc *mongo.Client
)

func initMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			fmt.Printf(logFormat, time.Now().Format(logTimeFormat), e.Command)
		},
	}
	opt := options.Client().ApplyURI("mongodb://root:123456@localhost:27017").
		SetMonitor(cmdMonitor).
		SetMaxPoolSize(30).
		SetMaxConnIdleTime(time.Minute)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatalf("mongo connect error: %v", err)
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("mongo ping error: %v", err)
	}

	mc = client
	log.Println("mongo connect success")
}

func closeMongo() {
	if err := mc.Disconnect(context.Background()); err != nil {
		log.Fatalf("mongo disconnect error: %v", err)
	}
}

type Inventory struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Item string             `json:"item" bson:"item"`
	Qty  int                `json:"qty" bson:"qty"`
	Tags []string           `json:"tags" bson:"tags"`
	Size *Size              `json:"size" bson:"size"`
}

type Size struct {
	H   int     `json:"h"`
	W   float64 `json:"w"`
	Uom string  `json:"uom"`
}

func TestQuickStart(t *testing.T) {
	initMongo()
	defer closeMongo()

	// insert
	coll := mc.Database("test").Collection("inventory")
	inventory := &Inventory{
		ID:   primitive.NewObjectID(),
		Item: "canvas",
		Qty:  100,
		Tags: []string{"cotton"},
		Size: &Size{
			H:   28,
			W:   35.5,
			Uom: "cm",
		},
	}
	idStr := inventory.ID.Hex()
	log.Printf("mongo insertOne before, id: %s", idStr)
	insertResult, err := coll.InsertOne(context.Background(), inventory)
	if err != nil {
		log.Fatalf("mongo insertOne error: %v", err)
	}
	log.Printf("mongo insertOne success, result: %v", insertResult)

	// find
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	filter := bson.M{"_id": id}
	var in Inventory
	if err = coll.FindOne(context.Background(), filter).Decode(&in); err != nil {
		log.Fatalf("mongo findOne error: %v", err)
	}
	log.Printf("mongo findOne success, result: %v", in)

	// update
	update := bson.M{"$set": bson.M{"qty": 300}}
	updateResult, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatalf("mongo updateOne error: %v", err)
	}
	log.Printf("mongo updateOne success, result: %v", updateResult)

	// delete
	deleteResult, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatalf("mongo deleteOne error: %v", err)
	}
	log.Printf("mongo deleteOne success, result: %v", deleteResult)

	// Output:
	// Without monitor
	// 2023/09/30 12:17:47 mongo connect success
	// 2023/09/30 12:17:47 mongo insertOne before, id: 6517a16b3055dffe5d84c080
	// 2023/09/30 12:17:47 mongo insertOne success, result: &{ObjectID("6517a16b3055dffe5d84c080")}
	// 2023/09/30 12:17:47 mongo findOne success, result: {ObjectID("6517a16b3055dffe5d84c080") canvas 100 [cotton] 0x14000090100}
	// 2023/09/30 12:17:47 mongo updateOne success, result: &{1 1 0 <nil>}
	// 2023/09/30 12:17:47 mongo deleteOne success, result: &{1}

	// With monitor
	// [mongo] [info]  2023/09/30 12:38:51.225478 [CMD] {"ping": {"$numberInt":"1"},"lsid": {"id": {"$binary":{"base64":"winKRDFEQEGx4ycpFiyYkw==","subType":"04"}}},"$db": "admin"}
	// 2023/09/30 12:38:51 mongo connect success
	// 2023/09/30 12:38:51 mongo insertOne before, id: 6517a65baf830321b1b4a090
	// [mongo] [info]  2023/09/30 12:38:51.226799 [CMD] {"insert": "inventory","ordered": true,"lsid": {"id": {"$binary":{"base64":"winKRDFEQEGx4ycpFiyYkw==","subType":"04"}}},"$db": "test","documents": [{"_id": {"$oid":"6517a65baf830321b1b4a090"},"item": "canvas","qty": {"$numberInt":"100"},"tags": ["cotton"],"size": {"h": {"$numberInt":"28"},"w": {"$numberDouble":"35.5"},"uom": "cm"}}]}
	// 2023/09/30 12:38:51 mongo insertOne success, result: &{ObjectID("6517a65baf830321b1b4a090")}
	// [mongo] [info]  2023/09/30 12:38:51.228045 [CMD] {"find": "inventory","filter": {"_id": {"$oid":"6517a65baf830321b1b4a090"}},"limit": {"$numberLong":"1"},"singleBatch": true,"lsid": {"id": {"$binary":{"base64":"winKRDFEQEGx4ycpFiyYkw==","subType":"04"}}},"$db": "test"}
	// 2023/09/30 12:38:51 mongo findOne success, result: {ObjectID("6517a65baf830321b1b4a090") canvas 100 [cotton] 0x14000294180}
	// [mongo] [info]  2023/09/30 12:38:51.229451 [CMD] {"update": "inventory","ordered": true,"lsid": {"id": {"$binary":{"base64":"winKRDFEQEGx4ycpFiyYkw==","subType":"04"}}},"$db": "test","updates": [{"q": {"_id": {"$oid":"6517a65baf830321b1b4a090"}},"u": {"$set": {"qty": {"$numberInt":"300"}}}}]}
	// 2023/09/30 12:38:51 mongo updateOne success, result: &{1 1 0 <nil>}
	// [mongo] [info]  2023/09/30 12:38:51.23098 [CMD] {"delete": "inventory","ordered": true,"lsid": {"id": {"$binary":{"base64":"winKRDFEQEGx4ycpFiyYkw==","subType":"04"}}},"$db": "test","deletes": [{"q": {"_id": {"$oid":"6517a65baf830321b1b4a090"}},"limit": {"$numberInt":"1"}}]}
	// 2023/09/30 12:38:51 mongo deleteOne success, result: &{1}
	// [mongo] [info]  2023/09/30 12:38:51.232269 [CMD] {"endSessions": [{"id": {"$binary":{"base64":"winKRDFEQEGx4ycpFiyYkw==","subType":"04"}}}],"$db": "admin"}
}
