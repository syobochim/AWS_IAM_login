package main

import (
  "github.com/sclevine/agouti"
  "fmt"
  "os"
  "bufio"
  "log"
  "errors"
  "strings"
)

type LoginInfo struct {
  Account string
  LoginID string
  Password string
}

func main() {
    file, err := os.Open("./login_info.tsv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var loginInfos []LoginInfo
    fmt.Println("どのアカウントでログインしますか？ idを入力してください。")

    for i := 0; scanner.Scan(); i++ {
      line := scanner.Text()
      splited := strings.Split(line, "\t")
      if len(splited) != 3 {
        errors.New("パースに失敗しました。ファイルはタブ区切りにしてください。")
      }

      var info LoginInfo
      info.Account = splited[0]
      info.LoginID = splited[1]
      info.Password = splited[2]
      loginInfos = append(loginInfos, info)
      fmt.Printf("id :  %d / accountId : %s / loginId : %s\n", i, info.Account, info.LoginID)
    }

    var n int
    fmt.Scan(&n)

    agoutiDriver := agouti.ChromeDriver()
    agoutiDriver.Start()
    page, _ := agoutiDriver.NewPage()

    page.Navigate("https://" + loginInfos[n].Account + ".signin.aws.amazon.com/console")
    page.FindByID("username").Fill(loginInfos[n].LoginID)
    page.FindByID("password").Fill(loginInfos[n].Password)
    page.FindByID("input_signin_button").Click()
}
