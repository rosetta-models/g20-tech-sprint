# REGnosys G20 API Demonstration

## Introduction
As part of the G20 work we have built an application to demonstrate how a client can interact with the REGnosys reporting API.
We have written the demonstration in a simple Go lang program. The program requests data from the exported REGnosys reporting API endpoint.
The program processes 7 input json files. These are available to review here: [input folder](./input)

## Usage
```go
go test apidemo/src/gtwenty 
go build
./apidemo -auth default-ro-api-key -project system/read-only-G20-TECH-SPRINT-0-1-0/latest/MAS/SFA/MAS_2013
```
Usage of `./apidemo`:
*  `-auth string`
    	The authorisation token
*  `-project string`
    	The username together with the project and report name. E.g. `joe.bloggs_auth0_123/g20/latest/MAS/SFA/MAS_2013`

## Output

This will produce a single HTML regulatory report: [Regulatory report](./output/report.html). 
