-- Количество успешных доставок, по определенному пользователю, где количество книг > 2 

SELECT COUNT(*) 
FROM "order" ord
JOIN (SELECT oo.order_id 
	  FROM order_info oo 
	  GROUP BY oo.order_id
	  HAVING COUNT(oo.order_id) > 2) oo ON oo.order_id = ord."id"
WHERE ord.date_done IS NOT NULL 
AND ord.customer_id=(SELECT cr."id" 
					 FROM customer cr
					 WHERE cr."name"='Pakowski')