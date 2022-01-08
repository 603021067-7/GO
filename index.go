
package main
import (
	"database/sql"
   _ "github.com/go-sql-driver/mysql"
	"log"
    "fmt"
    "net/http"
    "html/template"
)

type ResultData struct{
    Id int
    Name string
    Price  int
    Amount  int
}
func main() {
    http.HandleFunc("/",result) 
    http.HandleFunc("/delete",delete) 


    http.HandleFunc("/insert",insert)


    http.HandleFunc("/add",func(res http.ResponseWriter,req *http.Request){
    http.ServeFile(res,req,"add.html")
    })

    http.HandleFunc("/updatesuccess",update)



    http.HandleFunc("/update",func(res http.ResponseWriter,req *http.Request){
        http.ServeFile(res,req,"update.html") 
         var db, err = sql.Open("mysql", "root:@(127.0.0.1:3306)/projectgo")
            if err != nil {
                log.Fatal(err)
            }
            defer db.Close()

            stmt,err:=db.Prepare("select id,name,price,amount from product where id =23")
            stmt.Exec(req.URL.Query().Get("id"))
            rows,err:=db.Query("select id,name,price,amount from product where id =23")
            if err!= nil{
                panic(err)
            }
            tRes:=ResultData{}
            var results[]ResultData
                for rows.Next(){
                    var id int
                    var name string
                    var price int
                    var amount int
                err=rows.Scan(&id,&name,&price,&amount)
                    tRes.Id =id
                    tRes.Name =name
                    tRes.Price =price
                    tRes.Amount =amount
                results=append(results,tRes)
                    if err!= nil{
                        panic(err)
                    }
                }
templates.Execute(res,results)
        })


    http.ListenAndServe(":3000",nil) 
}
var templates =template.Must(template.ParseFiles("index.html"))

        func insert(res http.ResponseWriter, req *http.Request){
            var db, err = sql.Open("mysql", "root:@(127.0.0.1:3306)/projectgo")
            if err != nil {
                log.Fatal(err)
            }
            defer db.Close()
            stmt,err:=db.Prepare("insert into product (name,price,amount) values(?,?,?)")
            name := req.FormValue("name")
            price := req.FormValue("price")
            amount := req.FormValue("amount")
             stmt.Exec(name,price ,amount)
            if err != nil {
                fmt.Println(err)
            }
            http.Redirect(res,req,"index.html",301) //เด้งกลับหน้าแรก
        }





        func update(res http.ResponseWriter, req *http.Request){
            var db, err = sql.Open("mysql", "root:@(127.0.0.1:3306)/projectgo")
            if err != nil {
                log.Fatal(err)
            }
            defer db.Close()
            stmt,err:=db.Prepare("update  product set  name=?,price=?,amount=? where id=23")
            name := req.FormValue("name")
            price := req.FormValue("price")
            amount := req.FormValue("amount")
             stmt.Exec(name,price ,amount)
            if err != nil {
                fmt.Println(err)
            }
            http.Redirect(res,req,"index.html",301) //เด้งกลับหน้าแรก
        }






        func delete(res http.ResponseWriter, req *http.Request){
            var db, err = sql.Open("mysql", "root:@(127.0.0.1:3306)/projectgo")
            if err != nil {
                log.Fatal(err)
            }
            stmt,err:=db.Prepare("delete from product where id =?")
            stmt.Exec(req.URL.Query().Get("id"))
            if err != nil {
                panic(err)
            }
            http.Redirect(res,req,"/",301) //เด้งกลับหน้าแรก
        }


     

            func result(res http.ResponseWriter, req *http.Request){
                var db, err = sql.Open("mysql", "root:@(127.0.0.1:3306)/projectgo")
                if err!= nil{
                    log.Fatal(err)
                }
                defer db.Close()
                rows,err:=db.Query("select id,name,price,amount from product ")
                        if err!= nil{
                            panic(err)
                        }
                        tRes:=ResultData{}
                        var results[]ResultData
                            for rows.Next(){
                                var id int
                                var name string
                                var price int
                                var amount int
                            err=rows.Scan(&id,&name,&price,&amount)
                                tRes.Id =id
                                tRes.Name =name
                                tRes.Price =price
                                tRes.Amount =amount
                            results=append(results,tRes)
                                if err!= nil{
                                    panic(err)
                                }
                            }
            templates.Execute(res,results)
            fmt.Println(results)

            }







