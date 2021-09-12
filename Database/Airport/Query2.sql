-- Выбрать всех пассажиров, у которых багаж превышал 10 кг

SELECT DISTINCT pr."id", pr."name", pr.age, pr.email, pr.phone_number 
FROM passenger as pr
JOIN ticket as tt ON tt.passenger_id = pr."id"
WHERE tt.weight > 10