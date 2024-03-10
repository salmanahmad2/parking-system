BEGIN;

CREATE TABLE IF NOT EXISTS book_slot (
    id UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    slot_id UUID NOT NULL,
    vehicle_number VARCHAR(25) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    ended_at TIMESTAMP,
    is_parked Boolean DEFAULT true,
    bill_amount numeric
);

INSERT INTO migrations (migration_number) VALUES (4);

COMMIT;
