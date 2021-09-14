CREATE TABLE IF NOT EXISTS "ticket" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "passenger_id"      int8 NOT NULL,
    "flight_id"         int8 NOT NULL,
    "class_id"          int8 NOT NULL,
    "lunch_id"          int8 NOT NULL,
    "weight"            int8 NOT NULL,
    "price"             numeric(15,6) not null default 0::numeric
);

CREATE TABLE IF NOT EXISTS "passenger" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL,
    "age"               int2 NOT NULL,
    "email"             VARCHAR (100) NOT NULL,
    "phone_number"      VARCHAR (18)
);

CREATE TABLE IF NOT EXISTS "flight" (
    "id"                 int8 NOT NULL PRIMARY KEY,
    "from"               int8 NOT NULL,
    "to"                 int8 NOT NULL,
    "date"               date not null
);

CREATE TABLE IF NOT EXISTS "country" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "city" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "country_id"        int8 NOT NULL,
    "name"              VARCHAR (100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "class" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);

CREATE TABLE IF NOT EXISTS "lunch" (
    "id"                int8 NOT NULL PRIMARY KEY,
    "name"              VARCHAR (100) NOT NULL
);
