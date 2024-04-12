-- +goose Up
-- +goose StatementBegin
CREATE TABLE "average_tax_rate_submissions"
(
    "id"               bigint GENERATED BY DEFAULT AS IDENTITY (INCREMENT 1 MINVALUE 0 START 0),
    "epoch_id"         bigint      NOT NULL,
    "transaction_hash" text        NOT NULL,
    "average_tax_rate" decimal     NOT NULL,
    "created_at"       timestamptz NOT NULL DEFAULT now(),
    "updated_at"       timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT "pkey" PRIMARY KEY ("epoch_id")
);

CREATE INDEX "average_tax_rate_submissions_epoch_id_idx" ON "average_tax_rate_submissions" ("epoch_id" DESC);
CREATE INDEX "average_tax_rate_submissions_id_idx" ON "average_tax_rate_submissions" ("id" DESC);
CREATE INDEX "average_tax_rate_submissions_transaction_hash_idx" ON "average_tax_rate_submissions" ("transaction_hash");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "average_tax_rate_submissions";
-- +goose StatementEnd