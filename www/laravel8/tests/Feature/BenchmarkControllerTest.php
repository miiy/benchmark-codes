<?php

namespace Tests\Feature;

use Illuminate\Foundation\Testing\RefreshDatabase;
use Tests\TestCase;

class BenchmarkControllerTest extends TestCase
{
    /**
     * A basic test example.
     *
     * @return void
     */
    public function testIndex()
    {
        $response = $this->get('/');
        $response->dump();
        $response->assertStatus(200);
        $response->assertSeeText("hello!");
    }

    public function testRedis()
    {
        $response = $this->get('/redis');
        $response->dump();
        $response->assertStatus(200);
        $response->assertSeeText("hello!");
    }
}
