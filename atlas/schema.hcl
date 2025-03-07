schema "tenants" {
  comment = "standard public schema"
}

table "tenants" {
  schema = schema.tenants
  column "id" {
    null = false
    type = text
  }
  column "name" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.id]
  }
}


table "events" {
  schema = schema.tenants
  column "stream_id" {
    null = false
    type = text
  }
  column "stream_name" {
    null = false
    type = text
  }
  column "stream_version" {
    null = false
    type = int
  }
  column "event_id" {
    null = false
    type = text
  }
  column "event_name" {
    null = false
    type = text
  }
  column "event_data" {
    null = false
    type = bytea
  }
  column "occurred_at" {
    null = false
    type = timestamptz
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.stream_id, column.stream_name, column.stream_version]
  }
}

table "snapshots" {
  schema = schema.tenants
  column "stream_id" {
    null = false
    type = text
  }
  column "stream_name" {
    null = false
    type = text
  }
  column "stream_version" {
    null = false
    type = int
  }
  column "snapshot_name" {
    null = false
    type = text
  }
  column "snapshot_data" {
    null = false
    type = bytea
  }
  column "updated_at" {
    null = false
    type = timestamptz
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.stream_id, column.stream_name]
  }
}