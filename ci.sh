#!/bin/bash
docker build -t  ccr.ccs.tencentyun.com/langwei/test-log:v2.0 ./ &&
docker login --username=100002743387 ccr.ccs.tencentyun.com -p isd@cloud &&
docker push ccr.ccs.tencentyun.com/langwei/test-log:v2.0

