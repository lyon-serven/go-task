# go-task

A lightweight scheduled task library for Go, with two example levels for learning and adoption.

## What is in this repo

- `github.com/lyon-serven/go-task` — importable library API exposed from the module root
- `task.go` — root aliases and constructors for the public library surface
- `conf/` — public configuration model for app, log, and tasks
- `scheduler/` — core scheduling primitives, including global job constructor registration
- `examples/basic/` — minimal demo focused on the public API only
- `examples/demo/` — richer demo with proto-based config and per-job packages
- `internal/bootstrap/` — runtime wiring and logger/bootstrap helpers

## Which example should I read?

- `examples/basic` — start here if you want the shortest path to understanding the API with explicit registry wiring
- `examples/demo` — read this if you want a more realistic project structure with config, per-job packages, and `init()`-based job registration

## Run the examples

```bash
go run ./examples/basic
go run ./examples/demo/cmd/demo
```

## Example design

### `examples/basic`
Shows the minimum public API surface:
- public `conf.Config`
- one demo job in the example itself
- explicit `Registry`
- `TaskSpec`
- `LoadAndSchedule`
- direct logger creation through bootstrap helpers

### `examples/demo`
Shows a fuller app-style layout:
- `conf/conf.proto`
- `conf/config.yaml`
- `jobs/register.go`
- `jobs/hello/`
- `jobs/report/`
- blank-import driven registration in the entrypoint
- multiple task instances reusing the same job key with different args

## Log configuration
The public `conf.Log` model supports both a minimal setup and extended file layout settings.

Example:

```yaml
log:
  level: "INFO"
  file_path: "./logs/"
  file_name: "task.log"
  file_pattern: "/%Y_%m/%d/task_%H.log"
  rotation_time: 24h
  file_age: 4320h
  environment: "production"
  system_name: "go-task"
```

If `file_name`, `file_pattern`, `file_age`, `environment`, or `system_name` are omitted, bootstrap applies sensible defaults.

## Notes

- Jobs can be registered explicitly through a `Registry`, or globally through `scheduler.Register(...)`.
- Public examples use the root package API directly.
- The richer demo keeps its config schema and job wiring inside `examples/demo/`.
- The scheduler core remains separate from app-specific runtime concerns.
