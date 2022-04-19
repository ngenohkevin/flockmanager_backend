CREATE TABLE "kuroiler" (
                            "id" BIGSERIAL PRIMARY KEY,
                            "title" varchar UNIQUE NOT NULL,
                            "house" varchar NOT NULL,
                            "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "production" (
                              "id" bigserial PRIMARY KEY,
                              "production_id" bigserial NOT NULL,
                              "created_at" timestamp NOT NULL DEFAULT (now()),
                              "eggs" bigint,
                              "dirty" bigint,
                              "wrong_shape" bigint,
                              "weak_shell" bigint,
                              "damaged" bigint,
                              "hatching_eggs" bigint
);

CREATE TABLE "hatchery" (
                            "id" bigserial PRIMARY KEY,
                            "hatchery_id" bigserial NOT NULL,
                            "created_at" timestamp NOT NULL DEFAULT (now()),
                            "infertile" bigint,
                            "early" bigint,
                            "middle" bigint,
                            "late" bigint,
                            "dead_chicks" bigint,
                            "alive_chicks" bigint
);

ALTER TABLE "production" ADD FOREIGN KEY ("production_id") REFERENCES "kuroiler" ("id");

ALTER TABLE "hatchery" ADD FOREIGN KEY ("hatchery_id") REFERENCES "kuroiler" ("id");

CREATE INDEX ON "kuroiler" ("title");

CREATE INDEX ON "production" ("production_id");

CREATE INDEX ON "hatchery" ("hatchery_id");