package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"                            //驱动名-这个名字其实就是数据库驱动注册到database/sql 时所使用的名字.
	userName   = "root"                             //用户名
	passWord   = "WW876001"                         //密码
	ip         = "127.0.0.1"                        //ip地址
	port       = "3306"                             //端口号
	dbName     = "test"                             //数据库名
	ipFormat   = "%s:%s@tcp(%s:%s)/%s?charset=utf8" //格式
)

type Person struct {
	name string
	sal  float64
	tel  int
	job  int
	id   int
}

func DBInit() (*sql.DB, error) {
	DB, err := sql.Open(driverName, fmt.Sprintf(ipFormat, userName, passWord, ip, port, dbName))
	if err != nil {
		return nil, err
	}

	DB.SetMaxOpenConns(100) //SetMaxOpenConns设置与数据库建立连接的最大数目。
	DB.SetMaxIdleConns(10)  //SetMaxIdleConns设置连接池中的最大闲置连接数。
	err = DB.Ping()
	if err != nil {
		return nil, err
	}
	return DB, nil
}
func Insert(DB *sql.DB, person *Person) error {
	_, err := DB.Exec("INSERT INTO person (`name`,sal,tel,job) VALUES (?,?,?,?)", person.name, person.sal, person.tel, person.job)
	if err != nil {
		return err
	}
	return nil
}

func Delete(DB *sql.DB, data interface{}, way int) error {
	var err error
	switch way {
	case 1:
		_, err = DB.Exec("DELETE FROM person WHERE `name` = ?", data)
	case 2:
		_, err = DB.Exec("DELETE FROM person WHERE sal = ?", data)
	case 3:
		_, err = DB.Exec("DELETE FROM person WHERE tel = ?", data)
	case 4:
		_, err = DB.Exec("DELETE FROM person WHERE job = ?", data)
	case 5:
		_, err = DB.Exec("DELETE FROM person WHERE id = ?", data)
	}
	if err != nil {
		return err
	}
	return nil
}
func Modify(DB *sql.DB, sal float64, name string) error {
	_, err := DB.Exec("UPDATE person SET sal = ? WHERE name = ?", sal, name)
	if err != nil {
		return err
	}
	return nil
}
func Find(DB *sql.DB, name string) ([]Person, error) {
	list := make([]Person, 0)
	rows, err := DB.Query("SELECT * FROM person WHERE `name` = ?", name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p Person
		err = rows.Scan(&p.name, &p.sal, &p.tel, &p.job, &p.id)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}
func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	DB, err := DBInit()
	handleErr(err)
	person := Person{
		name: "王麻子",
		sal:  123321,
		tel:  999,
		job:  2,
		id:   007,
	}
	err = Insert(DB, &person)
	handleErr(err)
	err = Delete(DB, "王麻子", 1)
	handleErr(err)
	err = Modify(DB, 20000, "种昭")
	handleErr(err)
	list, err := Find(DB, "种昭")
	handleErr(err)
	for _, p := range list {
		fmt.Println(p.name, p.sal, p.tel, p.job, p.id)
	}
}
