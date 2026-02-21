# 1. 差分を生成 (例: make diff NAME=add_users)
diff:
	docker compose exec api atlas migrate diff --env local

# 2. DBへ適用
apply:
	docker compose exec api atlas migrate apply --env local

# 3. 1つ前の状態に戻す
down:
	docker compose exec api atlas migrate down --env local

# 4. ハッシュの再計算（手修正した時用）
hash:
	docker compose exec api atlas migrate hash --env local

boiler:
	docker compose exec api sqlboiler psql -o infra/model -p model --no-tests