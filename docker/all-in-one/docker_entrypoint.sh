#!/bin/bash

OPENFGA_LOG_FORMAT=json OPENFGA_PLAYGROUND_ENABLED=true /bin/openfga run --experimentals check-query-cache --check-query-cache-enabled &

FGACHECK=1
while [ $FGACHECK -ne 0 ]; do
	grpc_health_probe -addr=:8081
	FGACHECK=$?
done

/bin/datum serve --dev --debug --pretty --auth=true --fga-host=localhost:8080 --fga-scheme=http 
