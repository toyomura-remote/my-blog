`defer`
その関数が終了(return)するまで実行を後回しにするもの

#### Atlas

**Docker内に入れる:**
RUN curl -sSf https://atlasgo.sh | sh

**curlも入れる**
RUN apk add --no-cache git curl libc6-compat

**dbのコンテナは2つ作る必要はない**
docker作成時に空のDBを作っておく
ー ./backend/docker/postgres/local/init:/docker-entrypoint-initdb.d

atlas.hcl内にsrc = "file://db/original" のようにばしょを記述すれば、自動で参照してくれる　

**init.sqlに記述**
CREATE DATABASE atlas_dev;
:/docker-entrypoint-initdb.dは初回起動時に実行される(psqlのみ？)

**N+1**
1つのクエリで複数の行を取得する場合(limit=10などで)、取得した後にループ内でさらにクエリを実行してしまうと、ループの分だけクエリが実行されてパフォーマンスが落ちる。例えば投稿を10件取得して、それに付随するユーザー情報を取得する場合、ループの中でユーザー取得のクエリがループの分発行される
**qm.Load**をすると自動で紐づけをしてクエリは2回のみの発行
ormを使わない場合はjoinやin句で対応

gqlでloadを使うと、例えばpostのtitleのみ欲しい場合とかでもジョインされたクエリが走ってしまう。
dataloaderならユーザーが指定した情報のみを纏めてひろえる