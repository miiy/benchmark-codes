<?php

namespace App\Http\Controllers;

use Illuminate\Support\Facades\Redis;

class BenchmarkController extends Controller
{
    public function index()
    {
        return "hello!";
    }

    public function redis()
    {
        $hello = Redis::get('hello');
        if ($hello == NULL) {
           Redis::setex('hello', 60, 'hello!');
           $hello = Redis::get('hello');
        }
        return $hello;
    }


}
