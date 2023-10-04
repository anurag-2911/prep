package lld

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMeeting(t *testing.T) {
	fmt.Println("tests")
	meetingScheduler()
}

/*
A meeting scheduler software allows organizations to schedule and book meetings for a group of participants.
The scheduler determines a meeting time and location depending on the availability of the participants.
It ensures that most of the intended participants can effectively meet on the specified date and interval.
The system allows users to book and cancel meetings. The invited participants promptly receive these notifications.
The organizer can also add new participants to a meeting after the meeting is scheduled.
*/

type Participants struct {
	ID     string
	Name   string
	Email  string
	Booked Xschedule
}

type Xschedule struct {
	StartTime time.Time
	EndTime   time.Time
}
type Notification struct {
	Message     string
	MeetingTime Xschedule
}

type Xscheduler struct {
	users []*Participants
}

func meetingScheduler() {
	
}

func getRandomID() string {
	return fmt.Sprintf("%x", rand.Int())
}
func(this *Participants)getparticipants(n int)([]Participants){
	participants:=make([]Participants,0)
	usr1:=&Participants{ID: getRandomID(),Name: "draupadi",Email: "murmu@g20.com"}
	participants = append(participants,*usr1)
	usr2:=&Participants{ID: getRandomID(),Name: "murmu",Email: "murmug@g20.com"}
	participants = append(participants,*usr2)
	usr3:=&Participants{ID: getRandomID(),Name: "kovind",Email: "kovind@g20.com"}
	participants = append(participants,*usr3)
	
	return participants
}
