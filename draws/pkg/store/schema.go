package store

import "time"

type Draw struct {
	UUID       string       `bson:"uuid"`
	number     int64        `bson:"number"`
	day        time.Weekday `bson:"day"`
	cycle      int32        `bson:"cycle"`
	forclusion time.Time    `bson:"forclusion"`
}
