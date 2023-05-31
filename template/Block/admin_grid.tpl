<?php

namespace {{ .fname }}\{{ .lname }}\Block\Adminhtml\{{ .ndname }};

class Grid extends \Magento\Backend\Block\Widget\Grid\Extended
{
    /**
     * @var \Magento\Framework\Module\Manager
     */
    protected $moduleManager;

    protected $_currentTableFactory;

    protected $_scopeConfig;

    /**
     * @var \Magento\Cms\Model\Config\Source\Page
     */
    private $_page;

    /**
     * @var \Magento\Cms\Model\BlockFactory
     */
    private $_blockFactory;

    /**
     * Grid constructor.
     * @param \Magento\Backend\Block\Template\Context $context
     * @param \Magento\Backend\Helper\Data $backendHelper
     * @param \{{ .fname }}\{{ .lname }}\Model\{{ .factory }} $currentFactory
     * @param \Magento\Framework\Module\Manager $moduleManager
     * @param \Magento\Framework\App\Config\ScopeConfigInterface $scopeConfig
     * @param \Magento\Cms\Model\Config\Source\Page $page
     * @param \Magento\Cms\Model\BlockFactory $blockFactory
     * @param array $data
     */
    public function __construct(
        \Magento\Backend\Block\Template\Context $context,
        \Magento\Backend\Helper\Data $backendHelper,
        \{{ .fname }}\{{ .lname }}\Model\{{ .factory }} $currentFactory,
        \Magento\Framework\Module\Manager $moduleManager,
        \Magento\Framework\App\Config\ScopeConfigInterface $scopeConfig,
        \Magento\Cms\Model\Config\Source\Page $page,
        \Magento\Cms\Model\BlockFactory $blockFactory,
        array $data = []
    ) {
        $this->_currentTableFactory = $currentFactory;
        $this->moduleManager = $moduleManager;
        $this->_scopeConfig = $scopeConfig;
        $this->_page = $page;
        $this->_blockFactory = $blockFactory;
        parent::__construct($context, $backendHelper, $data);
    }

    /**
     * @return void
     */
    protected function _construct()
    {
        parent::_construct();
        /* ******************************** */
        // TODO replace there
        $this->setId('postGrid');
        $this->setDefaultSort('id');
        $this->setDefaultDir('ASC');
        $this->setSaveParametersInSession(true);
        $this->setUseAjax(false);
        $this->setVarNameFilter('post_filter');
        /* ******************************** */
    }

    /**
     * @return $this
     */
    protected function _prepareCollection()
    {
        $collection = $this->_currentTableFactory->create()->getCollection();
        $this->setCollection($collection);
        parent::_prepareCollection();
        return $this;
    }

    /**
     * @return $this
     * @SuppressWarnings(PHPMD.ExcessiveMethodLength)
     */
    protected function _prepareColumns()
    {

        {{ .viewCols }}

        {{ .viewExports }}

        return parent::_prepareColumns();
    }

    /**
     * @return $this|Grid
     */
    protected function _prepareMassaction()
    {
        /* ******************************** */
        // TODO replace there
        $this->setMassactionIdField('id');
        $this->getMassactionBlock()->setFormFieldName('remind');
        /* ******************************** */
        
        {{ .viewMass }}
        
        return $this;
    }

    /**
     * @return string
     */
    public function getGridUrl()
    {
        return $this->getUrl('{{ .lnameL }}/{{ .cnameL }}/index', ['_current' => true]);
    }

}
