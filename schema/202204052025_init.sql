-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS public.account
(
    id            SERIAL PRIMARY KEY,
    created_at    timestamp without time zone NOT NULL DEFAULT NOW(),
    user_id       bigint                      NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS public.post
(
    id          bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    title       varchar(255)                not null,
    description varchar(1000),
    created_at  timestamp without time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp without time zone NOT NULL DEFAULT NOW(),
    account_id  bigint REFERENCES public.account (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS public.comment
(
    id           bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    comment_text varchar(500)                not null,
    created_at   timestamp without time zone NOT NULL DEFAULT NOW(),
    updated_at   timestamp without time zone NOT NULL DEFAULT NOW(),
    post_id      bigint REFERENCES public.post (id) ON DELETE CASCADE,
    account_id      bigint                      NOT NULL
);

CREATE TABLE IF NOT EXISTS public.tag
(
    id    bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    title varchar(100) not null
);

CREATE TABLE IF NOT EXISTS public.category
(
    id    bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    title varchar(100) not null unique
);

CREATE TABLE IF NOT EXISTS image
(
    id      bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    link    varchar(500) not null,
    post_id bigint REFERENCES public.post (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posts_have_tags
(
    post_id int REFERENCES public.post (id) ON DELETE CASCADE,
    tag_id  int REFERENCES public.tag (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS posts_have_categories
(
    post_id     int REFERENCES public.post (id) ON DELETE CASCADE,
    category_id int REFERENCES public.category (id) ON DELETE CASCADE
);

INSERT INTO category(title)
VALUES ('none'),
       ('web-design'),
       ('graphic'),
       ('photography'),
       ('sketch');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS posts_have_categories CASCADE;
DROP TABLE IF EXISTS posts_have_tags CASCADE;
DROP TABLE IF EXISTS category CASCADE;
DROP TABLE IF EXISTS category CASCADE;
DROP TABLE IF EXISTS image CASCADE;
DROP TABLE IF EXISTS comment CASCADE;
DROP TABLE IF EXISTS post CASCADE;
DROP TABLE IF EXISTS account CASCADE;