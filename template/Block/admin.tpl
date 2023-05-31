<?php

namespace {{ .fname }}\{{ .lname }}\Block\Adminhtml;

class {{ .ndname }} extends \Magento\Backend\Block\Widget\Container
{
    /**
     * @var string
     */
    protected $_template = '{{ .nd }}.phtml';

    /**
     * @param \Magento\Backend\Block\Widget\Context $context
     * @param array $data
     */
    public function __construct(\Magento\Backend\Block\Widget\Context $context,array $data = [])
    {
        parent::__construct($context, $data);

        /*
        $this->buttonList->remove('Todo');
        $this->addButton(
            'Todo',
            [
                'label' => __('Todo'),
                'on_click' => "require(['jquery', 'Magento_Ui/js/modal/prompt'], function ($, prompt) {prompt({
                        title: 'Todo',
                        content: '<span>Input Your email:</span>',
                        actions: {
                            confirm: function () {
                                let c_email = $('input[data-role=\"promptField\"]').val()
                                console.log(c_email)
                                let url =  '".$this->getUrl('{{ .st }}/{{ .nd }}/massTodo')."';
                                location.href = url + '?email=' + c_email;
                            },
                            cancel: function () {
                            },
                            always: function () {
                            }
                        }
                    });})",
                'class' => 'primary',
                'level' => 1
            ]
        );
        */
        
    }

    /**
     * @return {{ .ndname }}
     * @throws \Magento\Framework\Exception\LocalizedException
     */
    protected function _prepareLayout()
    {
        $this->setChild(
            'grid',
            $this->getLayout()->createBlock('\{{ .fname }}\{{ .lname }}\Block\Adminhtml\{{ .ndname }}\Grid', '{{ .st }}.{{ .nd }}.grid')
        );
        return parent::_prepareLayout();
    }

    /**
     * Render grid
     *
     * @return string
     */
    public function getGridHtml()
    {
        return $this->getChildHtml('grid');
    }

}
