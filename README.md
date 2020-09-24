## nginx

http://static.test/

## php

http://php.test/hello.php

http://php.test/redis.php

## laravel8

```bash
cp .env.example .env
composer install

composer install --optimize-autoloader --no-dev
php artisan config:cache
php artisan route:cache
```

http://laravel8.test/

http://laravel8.test/redis

## go

http://go.test/

http://go.test/redis

http://go.test/redis-pool

## gin

http://gin.test/

http://gin.test/redis

http://gin.test/redis-pool