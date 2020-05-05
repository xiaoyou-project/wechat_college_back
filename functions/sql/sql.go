package sql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL" //这里我们导入驱动
	"github.com/Unknwon/goconfig"
	"strconv"
	"strings"
)

//数据库连接池
var DB *sql.DB

/**
读取配置文件
*/
func GetConfig(moudle string) map[string]string {
	//载入文件
	config, err := goconfig.LoadConfigFile("app.ini")
	if err != nil {
		return nil
	}
	//获取配置文件
	glob, _ := config.GetSection(moudle)
	//返回配置文件
	return glob
}

/**
初始化数据库连接池
*/
func Init_Db() bool {
	//读取配置
	data := GetConfig("mysql")
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{data["username"], ":", data["password"], "@tcp(", data["ip"], ":", data["port"], ")/", data["DBname"], "?charset=utf8mb4&collation=utf8mb4_unicode_ci"}, "")
	////打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sqltool-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	////设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	////设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	////验证连接
	if err := DB.Ping(); err != nil {
		return false
	}
	return true
}

/**
数据库重连函数
*/
func DB_ReConnect() {
	if err := DB.Ping(); err != nil {
		Init_Db()
	}
}

/**
关闭数据库
*/
func DB_close() error {
	err := DB.Close()
	if err != nil {
		return err
	}
	return nil
}

/**
数据库查询语句
*/
func Sql_dql(sql string) ([][]string, error) {
	fmt.Println(sql)
	//数据库自动重连
	DB_ReConnect()
	var result [][]string
	rows, err := DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//这边是利用反射还有接口来获取所有内容
	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.New("查询行数失败")
	}
	pointers := make([]interface{}, len(cols))
	_ = rows.Scan(pointers...)
	container := make([]string, len(cols))
	for i, _ := range pointers {
		pointers[i] = &container[i]
	}
	result = append(result, container)
	for rows.Next() {
		err = rows.Scan(pointers...)
		container := make([]string, len(cols))
		for i, _ := range pointers {
			pointers[i] = &container[i]
		}
		result = append(result, container)
	}
	if len(result) != 1 {
		result = result[:len(result)-1]
	}
	return result, nil
}

/**
获取数据库的类型并转换为map
*/
func Sql_map(sql string) (map[string]string, error) {
	result, err := Sql_dql(sql)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	data := make(map[string]string)
	for i := 0; i < len(result); i++ {
		data[result[i][0]] = result[i][1]
	}
	return data, nil
}

/**
sql执行语句
*/
func Sql_dml(sql string) bool {
	fmt.Println(sql)
	//数据库自动重连
	DB_ReConnect()
	tx, err := DB.Begin()
	if err != nil {
		return false
	}
	_, err = tx.Exec(sql)
	if err == nil {
		err = tx.Commit()
		if err == nil {
			return true
		}
	}
	return false
}

/**
sql执行语句插入后返回id
*/
func Sql_dml_id(sql string) (bool, string) {
	fmt.Println(sql)
	//数据库自动重连
	DB_ReConnect()
	tx, err := DB.Begin()
	if err != nil {
		return false, ""
	}
	res, err := tx.Exec(sql)
	if err == nil {
		err = tx.Commit()
		if err == nil {
			id, err := res.LastInsertId()
			//这里再执行语句，获取自增id
			if err == nil {
				return true, strconv.Itoa(int(id))
			}
		}
	}
	return false, ""
}
