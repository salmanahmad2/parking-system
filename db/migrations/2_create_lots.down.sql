BEGIN;

DROP TABLE IF EXISTS lots;

DELETE FROM migrations WHERE migration_number = 2;

COMMIT;
