package store

import "time"

type Draw struct {
	Uuid       int64        `bson:"uuid,omitempty"`
	Number     int64        `bson:"number,omitempty"`
	Day        time.Weekday `bson:"day,omitempty"`
	Cycle      int32        `bson:"cycle,omitempty"`
	Forclusion time.Time    `bson:"forclusion,omitempty"`
}
