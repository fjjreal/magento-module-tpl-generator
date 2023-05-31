{
    "name": "{{ .vname }}",
    "version": "{{ .version }}",
    "type": "magento2-module",
    "require": {
        "php": "~5.5.0|~5.6.0|7.0.2|7.0.4|~7.0.6|7.1.*",
        "ext-ctype": "*"
    },
    "autoload": {
        "psr-4": {
            "{{ .psr }}": ""
        },
        "files": [
            "registration.php"
        ]
    }
}
