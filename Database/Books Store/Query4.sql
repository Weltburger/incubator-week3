-- Вывести все заказы по определенному пользователю где были только электронные книги

SELECT *
FROM "order" ord
JOIN (SELECT oo.order_id, COUNT(oo.order_id) as cnt
			FROM order_info oo 
			JOIN (SELECT bk."id" 
			      FROM book bk 
				  WHERE bk."type_id"=2) bk ON bk."id" = oo.book_id
			GROUP BY oo.order_id
			HAVING (COUNT(oo.order_id) = (SELECT COUNT(oi.order_id) 
										  FROM order_info oi 
										  WHERE oi.order_id = oo.order_id))) oo 
ON oo.order_id = ord."id"
WHERE ord.customer_id=(SELECT cr."id" 
					   FROM customer cr
					   WHERE cr."name"='Jeka')