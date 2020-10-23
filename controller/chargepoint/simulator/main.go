// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2020 VMware, Inc. Tzvetomir Stoyanov (VMware) <tz.stoyanov@gmail.com>

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/CamusEnergy/kinney/controller/chargepoint/api"
	"github.com/CamusEnergy/kinney/controller/chargepoint/simulator/sim"
)

var (
	file    = flag.String("file", "", "excel file with recorded EV chargers data")
	dir     = flag.String("dir", "", "directory with json files, with recorded EV chargers data")
	addr    = flag.String("addr", ":8080", "IP address and port in format IP:port, used for listening for incoming API requests.")
	url     = flag.String("url", "/", "API endpoint")
	rgen    = flag.String("rand", "", "input json file with parameters for random simulator")
	forward = flag.Float64("forward", 1.0, "fast forward time multiplier")
)

func main() {
	flag.Parse()
	ev := sim.NewEvChargers(float32(*forward))
	var count int

	if *file != "" {
		c, err := sim.DataLoadExFile(file, ev)
		if err != nil {
			log.Fatal(err)
		}
		count += c
	}

	if *dir != "" {
		c, err := sim.DataLoadJsonDir(dir, ev)
		if err != nil {
			log.Fatal(err)
		}
		count += c
	}
	if count > 0 {
		fmt.Println("Loaded ", count, " samples")
	}

	if *rgen != "" {
		err := sim.DataLoadRandom(rgen, ev)
		if err != nil {
			log.Fatal(err)
		}
	}

	sim.DataPrint(ev)

	mux := http.NewServeMux()
	mux.Handle(*url, api.NewHandler(sim.SimulatorServer{Ev: ev}))
	err := http.ListenAndServe(*addr, mux)

	fmt.Println(err)
}
