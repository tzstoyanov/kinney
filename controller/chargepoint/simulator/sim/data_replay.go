// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2020 VMware, Inc. Tzvetomir Stoyanov (VMware) <tz.stoyanov@gmail.com>

package sim

import (
	"fmt"
	"sort"
	"time"
)

type getLoadReplay struct {
	firstRecord time.Time
	lastRecord  time.Time
	timeOffset  time.Duration
}

// Match the current time to times of the recorded charge sessions
func (g getLoadReplay) calcTime(group *chargeGroup) time.Time {
	now := time.Now()
	stime := now.Add(g.timeOffset)

	if stime.Equal(g.firstRecord) || stime.Equal(g.lastRecord) ||
		(stime.After(g.firstRecord) && stime.Before(g.lastRecord)) {
		return stime
	}
	stime = g.firstRecord
	g.timeOffset = g.firstRecord.Sub(now)

	return stime
}

func (g getLoadReplay) getPortLoad(port *chargePort, t time.Time) (float32, *vehicle, error) {
	var session *chargeSession
	var pload float32

	for _, s := range port.recorded {
		if len(s.samples) < 1 {
			continue
		}
		if t.Equal(s.samples[0].time) || t.Equal(s.samples[len(s.samples)-1].time) ||
			(t.After(s.samples[0].time) && t.Before(s.samples[len(s.samples)-1].time)) {
			session = s
			break
		}
	}
	if session != nil {
		for i := len(session.samples) - 1; i >= 0; i-- {
			if t.Equal(session.samples[i].time) {
				pload = session.samples[i].power
				break
			} else if t.After(session.samples[i].time) {
				if i == (len(session.samples) - 1) {
					pload = session.samples[i].power
					break
				}
				pload = (session.samples[i].power + session.samples[i+1].power) / 2
				break
			}
		}
		return pload, &session.vehicle, nil
	}

	return pload, nil, nil
}

func (g getLoadReplay) printGetLoadParams() {
	fmt.Println("\t\tFirst record", g.firstRecord)
	fmt.Println("\t\tLast record", g.lastRecord)
}

func sortSessionsSlice(sessions []*chargeSession) {
	sort.SliceStable(sessions, func(i, j int) bool {
		if len(sessions[i].samples) < 1 {
			return true
		}
		if len(sessions[j].samples) < 1 {
			return false
		}

		return sessions[i].samples[0].time.Before(sessions[j].samples[0].time)
	})
}

func sortProbesSlice(probes []*chargeSample) {
	sort.SliceStable(probes, func(i, j int) bool {
		return probes[i].time.Before(probes[j].time)
	})
}

// Sort recorded charge sessions according to their time
// and save the times of first and last one, for each charger group
func sortReplaySamples(e *EVChargers) {
	for _, org := range e.facilities {
		for _, gr := range org.groups {
			if gr.getLoad == nil {
				continue
			}

			replay, ok := gr.getLoad.(*getLoadReplay)
			if !ok {
				continue
			}

			for _, st := range gr.stations {
				for _, pr := range st.ports {
					for _, k := range pr.recorded {
						sortProbesSlice(k.samples)
					}
					sortSessionsSlice(pr.recorded)
					if len(pr.recorded) > 0 && len(pr.recorded[0].samples) > 0 {
						f := pr.recorded[0].samples[0].time
						a := pr.recorded[len(pr.recorded)-1]
						l := a.samples[len(a.samples)-1].time
						if replay.firstRecord.IsZero() {
							replay.firstRecord = f
							replay.lastRecord = l
						} else {
							if f.Before(replay.firstRecord) {
								replay.firstRecord = f
							}
							if l.After(replay.lastRecord) {
								replay.lastRecord = l
							}
						}
					}
				}
			}
		}
	}
}
