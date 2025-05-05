CREATE TABLE IF NOT EXISTS appointment (
  id TEXT NOT NULL PRIMARY KEY,
  user_id TEXT NOT NULL,
  service_id TEXT NOT NULL,
  team_id TEXT NOT NULL,
  start_time TEXT NOT NULL, -- ISO 8601 
  status TEXT NOT NULL,
  canceled_reason TEXT,
  reschedule_count INTEGER not NULL DEFAULT 0,
  priority BOOLEAN not NULL DEFAULT 0,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);
