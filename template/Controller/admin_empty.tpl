<?php

namespace {{ .fname }}\{{ .lname }}\Controller\Adminhtml\{{ .pname }};

use Magento\Backend\App\Action\Context;
use Magento\Framework\View\Result\PageFactory;

class {{ .cname }} extends {{ .extend }}
{
    /**
     * @var PageFactory
     */
    protected $resultPageFactory;

    /**
     * @param Context $context
     * @param PageFactory $resultPageFactory
     */
    public function __construct(
        Context $context,
        PageFactory $resultPageFactory
    ) {
        parent::__construct($context);
        $this->resultPageFactory = $resultPageFactory;
    }

    /**
     * {@inheritdoc}
     */
    protected function _isAllowed()
    {
        return true;
    }

    /**
     * {{ .cname }} action
     *
     * @return \Magento\Backend\Model\View\Result\Page
     */
    public function execute()
    {
        /** @var \Magento\Backend\Model\View\Result\Page $resultPage */
        $resultPage = $this->resultPageFactory->create();
        $resultPage->addBreadcrumb(__('{{ .fname }} {{ .lname }} {{ .pname }}'), __('{{ .cname }}'));
        $resultPage->addBreadcrumb(__('{{ .fname }} {{ .lname }} {{ .pname }}'), __('{{ .cname }}'));
        $resultPage->getConfig()->getTitle()->prepend(__('{{ .fname }} {{ .lname }}'));
        return $resultPage;
    }
}
