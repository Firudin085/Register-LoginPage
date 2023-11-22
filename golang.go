package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)


type Person struct{
	DataUsername string 
	DataPassword string 
}

var p Person 


func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        fmt.Println("Method : ", r.Method)
        tmpl, err := template.ParseFiles("./template/register.html")

        if err != nil {
            log.Fatal("Register template acilamadi!")
        }
        tmpl.Execute(w, nil)
    } else if r.Method == "POST" {
        fmt.Println("Method : ", r.Method)

        r.ParseForm()

        username := r.FormValue("username")
        password := r.FormValue("password")

        db, err := sql.Open("mysql", "firudin:123456@tcp(127.0.0.1)/db1")

        if err != nil {
            log.Fatal("Register database acilamadi!")
        }
        defer db.Close()

        // Check if the username already exists
        var count int
		err = db.QueryRow("SELECT COUNT(*) FROM db3 WHERE Username = ? AND Password = ?", username, password).Scan(&count)

        if err != nil {
            log.Fatal("Register datalar cekilemedi!")
        }

        if count > 0 {
            fmt.Println("Bu kullanici adi zaten var!")
			http.Redirect(w, r, "/register", http.StatusSeeOther)


			} else {
            _, err := db.Exec("INSERT INTO db3 (Username, Password) VALUES(?,?)", username, password)

            if err != nil {
                log.Fatal("Yeni datalar kaydedilemedi!")
            }

            fmt.Println("Veriler kaydedildi! Login sayfasına yönlendiriliyorsunuz!")
				http.Redirect(w, r, "/login", http.StatusSeeOther)
        }
    }
}




func MainPage(w http.ResponseWriter,r * http.Request){
	tmpl,err := template.ParseFiles("./template/mainpage.html")

	if err != nil{
		log.Fatal("Mainpage sayfasi acilamadi!")	
	}

	tmpl.Execute(w,nil)


}

func Login(w http.ResponseWriter, r * http.Request){

	if r.Method == "GET"{
		tmpl,err := template.ParseFiles("./template/login.html")

		if err !=nil{
			log.Fatal("Loginpage template acilamadi!")
		}
	
		tmpl.Execute(w,nil)
	}else if r.Method == "POST"{fmt.Println("Method : ", r.Method)

		r.ParseForm()

		Loginusername := r.FormValue("username")
		Loginpassword := r.FormValue("password")

		db, err := sql.Open("mysql", "firudin:123456@tcp(127.0.0.1)/db1")

		if err != nil {
            log.Fatal("Loginpage database acilamadi!!")
        }
		defer db.Close()

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM db3 WHERE Username = ? AND Password = ?", Loginusername, Loginpassword).Scan(&count)

		if err != nil {
            log.Fatal("Loginpage datalar cekilemedi!")
        }

		if count > 0 {
            fmt.Println("Bu kullanici login yapiyor ve datasi bulundu!!")
			UserPage(w,r)
			} else {
            fmt.Println("Boyle bir kullanici yok! Kayit sayfasina yönlendiriliyorsunuz!")
			http.Redirect(w, r, "/register", http.StatusSeeOther)
        }
		

	}


}

type Data struct{
	Username string 
}


func UserPage( w http.ResponseWriter, r*http.Request){

	tmpl,err := template.ParseFiles("./template/userpage.html")
	username := p.DataUsername
	if err != nil{
		log.Fatal("Userpage sayfasi acilamadi!")
	}

	
	data := Data{
		Username: username,
	}


	tmpl.Execute(w, data)

}




func main() {

	http.HandleFunc("/register",Register)
	http.HandleFunc("/",MainPage)
	http.HandleFunc("/login",Login)
	http.ListenAndServe(":9090",nil)
	fmt.Println(":9090 port listening...")
}