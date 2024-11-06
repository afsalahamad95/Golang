package main

import (
	"strings"
	"testing"
	"time"
)

type EventNotifications struct {
	Content string    `json:"content" bson:"content"`
	Timings time.Time `json:"timings" bson:"timings"`
}

func storeDb(document EventNotifications) string {
	return "insert success"
}

func TestStoreDb(t *testing.T) {
	marker := "success"
	document := EventNotifications{Content: "test", Timings: time.Now()}
	res := storeDb(document)
	if strings.Contains(res, marker) {
		t.Log("test success")
	} else {
		t.Errorf("expected %s not in %s", marker, res)
	}

}
