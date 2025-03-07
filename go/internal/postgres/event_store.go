package postgres

import (
	"app/internal/ddd"
	"app/internal/errors"
	"app/internal/es"
	"app/internal/registry"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type (
	EventStore struct {
		tableName string
		db        *pgx.Conn
		registry  registry.Registry
	}
	aggregateEvent struct {
		id         string
		name       string
		payload    ddd.EventPayload
		occurredAt time.Time
		aggregate  es.EventSourcedAggregate
		version    int
	}
)

var _ es.AggregateStore = (*EventStore)(nil)
var _ ddd.AggregateEvent = (*aggregateEvent)(nil)

func NewEventStore(tableName string, db *pgx.Conn, registry registry.Registry) EventStore {
	return EventStore{
		tableName: tableName,
		db:        db,
		registry:  registry,
	}
}

// TODO: ORMを利用する
func (s EventStore) Load(ctx context.Context, aggregate es.EventSourcedAggregate) (err error) {
	const query = `SELECT stream_version, event_id, event_name, event_data, occurred_at FROM %s WHERE stream_id = $1 AND stream_name = $2 AND stream_version > $3 ORDER BY stream_version ASC`

	aggregateID := aggregate.ID()
	aggregateName := aggregate.AggregateName()

	rows, err := s.db.Query(ctx, s.table(query), aggregateID, aggregateName, aggregate.Version())
	if err != nil {
		return err
	}
	defer rows.Close()

	_, err = pgx.ForEachRow(rows, []any{new(int), new(string), new(string), new([]byte), new(time.Time)}, func() error {
		var aggregateVersion int
		var eventID, eventName string
		var payloadData []byte
		var occurredAt time.Time

		err := rows.Scan(&aggregateVersion, &eventID, &eventName, &payloadData, &occurredAt)
		if err != nil {
			return err
		}

		var payload any
		payload, err = s.registry.Deserialize(eventName, payloadData)
		if err != nil {
			return err
		}

		event := aggregateEvent{
			id:         eventID,
			name:       eventName,
			payload:    payload,
			aggregate:  aggregate,
			version:    aggregateVersion,
			occurredAt: occurredAt,
		}

		return es.LoadEvent(aggregate, event)
	})

	return err
}

func (s EventStore) Save(ctx context.Context, aggregate es.EventSourcedAggregate) (err error) {
	const query = `INSERT INTO %s (stream_id, stream_name, stream_version, event_id, event_name, event_data, occurred_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	tx, err := s.db.Begin(ctx)

	if err != nil {
		return err
	}
	defer func() {
		p := recover()
		switch {
		case p != nil:
			_ = tx.Rollback(ctx)
			panic(p)
		case err != nil:
			rErr := tx.Rollback(ctx)
			if rErr != nil {
				err = errors.Wrap(err, rErr.Error())
			}
		default:
			err = tx.Commit(ctx)
		}
	}()

	aggregateID := aggregate.ID()
	aggregateName := aggregate.AggregateName()

	for _, event := range aggregate.Events() {
		var payloadData []byte

		payloadData, err = s.registry.Serialize(event.EventName(), event.Payload())
		if err != nil {
			return err
		}
		if _, err = tx.Exec(
			ctx, s.table(query), aggregateID, aggregateName, event.AggregateVersion(), event.ID(), event.EventName(), payloadData, event.OccurredAt(),
		); err != nil {
			return err
		}
	}

	return nil
}

func (s EventStore) table(query string) string {
	return fmt.Sprintf(query, s.tableName)
}

func (e aggregateEvent) ID() string                { return e.id }
func (e aggregateEvent) EventName() string         { return e.name }
func (e aggregateEvent) Payload() ddd.EventPayload { return e.payload }
func (e aggregateEvent) Metadata() ddd.Metadata    { return ddd.Metadata{} }
func (e aggregateEvent) OccurredAt() time.Time     { return e.occurredAt }
func (e aggregateEvent) AggregateName() string     { return e.aggregate.AggregateName() }
func (e aggregateEvent) AggregateID() string       { return e.aggregate.ID() }
func (e aggregateEvent) AggregateVersion() int     { return e.version }
