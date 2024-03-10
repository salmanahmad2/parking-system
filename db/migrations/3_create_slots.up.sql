BEGIN;

CREATE TABLE IF NOT EXISTS slots (
    id UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    lot_id UUID NOT NULL,
    number INTEGER NOT NULL,
    status VARCHAR(15) NOT NULL DEFAULT 'available',
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);

-- Create a trigger to update updated_at column on every update
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON slots 
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();


INSERT INTO migrations (migration_number) VALUES (3);

COMMIT;
