package main

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  // 数据库连接信息（请替换为你的实际信息）
  dsn := "username:password@tcp(127.0.0.1:3306)/dbname"
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  // 验证数据库连接
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  // 安全的查询操作
  userID := 1
  var username string
  err = db.QueryRow("SELECT username FROM users WHERE id=?", userID).Scan(&username)
  if err != nil {
    if err == sql.ErrNoRows {
      fmt.Println("No user found with ID", userID)
    } else {
      log.Fatal(err)
    }
  } else {
    fmt.Println("User found:", username)
  }

  // 安全的更新操作
  newUsername := "new_username"
  _, err = db.Exec("UPDATE users SET username=? WHERE id=?", newUsername, userID)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("User updated successfully")
}