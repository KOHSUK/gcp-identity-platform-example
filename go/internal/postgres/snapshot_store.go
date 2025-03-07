package postgres

import (
	"app/internal/es"
	"app/internal/registry"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type SnapshotStore struct {
	es.AggregateStore
	tableName string
	db        *pgx.Conn
	registry  registry.Registry
}

var _ es.AggregateStore = (*SnapshotStore)(nil)

func NewSnapshotStore(tableName string, db *pgx.Conn, registry registry.Registry) es.AggregateStoreMiddleware {
	snapshots := SnapshotStore{
		tableName: tableName,
		db:        db,
		registry:  registry,
	}

	return func(store es.AggregateStore) es.AggregateStore {
		snapshots.AggregateStore = store
		return snapshots
	}
}

// Load retrieves the latest snapshot for the given aggregate from the database.
// If no snapshot is found, it falls back to loading events from the aggregate store.
//
// Parameters:
// - ctx: The context for managing request-scoped values, cancellation, and deadlines.
// - aggregate: The event-sourced aggregate to load the snapshot for.
//
// Returns:
// - error: An error if the snapshot or events could not be loaded, otherwise nil.
func (s SnapshotStore) Load(ctx context.Context, aggregate es.EventSourcedAggregate) error {
	const query = `SELECT stream_version, snapshot_name, snapshot_data FROM %s WHERE stream_id =$1 AND stream_name = $2 LIMIT 1`

	var entityVersion int
	var snapshotName string
	var snapshotData []byte

	if err := s.db.QueryRow(ctx, s.table(query), aggregate.ID(), aggregate.AggregateName()).Scan(&entityVersion, &snapshotName, &snapshotData); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s.AggregateStore.Load(ctx, aggregate)
		}
		return err
	}

	v, err := s.registry.Deserialize(snapshotName, snapshotData, registry.ValidateImplements((*es.Snapshot)(nil)))
	if err != nil {
		return err
	}

	if err := es.LoadSnapshot(aggregate, v.(es.Snapshot), entityVersion); err != nil {
		return err
	}

	return s.AggregateStore.Load(ctx, aggregate)
}

func (s SnapshotStore) Save(ctx context.Context, aggregate es.EventSourcedAggregate) error {
	const query = `INSERT INTO %s (stream_id, stream_name, stream_version, snapshot_name, snapshot_data)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (stream_id, stream_name) DO
UPDATE SET stream_version = EXCLUDED.stream_version, snapshot_name = EXCLUDED.snapshot_name, snapshot_data = EXCLUDED.snapshot_data`

	if err := s.AggregateStore.Save(ctx, aggregate); err != nil {
		return err
	}

	if !s.shouldSnapshot(aggregate) {
		return nil
	}

	sser, ok := aggregate.(es.Snapshotter)
	if !ok {
		return fmt.Errorf("%T does not implement es.Snapshotter", aggregate)
	}

	snapshot := sser.ToSnapshot()

	data, err := s.registry.Serialize(snapshot.SnapshotName(), snapshot)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, s.table(query), aggregate.ID(), aggregate.AggregateName(), aggregate.PendingVersion(), snapshot.SnapshotName(), data)

	return err
}

func (SnapshotStore) shouldSnapshot(aggregate es.EventSourcedAggregate) bool {
	var maxChanges = 50
	var pendingVersion = aggregate.PendingVersion()
	var pendingChanges = len(aggregate.Events())

	return pendingVersion >= maxChanges && ((pendingChanges >= maxChanges) || (pendingVersion%maxChanges < pendingChanges) || (pendingVersion%maxChanges == 0))
}

func (s SnapshotStore) table(query string) string {
	return fmt.Sprintf(query, s.tableName)
}
