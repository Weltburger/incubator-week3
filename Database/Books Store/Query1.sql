-- Посчитать количество "смешанных" заказов за день

SELECT COUNT(*)
FROM "order" ord
JOIN (SELECT oo.order_id, COUNT(oo.order_id) as cnt
	  FROM order_info oo 
	  JOIN (SELECT bk."id" 
			FROM book bk 
			WHERE bk."id"=2 OR bk."id"=5) bk ON bk."id" = oo.book_id
	  GROUP BY oo.order_id
	  HAVING (COUNT(oo.order_id) != (SELECT COUNT(oi.order_id) 
									 FROM order_info oi 
									 WHERE oi.order_id = oo.order_id))) oo 
ON oo.order_id = ord."id"
WHERE ord.date_made BETWEEN '2021-09-09' AND '2021-09-10'