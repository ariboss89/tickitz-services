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
SELECT id, title,
       poster_url,
       duration,
       status,
       rating
FROM movies 
WHERE title LIKE '%a%' 
ORDER BY id ASC 
LIMIT 2 
OFFSET 2;

SELECT m.id, m.title,
       m.poster_url,
       m.duration,
       m.status,
       m.rating,
       STRING_AGG(g.name, ', ') AS "genres"
FROM movie_genres mg
JOIN movies m ON m.id = mg.movie_id 
JOIN genres g ON g.id = mg.genre_id
WHERE m.title LIKE '%'||'a'||'%' 
-- AND g.name ILIKE 'adventure' 
-- OR g.name ILIKE 'sci-fi'
GROUP BY m.id
ORDER BY m.id ASC
LIMIT 2 
OFFSET 2;

SELECT *FROM movies;

SELECT m.id, m.title,
       m.poster_url,
       m.duration,
       m.status,
       m.rating,
       STRING_AGG(g.name, ', ') AS "genres"
FROM movie_genres mg
JOIN movies m ON m.id = mg.movie_id 
JOIN genres g ON g.id = mg.genre_id
WHERE m.title LIKE '%a%'
GROUP BY m.id
HAVING STRING_AGG(g.name, ', ') LIKE '%adventure%'
ORDER BY m.id ASC;

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


SELECT m.id, m.title,
       m.poster_url,
       m.duration,
       m.status,
       m.rating,
       STRING_AGG(g.name, ', ') AS "genres"
		FROM movie_genres mg
		JOIN movies m ON m.id = mg.movie_id 
		JOIN genres g ON g.id = mg.genre_id
		-- WHERE m.title LIKE '%%'
		GROUP BY m.id
		ORDER BY m.id ASC
		LIMIT 10
		OFFSET 2;

    SELECT 
        m.id, m.title, m.poster_url, m.duration, m.status, m.rating, STRING_AGG(g.name, ', ') 
        FROM movie_genres mg 
        JOIN movies m ON m.id = mg.movie_id 
        JOIN genres g ON g.id = mg.genre_id 
        WHERE m.title ILIKE '%%' AND g.name = 'Action' OR g.name = 'Adventure'
        GROUP BY m.id 
        ORDER BY m.id ASC 
        LIMIT 5 
        OFFSET 0;