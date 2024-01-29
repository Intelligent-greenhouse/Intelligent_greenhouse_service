# cypunsource-auth


### 克隆

#### 克隆项目与子模块

```bash
git clone --recurse-submodules git@github.com:cypunsource/cypunsource-auth.git
```

#### 更新子模块
```bash
git pull
git submodule update --remote
git add proto
git commit -m "feat(proto): updated submodule to latest commit"
git push
```


### 构建

#### 构建api

```bash
go run generate_proto.go api
```

#### 构建配置

```bash
go run generate_proto.go conf
```

#### 构建项目为可执行文件
```bash
go build -ldflags "-X main.Version=`git describe --tags --always`" -o ./bin/app ./cmd
```

#### 构建容器
```bash
docker build -t cypunsource/cypunsource-tool:`git describe --tags --always` .
```
