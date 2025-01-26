# Skinny URL

Skinny URL is a backend for a URL Shortener. 

The purpose of this project was to learn GoLang, but also learn concurrency and backend development.

When a user inputs a large URL, a hash is created by using MD5,converting it to base62, and then using only the first seven characters. The created hash will serve as the short URL when it is appended to the domain name of the hosted website. The short url and long url are both stored in a Redis cache as well as a CassandraDB Database. Along with both urls, a created_date and expiration_date are both stored in conjunction with a clicked_count integer to keep track of the number of times the short_url has redirected to the long_url. 

Concurrency is also implemented in Skinny URL. Using GoRoutines, concurrent cache and database lookups are possible. Asynchronous logging is also used for click tracking.
