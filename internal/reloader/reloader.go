package reloader

import "sync"

// LoaderFunc handles reloading a service for an individual service. LoaderFunc must return an error
// if a failure occurs and the application should be considered unstable.
//
// Each LoaderFunc should check if an existing service instance exists and reconfigure the instance if
// it exists, otherwise a new service instance should be configured as if this is a fresh startup.
type LoaderFunc func() error

// Reloader is implemented by services that manage LoaderFunc's.
type Reloader interface {
	// Register adds a LoaderFunc to be managed by Reloader
	Register(f LoaderFunc)
	// Reload calls all registered LoaderFunc's allowing them to re-configure what is necessary
	// when Reload is called.
	Reload() error
}

type reloader struct {
	sync.RWMutex

	loaders []LoaderFunc
}

// New creates a new instance of a Reloader.
func New() Reloader {
	return &reloader{}
}

func (r *reloader) Register(f LoaderFunc) {
	r.Lock()
	defer r.Unlock()

	r.loaders = append(r.loaders, f)
}

func (r *reloader) Reload() error {
	r.RLock()
	defer r.RUnlock()

	for _, f := range r.loaders {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}
