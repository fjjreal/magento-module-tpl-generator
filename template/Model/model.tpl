<?php

namespace {{ .fname }}\{{ .lname }}\Model;

use Magento\Framework\Model\AbstractModel;

class {{ .cname }} extends AbstractModel
{
    protected function _construct()
    {
        $this->_init('{{ .fname }}\{{ .lname }}\Model\ResourceModel\{{ .cname }}');
    }
}
