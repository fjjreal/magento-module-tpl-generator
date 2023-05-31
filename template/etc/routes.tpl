<?xml version="1.0"?>
<config xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xsi:noNamespaceSchemaLocation="urn:magento:framework:App/etc/routes.xsd">
    <router id="{{ .area }}">
        <route id="{{ .rname }}" frontName="{{ .rname }}">
            <module name="{{ .mname }}"/>
        </route>
    </router>
</config>
