<?php

namespace {{ .fname }}\{{ .lname }}\Model\ResourceModel;

use Magento\Framework\Model\ResourceModel\Db\AbstractDb;

class {{ .cname }} extends AbstractDb
{
    protected function _construct()
    {
        $this->_init('{{ .table }}', '{{ .pk }}');
    }
}
