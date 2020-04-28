package models

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"golang.org/x/sync/semaphore"
)

const (
	DefaultSpeed      = 10
	WaitDroneStartSec = 5
)

type DroneManager struct {
	*tello.Driver
	Speed int
	// Weighted 1リソースに対しての並列アクセスの規制権を付与
	patrolSem    *semaphore.Weighted
	patrolQuit   chan bool
	isPatrolling bool
}

func NewDroneManager() *DroneManager {
	drone := tello.NewDriver("8889")
	droneManager := &DroneManager{
		Driver:       drone,
		Speed:        DefaultSpeed,
		patrolSem:    semaphore.NewWeighted(1),
		patrolQuit:   make(chan bool),
		isPatrolling: false,
	}
	work := func() {
		// TODO

	}
	robot := gobot.NewRobot("tello", []gobot.Connection{}, []gobot.Device{drone}, work)
	go robot.Start()
	time.Sleep(WaitDroneStartSec * time.Second)
	return droneManager
}

func (d *DroneManager) Patrol() {
	go func() {
		// 制限なしでセマフォを取得できる数は1
		isAquire := d.patrolSem.TryAcquire(1)
		if !isAquire {
			d.patrolQuit <- true
			d.isPatrolling = false
			return
		}
		defer d.patrolSem.Release(1)
		d.isPatrolling = true
		status := 0
		t := time.NewTicker(3 * time.Second)
		for {
			select {
			// tがクロックされた時，今回は3秒ごと
			case <-t.C:
				d.Hover()
				switch status {
				case 1:
					d.Forward(d.Speed)
				case 2:
					d.Right(d.Speed)
				case 3:
					d.Backward(d.Speed)
				case 4:
					d.Left(d.Speed)
				case 5:
					// 次，1にするため
					status = 0
				}
				status++
			case <-d.patrolQuit:
				// 2回目にPatrol()を呼び出すと，セマフォを取得できないのでif !isAquireでdroneが定石する
				t.Stop()
				d.Hover()
				d.isPatrolling = false
				return
			}
		}
	}()
}
