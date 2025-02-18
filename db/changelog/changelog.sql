-- liquibase formatted sql

-- changeset azharrohiman:1
CREATE TABLE "TB_CUSTOMER" (
    "ID" UUID NOT NULL,
    "USERNAME" VARCHAR(255) NOT NULL,
    "PASSWORD_HASH" BYTEA NOT NULL,
    "PHONE_NUM" VARCHAR(15) NOT NULL,
    CONSTRAINT "pk_tb_customer" PRIMARY KEY ("ID"),
    CONSTRAINT "uq_tb_customer_username" UNIQUE ("USERNAME"),
    CONSTRAINT "uq_tb_customer_phone_num" UNIQUE ("PHONE_NUM")
);

-- changeset azharrohiman:2
CREATE TABLE "TB_ACCOUNT" (
    "ID" UUID NOT NULL,
    "ACCOUNT_NUM" VARCHAR(255) NOT NULL,
    "BALANCE" NUMERIC(19, 2) NOT NULL,
    "CUSTOMER_ID" UUID NOT NULL,
    "IS_MAIN_ACCOUNT" BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT "pk_tb_account" PRIMARY KEY ("ID"),
    CONSTRAINT "fk_tb_account_customer" FOREIGN KEY ("CUSTOMER_ID") REFERENCES "TB_CUSTOMER"("ID"),
    CONSTRAINT "uq_tb_account_account_num" UNIQUE ("ACCOUNT_NUM")
);

-- changeset azharrohiman:3
CREATE INDEX "idx_tb_account_customer_id" ON "TB_ACCOUNT"("CUSTOMER_ID");