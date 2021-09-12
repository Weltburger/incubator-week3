-- Выбрать все билеты в пункт Б, за 1 месяц

SELECT tt."id", tt.passenger_id, tt.flight_id, tt.class_id, tt.lunch_id, tt.price 
FROM ticket tt
JOIN (SELECT ft."id" 
	  FROM flight ft
	  JOIN (SELECT cy."id" 
			FROM city cy
			WHERE cy."name"='Berlin') cy ON cy."id" = ft."to"
	 WHERE ft."date" BETWEEN '2021-08-12' AND '2021-09-12') ft ON ft.id = tt.flight_id