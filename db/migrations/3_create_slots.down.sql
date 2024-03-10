BEGIN;

DROP TABLE IF EXISTS slots;

DELETE FROM migrations WHERE migration_number = 3;

COMMIT;
