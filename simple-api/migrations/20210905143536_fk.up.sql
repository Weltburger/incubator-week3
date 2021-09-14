ALTER TABLE "public"."ticket"
    ADD CONSTRAINT fk_ticket_passenger
        FOREIGN KEY ("passenger_id")
            REFERENCES "public"."passenger" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE "public"."ticket"
    ADD CONSTRAINT fk_ticket_lunch
        FOREIGN KEY ("lunch_id")
            REFERENCES "lunch" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE "public"."ticket"
    ADD CONSTRAINT fk_ticket_class
        FOREIGN KEY ("class_id")
            REFERENCES "public"."class" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE "public"."ticket"
    ADD CONSTRAINT fk_ticket_flight
        FOREIGN KEY ("flight_id")
            REFERENCES "public"."flight" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE "public"."city"
    ADD CONSTRAINT fk_city_country
        FOREIGN KEY ("country_id")
            REFERENCES "public"."country" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE "public"."flight"
    ADD CONSTRAINT fk_flight_city_from
        FOREIGN KEY ("from")
            REFERENCES "public"."city" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

ALTER TABLE "public"."flight"
    ADD CONSTRAINT fk_flight_city_to
        FOREIGN KEY ("to")
            REFERENCES "public"."city" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;

INSERT INTO "lunch" ("id", "name")
VALUES
(1, 'minimum'),
(2, 'medium'),
(3, 'maximum');

INSERT INTO "class" ("id", "name")
VALUES
(1, 'classic'),
(2, 'comfort'),
(3, 'premium');

INSERT INTO "country" ("id", "name")
VALUES
(1, 'Ukraine'),
(2, 'Poland'),
(3, 'Germany'),
(4, 'France'),
(5, 'Spain');

INSERT INTO "passenger" ("id", "name", "age", "email", "phone_number")
VALUES
(1, 'Pakowski', 21, 'weltburger@ua.ua', '+380952940097'),
(2, 'Jeka', 21, 'jeka@ua.ua', '+380509854756'),
(3, 'Miha', 22, 'miha@ua.ua', '+380994986541'),
(4, 'Kate', 20, 'kate@ua.ua', '+380994563214');

INSERT INTO "city" ("id", "country_id", "name")
VALUES
(1, 1, 'Kiev'),
(2, 2, 'Warszawa'),
(3, 3, 'Berlin'),
(4, 4, 'Paris'),
(5, 5, 'Madrid');

INSERT INTO "flight" ("id", "from", "to", "date")
VALUES
(1, 1, 2, '2021-09-11'),
(2, 1, 3, '2021-09-12'),
(3, 1, 4, '2021-09-13'),
(4, 1, 5, '2021-09-14'),
(5, 2, 3, '2021-08-20'),
(6, 4, 3, '2021-08-02');

INSERT INTO "ticket" ("id", "passenger_id", "flight_id", "class_id", "lunch_id", "price", "weight")
VALUES
(1, 1, 1, 3, 2, 5000, 25),
(2, 1, 2, 3, 2, 3500, 20),
(3, 2, 1, 2, 1, 4000, 15),
(4, 3, 2, 3, 3, 6500, 37),
(5, 2, 3, 1, 1, 2000, 7),
(6, 4, 5, 2, 2, 3500, 9),
(7, 4, 6, 1, 3, 2500, 12);

