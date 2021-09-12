-- Высчитать общую массу перевезенного багажа за месяц

SELECT SUM(tt.weight) 
FROM ticket tt
JOIN (SELECT ft."id" 
			FROM flight ft
			WHERE ft."date" BETWEEN '2021-08-12' AND '2021-09-12') ft ON ft."id" = tt.flight_id