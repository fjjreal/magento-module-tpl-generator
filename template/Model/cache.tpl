<?php

namespace {{ .fname }}\{{ .lname }}\Model\Cache;

use Magento\Framework\App\Cache\Type\FrontendPool;
use Magento\Framework\Cache\Frontend\Decorator\TagScope;

class {{ .cname }} extends TagScope
{
    const TYPE_IDENTIFIER = '{{ .id }}';

    const CACHE_TAG = '{{ .tag }}';

    public function __construct(FrontendPool $cacheFrontendPool)
    {
        parent::__construct(
            $cacheFrontendPool->get(self::TYPE_IDENTIFIER),
            self::CACHE_TAG
        );
    }

}
