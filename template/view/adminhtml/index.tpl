<?xml version="1.0" ?>
<page xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
    xsi:noNamespaceSchemaLocation="urn:magento:framework:View/Layout/etc/page_configuration.xsd">
    <body>
        <referenceContainer name="content">
            <block class="{{ .fname }}\{{ .lname }}\Block\Adminhtml\{{ .ndname }}" name="{{ .st }}_{{ .nd }}_listing"/>
        </referenceContainer>
        <referenceContainer name="admin.{{ .st }}.{{ .nd }}.grid"></referenceContainer>
    </body>
</page>
