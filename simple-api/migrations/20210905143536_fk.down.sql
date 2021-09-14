ALTER TABLE "public"."ticket"
    DROP CONSTRAINT fk_ticket_passenger;

ALTER TABLE "public"."ticket"
    DROP CONSTRAINT fk_ticket_lunch;

ALTER TABLE "public"."ticket"
    DROP CONSTRAINT fk_ticket_class;

ALTER TABLE "public"."ticket"
    DROP CONSTRAINT fk_ticket_flight;

ALTER TABLE "public"."city"
    DROP CONSTRAINT fk_city_country;

ALTER TABLE "public"."flight"
    DROP CONSTRAINT fk_flight_city_from;

ALTER TABLE "public"."flight"
    DROP CONSTRAINT fk_flight_city_to;
