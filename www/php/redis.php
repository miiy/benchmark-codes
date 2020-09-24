<?php

require "./vendor/autoload.php";

$client = new Predis\Client([
    'scheme' => 'tcp',
    'host'   => '127.0.0.1',
    'port'   => 6379,
]);

$cacheKey = "hello";

$hello = $client->get($cacheKey);
if ($hello == NULL) {
    $client->setex($cacheKey, 60,"hello!");
    $hello = $client->get($cacheKey);
}
echo $hello;
