package response

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
	"github.com/oklog/run"
	"github.com/responserms/response/internal/config"
	"github.com/responserms/response/internal/ent"
	"github.com/rs/zerolog"
)

const (
	VersionTag    = ""
	VersionCommit = ""
)

type InitializationFunc func(ctx context.Context) error

type Core struct {
	config    *config.Config
	startTime time.Time

	services  run.Group
	ent       *ent.Client
	logger    zerolog.Logger
	validator *validator.Validate

	// initialization
	initFuncs      []InitializationFunc
	initFuncsMutex sync.RWMutex
	initOnce       sync.Once
}

func New(cfg *config.Config) (*Core, error) {
	if cfg == nil {
		return nil, errors.New("provided Config is nil")
	}

	if !cfg.Valid() {
		return nil, errors.New("provided Config has not been properly initialized")
	}

	core := &Core{
		config:    cfg,
		logger:    zerolog.New(os.Stderr),
		initFuncs: []InitializationFunc{},
		validator: validator.New(),
	}

	core.registerInitializationFunc(core.initEnt)

	return core, nil
}

// registerInitializationFunc registers a function to be called during the initialization phase of Response.
// This allows a standardized way to ensure Response is fully-initialized and all components are able to setup
// their particular piece.
func (c *Core) registerInitializationFunc(fn InitializationFunc) {
	c.initFuncsMutex.Lock()
	defer c.initFuncsMutex.Unlock()

	c.initFuncs = append(c.initFuncs, fn)
}

func (c *Core) init(ctx context.Context) error {
	var initErr multierror.Error

	c.initOnce.Do(func() {
		c.initFuncsMutex.RLock()
		defer c.initFuncsMutex.RUnlock()

		for _, initFunc := range c.initFuncs {
			if err := initFunc(ctx); err != nil {
				initErr.Errors = append(initErr.Errors, err)
			}
		}
	})

	return initErr.ErrorOrNil()
}

func (c *Core) RunningConfig() *config.Config {
	return c.config
}

func (c *Core) Logger(name string) zerolog.Logger {
	return c.logger.With().
		Str("service", name).
		Logger()
}

func (c *Core) Init(ctx context.Context) error {
	return c.init(ctx)
}

func (c *Core) RegisterMonitoredService(run func() error, int func(error)) {
	c.services.Add(run, int)
}

func (c *Core) Run(ctx context.Context) error {
	if !c.startTime.IsZero() {
		return errors.New("call to Run has already happened, Run exiting should terminate the program so that Core is not reused")
	}

	if err := c.init(ctx); err != nil {
		return fmt.Errorf("init error: %w", err)
	}

	changes, err := c.HasSchemaChanges(ctx)
	if err != nil {
		return fmt.Errorf("unable to check for schema changes: %w", err)
	}

	if changes {
		return ErrSchemaChangesAvailable
	}

	//if err := c.cacheSettings(ctx); err != nil {
	//	return fmt.Errorf("cacheSettings: %w", err)
	//}

	c.startTime = time.Now()
	c.services.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))

	return c.services.Run()
}

func (c *Core) ApplyContext(ctx context.Context) context.Context {
	return ent.NewContext(ctx, c.ent)
}
