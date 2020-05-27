package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	userName  string = "root"
	password  string = "111"
	ipAddrees string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "stu"
	charset   string = "utf8"
)

func connectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}
func addRecord(Db *sqlx.DB) {
	//for i := 0; i < 2; i++ {
	result, err := Db.Exec("insert into user (age,username,password,email,sex) values(?,?,?,?,?)", 10, "cq", "123", "123456@11.com", "man")
	if err != nil {
		fmt.Printf("data insert faied, error:[%v]", err.Error())
		return
	}
	id, _ := result.LastInsertId()
	fmt.Printf("insert success, last id:[%d]\n", id)
	//}
}

func updateRecord(Db *sqlx.DB) {
	//更新uid=1的username
	result, err := Db.Exec("update user set username = 'cqqqqc' where id = 3")
	if err != nil {
		fmt.Printf("update faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("update success, affected rows:[%d]\n", num)
}

func deleteRecord(Db *sqlx.DB) {
	//删除uid=2的数据
	result, err := Db.Exec("delete from user where id = 4")
	if err != nil {
		fmt.Printf("delete faied, error:[%v]", err.Error())
		return
	}
	num, _ := result.RowsAffected()
	fmt.Printf("delete success, affected rows:[%d]\n", num)
}

/*Query() 方法使用（查询单个字段数据）


func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
Query() 方法返回的是一个 sql.Rows 类型的结果集

也可以用来查询多个字段的数据，不过需要定义多个字段的变量进行接收

迭代后者的 Next() 方法，然后使用 Scan() 方法给对应类型变量赋值，以便取出结果，最后再把结果集关闭（释放连接）*/
func queryData(Db *sqlx.DB) {
	rows, err := Db.Query("select * from user")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	for rows.Next() {
		//定义变量接收查询数据
		var uid int
		var create_time, username, password, department, email string

		err := rows.Scan(&uid, &create_time, &username, &password, &department, &email)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		fmt.Println(uid, create_time, username, password, department, email)
	}

	//关闭结果集（释放连接）
	rows.Close()
}

// Get() 方法使用
//func (db *DB) Get(dest interface{}, query string, args ...interface{}) error
//是将查询到的一条记录，保存到结构体 结构体的字段名首字母必须大写，不然无法寻址
func getData(Db *sqlx.DB) {

	type userInfo struct {
		Id       int    `db:"id"`
		UserName string `db:"username"`
		Age      string `db:"age"`
		Password string `db:"password"`
		Sex      string `db:"sex"`
		Email    string `db:"email"`
	}

	//初始化定义结构体，用来存放查询数据
	var userData *userInfo = new(userInfo)
	err := Db.Get(userData, "select *from user where id = 1")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}
	//打印结构体内容
	fmt.Println(userData.Id, userData.UserName,
		userData.Password, userData.Age, userData.Email, userData.Sex)
}

/*Select() 方法使用
func (db *DB) Select(dest interface{}, query string, args ...interface{}) error
将查询的多条记录，保存到结构体的切片中
结构体的字段名首字母必须大写，不然无法寻址*/
func selectData(Db *sqlx.DB) {
	type userInfo struct {
		Id       int    `db:"id"`
		UserName string `db:"username"`
		Age      string `db:"age"`
		Password string `db:"password"`
		Sex      string `db:"sex"`
		Email    string `db:"email"`
	}

	//定义结构体切片，用来存放多条查询记录
	var userInfoSlice []userInfo
	err := Db.Select(&userInfoSlice, "select * from user where username like '%cqq%'")
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return
	}

	//遍历结构体切片
	for _, userData := range userInfoSlice {
		fmt.Println(userData.Id, userData.UserName,
			userData.Password, userData.Age, userData.Email, userData.Sex)
	}

}

func main() {
	var Db *sqlx.DB = connectMysql()
	defer Db.Close()

	//addRecord(Db)
	//updateRecord(Db)
	//deleteRecord(Db)
	//queryData(Db)
	//getData(Db)
	selectData(Db)
}
