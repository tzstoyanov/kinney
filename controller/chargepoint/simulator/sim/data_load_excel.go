// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2020 VMware, Inc. Tzvetomir Stoyanov (VMware) <tz.stoyanov@gmail.com>

package sim

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	ex "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func DataLoadExFile(file *string, e *EVChargers) (int, error) {
	var samples int
	credential := ""

	f, err := ex.OpenFile(*file)
	if err != nil {
		return samples, err
	}

	cpnID := "1"
	cpnName := cpnDefName
	cpnDesc := cpnDefDesc
	e.getCPNetwork(&cpnID, &cpnName, &cpnDesc)

	vmGroup := vmwareOrganizationID
	vmName := vmwareOrganizationName
	e.getChargeFacility(&vmGroup, &vmName)

	sheets := f.GetSheetMap()
	for _, name := range sheets {
		rows, err := f.GetRows(name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, col := range rows[1:] {
			var i int

			if len(col) < 4 || len(col) > 6 {
				continue
			}
			/*
				col[0]: Timestamp,		"1583550049.76136"
				col[1]: VehicleID,		"HNA3BC734CE51"
				col[2]: Charge,			"5.661"
				col[3]: Full-Port-ID	"238421*1:569591*2"
			*/
			for i = 0; i < 4; i++ {
				if col[i] == "" {
					break
				}
			}
			if i < 4 {
				continue
			}
			vehicle := strings.TrimSpace(col[1])
			port := strings.TrimSpace(col[3])

			times := strings.SplitN(col[0], ".", -1)
			t1, err := strconv.ParseInt(times[0], 10, 64)
			if err != nil {
				continue
			}
			t2, err := strconv.ParseInt(times[1], 10, 64)
			if err != nil {
				continue
			}
			p, err := strconv.ParseFloat(col[2], 32)
			if err != nil {
				continue
			}

			ids := strings.SplitN(port, "*", -1)
			if ids[0] == "" || ids[1] == "" || ids[2] == "" {
				continue
			}
			empty := ""
			e.getChargeGroup(&vmGroup, &ids[0], &empty, &getLoadReplay{})
			e.getChargeStation(&vmGroup, &ids[0], &ids[1], &empty, &empty,
				&locGeo{lat: geoDefLat, long: geoDefLong})
			e.addChargeProbe(&vmGroup, &ids[0], &ids[1], &ids[2], &vehicle, &credential, time.Unix(t1, t2), float32(p))

			samples++
		}
	}

	sortReplaySamples(e)

	return samples, nil
}
