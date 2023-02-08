CREATE TABLE "ingredients" (
  "ingredient_id" bigserial PRIMARY KEY,
  "ingredient_name" varchar(250) NOT NULL
);

CREATE TABLE "measurement_qty" (
  "measurement_qty_id" bigserial PRIMARY KEY,
  "qty_amount" bigint
);

CREATE TABLE "measurement_units" (
  "measurement_units_id" bigserial PRIMARY KEY,
  "unit" varchar(250)
);

CREATE TABLE "recipe" (
  "recipe_id" bigserial PRIMARY KEY,
  "drink_name" varchar(250) NOT NULL,
  "instructions" varchar(250) NOT NULL,
  "image_url" varchar(250),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "cocktail" (
  "cocktail_id" bigserial PRIMARY KEY,
  "recipe_id" bigint NOT NULL,
  "ingredient_id" bigint NOT NULL,
  "measurement_qty_id" bigint DEFAULT null,
  "measurement_units_id" bigint DEFAULT null
);

ALTER TABLE "cocktail" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipe" ("recipe_id");

ALTER TABLE "cocktail" ADD FOREIGN KEY ("ingredient_id") REFERENCES "ingredients" ("ingredient_id");

ALTER TABLE "cocktail" ADD FOREIGN KEY ("measurement_qty_id") REFERENCES "measurement_qty" ("measurement_qty_id");

ALTER TABLE "cocktail" ADD FOREIGN KEY ("measurement_units_id") REFERENCES "measurement_units" ("measurement_units_id");