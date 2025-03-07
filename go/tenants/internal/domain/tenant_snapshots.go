package domain

type TenantV1 struct {
	Name string
}

func (TenantV1) SnapshotName() string { return "tenants.TenantV1" }
