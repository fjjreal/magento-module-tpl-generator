<?php

namespace {{ .fname }}\{{ .lname }}\Cron;

use Magento\Framework\ObjectManagerInterface;

class {{ .cname }}
{
    private $objectManager;

    public function __construct(
        ObjectManagerInterface $objectManager
    )
    {
        $this->objectManager = $objectManager;
    }

    public function {{ .method }}()
    {
        // TODO sth
    }
}
