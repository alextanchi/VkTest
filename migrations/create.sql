CREATE TABLE movie
(
    id           uuid          NOT NULL,
    title        varchar(150)  NOT NULL,
    description  varchar(1000) NOT NULL,
    release_date date          NOT NULL,
    rating       FLOAT         NOT NULL,



    CONSTRAINT movies_pkey PRIMARY KEY (id)

);

CREATE TABLE actor
(
    id         uuid    NOT NULL,
    name       varchar NOT NULL,
    surname    varchar NOT NULL,
    sex        varchar NOT NULL,
    birth_date date    NOT NULL,

    CONSTRAINT actors_pkey PRIMARY KEY (id)

);
CREATE TABLE movie_actor
(
    movie_id uuid REFERENCES movie (id) ON DELETE CASCADE,
    actor_id uuid REFERENCES actor (id) ON DELETE CASCADE,
    PRIMARY KEY (movie_id, actor_id)
);

CREATE TABLE role
(
    id   uuid    NOT NULL,
    role varchar NOT NULL,

    CONSTRAINT roles_pkey PRIMARY KEY (id)

);

CREATE TABLE account
(
    id       uuid    NOT NULL,
    login    varchar NOT NULL,
    password varchar NOT NULL,
    role_id  uuid    NOT NULL,
    CONSTRAINT accounts_pkey PRIMARY KEY (id),
    CONSTRAINT account_roleid_fk FOREIGN KEY (role_id)
        REFERENCES role (id) ON DELETE RESTRICT ON UPDATE RESTRICT

);












