-- Посчитать количество заказов в которых содержится книга "N"

SELECT COUNT(*) FROM (SELECT oo.order_id
FROM order_info oo 
JOIN (SELECT bk."id" 
			FROM book bk 
			WHERE bk."title"='Book3') bk ON bk."id" = oo.book_id
GROUP BY oo.order_id) as "count"