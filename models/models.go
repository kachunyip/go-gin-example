package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/kachunyip/go-gin-example/pkg/setting"
	"log"
)

var db *gorm.DB

//CREATE TABLE `blog_tag` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`name` varchar(100) DEFAULT '' COMMENT '标签名称',
//`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
//`created_by` varchar(100) DEFAULT '' COMMENT '创建人',
//`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
//`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
//`deleted_on` int(10) unsigned DEFAULT '0',
//`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

//CREATE TABLE `blog_article` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
//`title` varchar(100) DEFAULT '' COMMENT '文章标题',
//`desc` varchar(255) DEFAULT '' COMMENT '简述',
//`content` text,
//`created_on` int(11) DEFAULT NULL,
//`created_by` varchar(100) DEFAULT '' COMMENT '创建人',
//`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
//`modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
//`deleted_on` int(10) unsigned DEFAULT '0',
//`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';


//CREATE TABLE `blog_auth` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`username` varchar(50) DEFAULT '' COMMENT '账号',
//`password` varchar(50) DEFAULT '' COMMENT '密码',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;
//
//INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');


type Model struct {
	ID	int	`gorm:"primary_key" json:"id"`
	CreatedOn	int	`json:"created_on"`
	ModifiedOn	int	`json:"modified_on"`
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
