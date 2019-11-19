# 雰囲気ですら理解しないDocker入門

理解はおいといてとりあえず使ってみようの精神で触ってみようなハンズオン
最低限使えるようになろう

- docker run 
    - コンテナを引っ張ってきて起動するコマンド
- docker exec
    - コンテナ内でコマンドを実行するコマンド
- docker ps
    - コンテナの起動状況を確認するコマンド
- docker stop
    - コンテナを止めるコマンド(ころころする)
- docker rm
    - コンテナを削除するコマンド(成仏してクレメンス)

# docker run　のオプション
- iオプション -i
    - ホスト -> コンテナ の入出力をつなげる
- tオプション -t
    - コンテナ -> ホスト の入出力をつなげる
- dオプション -d
    - バックグラウンド実行
- vオプション -v
    - ファイルを共有する
- pオプション -p
    - ポートをつなげる
- wオプション -w
    - 作業ディレクトリを指定する
- nameオプション --name
    - コンテナに名前をつける
- rmオプション --rm
    - コンテナをstopしたら自動でrmする

**オプションitdはハッピーセット**

# docker run の実際の例
```
$ docker run -itd --name python -v $(PWD):/py -p 8000:8000 -w="/py" python

$ docker run -itd --name [CONTAINER NAME] -v $(PWD):[CONTAINER DIRCTORY] -p [LOCAL PORT]:[CONTAINER PORT] -w="[CONTAINER DIRCTORY]" [IMAGE NAME]
```

$(PWD) はローカルのカレントディレクトリを展開(pwdコマンドの実行)
# docker-compose
docker-compose.ymlにコンテナの概要を記述し、Dockerファイルを用いてコンテナの詳細を定義する。
とりあえず触ってみようか

## こまんど

- docker-compose run CONTAINER_NAME COMMAND
    - 指定したコンテナでコマンドを実行する (up より前に)
- docker-compose up
    - コンテナを全て起動する、ログ一生吐き続けるので -d をつけよう
- docker-compose exec CONTAINER_NAME COMMAND
    - 指定したコンテナでコマンドを実行する (up より後に)
- docker-compose down
    - docker-composeで管理している全てのコンテナに対してstopとrmをする
- docker-compose ps 
    - コンテナの状態の確認
- docker-compose logs CONTAINER_NAME
    - 指定コンテナのログをみる


## つかう

- docker-compose up -d
- docker-compose ps
- docker-compose exec nginx bash
- docker-compose down

# おまけ

便利なやつ

- docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q)
    - 全てのコンテナをなきものにする




