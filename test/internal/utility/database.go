package utility

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	user := os.Getenv("MYSQL_USER")
	// 環境変数"~~"の値をuserに格納
	// if user == "" {
	// 	user = "sharin" // デフォルトの値
	// }
	pw := os.Getenv("MYSQL_PASSWORD")
	// if pw == "" {
	// 	pw = "0530" // デフォルトの値
	// }
	db_name := os.Getenv("MYSQL_DATABASE")
	// if db_name == "" {
	// 	db_name = "love" // デフォルトの値
	// }
	//以上、コメントアウト部は環境変数をプログラム内で指定したい場合の記述方法
	//ハードコーティングなのでやめましょう
	//(テスト時ならokだろうけど忘れて公開されたら
	// コメントアウトしてあっても終わりだからやめた方が良さげ？)

	var path string = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
	var err error

	fmt.Printf("%s\n%s\n", path, err)

	if Db, err = sql.Open("mysql", path); err != nil {
		fmt.Printf("database.goのinitでエラー発生:err=%s, path=%s", err, path)
		// log.Fatal("Db open error:", err.Error())
	}
	fmt.Printf("%s\n%s\n", path, err)
	checkConnect(1)

}

// mysqlを使っています。golangからSQLへ接続する際、Database名の指定はしなくて良いのでしょうか？
// > GolangからMySQLへ接続する際には、Database名を指定する必要があります。
// >データベース名は接続情報の一部であり、MySQLサーバー上の特定のデータベースに接続するために必要です。
//
// user := "ユーザー名"
// password := "パスワード"
// host := "ホスト名"
// port := "ポート番号"
// database := "データベース名"

// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

// db, err := sql.Open("mysql", dsn)
// if err != nil {
//     log.Fatal(err)
// }
// defer db.Close()

func checkConnect(count uint) {
	var err error
	if err = Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		if count > 0 {
			checkConnect(count)
		} else {
			fmt.Println("Connection retries exhausted err")
			fmt.Printf("err=%s", err)
			return
		}
	} else {
		fmt.Println("db connected!!")
	}
}
