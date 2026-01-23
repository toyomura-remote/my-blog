
- **Dockerコンテナの強制停止**
	docker rm -f $(docker ps -aq)

- コンテナ内は独自のディレクトリ構成が存在する

- GOのコンテナを使う際はbuildにDockerfileを指定して、独自のコンテナを作る
	- `build: ./backend` ← Dockerfileが存在する階層

- volumesはローカルとコンテナ間の差分を埋める

- Airはファイルの変更（差分）を検知すると、コンテナを「立ち上げ直す」のではなく、中の「プログラムをビルド・実行し直す
	- コンテナ自体は壊さず、起動したまま中のプロセス（Goの実行ファイル）だけを高速で入れ替える
	- 変更を検知した瞬間、Airはコンテナ内で「ビルド（コンパイル）」を走らせる
		- Goは「書いたコードをそのまま実行」できない言語なので、Airが裏側で `go build` を瞬時に実行し、新しい「実行ファイル」を作ってから動かす
- `go.mod` が書き換えられれば、自動でインストール（ダウンロード）される

- PostgreSQLのイメージを作成すると、コンテナ内に`/docker-entrypoint-initdb.d/`が存在しており、そこにsqlファイルを入れると、初回のみそのSQL文が走る
	- dbのvolumesに下記のように書く
	- `- ./db/init.sql:/docker-entrypoint-initdb.d/init.sql`

