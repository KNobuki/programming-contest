# 概要
Go 競プロ用のリポジトリ

# ディレクトリ構成
```
.
├── README.md
# コマンド置き場
├── cmd
│   └── makedir
│       └── main.go
# コンテストの解答置き場
├── contests
# contestgenコマンド 実行ファイル
├── contestgen
# テンプレート置き場
└── template
    └── main.go
```

# ContestGen
ContestGenは、競技プログラミング用のコンテストディレクトリと問題ディレクトリを自動生成するためのCLIツールです。ディレクトリ構成とテンプレートファイルをもとに、指定されたコンテスト名と問題数に対応するディレクトリを生成します。

## 使い方
以下のようにコマンドラインから実行できます。

``` shell
./contestgen -contest example_contest -problems 30
```
-contest オプションにコンテスト名を、-problems オプションに問題数を指定してください。contestsディレクトリ配下に指定されたコンテスト名のディレクトリが作成され、問題名ごとにサブディレクトリが作成され、テンプレートファイルがコピーされます。

デフォルトでは、問題数は7に設定されています。-problems オプションを省略した場合、自動的に7問分のディレクトリが作成されます。

