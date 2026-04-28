package scheduler

import (
	"fmt"
	"reflect"
)

var (
	registeredJobType = reflect.TypeOf((*Job)(nil)).Elem()
	registeredCtors   []func() Job
)

func Register(fns ...any) {
	for _, fn := range fns {
		rv := reflect.ValueOf(fn)
		rt := rv.Type()
		if rt.Kind() != reflect.Func || rt.NumIn() != 0 || rt.NumOut() != 1 {
			panic(fmt.Sprintf("scheduler.Register: %T must be a zero-arg constructor", fn))
		}
		if !rt.Out(0).Implements(registeredJobType) {
			panic(fmt.Sprintf("scheduler.Register: %T return type does not implement Job", fn))
		}
		captured := rv
		registeredCtors = append(registeredCtors, func() Job {
			return captured.Call(nil)[0].Interface().(Job)
		})
	}
}

func Build() []Job {
	jobs := make([]Job, 0, len(registeredCtors))
	for _, ctor := range registeredCtors {
		jobs = append(jobs, ctor())
	}
	return jobs
}

func PopulateRegistry(registry *Registry) {
	if registry == nil {
		panic("scheduler: registry is nil")
	}

	for _, ctor := range registeredCtors {
		ctor := ctor
		job := ctor()
		if job == nil {
			panic("scheduler: registered constructor returned nil")
		}
		registry.MustRegister(job.Name(), func() Job {
			return ctor()
		})
	}
}

func NewRegisteredRegistry() *Registry {
	registry := NewRegistry()
	PopulateRegistry(registry)
	return registry
}

type Registry struct {
	factories map[string]Factory
}

func NewRegistry() *Registry {
	return &Registry{factories: make(map[string]Factory)}
}

func (r *Registry) Register(key string, factory Factory) error {
	if r == nil {
		return fmt.Errorf("scheduler: registry is nil")
	}
	if key == "" {
		return fmt.Errorf("scheduler: job key is required")
	}
	if factory == nil {
		return fmt.Errorf("scheduler: factory is required for %q", key)
	}
	if _, exists := r.factories[key]; exists {
		return fmt.Errorf("scheduler: duplicate job key %q", key)
	}
	r.factories[key] = factory
	return nil
}

func (r *Registry) MustRegister(key string, factory Factory) {
	if err := r.Register(key, factory); err != nil {
		panic(err)
	}
}

func (r *Registry) Build(key string) (Job, error) {
	if r == nil {
		return nil, fmt.Errorf("scheduler: registry is nil")
	}
	factory, exists := r.factories[key]
	if !exists {
		return nil, fmt.Errorf("scheduler: job %q is not registered", key)
	}
	job := factory()
	if job == nil {
		return nil, fmt.Errorf("scheduler: factory for %q returned nil", key)
	}
	return job, nil
}
