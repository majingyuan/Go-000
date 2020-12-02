package Week02

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func queryUser(id int) (user,error){
	var uInfo user
	db, err := sql.Open("mysql", "root:root@tcp(192.168.67.129:3306)/channel_sharding?charset=utf8")
	if err != nil {
		return uInfo,err
    }

	fmt.Println("连接数据库成功")
	
	
	rows, err := db.Query("select id, `user_name`, password from user where id = ?", id);
    
    if rows.Next() {
        var id int
        var userName string
        var passwrod string
        if err:= rows.Scan(&id, &userName, &passwrod); err != nil {
            return uInfo,err
        }
        uInfo = user{id, userName, passwrod};
    }
	
	defer db.Close()
    return uInfo,nil;

}