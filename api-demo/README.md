# REGnosys G20 API Demonstration

## Introduction
As part of the G20 work we have built an application to demonstrate how a client can interact with the REGnosys reporting API.
We have written the demonstration in a simple Go lang program. The program requests data from the exported REGnosys reporting API endpoint.
The program processes 7 input json files. These are available to review here: [input folder](./input)

## Usage
```go
go test apidemo/src/gtwenty 
go build
./apidemo
```
Usage of `./apidemo`:
*  `-auth string`
    	The authorisation token
*  `-cookie string`
    	The affinity cookie. E.g. `ROSETTA_SESSION=0123456789.123.1234.123456`
*  `-project string`
    	The username together with the project and report name. E.g. `joe.bloggs_auth0_123/g20/latest/MAS/SFA/MAS_2013`

## Output

This will produce a single HTML regulatory report: [Regulatory report](./output/report.html). 
