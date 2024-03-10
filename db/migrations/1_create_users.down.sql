BEGIN;

DROP TABLE IF EXISTS users;

DELETE FROM migrations WHERE migration_number = 1;

COMMIT;
