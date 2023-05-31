package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/spf13/viper"
)

func main() {
	NewModuleConfig("")
	MConf.Start()
	MConf.Routef()
	MConf.Contorller()
	MConf.Systemf()
	MConf.Cronf()
	MConf.Emailf()
	MConf.Menuf()
	MConf.Modelf()
	MConf.SetUpf()
	MConf.Cachef()
}

//  -----------------------------------------------------------------  //

// routers.xml 配置
type Route struct {
	Admin string `mapstructure:"adminhtml"`
	Front string `mapstructure:"frontend"`
}

type Mass struct {
	Id     string `mapstructure:"id"`
	Action string `mapstructure:"action"`
	Label  string `mapstructure:"label"`
	Sure   string `mapstructure:"sure"`
}

type Export struct {
	Id    string `mapstructure:"id"`
	Label string `mapstructure:"label"`
}

type GridCol struct {
	Id    string `mapstructure:"id"`
	Label string `mapstructure:"label"`
	Type  string `mapstructure:"type"`
	Hcss  string `mapstructure:"hcss"`
	Ccss  string `mapstructure:"ccss"`
}

// 控制器参数配置
type Ctl struct {
	St      string    `mapstructure:"st"`
	Nd      string    `mapstructure:"nd"`
	Rd      string    `mapstructure:"rd"`
	Extend  string    `mapstructure:"extend"`
	View    bool      `mapstructure:"view"`
	Factory string    `mapstructure:"factory"`
	Cols    []GridCol `mapstructure:"columns"`
	Exports []Export  `mapstructure:"exports"`
	Mass    []Mass    `mapstructure:"mass"`
}

// 控制器类型配置
type Area struct {
	Id   string `mapstructure:"id"`
	Ctls []Ctl  `mapstructure:"ctls"`
}

// 控制器配置
type Router struct {
	Areas []Area `mapstructure:"areas"`
}

// 缓存相关
type SCache struct {
	Id          string `mapstructure:"id"`
	Tag         string `mapstructure:"tag"`
	Instance    string `mapstructure:"instance"`
	Label       string `mapstructure:"label"`
	Description string `mapstructure:"description"`
	Trans       string `mapstructure:"trans"`
}

// 表列名
type SColumn struct {
	Id       string `mapstructure:"id"`
	Type     string `mapstructure:"type"`
	Length   int    `mapstructure:"length"`
	Nullable bool   `mapstructure:"nullable"`
	Default  string `mapstructure:"default"`
	Comment  string `mapstructure:"comment"`
	Primary  bool   `mapstructure:"primary"`
	Unsign   bool   `mapstructure:"unsign"`
	Identity bool   `mapstructure:"identity"`
}

// 表索引
type SIndex struct {
	Id   string `mapstructure:"id"`
	Type string `mapstructure:"type"`
}

// 表相关（setup）
type STable struct {
	Id     string    `mapstructure:"id"`
	Mname  string    `mapstructure:"mname"`
	TMod   bool      `mapstructure:"tmod"`
	Colums []SColumn `mapstructure:"columns"`
	Index  []SIndex  `mapstructure:"index"`
	Pk     string    `mapstructure:"pk"`
}

// DB 模型
type SModel struct {
	Name  string `mapstructure:"name"`
	Table string `mapstructure:"table"`
	Pk    string `mapstructure:"pk"`
}

// 管理端菜单
type SMenu struct {
	Id       string `mapstructure:"id"`
	Resource string `mapstructure:"resource"`
	Title    string `mapstructure:"title"`
	Action   string `mapstructure:"action"`
	Module   string `mapstructure:"module"`
	Parent   string `mapstructure:"parent"`
}

// 邮件模板参数
type ETpls struct {
	Id      string `mapstructure:"id"`
	Label   string `mapstructure:"label"`
	File    string `mapstructure:"file"`
	Type    string `mapstructure:"type"`
	Module  string `mapstructure:"module"`
	Area    string `mapstructure:"area"`
	Subject string `mapstructure:"subject"`
	Theader string `mapstructure:"theader"`
	Tfooter string `mapstructure:"tfooter"`
}

// 邮件模板配置
type EmailTpl struct {
	Theader string  `mapstructure:"theader"`
	Tfooter string  `mapstructure:"tfooter"`
	Tpls    []ETpls `mapstructure:"templates"`
}

// 定时任务 - 任务配置
type CJob struct {
	Name     string `mapstructure:"name"`
	Instance string `mapstructure:"instance"`
	Method   string `mapstructure:"method"`
	Schedule string `mapstructure:"schedule"`
}

// 定时任务 - 任务组参数配置
type CGroup struct {
	Id   string `mapstructure:"id"`
	Jobs []CJob `mapstructure:"jobs"`
}

// 定时任务
type Cron struct {
	Groups []CGroup `mapstructure:"groups"`
}

// system.xml section - 组中的配置
type SField struct {
	Id      string `mapstructure:"id"`
	Type    string `mapstructure:"type"`
	Label   string `mapstructure:"label"`
	Comment string `mapstructure:"comment"`
}

// system.xml section 中的 组配置
type SGroup struct {
	Id     string   `mapstructure:"id"`
	Label  string   `mapstructure:"label"`
	Fields []SField `mapstructure:"fields"`
}

// system.xml section 配置
type SSection struct {
	Id       string   `mapstructure:"id"`
	Css      string   `mapstructure:"css"`
	Label    string   `mapstructure:"label"`
	Resource string   `mapstructure:"resource"`
	Groups   []SGroup `mapstructure:"groups"`
}

// system.xml 组配置
type STab struct {
	Id       string     `mapstructure:"id"`
	Label    string     `mapstructure:"label"`
	Sections []SSection `mapstructure:"sections"`
}

// system.xml configuration struct
type SConfig struct {
	Tabs []STab `mapstructure:"tabs"`
}

// Yaml模块配置
type ModuleConfig struct {
	Fname    string
	Lname    string
	Name     string   `mapstructure:"mname"`
	Version  string   `mapstructure:"version"`
	Author   string   `mapstructure:"author"`
	Route    Route    `mapstructure:"route"`
	Routers  Router   `mapstructure:"routers"`
	SConfig  SConfig  `mapstructure:"system"`
	Cron     Cron     `mapstructure:"cron"`
	EmailTpl EmailTpl `mapstructure:"email"`
	Menu     []SMenu  `mapstructure:"menu"`
	Model    []SModel `mapstructure:"model"`
	Tables   []STable `mapstructure:"table"`
	Cache    []SCache `mapstructure:"cache"`
}

func (b *ModuleConfig) Start() {
	mData := map[string]interface{}{
		"mname":   b.Name,
		"author":  b.Author,
		"vname":   toVendorName(b.Name),
		"version": b.Version,
		"psr":     toPsr(b.Fname, b.Lname),
	}
	// 1.module directory
	// 2.registeration.php
	fp.TouchFile("/registration.php", mData, "/registration.tpl")
	// 3.composer.json
	fp.TouchFile("/composer.json", mData, "/composer.tpl")
	// 4.etc/module.xml
	fp.MakeDir("/etc")
	fp.TouchFile("/etc/module.xml", mData, "/etc/module.tpl")
	// 5.readme.md
	fp.TouchFile("/README.md", mData, "/README.tpl")
}

func (b *ModuleConfig) Routef() {
	rData := map[string]interface{}{
		"mname": b.Name,
		"area":  "admin",
		"rname": b.Route.Admin,
	}
	if b.Route.Admin != "" {
		fp.MakeDir("/etc" + "/adminhtml")
		fp.TouchFile("/etc"+"/adminhtml/routes.xml", rData, "/etc/routes.tpl")
	}
	if b.Route.Front != "" {
		rData["area"] = "standard"
		rData["rName"] = b.Route.Front
		fp.MakeDir("/etc" + "/frontend")
		fp.TouchFile("/etc"+"/frontend/routes.xml", rData, "/etc/routes.tpl")
	}
}

func getViewMass(url string, m Mass) string {
	return `
	$this->getMassactionBlock()->addItem(
		'` + m.Id + `',
		[
			'label' => __('` + m.Label + `'),
			'url' => $this->getUrl('` + url + `'),
			'confirm' => __('Are you sure to ` + m.Id + `?')
		]
	);
	`
}

func getViewColStr(col GridCol) string {
	var d string
	d += `'index' => '` + col.Id + `',`
	if col.Label == "" {
		col.Label = col.Id
	}
	if col.Type != "" {
		d += `'type' => '` + col.Type + `',`
	}
	d += `'header' => __('` + col.Label + `'),`
	if col.Hcss != "" {
		d += `'header_css_class' => '` + col.Hcss + `',`
	}
	if col.Ccss != "" {
		d += `'column_css_class' => '` + col.Ccss + `',`
	}
	return `
	$this->addColumn(
		'` + col.Id + `',
		[
			` + d + `
		]
	);
	`
}

func getViewExport(url string, label string) string {
	return `
	$this->addExportType($this->getUrl('` + url + `', ['_current' => true]),__('` + label + `'));
	`
}

func getAdminView(c Ctl) {
	data := map[string]interface{}{
		"fname":       MConf.Fname,
		"lname":       MConf.Lname,
		"pname":       MConf.Lname,
		"cname":       strings.Title(c.Rd),
		"cnameL":      c.Rd,
		"st":          c.St,
		"nd":          c.Nd,
		"ndname":      strings.Title(c.Nd),
		"viewCols":    "",
		"viewExports": "",
		"viewMass":    "",
		"factory":     c.Factory + "Factory",
	}
	// fmt.Println(data)
	// 1.layout
	fp.MakeDir("/view")
	fp.MakeDir("/view/adminhtml")
	fp.MakeDir("/view/adminhtml/layout")
	// 1.1 index.xml
	fp.TouchFile("/view/adminhtml/layout/"+c.St+"_"+c.Nd+"_"+c.Rd+".xml", data, "/view/adminhtml/index.tpl")
	// 1.2 grid.xml
	fp.TouchFile("/view/adminhtml/layout/"+c.St+"_"+c.Nd+"_grid.xml", data, "/view/adminhtml/grid.tpl")
	// 2.template phtml
	fp.MakeDir("/view/adminhtml/templates")
	fp.TouchFile("/view/adminhtml/templates/"+c.Rd+".phtml", data, "/view/adminhtml/templates.tpl")
	// 3.Block
	fp.MakeDir("/Block")
	fp.MakeDir("/Block/Adminhtml")
	// 3.1 block
	fp.TouchFile("/Block/Adminhtml/"+strings.Title(c.Nd)+".php", data, "/Block/admin.tpl")
	// 3.2 grid
	fp.MakeDir("/Block/Adminhtml/" + strings.Title(c.Nd))
	var viewColStr, viewEportStr, viewMassStr string
	for _, col := range c.Cols {
		viewColStr += getViewColStr(col)
	}
	for _, e := range c.Exports {
		viewEportStr += getViewExport(c.St+"/"+c.Nd+"/"+e.Id, e.Label)
		// TODO 生成控制器
		data["cname"] = e.Id
		fp.TouchFile("/Contorller/Adminhtml/"+strings.Title(c.Nd)+"/"+e.Id+".php", data, "/Controller/admin_export.tpl")
	}
	for _, m := range c.Mass {
		viewMassStr += getViewMass(c.St+"/"+c.Nd+"/"+m.Action, m)
		// TODO 生成控制器
		data["cname"] = m.Action
		fp.TouchFile("/Contorller/Adminhtml/"+strings.Title(c.Nd)+"/"+m.Action+".php", data, "/Controller/admin_mass.tpl")
	}
	data["viewCols"] = viewColStr
	data["viewExports"] = viewEportStr
	data["viewMass"] = viewMassStr
	fp.TouchFile("/Block/Adminhtml/"+strings.Title(c.Nd)+"/Grid.php", data, "/Block/admin_grid.tpl")
}

func (c Ctl) IsView(area string) {
	if c.View {
		switch area {
		case "admin":
			getAdminView(c)
		}
	}
}

func (b *ModuleConfig) Contorller() {
	if len(b.Routers.Areas) == 0 {
		return
	}
	fp.MakeDir("/Contorller")
	fp.MakeDir("/Contorller/Adminhtml")

	ctlData := map[string]interface{}{
		"fname":  b.Fname,
		"lname":  b.Lname,
		"extend": "\\Magento\\Framework\\App\\Action\\Action implements \\Magento\\Framework\\App\\CsrfAwareActionInterface",
	}
	var pName, cName, cTpl, fn string
	for _, r := range b.Routers.Areas {
		// fmt.Println("-------------------------------------------------------")
		// fmt.Println(r.Id)
		ctlTpl := "front"
		areaTargetDir := "/Contorller"
		if r.Id == "admin" {
			areaTargetDir += "/Adminhtml"
			ctlTpl = "admin"
			ctlData["extend"] = "\\Magento\\Backend\\App\\Action"
		}
		if r.Id == "base" {
			ctlTpl = "base"
		}
		for _, c := range r.Ctls {
			cTpl = ctlTpl
			if c.Extend == "" {
				cTpl = ctlTpl + "_empty"
			}
			pName = strings.Title(c.Nd)
			cName = strings.Title(c.Rd)
			ctlData["pname"] = pName
			ctlData["cname"] = cName
			if c.Extend != "" {
				ctlData["extend"] = c.Extend
			}
			// fmt.Println("-----------------------")
			fp.MakeDir(areaTargetDir + "/" + pName)
			fn = areaTargetDir + "/" + pName + "/" + cName + ".php"
			if pName == "" {
				fn = areaTargetDir + "/" + cName + ".php"
			}
			fp.TouchFile(fn, ctlData, "/Controller/"+cTpl+".tpl")
			c.IsView(r.Id)
		}
	}
}

func generateSystemFile(sTab STab) string {
	tabAllStr := `
		<tab id="{{ .Id }}" translate="label" sortOrder="1">
			{{ .LabelStr }}
		</tab>
	`
	sectionAllStr := `
		<section id="{{ .Id }}" translate="label" sortOrder="1" showInDefault="1" showInWebsite="1" showInStore="1">
			{{ .LabelStr }}
			{{ .CssStr }}
			{{ .TabStr }}
			{{ .ResourceStr }}
			{{ .GroupAllStr }}
		</section>
	`
	groupAllStr := `
		<group id="{{ .Id }}" translate="label" sortOrder="1" showInDefault="1" showInWebsite="1" showInStore="1">
			{{ .LabelStr }}
			{{ .FieldAllStr }}
		</group>
	`
	fieldAllStr := `
		<field id="{{ .Id }}" showInDefault="1" showInStore="1" showInWebsite="1" sortOrder="1" translate="label" type="{{ .Type }}">
			{{ .LabelStr }}
			{{ .CommentStr }}
		</field>
	`
	var tmpSections, tmpGroups, tmpFields string
	var fss, fgs, ffs string
	// Father Tab
	fts := strings.ReplaceAll(tabAllStr, "{{ .Id }}", sTab.Id)
	if sTab.Label != "" {
		fts = strings.ReplaceAll(fts, "{{ .LabelStr }}", getSysLabel(sTab.Label))
	}
	ts := getSysTab(sTab.Id)
	tmpSections = ""
	for _, s := range sTab.Sections {
		fss = strings.ReplaceAll(sectionAllStr, "{{ .Id }}", s.Id)
		fss = strings.ReplaceAll(fss, "{{ .TabStr }}", ts)
		if s.Label != "" {
			fss = strings.ReplaceAll(fss, "{{ .LabelStr }}", getSysLabel(s.Label))
		}
		if s.Css != "" {
			fss = strings.ReplaceAll(fss, "{{ .CssStr }}", getSysCss(s.Css))
		}
		if s.Resource != "" {
			fss = strings.ReplaceAll(fss, "{{ .ResourceStr }}", getSysResource(s.Resource))
		}

		tmpGroups = ""
		for _, g := range s.Groups {
			fgs = strings.ReplaceAll(groupAllStr, "{{ .Id }}", g.Id)
			if g.Label != "" {
				fgs = strings.ReplaceAll(fgs, "{{ .LabelStr }}", getSysLabel(g.Label))
			}
			tmpFields = ""
			for _, f := range g.Fields {
				ffs = strings.ReplaceAll(fieldAllStr, "{{ .Id }}", f.Id)
				if f.Label != "" {
					ffs = strings.ReplaceAll(ffs, "{{ .LabelStr }}", getSysLabel(f.Label))
				}
				if f.Comment != "" {
					ffs = strings.ReplaceAll(ffs, "{{ .CommentStr }}", getSysComment(f.Comment))
				}
				ffs = strings.ReplaceAll(ffs, "{{ .Type }}", f.Type)
				tmpFields += ffs
			}
			fgs = strings.ReplaceAll(fgs, "{{ .FieldAllStr }}", tmpFields)
			tmpGroups += fgs
		}
		fss = strings.ReplaceAll(fss, "{{ .GroupAllStr }}", tmpGroups)
		tmpSections += fss
	}
	return fts + tmpSections
}

func (b *ModuleConfig) Systemf() {
	if len(b.SConfig.Tabs) == 0 {
		return
	}
	fp.MakeDir("/etc/adminhtml")
	systemAllStr := `
		<system>
			{{ .System }}
		</system>
	`
	tmpSystem := ""
	for _, t := range b.SConfig.Tabs {
		// fmt.Println(t.Id)
		// fmt.Println(t.Label)
		tmpSystem += generateSystemFile(t)
	}
	systemAllStr = strings.ReplaceAll(systemAllStr, "{{ .System }}", tmpSystem)
	fp.TouchFile("/etc/adminhtml/system.xml", map[string]interface{}{
		"systemConfig": systemAllStr,
	}, "/etc/adminhtml/system.tpl")
}

func (b *ModuleConfig) Cronf() {
	if len(b.Cron.Groups) == 0 {
		return
	}
	fp.MakeDir("/Cron")
	groupStr := `
		<group id="{{ .id }}">
			{{ .jobs }}
		</group>
	`
	var groups, jobs, gs, js string
	jobStr := `
		<job instance="{{ .fname }}\{{ .lname }}\Cron\{{ .cname }}" method="{{ .method }}" name="{{ .name }}">
			<schedule>{{ .schedule }}</schedule>
		</job>
	`
	jobStr = strings.ReplaceAll(jobStr, "{{ .fname }}", b.Fname)
	jobStr = strings.ReplaceAll(jobStr, "{{ .lname }}", b.Lname)

	for _, g := range b.Cron.Groups {
		for _, j := range g.Jobs {

			js = replaceAll(jobStr, map[string]interface{}{
				"fname":    b.Fname,
				"lname":    b.Lname,
				"name":     j.Name,
				"cname":    j.Instance,
				"method":   j.Method,
				"schedule": j.Schedule,
			})

			jobs += js
			fp.TouchFile("/Cron/"+j.Instance+".php", map[string]interface{}{
				"fname":  b.Fname,
				"lname":  b.Lname,
				"cname":  j.Instance,
				"method": j.Method,
			}, "/Cron/cron.tpl")
		}
		gs = replaceAll(groupStr, map[string]interface{}{
			"id":   g.Id,
			"jobs": jobs,
		})
		groups += gs
	}
	fp.TouchFile("/etc/crontab.xml", map[string]interface{}{
		"cronData": groups,
	}, "/etc/crontab.tpl")
}

func (b *ModuleConfig) Emailf() {
	if len(b.EmailTpl.Tpls) == 0 {
		return
	}
	fp.MakeDir("/view")
	fp.MakeDir("/view/adminhtml")
	fp.MakeDir("/view/adminhtml/email")
	fp.MakeDir("/view/frontend")
	fp.MakeDir("/view/frontend/email")
	tStr := `
		<template id="{{ .id }}" label="{{ .label }}" file="{{ .file }}" type="{{ .type }}" module="{{ .module }}" area="{{ .area }}"/>
	`
	tStr = strings.ReplaceAll(tStr, "{{ .module }}", b.Name)
	var allTs, ts string
	data := map[string]interface{}{
		"subject": "",
		"header":  b.EmailTpl.Theader,
		"content": "<div>email content</div>",
		"footer":  b.EmailTpl.Tfooter,
	}
	for _, t := range b.EmailTpl.Tpls {

		ts = replaceAll(tStr, map[string]interface{}{
			"id":    t.Id,
			"label": t.Label,
			"file":  t.File,
			"type":  t.Type,
			"area":  t.Area,
		})

		allTs += ts
		data["subject"] = t.Subject
		if t.Theader != "" {
			data["theader"] = t.Theader
		}
		if t.Tfooter != "" {
			data["tfooter"] = t.Tfooter
		}
		fp.TouchFile("/view/"+t.Area+"/email/"+t.File, data, "/view/email.tpl")
	}
	fp.TouchFile("/etc/email_templates.xml", map[string]interface{}{
		"templates": allTs,
	}, "/etc/email_templates.tpl")
}

func (c *SMenu) GenerateMenuStr() string {
	menuStr := `
		<add id="{{ .mname }}::{{ .id }}" title="{{ .title }}" module="{{ .mname }}" sortOrder="1" parent="{{ .parent }}" {{ .action }} resource="{{ .resource }}"/>
	`
	if c.Action != "" {
		c.Action = `action="` + c.Action + `"`
	}

	var ms string
	ms = replaceAll(menuStr, map[string]interface{}{
		"mname":    c.Module,
		"id":       c.Id,
		"title":    c.Title,
		"parent":   c.Parent,
		"resource": c.Resource,
		"action":   c.Action,
	})

	return ms
}

func gMenuStr(c *SMenu) string {
	return c.GenerateMenuStr()
}

func (b *ModuleConfig) Menuf() {
	var ms string
	for _, m := range b.Menu {
		if m.Module == "" {
			m.Module = b.Name
		}
		ms += gMenuStr(&m)
	}
	if ms != "" {
		fp.MakeDir("/etc")
		fp.MakeDir("/etc/adminhtml")
		fp.TouchFile("/etc/adminhtml/menu.xml", map[string]interface{}{
			"menu": ms,
		}, "/etc/adminhtml/menu.tpl")
	}
}

func (m *SModel) File(data map[string]interface{}) {
	data["pk"] = m.Pk
	data["cname"] = m.Name
	data["table"] = m.Table
	fp.TouchFile("/Model/"+m.Name+".php", data, "/Model/model.tpl")
	fp.TouchFile("/Model/ResourceModel/"+m.Name+".php", data, "/Model/resource.tpl")
	fp.MakeDir("/Model/ResourceModel/" + m.Name)
	fp.TouchFile("/Model/ResourceModel/"+m.Name+"/Collection.php", data, "/Model/collection.tpl")
}

func (b *ModuleConfig) Modelf() {
	if len(b.Model) == 0 {
		return
	}
	fp.MakeDir("/Model")
	fp.MakeDir("/Model/ResourceModel")
	data := map[string]interface{}{
		"fname": b.Fname,
		"lname": b.Lname,
		"pk":    "",
		"cname": "",
		"table": "",
	}
	for _, m := range b.Model {
		if m.Pk == "" {
			m.Pk = "id"
		}
		m.File(data)
	}
}

func getColumnStr(c SColumn) string {
	extra := "["
	if !c.Primary {
		na := "true"
		if c.Nullable {
			na = "true"
		}
		extra += "'nullable' => " + na
		extra += ", 'default' => " + c.Default
		if c.Unsign {
			extra += ", 'unsigned' => true"
		}
		if c.Identity {
			extra += ", 'identity' => true"
		}
	} else {
		extra += "'identity' => true, 'unsigned' => true, 'nullable' => false, 'primary' => true"
	}
	extra += "]"
	var cType string
	switch c.Type {
	case "int":
		cType = "Table::TYPE_INTEGER"
	case "big_int":
		cType = "Table::TYPE_BIGINT"
	case "varchar":
		cType = "Table::TYPE_TEXT"
	case "ts":
		cType = "Table::TYPE_TIMESTAMP"
	default:
		cType = c.Type
	}
	lenStr := "null"
	if c.Length > 0 {
		lenStr = strconv.Itoa(c.Length)
	}
	return `
		->addColumn(
			'` + c.Id + `',
			` + cType + `,
			` + lenStr + `,
			` + extra + `,
			'` + c.Comment + `'
		)
	`
}

func getIndexStr(i SIndex) string {
	cols := strings.Split(i.Id, ",")
	if len(cols) == 0 {
		return ""
	}
	colStr := "["
	for _, c := range cols {
		colStr += "'" + c + "',"
	}
	colStr = colStr[:len(colStr)-1]
	colStr += "]"
	var iType string
	switch i.Type {
	case "unique":
		iType = ",['type' => \\Magento\\Framework\\DB\\Adapter\\AdapterInterface::INDEX_TYPE_UNIQUE]"
	case "pk":
		iType = ",['type' => \\Magento\\Framework\\DB\\Adapter\\AdapterInterface::INDEX_TYPE_PRIMARY]"
	case "full":
		iType = ",['type' => \\Magento\\Framework\\DB\\Adapter\\AdapterInterface::INDEX_TYPE_FULLTEXT]"
	}
	return `
		->addIndex(
			$setup->getIdxName($tableName, ` + colStr + `),
			` + colStr + iType + `
		)
	`
}

func (b *ModuleConfig) SetUpf() {
	if len(b.Tables) == 0 {
		return
	}
	hasDir := false
	fp.MakeDir("/Setup")
	var ts string
	var tMod SModel
	tData := map[string]interface{}{
		"fname": b.Fname,
		"lname": b.Lname,
		"pk":    "",
		"cname": "",
		"table": "",
	}
	for _, t := range b.Tables {
		ts += `$tableName = "` + t.Id + `";`
		ts += `if (!$setup->tableExists($setup->getTable($tableName))) {`
		ts += `$table = $setup->getConnection()
					->newTable($setup->getTable($tableName))`
		for _, c := range t.Colums {
			ts += getColumnStr(c)
		}
		for _, i := range t.Index {
			ts += getIndexStr(i)
		}
		ts += `;`
		ts += `$setup->getConnection()->createTable($table);`
		ts += `}`
		if t.TMod {
			if !hasDir {
				fp.MakeDir("/Model")
				fp.MakeDir("/Model/ResourceModel")
			}
			hasDir = true
			tMod.Name = t.Mname
			tMod.Table = t.Id
			tMod.Pk = t.Pk
			tMod.File(tData)
		}
	}
	fp.TouchFile("/Setup/UpgradeSchema.php", map[string]interface{}{
		"fname":       b.Fname,
		"lname":       b.Lname,
		"version":     b.Version,
		"versionData": ts,
	}, "/Setup/up.tpl")
}

func (c *SCache) GetCacheStr() string {
	var lb, des string
	if c.Label != "" {
		lb = "<label>" + c.Label + "</label>"
	}
	if c.Description != "" {
		des = "<description>" + c.Description + "</description>"
	}
	return `
	<type name="` + c.Id + `" translate="` + c.Trans + `" instance="` + MConf.Fname + `\` + MConf.Lname + `\Model\Cache\` + c.Instance + `">
        ` + lb + `
        ` + des + `
    </type>
	`
}

func (b *ModuleConfig) Cachef() {
	if len(b.Cache) == 0 {
		return
	}
	fp.MakeDir("/Model")
	fp.MakeDir("/Model/Cache")
	var cs string
	data := map[string]interface{}{
		"fname": b.Fname,
		"lname": b.Lname,
		"id":    "",
		"tag":   "",
		"cname": "",
	}
	for _, c := range b.Cache {
		cs += c.GetCacheStr()
		data["id"] = c.Id
		data["tag"] = c.Tag
		data["cname"] = c.Instance
		fp.TouchFile("/Model/Cache/"+c.Instance+".php", data, "/Model/cache.tpl")
	}
	fp.TouchFile("/etc/cache.xml", map[string]interface{}{
		"cacheData": cs,
	}, "/etc/cache.tpl")
}

func getSysLabel(s string) string {
	return makeLabel("label", s)
}
func getSysCss(s string) string {
	return makeLabel("class", s)
}
func getSysTab(s string) string {
	return makeLabel("tab", s)
}
func getSysResource(s string) string {
	return makeLabel("resource", s)
}
func getSysComment(s string) string {
	return makeLabel("comment", s)
}
func makeLabel(label string, val string) string {
	return "<" + label + ">" + val + "</" + label + ">"
}
func toVendorName(mname string) string {
	return strings.ToLower(strings.ReplaceAll(mname, "_", "/"))
}
func toPsr(f string, l string) string {
	return strings.ReplaceAll(strings.Title(f+" "+l+" "), " ", "\\\\")
}

var tt1 *template.Template
var tt1Buf bytes.Buffer

func replaceAll(targetStr string, data map[string]interface{}) string {
	tt1Buf.Reset()
	tt1 = template.Must(tt1.Parse(targetStr))
	tt1.Execute(&tt1Buf, data)
	return tt1Buf.String()
}

var fp FilePath
var MConf ModuleConfig

func NewModuleConfig(outDir string) {
	viper.SetConfigFile("./module.yaml")
	viper.SetConfigName("module")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	viper.Unmarshal(&MConf)

	names := getNames(MConf.Name)
	if len(names) != 2 {
		panic("module name error")
	}
	MConf.Fname = names[0]
	MConf.Lname = names[1]

	tt1 = template.New("")
	fp = NewFilePath(outDir)
	fp.OutputDir = fp.PWD + "/_output"
	fp.MakeDir(fp.OutputDir + "/" + MConf.Fname)
	fp.MakeDir(fp.OutputDir + "/" + MConf.Fname + "/" + MConf.Lname)
	fp.TplDir = fp.PWD + "/template"
	fp.TargetDir = fp.OutputDir + "/" + MConf.Fname + "/" + MConf.Lname
}

func getNames(n string) []string {
	return strings.Split(n, "_")
}

//  -----------------------------------------------------------------  //

type FilePath struct {
	OutputDir string
	TplDir    string
	TargetDir string
	PWD       string
}

func NewFilePath(baseDir string) FilePath {
	if baseDir == "" {
		dir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}
		baseDir = dir
	}
	return FilePath{
		PWD: baseDir,
	}
}

func (f *FilePath) MakeDir(path string) {
	path = f.TargetDir + path
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return
	}
	ok := os.Mkdir(path, os.FileMode(0755))
	if ok != nil {
		panic(ok.Error())
	}
}

func (f *FilePath) TouchFile(filename string, data map[string]interface{}, tmpPath string) {
	filename = f.TargetDir + filename
	tmpPath = f.TplDir + tmpPath
	fmt.Printf("generate file: %s\n", filename)
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return
	}
	tpl, err := template.ParseFiles(tmpPath)
	if err != nil {
		panic(err.Error())
	}
	file, _ := os.Create(filename)
	defer file.Close()
	tpl.Execute(file, data)
}
