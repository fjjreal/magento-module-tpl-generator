<?php

namespace {{ .fname }}\{{ .lname }}\Controller\{{ .pname }};

use Magento\Framework\App\Action\Action;
use Magento\Framework\App\Action\Context;
use Magento\Framework\App\Request\InvalidRequestException;
use Magento\Framework\App\RequestInterface;
use Magento\Framework\App\CsrfAwareActionInterface;
use Magento\Framework\Controller\Result\JsonFactory;
use Magento\Framework\Controller\Result\Json;

class {{ .cname }} extends Action implements CsrfAwareActionInterface
{
    /**
     * @var Json
     */
    protected $json;

    public function __construct(
        Context     $context,
        JsonFactory $jsonFactory
    )
    {
        parent::__construct($context);
        $this->json = $jsonFactory->create();
    }

    public function execute()
    {
        // TODO sth
        return $this->err("Hello World!");
    }

    /**
     * @param $param
     * @param $checkFunc
     * @return mixed|void
     */
    protected function mustHasParam($param,$checkFunc = "")
    {
        $p = $this->getRequest()->getParam($param);
        if(empty($p)){
            $this->jsonErr("Param[{$param}] Empty Error");
        }
        if (($checkFunc instanceof \Closure) && !call_user_func($checkFunc,$p)) {
            $this->jsonErr("Param[{$param}] Format Error");
        }
        return $p;
    }

    protected function isLogin()
    {
        return $this->_objectManager->create(\Magento\Customer\Model\Session::class)->getCustomer();
    }

    protected function succ($data, $code = 0)
    {
        return $this->json->setData(['code' => $code, 'message' => "", 'data' => $data]);
    }

    protected function err($message = "", $code = 1)
    {
        return $this->json->setData(['code' => $code, 'message' => __($message), 'data' => null]);
    }

    protected function jsonErr($message, $code = 1)
    {
        header('Content-Type: application/json; charset=utf-8');
        echo json_encode(['code' => $code, 'message' => __($message), 'data' => null]);
        exit();
    }

    public function createCsrfValidationException(RequestInterface $request): ?InvalidRequestException
    {
        return null;
    }

    public function validateForCsrf(RequestInterface $request): ?bool
    {
        return true;
    }
}
