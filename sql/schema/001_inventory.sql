-- +goose Up
-- Dependancy Free Tables
CREATE TABLE "suppliers" (
  "supplier_id" integer PRIMARY KEY,
  "supplier_name" varchar NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "mail" varchar UNIQUE NOT NULL
);

CREATE TABLE "addresses" (
  "address_id" integer PRIMARY KEY,
  "latitude" float  NOT NULL,
  "longitude" float  NOT NULL
);

-- Single Dependancy Tables

CREATE TABLE "products" (
  "product_id" integer PRIMARY KEY,
  "product_name" varchar  NOT NULL,
  "description" text,
  "unit_price" float  NOT NULL CONSTRAINT postive_unit_price CHECK (unit_price >= 0),
  "supplier_id" integer NOT NULL REFERENCES suppliers ON DELETE RESTRICT
);

CREATE TABLE "locations" (
  "location_id" integer PRIMARY KEY,
  "location_name" varchar  NOT NULL,
  "address_id" integer  NOT NULL REFERENCES addresses ON DELETE RESTRICT
);

-- Two Dependancy Tables

CREATE TABLE "stocks" (
  "product_id" integer REFERENCES products ON DELETE RESTRICT,
  "location_id" integer REFERENCES locations ON DELETE RESTRICT,
  "quantity" integer  NOT NULL CONSTRAINT postive_quantity CHECK (quantity >= 0),
  PRIMARY KEY ("product_id", "location_id")
);

-- +goose Down
DROP TABLE stocks;
DROP TABLE locations;
DROP TABLE products;
DROP TABLE addresses;
DROP TABLE suppliers;