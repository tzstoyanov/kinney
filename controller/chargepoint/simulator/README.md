# EV chargers simulator

This simulator emulates ChargePoint web API for accessing EV chargers.
At the current stage it works only in "replay" mode:
 - Load historical data from excel or json files
 - Reply these data via the ChargePoint APIs:
  - `GetCPNInstances`
  - `GetStationGroups`
  - `GetStations`
  - `GetLoad`
 	
## Running the simulator:
1. Read the data from excel file and serve them:
   ```bash
   go run main.go --file <excel file with recorded charging data>
   ```
2. Read the data from json files and serve them:
   ```bash
   go run main.go --dir <directory with json files with recorded charging data>
   ```