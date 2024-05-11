CREATE SCHEMA auth;
CREATE SCHEMA profiles;

CREATE TABLE auth.users
(
    id         uuid,
    email      TEXT UNIQUE NOT NULL,
    password   TEXT,
    PRIMARY KEY (id)
);

CREATE TYPE HABIT AS ENUM ('no', 'sometimes', 'yes', 'prefer not to say');
CREATE TYPE INTENTION AS ENUM (
    'life partner',
    'long-term relationship',
    'short-term relationship',
    'friendship',
    'figuring it out',
    'prefer not to say'
    );
CREATE TYPE GENDER AS ENUM ('man', 'woman');
CREATE TYPE PREFERRED_PARTNER AS ENUM ('man', 'woman', 'anyone');
CREATE TYPE FAMILY_PLANS AS ENUM ('don''t want children', 'want children', 'open to children', 'not sure yet', 'prefer not to say');

CREATE TABLE profiles.profiles
(
    user_id           uuid,
    first_name        TEXT,
    last_name         TEXT,
    birth_date        DATE,
    sex               GENDER,
    preferred_partner PREFERRED_PARTNER,
    intention         INTENTION,
    height            INTEGER,
    has_children      BOOLEAN,
    family_plans      TEXT,
    location          TEXT,
    drinks_alcohol    HABIT,
    smokes            HABIT,
    PRIMARY KEY (user_id)
);

CREATE TYPE PROMPT_TYPE AS ENUM ('text', 'image');

CREATE TABLE profiles.prompts
(
    id       uuid,
    user_id  uuid,
    position INTEGER,
    question TEXT,
    content  TEXT,
    type     PROMPT_TYPE,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES profiles.profiles (user_id)
);

ALTER TYPE FAMILY_PLANS RENAME VALUE 'don''t want children' TO 'do not want children';

ALTER TABLE profiles.prompts
    ADD CONSTRAINT unique_prompt_pair UNIQUE (user_id, question, type);

ALTER TABLE profiles.prompts
    DROP CONSTRAINT prompts_user_id_fkey;