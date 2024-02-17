# Public API 

Place a transcoding job in the queue, request the status of a given task.

## HTTP

HTTP POST /api/{version}/transcode
Returns: HTTP 200 / task GUID 

HTTP GET /api/{version}/transcode/{taskGuid}
Returns: HTTP 200 / Task Status / Result URI

## gRPC

*Todo: gRPC section*