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
3. Generate random data, based on configuration parameters defined in a json file :
   ```bash
   go run main.go --rand <json file with configuration parameters>
   ```
   There is a sample file with configuration parameters for the random generator
   `res/grand.json` :

	`maxCPNs` - generate up to maxCPNs ChargePoint Networks
	
	`maxFacilities` - generate up to maxFacilities charging facilities
	
	`maxChargeGroups` - generate up to maxChargeGroups charging groups in each facility
	
	`maxChargeStations` - generate up to maxChargeStations charging stations in each group
	
	`maxChargePorts` - generate up to maxChargePorts charging ports in each station
	
	`maxVehicleBattery` - the maximum capacity of a electric vehicle battery.
				Generated vehicles have battery with random capacity
				in the interval [maxVehicleBattery/2, maxVehicleBattery]
				
	`RandomSeed` - Random seed to be used, to reporduce the same random generated numbers
	
	`CPNs` - generate CPNs count of ChargePoint Networks
	
	`Facilities` - generate Facilities count of ChargePoint facilities
	
	`ChargeGroups` - generate ChargeGroups count of charging groups in each facility
	
	`ChargeStations` - generate ChargeStations count of charging stations in each group
	
	`ChargePorts` - generate ChargePorts count of charging ports in each station
	
	`PortLoad` - probability of a port to have a charging session, in %

    both exact and max numbers can be specified, in this case max numbers have a higher priority.