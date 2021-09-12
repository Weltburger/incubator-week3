-- Количество успешных доставок, по определенному пользователю, где количество книг > 2 

SELECT COUNT(*) 
FROM "order" ord
WHERE ord.date_done IS NOT NULL 
AND ord.customer_id=(SELECT cr."id" 
					 FROM customer cr
					 WHERE cr."name"='Pakowski')