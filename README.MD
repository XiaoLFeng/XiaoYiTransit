# 公交车辆信息管理系统 (Bus Vehicle Information Management System)

## 1. 项目简介

本项目是一个基于 GoFrame 2.9.X 开发的公交车辆信息管理系统。旨在提供一个高效、稳定、易用的平台，用于管理公交公司的车辆信息、线路信息、司机信息、调度信息以及相关的维护记录等。系统致力于简化公交运营管理流程，提高数据处理效率和准确性。

## 2. 主要功能

* **车辆管理:**
    * [ ] 车辆信息录入（车牌号、型号、购买日期、状态等）
    * [ ] 车辆信息查询、修改、删除
    * [ ] 车辆状态跟踪（运营中、维修中、报废等）
    * [ ] 车辆年检记录
    * [ ] 车辆保险记录
* **线路管理:**
    * [ ] 公交线路信息录入（线路编号、起止站点、途经站点、票价等）
    * [ ] 线路信息查询、修改、删除
    * [ ] 线路站点管理
* **司机管理:**
    * [ ] 司机信息录入（工号、姓名、联系方式、驾照信息等）
    * [ ] 司机信息查询、修改、删除
    * [ ] 司机排班查看
* **调度管理:**
    * [ ] 车辆调度分配（为特定线路分配车辆和司机）
    * [ ] 调度历史记录查询
    * [ ] 实时车辆位置监控 (未来扩展)
* **维修管理:**
    * [ ] 车辆维修申请
    * [ ] 维修记录登记（维修内容、费用、日期等）
    * [ ] 维修历史查询
* **用户与权限管理 (RBAC):**
    * [ ] 用户管理（管理员、调度员、普通用户等）
    * [ ] 角色管理
    * [ ] 权限分配
* **统计报表:**
    * [ ] 车辆使用率统计
    * [ ] 线路运营情况分析
    * [ ] 维修成本统计

*(注：带 `[ ]` 的为计划功能，可根据实际开发进度勾选)*

## 3. 技术选型

* **后端框架:** GoFrame `v2.9.x`
* **数据库:** MySQL (推荐 `8.0+`) / PostgreSQL (可选)
* **前端框架:** React
* **API 文档:** OpenAPI (Swagger) 自动生成 (GoFrame 内置支持)
* **配置文件:** `manifest/config/config.yaml`
* **日志管理:** GoFrame Logging

## 4. 环境要求

* Go `1.24.1`
* MySQL `8.3`
* Git

## 5. 安装与运行

### 5.1. 克隆项目

```bash
git clone [你的项目 Git 仓库地址]
cd [项目目录名称]
```

### 5.2. 配置环境

1.  **修改配置文件:**
    复制 `manifest/config/config.yaml.example` 为 `manifest/config/config.yaml`。
    根据你的本地环境修改 `config.yaml` 文件中的数据库连接信息 (MySQL, Redis等)、服务器端口等配置。

    ```yaml
    # 示例: manifest/config/config.yaml
    # MySQL数据库配置
    database:
      default:
        link: "mysql:user:password@tcp(127.0.0.1:3306)/your_database_name?charset=utf8mb4&parseTime=true&loc=Local"
        debug: true # 开发环境建议开启

    # Redis 配置 (如果使用)
    redis:
      default:
        address: "127.0.0.1:6379"
        db: 0
        # pass: "your_redis_password" # 如果有密码

    # 服务器配置
    server:
      address: ":8000" # API服务监听地址和端口
      # ... 其他配置
    ```

2.  **数据库迁移:**
    * **自动迁移 (推荐):** GoFrame 提供了强大的 migration 工具。
        ```bash
        # 安装 gf-cli (如果尚未安装)
        go install [github.com/gogf/gf/cmd/gf/v2@latest](https://github.com/gogf/gf/cmd/gf/v2@latest)

        # 在项目根目录下执行 migration (假设你的 migration 文件已创建在 internal/packed 目录下或通过 gf gen dao 生成)
        # 如果你是手动创建的 SQL 文件，请先导入到数据库
        # 如果使用 gf gen service 生成了 service，通常会自动处理表结构（如果配置了 autoMigrateTables）
        # 或者，你可以在启动时通过代码逻辑进行自动建表和初始化数据。
        # 例如，在 main.go 中添加:
        # g.DB().AutoMigrate(ctx, &model.YourTable{})
        ```
    * **手动导入:** 将项目提供的 SQL 文件 (例如 `manifest/sql/schema.sql`) 导入到你的数据库中。

### 5.3. 依赖安装

```bash
# 使用 Go Modules 管理依赖
go mod tidy
```

### 5.4. 编译与运行

* **开发环境运行 (推荐):**
  使用 `gf run main.go` 命令，此命令会监听文件变动并自动编译运行。
    ```bash
    gf run main.go
    ```

* **编译后运行:**
    ```bash
    # 编译 (生成可执行文件，例如 bus_admin)
    gf build main.go -n bus_admin -a amd64 -s linux # 示例：编译为 Linux amd64
    # 或直接 go build
    go build -o bus_admin main.go

    # 运行可执行文件
    ./bus_admin
    ```

### 5.5. 访问系统

启动成功后，可以通过浏览器访问：
* **API 服务:** `http://localhost:8000` (或你在 `config.yaml` 中配置的地址和端口)
* **API 文档 (Swagger):** `http://localhost:8000/swagger`

## 6. 目录结构 (基于 GoFrame 推荐结构)

```
.
├── api                     // API 定义 (proto 文件或 Go 定义)
│   └── v1                  // API 版本 v1
│       └── bus_vehicle.go  // 示例：车辆管理 API 定义
├── hack                    // 脚本和工具 (例如 Dockerfile, Makefile)
├── internal                // 内部业务逻辑 (不对外暴露)
│   ├── cmd                 // 应用启动命令相关
│   │   └── server
│   │       └── server.go   // 服务启动逻辑
│   ├── consts              // 常量定义
│   ├── controller          // 控制器层 (HTTP/GRPC Handler)
│   │   └── bus_vehicle.go  // 示例：车辆管理控制器
│   ├── dao                 // 数据访问对象 (与数据库表结构对应)
│   │   └── bus_vehicle.go
│   ├── logic               // 业务逻辑层 (核心业务处理)
│   │   └── bus_vehicle
│   │       └── bus_vehicle.go
│   ├── model               // 数据模型 (请求/响应结构体，内部数据结构)
│   │   ├── do              // 数据库实体对象 (通常由 ORM 生成)
│   │   │   └── bus_vehicle.go
│   │   └── entity          // 数据库实体对象 (通常由 ORM 生成，与 do 类似，看 gf 版本和习惯)
│   │       └── bus_vehicle.go
│   │   └── input           // 输入参数结构体
│   │   └── output          // 输出结果结构体
│   ├── packed              // 静态资源打包 (例如 swagger.json, SQL 迁移文件)
│   ├── service             // 服务接口定义与实现 (对 logic 的封装和组合)
│   │   └── bus_vehicle.go
│   └── utility             // 工具类
├── manifest                // 项目清单和配置
│   ├── config              // 配置文件
│   │   ├── config.yaml
│   │   └── config.yaml.example
│   ├── deploy              // 部署相关 (例如 k8s yaml)
│   ├── document            // 项目文档
│   └── sql                 // SQL 文件 (例如初始化 schema, 种子数据)
├── main.go                 // 程序入口文件
├── go.mod                  // Go Modules 文件
├── go.sum
└── README.md               // 就是你现在看到的这个文件
```

## 7. API 接口文档

API 接口文档通过 Swagger UI 自动生成。当应用启动后，可以通过访问 `/swagger`路径查看和测试 API 接口。
例如：`http://localhost:8000/swagger`

## 8. 数据库设计 (简要)

* **`bv_vehicle` (车辆表):** `id`, `plate_number` (车牌), `model` (型号), `purchase_date`, `status` (状态: 1-运营, 2-维修, 3-停运, 4-报废), `created_at`, `updated_at`, ...
* **`bv_route` (线路表):** `id`, `route_number` (线路编号), `name` (线路名称), `start_station`, `end_station`, `stops` (途经站点, JSON 或关联表), `fare` (票价), `created_at`, `updated_at`, ...
* **`bv_driver` (司机表):** `id`, `employee_id` (工号), `name`, `phone`, `license_number`, `created_at`, `updated_at`, ...
* **`bv_schedule` (调度表):** `id`, `route_id`, `vehicle_id`, `driver_id`, `schedule_date`, `start_time`, `end_time`, `created_at`, `updated_at`, ...
* **`bv_maintenance` (维修记录表):** `id`, `vehicle_id`, `description`, `cost`, `maintenance_date`, `created_at`, `updated_at`, ...
* **`sys_user` (用户表):** ...
* **`sys_role` (角色表):** ...
* **`sys_permission` (权限表):** ...
* **(其他表根据需求添加)**

*(详细的表结构和关系请参考 `manifest/sql/schema.sql` 或对应的数据库迁移文件)*

## 9. 参与贡献

欢迎提交 Pull Request 或 Issue 来改进本项目。

1.  Fork 本仓库
2.  创建你的 Feature 分支 (`git checkout -b feature/AmazingFeature`)
3.  提交你的更改 (`git commit -m 'Add some AmazingFeature'`)
4.  将更改推送到分支 (`git push origin feature/AmazingFeature`)
5.  打开一个 Pull Request

## 10. 许可证

本项目采用 [MIT License](LICENSE) (如果选择 MIT，请添加 LICENSE 文件)。

```

这个 README 模板提供了一个比较全面的起点。你可以根据项目的实际情况进行调整和补充，例如添加更详细的 API 端点列表、更具体的数据库表结构说明，或者前端部分的说明。

希望这个对你有