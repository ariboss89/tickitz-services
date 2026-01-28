-- CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- 1. Login
SELECT first_name, last_name, role
  FROM users
 WHERE email = 'ramadhan89.ari@gmail.com' 
   AND password = crypt('123456', password);

-- 2. Register user
INSERT INTO users (first_name, last_name, email, password, role, created_at) 
VALUES (('Ari'), ('Ramadhan'), ('ramadhan89.ari@gmail.com'),
(crypt('123456', gen_salt('bf'))),'user', NOW());

-- 3. Get upcoming movie
SELECT title,
       synopsis,
       poster_url,
       background_url,
       release_date,
       duration,
       status,
       rating,
       created_at,
       updated_at
FROM movies WHERE status = 'upcoming';

-- 4. popular movie
SELECT title,
       synopsis,
       poster_url,
       background_url,
       release_date,
       duration,
       status,
       rating,
       created_at,
       updated_at
FROM movies WHERE rating > 7.5;

-- 5. Pagination 
SELECT 
       "Id", title,
       synopsis,
       poster_url,
       background_url,
       release_date,
       duration,
       status,
       rating
FROM movies 
ORDER BY "Id" ASC 
LIMIT 2 
OFFSET 2;

-- 6. FILTER Movie By name and pagination
SELECT "Id", title,
       synopsis,
       poster_url,
       background_url,
       release_date,
       duration,
       status,
       rating
FROM movies 
WHERE title LIKE '%a%' 
ORDER BY "Id" ASC 
LIMIT 2 
OFFSET 2;

-- b. FILTER Movie By genre and pagination
-- SELECT g.name, 
--        m.title,
--        m.synopsis,
--        m.poster_url,
--        m.background_url,
--        m.release_date,
--        m.duration
-- FROM movie_genres mg
-- JOIN movies m ON m."Id" = mg.movie_id
-- JOIN genres g ON g."Id" = mg.genre_id
-- WHERE g.name LIKE '%%'
-- ORDER BY mg.id ASC 
-- LIMIT 2 
-- OFFSET 2;

SELECT m.title,
       m.synopsis,
       m.poster_url,
       m.background_url,
       m.release_date,
       m.duration,
       STRING_AGG(g.name, ', ')
FROM movie_genres mg
JOIN movies m ON m.id = mg.movie_id
JOIN genres g ON g.id = mg.genre_id
WHERE g.name LIKE '%%'
GROUP BY
       m.title,
       m.synopsis,
       m.poster_url,
       m.background_url,
       m.release_date,
       m.duration
HAVING
    STRING_AGG(g.name, ', ') LIKE '%Action%%Action%'
ORDER BY m.title ASC;
-- LIMIT 2 
-- OFFSET 2;

SELECT
       g.name, 
       m.title,
       m.synopsis,
       m.poster_url,
       m.background_url,
       m.release_date,
       m.duration,
       STRING_AGG(g.name, ', ') AS aggregated_result
FROM movie_genres mg
JOIN movies m ON m.id = mg.movie_id
JOIN genres g ON g.id = mg.genre_id
WHERE g.name LIKE '%%'
GROUP BY
    g.name, m.title, m.synopsis,
       m.poster_url,
       m.background_url,
       m.release_date,
       m.duration
HAVING
    STRING_AGG(g.name, ', ') LIKE '%Action%%Sci-Fi%';


-- 7. GET Schedule
SELECT s.show_date,
       s.show_time,
       s.price,
       m.title,
       c.name,
       c.address,
       c.city,
       st.studio_name
FROM schedules s
JOIN movies m ON s.movie_id = m."Id"
JOIN cinemas c ON s.cinema_id = c."Id"
JOIN studios st ON s.studio_id = st."Id";

-- 8. Get Seat Sold/ Available
-- SELECT s.seat_number,
--        s.seat_type,
--        -- schedule
--        sd.movie_id,
--        sd.cinema_id,
--        sd.show_date,
--        sd.show_time
-- FROM schedules sd
-- JOIN seats s ON s.schedule_id = sd.Id
-- JOIN movies m ON sd.movie_id = m.Id
-- JOIN cinemas c ON sd.cinema_id = c.Id
-- WHERE c.Id = 1 AND m.Id = 1 AND sd.show_date= '2026-01-01' AND sd.show_time='16:15';

SELECT 
      o.Id,
      dt.seat_id
FROM orders o
JOIN dt_orders dt ON dt.seat_id = s."Id"
JOIN seats s ON s."Id" = dt.seat_id
WHERE o.Id = 1;

--9. Get Movie Detail
--first we need to take the genre first
SELECT g.name, m.title
FROM movie_genres mg
JOIN movies m ON m."Id" = mg.movie_id
JOIN genres g ON g."Id" = mg.genre_id
WHERE mg.movie_id = 1;

--second we need to take actors name
SELECT a.name, m.title
FROM movie_actors ma
JOIN actors a ON a.id = ma.actor_id
JOIN movies m ON m.id = ma.movie_id
WHERE ma.movie_id = 1;

-- second we get the data for detail
SELECT m.title,
       m.synopsis,
       m.poster_url,
       m.background_url,
       m.release_date,
       m.duration,
       d.name AS "directors"
       -- ambil data movie genres
FROM movies m
JOIN directors d ON d.movie_id = m."Id"
WHERE m."Id" = 1;

-- GET DETAIL OTHER WAYS
-- SELECT g.name, 
--        m.title,
--        m.synopsis,
--        m.poster_url,
--        m.background_url,
--        m.release_date,
--        m.duration,
--        d.name AS "director",
--        a.name AS "actor"
-- FROM movie_genres mg
-- JOIN movies m ON m.id = mg.movie_id
-- JOIN genres g ON g.id = mg.genre_id
-- JOIN directors d ON d.movie_id = mg.movie_id
-- JOIN actors a ON a.movie_id = mg.movie_id
-- WHERE mg.movie_id = 1;

-- 10. CREATE ORDER

-- SELECT 
-- FROM orders o
-- JOIN users u ON u.id = o.user_id
-- JOIN payments p ON p.id = o.payment_id;

-- SAVE DETAIL FIRST
-- SELECT id, 
--        order_id, 
--        seat_id, 
--        movies_id, 
--        price 
-- FROM dt_orders do
-- JOIN orders o ON o.id = do.order_id;


INSERT INTO orders (sub_total, total_price, tax, status, created_at, user_id, payment_id, schedule_id) 
VALUES (105000, 125500, 10500, 'pending', NOW(), 1, 1, 16);

-- CREATE DT ORDER
INSERT INTO dt_orders (order_id, seat_id) VALUES (2, 1);

-- 11. Get profile
SELECT first_name,
       last_name,
       email,
       password,
       role
FROM users WHERE email = 'ramadhan89.ari@gmail.com';

-- 12. Get History
SELECT
        o."Id",
        o.total_price,
        o.status,
        o.booking_code,
        p.payment_method
FROM orders o
JOIN users u ON o.user_id = u."Id"
JOIN payments p ON p."Id" = o.payment_id
WHERE o.status = 'paid' AND o.user_id = 1;

-- 13. Edit Profile
UPDATE users SET first_name = 'Arri', last_name = 'Ramadhann', 
password = (crypt('12345678', gen_salt('bf'))), updated_at = NOW()
WHERE email = 'ramadhan89.ari@gmail.com';

-- 14. Get All Movie
SELECT title,
       synopsis,
       poster_url,
       background_url,
       release_date,
       duration,
       status,
       rating,
       deleted_at
FROM movies WHERE deleted_at is null;

-- 15. DELETE MOVIE
UPDATE movies SET deleted_at = NOW() WHERE Id = 2;

-- 16. EDIT MOVIE
UPDATE movies SET title = 'Brain Dead', synopsis = 'Reel Rock 8', poster_url = 'http://dummyimage.com/244x100.png/cc0000/ffffff',
background_url = 'http://dummyimage.com/186x100.png/dddddd/000000', release_date = '2025-01-27 17:50:02',
duration = 100, status = 'upcoming', rating= '7.0', updated_at = NOW() WHERE Id = 2;

-- SCHEDULE
SELECT s.id,
       s.show_date,
       s.show_time,
       s.price,
       m.title,
       st.studio_name,
       c.name,
			 c.address,
			 c.city
	FROM schedules s
	JOIN cinemas c ON c.id = s.cinema_id
	JOIN studios st ON st.id = c.id
  JOIN movies m ON m.id = s.movie_id
  WHERE s.movie_id = 8 AND c.city = 'DKI JAKARTA' AND s.show_time = '12:00' AND s.show_date='NOW()'
  ORDER BY id ASC;

  --OR s.show_date = (NOW() + interval '3 day')



  SELECT id,
       order_id,
       seat_id,
       movies_id,
       price,
       created_at,
       updated_at,
       deleted_at
FROM public.dt_orders
LIMIT 1000;

--history
ALTER TABLE history DROP CONSTRAINT history_order_id_fkey;
ALTER TABLE history ALTER COLUMN order_id TYPE VARCHAR(255);

--order
ALTER TABLE orders ALTER COLUMN id TYPE VARCHAR(255);
ALTER TABLE orders DROP COLUMN order_id;

--constraint fkey
ALTER TABLE dt_orders
ADD CONSTRAINT dt_orders_order_id_fkey
FOREIGN KEY (order_id)
REFERENCES orders (id);

SELECT id,
       order_id,
       seat_id,
       schedule_id,
       created_at,
       updated_at,
       deleted_at
FROM public.dt_orders
LIMIT 1000;

ALTER TABLE dt_orders DROP CONSTRAINT dt_orders_order_id_fkey;

ALTER TABLE dt_orders DROP COLUMN movies_id;
ALTER TABLE dt_orders ADD CONSTRAINT dt_orders_order_id_fkey
FOREIGN KEY (order_id)
REFERENCES orders (id);

INSERT into dt_orders (order_id, seat_id, schedule_id) VALUES ('ORDER00001', 91, 1);

ALTER TABLE orders ALTER COLUMN id TYPE VARCHAR(255);

DROP TABLE dt_orders CASCADE;

-- aaaa


CREATE TABLE orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  schedule_id int,
  total_ticket int,
  sub_total FLOAT,
  tax FLOAT,
  total_price FLOAT,
  status VARCHAR(255) DEFAULT 'pending',
  booking_code VARCHAR(255),
  point int,
  user_id int,
  payment_id int, 
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_id) REFERENCES payments(id),
  FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);

CREATE TABLE dt_orders (
  id serial PRIMARY KEY,
  order_id UUID,
  seat_id int, 
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (seat_id) REFERENCES seats(id),
  FOREIGN KEY (order_id) REFERENCES orders(id)
);

drop table dt_orders CASCADE;

-- schedule
-- schedule_id int,
--   total_ticket int,
--   sub_total FLOAT,
--   tax FLOAT,
--   total_price FLOAT,
--   status VARCHAR(255) DEFAULT 'pending',
--   booking_code VARCHAR(255),
--   point int,
--   user_id int,
--   payment_id int, 

SELECT
    s.
    o.id, 
    o.total_price, 
    o.status, 
    o.booking_code, 
    p.payment_method 
FROM orders o 
JOIN dt_orders do
JOIN users u ON o.user_id = u.id
JOIN schedules s ON s.id = o.schedule_id 
JOIN payments p ON p.id = o.payment_id 
WHERE o.status = 'paid' 
AND o.user_id = $1

BEGIN
INSERT INTO dt_orders (seat_id, order_id) VALUES (90, 'ORDER00001');
INSERT INTO dt_orders (seat_id, order_id) VALUES (91, 'ORDER00001');

INSERT INTO orders (id, schedule_id, sub_total, tax, total_price, booking_code, point, user_id, payment_id)
VALUES ('ORDER00001', 1, 80000, 8000, 88000, 'TIX001826', 80, 19, 1);

UPDATE orders SET status = 'paid' WHERE id = 'ORDER00001';

COMMIT

ROLLBACK