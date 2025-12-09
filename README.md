# 留言板

一个使用 Go 语言和 [Fiber](https://gofiber.io/) 框架编写的，用于发布和检索留言的简单 RESTful API 服务。

## 功能特性

-   发布新留言。
-   获取分页的留言列表。
-   支持 Docker 进行容器化部署。
-   支持开发环境下的实时热重载。

## 环境要求

-   [Go](https://go.dev/dl/) (版本 1.25 或更高)
-   [Docker](https://www.docker.com/products/docker-desktop/) 和 [Docker Compose](https://docs.docker.com/compose/install/)
-   [`air`](https://github.com/cosmtrek/air) 用于开发环境热重载 (可选)
-   [`wrk`](https://github.com/wg/wrk) 用于性能测试 (可选)

## 快速开始

请按照以下步骤在您的本地机器上运行本项目。

### 1. 克隆代码仓库

```bash
git clone https://github.com/your-username/leave-a-message.git
cd leave-a-message
```
**注意:** 请将 `your-username` 替换为您的 GitHub 用户名。

### 2. 启动数据库

本项目使用 MySQL 数据库。为了方便，项目提供了一个 `docker-compose` 文件。

```bash
docker-compose -f compose-mysql.yml up -d
```

```mysql
create database message;
```

该命令将在后台启动一个 MySQL 容器。

### 3. 配置环境变量

复制环境变量示例文件，并根据需要进行修改。默认值已配置为可与 `docker-compose` 启动的数据库配合使用。

```bash
cp .env.example .env
```

### 4. 安装依赖

安装项目所需的 Go 模块。

```bash
go mod tidy
```

### 5. 运行应用

您可以通过两种模式运行本应用：

**开发模式 (带热重载):**

此模式需要先安装 [`air`](https://github.com/cosmtrek/air)。

```bash
air
```

当您修改代码后，服务会自动重启。

**生产模式:**

```bash
go run main.go
```

API 服务将会运行在 `http://localhost:3000` (或您在 `.env` 文件中指定的端口)。

### 6. 填充数据 (可选)

项目提供了一个 shell 脚本，可以向数据库中填充 1000 条示例留言。

```bash
./populate_messages.sh
```

## API 接口

API 的基础 URL 为 `/api/v1`。

### 创建一条留言

-   **接口:** `POST /message`
-   **描述:** 创建一条新的留言。
-   **请求体:**

    ```json
    {
        "message": "你好，世界！"
    }
    ```

-   **`curl` 示例:**

    ```bash
    curl -X POST -H "Content-Type: application/json" \
         -d '{"message":"这是一条测试留言"}' \
         http://localhost:3000/api/v1/message
    ```

### 获取留言列表

-   **接口:** `GET /message`
-   **描述:** 获取分页的留言列表。
-   **查询参数:**
    -   `page` (可选): 要获取的页码。默认值: `1`。
    -   `limit` (可选): 每页的留言数量。默认值: `10`。
-   **`curl` 示例:**

    ```bash
    curl "http://localhost:3000/api/v1/message?page=2&limit=20"
    ```

## 编译与运行

您也可以编译应用为二进制文件并直接运行。

### 1. 编译应用

```bash
go build -o leave-a-message main.go
```

此命令会生成一个名为 `leave-a-message` 的可执行文件。

### 2. 运行应用

请确保您已经创建了 `.env` 文件。

```bash
./leave-a-message
```

## 性能测试

QPS=并发/RT

使用 `wrk` 进行了性能基准测试。部分结果如下：

<details>
<summary>点击查看 wrk 基准测试结果</summary>

**并发数: 1**
平均 RT: ~13.91ms, QPS: ~72

```
wrk -c1 -t1 -d10s http://x.x.x.x:31532/api/v1/message?limit=10&page=1
Running 10s test @ http://x.x.x.x:31532/api/v1/message?limit=10&page=1
  1 threads and 1 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    13.91ms    3.46ms  48.24ms   91.04%
    Req/Sec    72.28      7.71    90.00     60.00%
  726 requests in 10.07s, 0.99MB read
Requests/sec:     72.11
Transfer/sec:    101.20KB
```

**并发数: 2**
平均 RT: ~13.08ms, QPS: ~156

```
wrk -c2 -t1 -d10s http://x.x.x.x:31532/api/v1/message?limit=10&page=1
Running 10s test @ http://x.x.x.x:31532/api/v1/message?limit=10&page=1
  1 threads and 2 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    13.08ms    5.25ms  58.30ms   89.70%
    Req/Sec   156.40     22.27   202.00     69.00%
  1567 requests in 10.06s, 2.15MB read
Requests/sec:    155.79
Transfer/sec:    218.62KB
```
</details>