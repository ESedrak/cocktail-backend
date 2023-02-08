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

CREATE TABLE "cocktail" (
  "cocktail_id" bigserial PRIMARY KEY,
  "drink_name" varchar(250) NOT NULL,
  "instructions" varchar(250) NOT NULL,
  "image_url" varchar(250),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ingredientsToCocktails" (
  "cocktail_id" bigint PRIMARY KEY NOT NULL,
  "ingredient_id" bigint NOT NULL,
  "measurement_qty_id" bigint,
  "measurement_units_id" bigint
);

ALTER TABLE "ingredientsToCocktails" ADD FOREIGN KEY ("cocktail_id") REFERENCES "cocktail" ("cocktail_id");

ALTER TABLE "ingredientsToCocktails" ADD FOREIGN KEY ("ingredient_id") REFERENCES "ingredients" ("ingredient_id");

ALTER TABLE "ingredientsToCocktails" ADD FOREIGN KEY ("measurement_qty_id") REFERENCES "measurement_qty" ("measurement_qty_id");

ALTER TABLE "ingredientsToCocktails" ADD FOREIGN KEY ("measurement_units_id") REFERENCES "measurement_units" ("measurement_units_id");