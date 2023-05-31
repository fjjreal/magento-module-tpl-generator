<?php

namespace {{ .fname }}\{{ .lname }}\Model\ResourceModel\{{ .cname }};

use Magento\Framework\Model\ResourceModel\Db\Collection\AbstractCollection;

class Collection extends AbstractCollection
{
    protected $_idFieldName = '{{ .pk }}';
    protected function _construct()
    {
        $this->_init('{{ .fname }}\{{ .lname }}\Model\{{ .cname }}', '{{ .fname }}\{{ .lname }}\Model\ResourceModel\{{ .cname }}');
    }
}
