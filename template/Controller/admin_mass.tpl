<?php
namespace {{ .fname }}\{{ .lname }}\Controller\Adminhtml\{{ .ndname }};

class {{ .cname }} extends \Magento\Backend\App\Action
{
    /**
     * @return \Magento\Backend\Model\View\Result\Redirect
     */
    public function execute()
    {
        $itemIds = $this->getRequest()->getParam('id');
        
        // TODO sth

        // $this->messageManager->addSuccess("Success.");
        // $this->messageManager->addError($e->getMessage());

        return $this->resultRedirectFactory->create()->setPath('*/*/index');
    }
}
