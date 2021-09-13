-- Посчитать продажи за месяц по автору(количество книг и денежный эквивалент)

SELECT SUM(inf.cnt) as "count", SUM(inf.sm) as "sum" FROM (SELECT oo.book_id, SUM(oo."count") as cnt, SUM(bk.price * oo."count") as sm
FROM (SELECT oo.book_id, oo."count" 
			FROM order_info oo
			JOIN "order" ord ON ord."id" = oo.order_id
			WHERE ord.date_made BETWEEN (NOW()-interval '1 month') AND NOW()) oo 
JOIN (SELECT bk."id", bk.price
			FROM book bk
			JOIN (SELECT ar."id" 
						FROM author ar
						WHERE ar."name"='Sapkowski') ar ON ar."id" = bk.author_id) bk
			ON bk."id" = oo.book_id
			GROUP BY oo.book_id) as inf

------------------------------------------------------------------------
SELECT SUM(inf.cnt), SUM(inf.sm) FROM (SELECT oo.book_id, SUM(oo."count") as cnt, SUM(bk.price * oo."count") as sm
FROM order_info oo 
JOIN (SELECT bk."id", bk.price
			FROM book bk
			JOIN (SELECT ar."id" 
						FROM author ar
						WHERE ar."name"='Sapkowski') ar ON ar."id" = bk.author_id) bk
			ON bk."id" = oo.book_id
			GROUP BY oo.book_id) as inf