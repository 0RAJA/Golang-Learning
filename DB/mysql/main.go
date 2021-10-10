package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //匿名导包
)

/*
上面的mysql驱动中引入的就是mysql包中各个init()方法，你无法通过包名来调用包中的其他函数。
导入时，驱动的初始化函数会调用sql.Register将自己注册在database/sql包的全局变量sql.drivers中，以便以后通过sql.Open访问。

func Open(driverName, dataSourceName string) (*DB, error)
●driverName: 使用的驱动名.这个名字其实就是数据库驱动注册到database/sql时所使用的名字.
●dataSourceName:数据库连接信息，这个连接包含了数据库的用户名,密码,数据库主机以及需要连接的数据库名等信息.
●db, err := sql.Open( "mysql", "用户名:密码@tcp(IP:端口)/数据库?charset=utf8")

DB
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
Exec执行一次命令（包括查询、删除、更新、插入等），不返回任何执行结果。参数args表示query中的占位参数。

func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
Query执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。

func (db *DB) QueryRow(query string, args ...interface{}) *Row
QueryRow执行一次查询，并期望返回最多一行结果（即Row）。
QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）

func (db *DB) Prepare(query string) (*Stmt, error)
Prepare创建一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。

func (db *DB) Begin() (*Tx, error)
Begin开始一个事务。隔离水平由数据库驱动决定。

Rows
Rows是查询的结果。它的游标指向结果集的第零行，使用Next方法来遍历各行结果

func (rs *Rows) Columns() ([]string, error)
Columns返回列名。如果Rows已经关闭会返回错误。

func (rs *Rows) Scan(dest ...interface{}) error
Scan将当前行各列结果填充进dest指定的各个值中

func (rs *Rows) Next() bool
Next准备用于Scan方法的下一行结果。如果成功会返回真，如果没有下一行或者出现错误会返回假。Err应该被调用以区分这两种情况。

func (rs *Rows) Close() error
Close关闭Rows，阻止对其更多的列举。
如果Next方法返回假，Rows会自动关闭，满足。检查Err方法结果的条件。
Close方法时幂等的（多次调用无效的成功），不影响Err方法的结果。
*/

type Person struct {
	name string
	sal  float64
	tel  int
	job  int
	id   int
}

func main() {
	DB, err := sql.Open("mysql", "root:WW876001@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	//DB.SetConnMaxLifetime(100) //数据库最大连接数
	//DB.SetMaxIdleConns(10) //数据库最大闲置连接数
	if false { //验证连接
		if err := DB.Ping(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Connect success")
	}
	//DML
	if false {
		//增,删,改数据
		_, err := DB.Exec("INSERT INTO person (`name`,sal,tel,job)VALUES (?,?,?,?)", "种昭", 99999, 566, 2) //增加
		_, err = DB.Exec("DELETE FROM person WHERE `name` = '种昭'")                                        //删除
		_, err = DB.Exec("UPDATE person SET sal = 99999 WHERE `name` = '种昭'")                             //修改
		if err != nil {
			fmt.Println(err)
		}
	}
	//DQL
	if false {
		name := "Raja"
		rows, err := DB.Query("SELECT * FROM person WHERE `name` = ? ", name) //多条数据查询
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() { //遍历查询结果
			var p Person
			err = rows.Scan(&p.name, &p.sal, &p.tel, &p.job, &p.id)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(p.name, p.sal, p.tel, p.job, p.id)
		}
		if err := rows.Err(); err != nil {
			fmt.Println(err)
		}
		rows.Close()
	}
	if false {
		name := "种昭"
		var (
			tel int
			sal float64
		)
		err = DB.QueryRow("SELECT tel,sal FROM person WHERE `name` = ?", name).Scan(&tel, &sal) //单条数据查询
		fmt.Println(tel, sal)
	}
	if true { //准备查询--将占位符作为参数传入
		stmt, err := DB.Prepare("SELECT `name` FROM person WHERE sal > ?")
		if err != nil {
			fmt.Println(err)
		}
		rows, err := stmt.Query(50000)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var name string
			_ = rows.Scan(&name)
			fmt.Println(name)
		}
		if err := rows.Err(); err != nil {
			fmt.Println(err)
		}
		rows.Close()
	}
}
