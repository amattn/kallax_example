BEGIN;

ALTER TABLE users ADD COLUMN company_id bigint REFERENCES companies(id);

COMMIT;
