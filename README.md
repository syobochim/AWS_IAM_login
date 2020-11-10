# AWS IAM Login Tool

ターミナルから実行することで、IAMアカウントを使ってブラウザログインできるツール。  
複数の種類のユーザーでログインする必要がある場合に便利。

動作確認はMacOSのみ実施。

## setup
ChromeDriverが必要。以下のコマンドを実行してインストールする。

```
$ brew cask install chromedriver
```

IAMのログイン情報は、exeファイルと同じディレクトリに `login_info.tsv` というファイル名で保存する。


```
├── aws_login
└── login_info.tsv
```

`login_info.tsv`はAccountID、LoginID、Passwordをタブ区切りで記載する。

```
<AccountID>    <LoginID>   <Password>
```

例 : それぞれの値の間はTab
```
123456789012	syobochim	Password!1234
```

## 実行方法
以下のコマンドで実行する。
`login_info.tsv`に記載しているアカウントが一覧で表示されるので、ログインしたいアカウントに対して先頭のIDをターミナルに入力すると、対象のアカウントでログインできる。

```
$ ./aws_login
どのアカウントでログインしますか？ idを入力してください。
id :  1 / accountId : 123456789012 / loginId : account1
id :  2 / accountId : 123456789012 / loginId : account2
1
```

## もしもchrome driverが開けないときは

以下のようなポップアップが表示されてChrome Driverが開けない時は、一度 **「キャンセル」** を押す。

<img width="420" alt="alert" src="https://user-images.githubusercontent.com/3622660/98622630-4b6fbc80-234d-11eb-96e9-48f55d741823.png">

システム環境設定の「セキュリティとプライバシー」を開き、一般タブに表示されている **「このまま許可」** をクリックする。

<img width="668" alt="setting" src="https://user-images.githubusercontent.com/3622660/98622720-7e19b500-234d-11eb-92cf-237656f7703d.png">

再度コマンドを実行すると以下のようなポップアップに変わっているので、 **「開く」** をクリックする。

<img width="490" alt="popup" src="https://user-images.githubusercontent.com/3622660/98622848-c89b3180-234d-11eb-93a4-a8b8f6d22d59.png">

## How to build

```
$ go get github.com/sclevine/agouti
$ go build aws_login.go
```
