CREATE TABLE "users"
(
    "id"         SERIAL UNIQUE PRIMARY KEY,
    "username"   varchar UNIQUE NOT NULL,
    "password"   varchar        NOT NULL,
    "active"     boolean        NOT NULL DEFAULT false,
    "created_at" timestamp      NOT NULL,
    "updated_at" timestamp      NOT NULL
);

CREATE TABLE "sessions"
(
    "id"                          SERIAL UNIQUE PRIMARY KEY,
    "user_id"                     int            NOT NULL,
    "session_id"                  varchar UNIQUE NOT NULL,
    "access_token"                varchar UNIQUE NOT NULL,
    "access_token_expiration_at"  timestamp      NOT NULL,
    "refresh_token"               varchar        NOT NULL,
    "refresh_token_expiration_at" timestamp      NOT NULL,
    "ip"                          inet           NOT NULL,
    "agent"                       varchar        NOT NULL,
    "created_at"                  timestamp      NOT NULL,
    "last_activity_at"            timestamp      NOT NULL
);

CREATE TABLE "user_attributes"
(
    "id"      SERIAL PRIMARY KEY NOT NULL,
    "user_id" int                NOT NULL,
    "key"     varchar            NOT NULL,
    "value"   varchar            NOT NULL
);

ALTER TABLE "sessions"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "user_attributes"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX "users_username_index" ON "users" ("username");

CREATE INDEX "users_password_index" ON "users" ("password");

CREATE UNIQUE INDEX "sessions_uid_sid_index" ON "sessions" ("user_id", "session_id");

COMMENT ON COLUMN "users"."username" IS 'имя пользователя';

COMMENT ON COLUMN "users"."password" IS 'пароль пользователя';

COMMENT ON COLUMN "users"."active" IS 'статус активности учетной записи';

COMMENT ON COLUMN "users"."created_at" IS 'дата создания';

COMMENT ON COLUMN "users"."updated_at" IS 'дата изменения';

COMMENT ON COLUMN "sessions"."user_id" IS 'ID пользователя сессии';

COMMENT ON COLUMN "sessions"."session_id" IS 'идентификатор сессии пользователя';

COMMENT ON COLUMN "sessions"."access_token" IS 'токен доступа';

COMMENT ON COLUMN "sessions"."access_token_expiration_at" IS 'дата истечения срока действия токена доступа';

COMMENT ON COLUMN "sessions"."refresh_token" IS 'тоекн обновления токена доступа';

COMMENT ON COLUMN "sessions"."refresh_token_expiration_at" IS 'дата истечения срока действия токена обновления';

COMMENT ON COLUMN "sessions"."ip" IS 'IP клиента';

COMMENT ON COLUMN "sessions"."agent" IS 'Агент пользователя';

COMMENT ON COLUMN "sessions"."created_at" IS 'дата создания сессии';

COMMENT ON COLUMN "sessions"."last_activity_at" IS 'дата последней активности';

COMMENT ON COLUMN "user_attributes"."user_id" IS 'ID пользователя';

COMMENT ON COLUMN "user_attributes"."key" IS 'ключь атрибута';

COMMENT ON COLUMN "user_attributes"."value" IS 'значение атрибута';
