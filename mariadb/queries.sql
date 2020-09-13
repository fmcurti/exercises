
SELECT * FROM Followers f
join Tweets tw on tw.userid = f.follows
join Users u on u.userid = f.follows
where f.userid = (SELECT userid from Users where username = 'Mark')
order by publication_date desc
LIMIT 30;


SELECT * FROM Followers f
join Tweets tw on tw.userid = f.follows
join Users u on u.userid = f.follows
where f.userid = (SELECT userid from Users where username = 'Mark')
order by publication_date desc
LIMIT 30 OFFSET  30;