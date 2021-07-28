package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

type Request struct {
	To  string
	FSM *fsm.FSM
}

type RequestData struct {
	guid     GUID
	tenantID string
	requestedDateTinme
}

type RequestObject struct {
	request    Request
	dataObject JSON
}

func NewRequest(to string) *Request {
	request := &Request{
		To: to,
	}

	request.FSM = fsm.NewFSM(
		"requested",
		fsm.Events{
			{Name: "review", Src: []string{"requested"}, Dst: "awaitingReview"},
			{Name: "reviewing", Src: []string{"awaitingReview"}, Dst: "inReview"},
			{Name: "approve", Src: []string{"inReview"}, Dst: "approved"},
			{Name: "decline", Src: []string{"inReview"}, Dst: "declined"},
			{Name: "expire", Src: []string{"inReview"}, Dst: "expired"},
			{Name: "complete", Src: []string{"approved", "declined", "expired"}, Dst: "complete"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { request.enterState(e) },
		},
	)

	return request
}

func (request *Request) enterState(e *fsm.Event) {
	fmt.Printf("The request (%s) is in state: %s\n", request.To, e.Dst)
}

func nonStruct() {

	fsm := fsm.NewFSM(
		"requested",
		fsm.Events{
			{Name: "review", Src: []string{"requested"}, Dst: "awaitingReview"},
			{Name: "reviewing", Src: []string{"awaitingReview"}, Dst: "inReview"},
			{Name: "approve", Src: []string{"inReview"}, Dst: "approved"},
			{Name: "decline", Src: []string{"inReview"}, Dst: "declined"},
			{Name: "expire", Src: []string{"inReview"}, Dst: "expired"},
			{Name: "complete", Src: []string{"approved", "declined", "expired"}, Dst: "complete"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event("review")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event("reviewing")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	// err = fsm.Event("approve")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(fsm.Current())

	err = fsm.Event("decline")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	// err = fsm.Event("expire")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(fsm.Current())

	err = fsm.Event("complete")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())
}

func main() {

	request := NewRequest("asdfasdf")

	err := request.FSM.Event("review")
	if err != nil {
		fmt.Println(err)
	}

	err = request.FSM.Event("reviewing")
	if err != nil {
		fmt.Println(err)
	}
	err = request.FSM.Event("approve")
	if err != nil {
		fmt.Println(err)
	}
	// err = request.FSM.Event("decline")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = request.FSM.Event("expire")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err = request.FSM.Event("complete")
	if err != nil {
		fmt.Println(err)
	}

	// fsm := fsm.NewFSM(
	// 	"requested",
	// 	fsm.Events{
	// 		{Name: "review", Src: []string{"requested"}, Dst: "awaitingReview"},
	// 		{Name: "reviewing", Src: []string{"awaitingReview"}, Dst: "inReview"},
	// 		{Name: "approve", Src: []string{"inReview"}, Dst: "approved"},
	// 		{Name: "decline", Src: []string{"inReview"}, Dst: "declined"},
	// 		{Name: "expire", Src: []string{"inReview"}, Dst: "expired"},
	// 		{Name: "complete", Src: []string{"approved", "declined", "expired"}, Dst: "complete"},
	// 	},
	// 	fsm.Callbacks{},
	// )

	// fmt.Println(fsm.Current())

	// err := fsm.Event("review")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(fsm.Current())

	// err = fsm.Event("reviewing")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(fsm.Current())

	// // err = fsm.Event("approve")
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }

	// // fmt.Println(fsm.Current())

	// err = fsm.Event("decline")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(fsm.Current())

	// // err = fsm.Event("expire")
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }

	// // fmt.Println(fsm.Current())

	// err = fsm.Event("complete")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(fsm.Current())
}
