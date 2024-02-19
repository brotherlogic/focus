package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brotherlogic/goserver/utils"

	pb "github.com/brotherlogic/focus/proto"
)

func main() {
	ctx, cancel := utils.ManualContext("focus-cli", time.Minute)
	defer cancel()

	conn, err := utils.LFDialServer(ctx, "focus")
	if err != nil {
		log.Fatalf("Bad dial: %v", err)
	}

	client := pb.NewFocusServiceClient(conn)
	focus, err := client.GetFocus(ctx, &pb.GetFocusRequest{})
	if err != nil {
		log.Fatalf("Bad focus: %v", err)
	}

	fmt.Printf("%v (%v) [%v]\n", focus.GetFocus().GetDetail(), focus.GetFocus().GetLink(), focus.Focus.GetType())
}
