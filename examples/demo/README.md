# demo

A richer example application for `go-task`.

## What it shows

- proto-based config schema in `conf/conf.proto`
- example config from `conf/config.yaml`
- startup wiring kept directly in `cmd/demo/main.go`
- `init()`-based job registration from `jobs/register.go`
- one package per job under `jobs/hello` and `jobs/report`
- one job key reused by multiple task instances with different args
- extended log configuration through the public `conf.Log` model

## Run

```bash
go run ./examples/demo/cmd/demo
```

## Log configuration example

```yaml
log:
  level: "INFO"
  file_path: ""
  file_name: "task.log"
  file_pattern: "/%Y_%m/%d/task_%H.log"
  rotation_time: 24h
  file_age: 4320h
  environment: "development"
  system_name: "go-task-demo"
```
