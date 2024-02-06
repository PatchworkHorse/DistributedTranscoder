## Lang(s)
- Go

## Frameworks / Platforms / Tools 
- k3s.io (Container orchestration) 
- KEDA (Scaling, use S3 bucket length or Redis queue length for scale param)
- ffmpeg (Transcoding) 
- Asynq (Task/queue lib for Golang)
- Redis (job qeue)
- gRPC (request/response)
- Fiber (HTTP routing if we support REST) 
- S3 storage (Linode or AWS)
  
## Demands/Requirements (high level) 
### RPC Endpoints (Should we also offer REST?)
- Transcode
    - Accept URL to media file OR
    - bin data (Todo: Chunked upload)
    - List of demands for the operation (Resolution, output format) 
    - Return a GUID for the job
- FetchResult
    - Accept the GUID
    - Return link to storage result (s3?)

## Inbound flow 
1) Transcode request (Go)
2) Assign GUID 
3) Place file in object store (S3) - id.media (I.e., 41ff77244c8346c5b700b8571cd581e4.media)
4) Append custom headers to the S3 object - Date uploaded, demands, other metadata

?? Add job to Redis ?? 

## Worker flow (unorder)
Todo: 
- Monitoring queue (Asynq)
- Container lifetime (Per transcode job so we start from a known state every time?)
- Scale-to-zero (?)
- Break up media files into smaller pieces, then append just before dumping in result bucket
- 
