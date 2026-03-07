CREATE TABLE IF NOT EXISTS endpoint_visits (
    endpoint TEXT PRIMARY KEY,
    visits INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS totals (
    id INTEGER PRIMARY KEY CHECK (id = 1),
    total_visits INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS visitors_daily (
    day TEXT NOT NULL,
    ip_hash TEXT NOT NULL,
    PRIMARY KEY(day, ip_hash)
);

INSERT OR IGNORE INTO totals (id, total_visits) VALUES (1, 0);