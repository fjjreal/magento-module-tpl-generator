<?php

namespace {{ .fname }}\{{ .lname }}\Controller\Adminhtml\{{ .ndname }};

class {{ .cname }} extends \Magento\Backend\App\Action
{

    public function execute()
    {
        $filename = time() . ".csv";
        $fileData = "Export sth.\n";

        // TODO fill data

        header("Content-type:text/csv");
        header("Content-Disposition:attachment;filename=" . $filename);
        header('Cache-Control:must-revalidate,post-check=0,pre-check=0');
        header('Expires:0');
        header('Pragma:public');
        echo $fileData;
        exit;
    }
}
