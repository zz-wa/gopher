-- +goose Up
-- +goose StatementBegin
CREATE  TABLE  url_Shorter(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    OriginalURL TEXT ,
    ShortURL  TEXT UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE url_Shorter
-- +goose StatementEnd
