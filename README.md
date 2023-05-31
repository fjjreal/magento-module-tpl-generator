# 使用

- 1.调整module.yaml
- 1.1 可适当调整template文件夹中的模板 
- 2.git bash执行 ./magento-module-generator.exe

# Yaml 

```yaml
mname: Exp_Mod
author: 'fjjreal@gmail.com'
version: 1.0.0
route:
  adminhtml: mod
  frontend: mod
routers:
  areas:
    - id: admin
      ctls:
        - st: mod
          nd: data
          rd: index
          # 生成后台列表视图
          view: true
          factory: ModData
          columns:
            - id: id
              label: ID
              type: number
              hcss: 'col-id'
              ccss: 'col-id'
            - id: atype
              label: 'Exp Mod Type'
            - id: customer_id
              label: 'Customer ID'
              type: number
          exports:
            - id: 'ExportModDataCsv'
              label: 'GetModDataDaily'
          mass:
            - id: delete
              action: massDelete
              label: delete
        - st: mod
          nd: data
          rd: massEnabled
    - id: front
      ctls:
        - st: mod
          nd: data
          rd: all
        - st: mod
          nd: data
          rd: top10
          extend: '\Exp\Mod\Controller\Common'
    - id: base
      ctls:
        - st: ''
          nd: ''
          rd: common
system:
  tabs:
    - id: exp_mod
      label: 'Exp Mod'
    - id: 'exp_default'
      label: 'Exp'
      sections:
        - id: 'mod'
          css: 'separator-top'
          label: 'Mod'
          resource: 'Exp_Default::mod'
          groups:
            - id: 'limit'
              label: 'Limiter'
              fields:
                - id: 'rate'
                  type: 'text'
                  label: 'api rate'
                  comment: 'api rate'
cron:
  groups:
    - id: default
      jobs:
        - name: exp_mod_email_sender
          instance: 'EmailSender'
          method: execute
          schedule: '* * * * *'
email:
  theader: '{{template config_path="design/email/header_template"}}'
  tfooter: '{{template config_path="design/email/footer_template"}}'
  templates:
    - id: mod_fast
      label: 'Send Email Code To User'
      file: 'mod_fast.html'
      type: 'html'
      area: 'frontend'
      subject: '<!--@subject {{trans "Mod Email Subject"}} @-->'
      theader: '{{template config_path="design/email/header_template2"}}'
menu:
  - id: 'mod'
    title: 'mod'
    resource: 'Exp_Mod::parent'
    action: ''
    parent: 'Magento_Backend::marketing'
    module: 'Exp_Mod'
  - id: 'mod_data'
    title: 'mod Activity'
    resource: 'Exp_Mod::data'
    action: 'mod/data/index'
    parent: 'Exp_Mod::mod'
    module: 'Exp_Mod'
model:
  - name: 'ModData'
    table: exp_mod_data
    pk: id
table:
  - id: 'exp_mod_data'
    mname: 'ModData'
    tmod: true
    pk: id
    columns:
      - id: 'id'
        type: big_int
        length: 20
        identity: true
        unsign: true
        nullable: true
        primary: true
        comment: 'Id'
      - id: atype
        type: varchar
        length: 32
        nullable: false
        default: '"46546"'
        comment: 'type'
      - id: customer_id
        type: big_int
        length: 20
        nullable: false
        default: 0
        comment: 'Customer ID'
    index:
      - id: atype
        type: unique
      - id: customer_id
      - id: 'atype,customer_id'
cache:
  - id: 'exp_mod_data_cache_id'
    tag: 'EXP_MOD_DATA_CACHE_TAG'
    instance: 'Data'
    label: 'Exp Mod Data Cache Label'
    description: 'Exp Mod Data Cache'
    trans: 'label,description'

```
