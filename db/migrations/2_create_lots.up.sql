BEGIN;

CREATE TABLE IF NOT EXISTS lots (
    id UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    name VARCHAR(50) NOT NULL,
    address VARCHAR(200),
    hourly_rate numeric NOT NULL,
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
BEFORE UPDATE ON lots 
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();


INSERT INTO migrations (migration_number) VALUES (2);

COMMIT;
