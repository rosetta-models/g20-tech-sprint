#!/bin/sh

go test apidemo/src/gtwenty 
go build
./apidemo -auth default-ro-api-key -project system/read-only-G20-TECH-SPRINT-0-1-1/latest/MAS/SFA/MAS_2013

