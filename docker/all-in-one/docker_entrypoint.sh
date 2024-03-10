#!/bin/bash

OPENFGA_LOG_FORMAT=json OPENFGA_PLAYGROUND_ENABLED=true /bin/openfga run --experimentals check-query-cache --check-query-cache-enabled &

FGACHECK=1
while [ $FGACHECK -ne 0 ]; do
	grpc_health_probe -addr=:8081
	FGACHECK=$?
done

/bin/redis-server --save 20 1 --loglevel warning --daemonize yes

/bin/datum serve --debug --pretty
