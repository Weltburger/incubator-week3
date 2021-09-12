-- Посчитать количество пассажиров, которые заказывали больше одного билета

SELECT COUNT(*) FROM (SELECT COUNT(tt.passenger_id)
FROM ticket as tt
GROUP BY tt.passenger_id
HAVING COUNT(tt.passenger_id) > 1) as res