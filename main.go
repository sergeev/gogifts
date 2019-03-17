package main

import (
	"fmt"
	"log"
	"net/http"

	helper "./helpers"
)

func main() {

	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()

	// Создаем форму регистрации
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		/*
			// Для тестирования.
			for key, value := range r.Form {
				fmt.Printf("%s = %s\n", key, value)
			}
		*/

		uName = r.FormValue("username")     // Имя или псевдоним
		email = r.FormValue("email")        // Почта
		pwd = r.FormValue("password")       // Пароль
		pwdConfirm = r.FormValue("confirm") // Подтверждение пароля

		// Проверяем пустые данные
		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		// Если ошибка то выводим ее
		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "There is empty data.")
			log.Printf("Server: There is empty data")
			return
		}

		if pwd == pwdConfirm {
			// Здесь сохраняем в базу данных (username, email and password)
			fmt.Fprintln(w, "Registration successful.")
			log.Printf("Server: Registration successful.")
		} else {
			fmt.Fprintln(w, "Password information must be the same.")
			log.Printf("Server: Password information must be the same.")
		}
	})

	// Создаем форму входа в систему
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")  // Data fr	om the form
		pwd = r.FormValue("password") // Data from the form

		// Проверяем пустые данные
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "There is empty data.")
			log.Printf("Server: There is empty data.")
			return
		}

		dbPwd := "12345"                // Симуляция данных в базе
		dbEmail := "vsm@gmail.com" 		// Симуляция данных в базе

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintln(w, "Login succesful!")
			log.Printf("Server: Login succesful!")
		} else {
			fmt.Fprintln(w, "Login failed!")
			log.Printf("Server: Login failed!")
		}
	})

	http.ListenAndServe(":8080", mux)
}