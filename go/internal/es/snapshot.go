package es

import "fmt"

type Snapshot interface {
	SnapshotName() string
}

type SnapshotApplier interface {
	ApplySnapshot(snapshot Snapshot) error
}

type Snapshotter interface {
	SnapshotApplier
	ToSnapshot() Snapshot
}

func LoadSnapshot(v any, snapshot Snapshot, version int) error {
	type loader interface {
		SnapshotApplier
		VersionSetter
	}

	agg, ok := v.(loader)
	if !ok {
		return fmt.Errorf("%T does not have the methods implemented to load snapshot", v)
	}

	if err := agg.ApplySnapshot(snapshot); err != nil {
		return err
	}
	agg.setVersion(version)

	return nil
}
