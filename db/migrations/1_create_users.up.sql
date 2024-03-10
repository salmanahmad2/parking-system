BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS migrations (id SERIAL PRIMARY KEY, migration_number INT UNIQUE, applied_at TIMESTAMP DEFAULT NOW());

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    role VARCHAR(50) NOT NULL,
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
BEFORE UPDATE ON users 
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();


INSERT INTO migrations (migration_number) VALUES (1);

COMMIT;
