# AWS Login Tool

ターミナルから実行することで、IAMアカウントを使ってブラウザログインできるツール。

動作確認はMacOSのみ実施。

## setup
ChromeDriverが必要。以下のコマンドを実行してインストールする。

```
$ brew cask install chromedriver
```

IAMのログイン情報は、exeファイルと同じディレクトリに `login_info.tsv` というファイル名で保存する。


```
├── README.md
├── aws_login.exe
├── aws_login.go
└── login_info.tsv
```

`login_info.tsv`はAccountID、LoginID、Passwordをタブ区切りで記載する。

```
<AccountID>    <LoginID>   <Password>
```

## 実行方法
以下のコマンドで実行する。
`login_info.tsv`に記載しているアカウントが一覧で表示されるので、ログインしたいアカウントに対して先頭のIDをターミナルに入力すると、対象のアカウントでログインできる。

```
$ ./aws_login.exe
どのアカウントでログインしますか？ idを入力してください。
id :  0 / accountId : 123456789012 / loginId : account1
id :  0 / accountId : 123456789012 / loginId : account2
1
```

