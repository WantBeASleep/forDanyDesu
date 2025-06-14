CREATE TABLE segment
(
    id        uuid PRIMARY KEY,
    node_id   uuid NOT NULL REFERENCES node (id),
    image_id  uuid NOT NULL REFERENCES image (id),
    contor    text NOT NULL,
    tirads_23 real NOT NULL,
    tirads_4  real NOT NULL,
    tirads_5  real NOT NULL
);

COMMENT ON TABLE segment IS 'Хранилище сегментов в узи';
COMMENT ON COLUMN segment.contor IS 'контур узла (JSON)';
COMMENT ON COLUMN segment.tirads_23 IS 'процент отношения к классу tirads_23';
COMMENT ON COLUMN segment.tirads_4 IS 'процент отношения к классу tirads_4';
COMMENT ON COLUMN segment.tirads_5 IS 'процент отношения к классу tirads_5';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS node CASCADE;
DROP TABLE IF EXISTS segment CASCADE;
DROP TABLE IF EXISTS image CASCADE;
DROP TABLE IF EXISTS device CASCADE;
DROP TABLE IF EXISTS uzi CASCADE;
-- +goose StatementEnd