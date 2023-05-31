<?xml version="1.0"?>
<layout xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
    xsi:noNamespaceSchemaLocation="urn:Magento/Framework/View/Layout/etc/layout_generic.xsd">
    <update handle="formkey"/>
    <container name="root" label="Root">
        <block class="{{ .fname }}\{{ .lname }}\Block\Adminhtml\{{ .ndname }}\Grid" name="admin.{{ .st }}.{{ .nd }}.grid"/>
    </container>
</layout>
