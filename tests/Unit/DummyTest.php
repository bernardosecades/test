<?php

namespace Tests\Unit;

use BernardoSecades\Test\Dummy;

class DummyTest extends \PHPUnit_Framework_TestCase
{
    /**
     * @test
     */
    public function dummy()
    {
        $dummy = new Dummy();
        $this->assertEquals('lololo', $dummy->execute());
    }
}
