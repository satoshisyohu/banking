package main

import (
	"flag"
	"go_bank/frameworks"
	"go_bank/interface/controllers"
	"log"
)

func main() {

	f := flag.String("db", "none", "createtable")
	flag.Parse()
	log.Println(f)
	if *f == "init" {
		log.Println("a")
	}
	frameworks.Config(f)
	r := controllers.Router()
	r.Run(":3000")

}

/*todo
1.db周りの設定をする
2.dbのテーブルを作成する
3.TBLの要素を考える
→顧客TBL

*/
