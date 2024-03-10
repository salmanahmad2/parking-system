BEGIN;

DROP TABLE IF EXISTS book_slot;

DELETE FROM migrations WHERE migration_number = 4;

COMMIT;
