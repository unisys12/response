// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// DeathCertificate is the client for interacting with the DeathCertificate builders.
	DeathCertificate *DeathCertificateClient
	// DeathCertifier is the client for interacting with the DeathCertifier builders.
	DeathCertifier *DeathCertifierClient
	// DeathManner is the client for interacting with the DeathManner builders.
	DeathManner *DeathMannerClient
	// DeathPlace is the client for interacting with the DeathPlace builders.
	DeathPlace *DeathPlaceClient
	// Ethnicity is the client for interacting with the Ethnicity builders.
	Ethnicity *EthnicityClient
	// GameServer is the client for interacting with the GameServer builders.
	GameServer *GameServerClient
	// Metadata is the client for interacting with the Metadata builders.
	Metadata *MetadataClient
	// OAuthConnection is the client for interacting with the OAuthConnection builders.
	OAuthConnection *OAuthConnectionClient
	// Person is the client for interacting with the Person builders.
	Person *PersonClient
	// Player is the client for interacting with the Player builders.
	Player *PlayerClient
	// PlayerIdentifier is the client for interacting with the PlayerIdentifier builders.
	PlayerIdentifier *PlayerIdentifierClient
	// Race is the client for interacting with the Race builders.
	Race *RaceClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// Setting is the client for interacting with the Setting builders.
	Setting *SettingClient
	// Sex is the client for interacting with the Sex builders.
	Sex *SexClient
	// State is the client for interacting with the State builders.
	State *StateClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Vehicle is the client for interacting with the Vehicle builders.
	Vehicle *VehicleClient
	// VehicleClass is the client for interacting with the VehicleClass builders.
	VehicleClass *VehicleClassClient
	// VehicleColor is the client for interacting with the VehicleColor builders.
	VehicleColor *VehicleColorClient
	// VehicleMake is the client for interacting with the VehicleMake builders.
	VehicleMake *VehicleMakeClient
	// VehicleModel is the client for interacting with the VehicleModel builders.
	VehicleModel *VehicleModelClient
	// VehicleRegistration is the client for interacting with the VehicleRegistration builders.
	VehicleRegistration *VehicleRegistrationClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Committer method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollbacker method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.DeathCertificate = NewDeathCertificateClient(tx.config)
	tx.DeathCertifier = NewDeathCertifierClient(tx.config)
	tx.DeathManner = NewDeathMannerClient(tx.config)
	tx.DeathPlace = NewDeathPlaceClient(tx.config)
	tx.Ethnicity = NewEthnicityClient(tx.config)
	tx.GameServer = NewGameServerClient(tx.config)
	tx.Metadata = NewMetadataClient(tx.config)
	tx.OAuthConnection = NewOAuthConnectionClient(tx.config)
	tx.Person = NewPersonClient(tx.config)
	tx.Player = NewPlayerClient(tx.config)
	tx.PlayerIdentifier = NewPlayerIdentifierClient(tx.config)
	tx.Race = NewRaceClient(tx.config)
	tx.Session = NewSessionClient(tx.config)
	tx.Setting = NewSettingClient(tx.config)
	tx.Sex = NewSexClient(tx.config)
	tx.State = NewStateClient(tx.config)
	tx.User = NewUserClient(tx.config)
	tx.Vehicle = NewVehicleClient(tx.config)
	tx.VehicleClass = NewVehicleClassClient(tx.config)
	tx.VehicleColor = NewVehicleColorClient(tx.config)
	tx.VehicleMake = NewVehicleMakeClient(tx.config)
	tx.VehicleModel = NewVehicleModelClient(tx.config)
	tx.VehicleRegistration = NewVehicleRegistrationClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: DeathCertificate.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
