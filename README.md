# CombinationSelect

Dynamic Programmingで組み合わせを見つける関数をGoで試しに書いてみる。
ついでにGoでAPIも書いてみる。

## dockerで環境作る
```
docker-compose up -d
docker-compose exec app bash
```

## test実行
```
go test -v ./...
```

## main実行
```
go run cmd/combination-select/main.go
```

## POSTで結果を返すようにした
```
curl -X POST -i localhost:8080/combination-select -d @testData.json
```