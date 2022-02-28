package store

import "time"

type Draw struct {
	Uuid       int64        `bson:"uuid"`
	Number     int64        `bson:"number"`
	Day        time.Weekday `bson:"day"`
	Cycle      int32        `bson:"cycle"`
	Forclusion time.Time    `bson:"forclusion"`
}
