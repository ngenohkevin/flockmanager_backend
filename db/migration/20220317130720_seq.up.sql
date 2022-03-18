CREATE TABLE "kuroiler" (
                            "id" SERIAL PRIMARY KEY,
                            "title" varchar UNIQUE NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rainbowrooster" (
                                  "id" SERIAL PRIMARY KEY,
                                  "title" varchar UNIQUE NOT NULL,
                                  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "broiler" (
                           "id" SERIAL PRIMARY KEY,
                           "title" varchar UNIQUE NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "layers" (
                          "id" SERIAL PRIMARY KEY,
                          "title" varchar UNIQUE NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "production" (
                              "production_id" bigserial,
                              "created_at" timestamptz NOT NULL DEFAULT (now()),
                              "eggs" bigint,
                              "dirty" int,
                              "wrong_shape" int,
                              "weak_shell" int,
                              "damaged" int,
                              "hatching_eggs" bigint,
                              "kuroiler_id" int,
                              "rainbowrooster_id" int,
                              "broiler_id" int,
                              "layers_id" int
);

CREATE TABLE "hatchery" (
                            "hatchery_id" bigserial,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "infertile" int,
                            "early" int,
                            "middle" int,
                            "late" int,
                            "dead_chicks" int,
                            "alive_chicks" bigint,
                            "kuroiler_id" int,
                            "rainbowrooster_id" int,
                            "broiler_id" int,
                            "layers_id" int
);

CREATE TABLE "premises" (
                            "premises_id" bigserial,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "farm" varchar,
                            "house" varchar,
                            "kuroiler_id" int,
                            "rainbowrooster_id" int,
                            "broiler_id" int,
                            "layers_id" int
);

ALTER TABLE "production" ADD FOREIGN KEY ("kuroiler_id") REFERENCES "kuroiler" ("id");

ALTER TABLE "production" ADD FOREIGN KEY ("rainbowrooster_id") REFERENCES "rainbowrooster" ("id");

ALTER TABLE "production" ADD FOREIGN KEY ("broiler_id") REFERENCES "broiler" ("id");

ALTER TABLE "production" ADD FOREIGN KEY ("layers_id") REFERENCES "layers" ("id");

ALTER TABLE "hatchery" ADD FOREIGN KEY ("kuroiler_id") REFERENCES "kuroiler" ("id");

ALTER TABLE "hatchery" ADD FOREIGN KEY ("rainbowrooster_id") REFERENCES "rainbowrooster" ("id");

ALTER TABLE "hatchery" ADD FOREIGN KEY ("broiler_id") REFERENCES "broiler" ("id");

ALTER TABLE "hatchery" ADD FOREIGN KEY ("layers_id") REFERENCES "layers" ("id");

ALTER TABLE "premises" ADD FOREIGN KEY ("kuroiler_id") REFERENCES "kuroiler" ("id");

ALTER TABLE "premises" ADD FOREIGN KEY ("rainbowrooster_id") REFERENCES "rainbowrooster" ("id");

ALTER TABLE "premises" ADD FOREIGN KEY ("broiler_id") REFERENCES "broiler" ("id");

ALTER TABLE "premises" ADD FOREIGN KEY ("layers_id") REFERENCES "layers" ("id");

CREATE INDEX ON "production" ("production_id");

CREATE INDEX ON "production" ("hatching_eggs");

CREATE INDEX ON "hatchery" ("hatchery_id");

CREATE INDEX ON "hatchery" ("alive_chicks");