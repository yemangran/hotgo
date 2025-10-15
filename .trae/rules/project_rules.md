# GoFrame & HotGo 开发指南

## 1. 项目架构概述

### 1.1 项目简介
这是一个基于 GoFrame 框架并使用 HotGo 进行二次开发的项目。项目采用了标准的 GoFrame 分层架构，并集成了 HotGo 的扩展功能。

### 1.2 核心入口
- @main.go - 项目主入口文件，负责初始化和启动应用
- @internal/cmd/cmd.go - 命令行入口处理

### 1.3 项目目录结构
```
项目结构
├── api/                # API接口定义
├── internal/           # 核心业务代码
│   ├── controller/     # 控制器层，处理请求响应
│   ├── service/        # 服务层（自动生成），实现业务接口定义
│   ├── logic/          # 业务逻辑层，实现具体业务逻辑
│   ├── dao/            # 数据访问层（自动生成），数据库操作封装
│   ├── model/          # 数据模型层，定义数据结构
│   │   ├── do/         # 数据对象（自动生成）
│   │   ├── entity/     # 实体对象（自动生成）
│   │   └── input/      # 输入参数对象（可手动修改）
│   ├── router/         # 路由配置，定义API路由规则
│   ├── consts/         # 常量定义
│   ├── library/        # 公共库函数
│   ├── websocket/      # WebSocket相关实现
│   ├── queues/         # 队列任务处理
│   ├── crons/          # 定时任务实现
│   └── global/         # 全局变量和初始化
├── addons/             # 插件模块目录
│   ├── xxx/            # 具体插件目录
│   │   ├── api/        # 插件API定义
│   │   ├── controller/ # 控制器
│   │   ├── service/    # 服务接口（自动生成）
│   │   ├── logic/      # 业务逻辑实现
│   │   └── ...         # 其他插件模块
│   └── modules/        # 插件模块注册目录
├── manifest/           # 项目配置清单
├── resource/           # 资源文件（配置文件、模板文件、静态资源）
├── storage/            # 存储目录
├── utility/            # 工具函数
└── hack/               # 开发工具和脚本
```

### 1.4 分层架构
```
Controller层 -> Service层(接口) -> Logic层(实现) -> Dao层
     ↓             ↓               ↓              ↓
  请求处理     接口定义        业务逻辑实现     数据访问
```

## 2. ⚠️ 自动生成代码规则

### 2.1 禁止手动修改的目录
以下目录由工具自动生成，禁止手动修改：
- `internal/service/` - 由 `gf gen service` 自动生成
- `internal/dao/` - 由 `gf gen dao` 自动生成
- `internal/model/do/` - 由 `gf gen dao` 自动生成
- `internal/model/entity/` - 由 `gf gen dao` 自动生成
- `addons/xxx/service/` - 由 `gf gen service` 自动生成
- `addons/xxx/dao/` - 由 `gf gen dao` 自动生成
- `addons/xxx/model/do/` - 由 `gf gen dao` 自动生成
- `addons/xxx/model/entity/` - 由 `gf gen dao` 自动生成
- `addons/xxx/api/xxx/xxx.go` - 由 `gf gen ctrl` 自动生成
- `addons/xxx/api/login/login.go` - 由 `gf gen ctrl` 自动生成

### 2.2 代码生成工具使用
1. 生成 dao 层代码：
```bash
gf gen dao
```

2. 生成控制器：
```bash
# 核心模块控制器
gf gen ctrl -m

# 插件模块控制器
gf gen ctrl -s=addons/xxx/api/api -d=addons/xxx/controller/api -m
```

3. 生成服务层：
```bash
# 核心模块服务
gf gen service -s=internal/logic -d=internal/service

# 插件模块服务
gf gen service -s=addons/xxx/logic -d=addons/xxx/service
```

4. 注意事项：
   - 生成前确保数据库连接配置正确
   - 修改数据库表结构后需要重新生成
   - 生成的代码不要手动修改，以免被覆盖

## 3. 核心开发指南 (internal/)

### 3.1 开发流程
1. 设计数据库表结构（如需要）
2. 使用 `gf gen dao` 生成数据访问层
3. 在 `api` 目录下定义接口
4. 使用 `gf gen ctrl` 生成控制器骨架
5. 在 `logic` 层实现业务逻辑
6. 使用 `gf gen service` 生成服务接口

### 3.2 控制器实现示例
```go
// internal/controller/user/user.go
func (c *cUser) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
    list, totalCount, err := service.User().GetList(ctx, &req.UserListInp)
    if err != nil {
        return nil, err
    }
    res = new(v1.ListRes)
    res.List = list
    res.PageRes.Pack(req, totalCount)
    return res, nil
}
```

### 3.3 逻辑层实现示例
```go
// internal/logic/user/user.go
type sUser struct{}

func init() {
    service.RegisterUser(NewUser())
}

func NewUser() *sUser {
    return &sUser{}
}

func (s *sUser) GetList(ctx context.Context, in *model.UserListInp) (res *model.UserListModel, totalCount int64, err error) {
    // 1. 业务逻辑处理
    
    // 2. 数据库查询
    m := dao.User.Ctx(ctx)
    if in.Username != "" {
        m = m.Where(dao.User.Columns.Username, in.Username)
    }
    
    // 3. 分页查询
    result, err := m.Page(in.Page, in.PageSize).ScanAndCount(&res,&totalCount,false)
    if err != nil {
        return nil, 0, gerror.Wrap(err, "获取用户列表失败")
    }
    
    // 4. 数据转换与返回
    return res, totalCount, nil
}
```

## 4. 插件开发指南 (addons/)

### 4.1 目录结构
```
addons/
  ├── xxx/                      # 具体插件目录
  │   ├── api/                  # 插件API定义
  │   │   ├── admin/            # 后台管理接口
  │   │   └── api/              # 前端API接口
  │   ├── controller/           # 控制器
  │   ├── service/              # 服务接口（自动生成）
  │   ├── logic/                # 业务逻辑实现
  │   │   ├── admin/            # 后台逻辑
  │   │   ├── api/              # 前端逻辑
  │   │   └── common/           # 公共逻辑
  │   ├── model/                # 数据模型
  │   │   └── input/            # 输入模型
  │   │       ├── adminin/      # 后台输入输出结构体模型
  │   │       └── apiin/        # 前端输入输出结构体模型
  │   ├── dao/                  # 数据访问对象（自动生成）
  │   └── router/               # 路由配置
  └── modules/                  # 插件模块注册目录
```

### 4.2 插件开发流程
1. 创建插件目录结构
2. 定义数据库表结构（如需要）
3. 使用 `gf gen dao` 生成数据访问层
4. 在 `api` 目录下定义接口
5. 使用 `gf gen ctrl` 生成控制器骨架
6. 在 `logic` 层实现业务逻辑
7. 使用 `gf gen service` 生成服务接口
8. 在 `modules` 中注册插件

### 4.3 插件注册示例
```go
// addons/modules/xxx.go
package modules
import _ "hotgo/addons/xxx"
```

### 4.4 命令生成
- 生成控制器: `gf gen ctrl -s=addons/xxx/api/api -d=addons/xxx/controller/api -m`
- 生成Service: `gf gen service -s=addons/xxx/logic -d=addons/xxx/service`

## 5. 数据模型与数据库操作

### 5.1 数据模型定义
```go
// 输入参数
type UserListInp struct {
    form.PageReq
    Username string `json:"username" dc:"用户名"`
    Status   int    `json:"status" dc:"状态"`
}

// 输出参数
type UserListModel struct {
    Username string `json:"username" dc:"用户名"`
    Status   int    `json:"status" dc:"状态"`
    ...
}
```

### 5.2 数据访问层结构
- `internal/dao/` - 数据访问对象（自动生成）
  - `internal/dao/internal/` - 自动生成的具体实现代码
- `internal/model/` - 数据模型定义
  - `do/` - 数据对象定义（自动生成）
  - `entity/` - 实体对象定义（自动生成）
  - `input/` - 输入参数对象（可手动修改）
    - `adminin/` - 管理后台参数
    - `apiin/` - 前端API参数

### 5.3 输入参数对象示例
```go
// 输入参数
type UserInp struct {
    Username string  `json:"username" dc:"用户名"`
}

// 输出参数
type UserModel struct {
    Username string `json:"username" dc:"用户名"`
    Mobile string `json:"mobile" dc:"手机号码"`
}
```

### 5.4 数据库操作规范
1. 所有字段查询必须使用 `dao.Table.Columns.FieldName`
2. 常用查询方法:
```go
// 单条记录查询
one, err := dao.User.Ctx(ctx).Where(dao.User.Columns.Id, id).One()

// 多条记录查询
list, err := dao.User.Ctx(ctx).Where(dao.User.Columns.Status, 1).All()

// 分页查询
list, totalCount, err := dao.User.Ctx(ctx).Page(page, pageSize).ScanAndCount(&result, &count, false)

// 条件构建查询
m := dao.User.Ctx(ctx)
if username != "" {
    m = m.Where(dao.User.Columns.Username, username)
}
if status > 0 {
    m = m.Where(dao.User.Columns.Status, status)
}
list, err := m.Order(dao.User.Columns.Id + " DESC").All()
```

3. 事务处理示例:
```go
err := dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
    // 事务操作1
    _, err := dao.User.Ctx(ctx).TX(tx).Data(data1).Insert()
    if err != nil {
        return err
    }
    
    // 事务操作2
    _, err = dao.UserLog.Ctx(ctx).TX(tx).Data(data2).Insert()
    if err != nil {
        return err
    }
    
    return nil
})
```

### 5.5 高级查询技巧
1. 关联查询:
```go
// 使用关联表查询
result, err := dao.User.Ctx(ctx).
    LeftJoin(dao.UserRole.Table+" r", "r.user_id="+dao.User.Table+"."+dao.User.Columns.Id).
    Fields(dao.User.Table+".*", "r.role_name").
    Where(dao.User.Columns.Status, 1).
    All()
```

2. 聚合查询:
```go
// 统计查询
count, err := dao.User.Ctx(ctx).Where(dao.User.Columns.Status, 1).Count()

// 聚合函数
result, err := dao.Order.Ctx(ctx).
    Fields("user_id, SUM(amount) as total_amount").
    Group("user_id").
    Having("SUM(amount) > ?", 1000).
    All()
```

## 6. API接口规范

### 6.1 API定义示例
```go
// api/v1/user.go
type ListReq struct {
    g.Meta `path:"/user/list" method:"get" tags:"用户管理" summary:"获取用户列表"`
    model.UserListInp
}

type ListRes struct {
    List  []*model.UserListModel `json:"list" dc:"数据列表"`
    *form.PageRes
}
```

### 6.2 API注解规范
- `path` - 接口路径
- `method` - 请求方法 (get, post, put, delete)
- `tags` - 接口分组
- `summary` - 接口描述
- `v` - 验证规则（如：required, min, max, length, range, email, phone）
- `dc` - 字段说明（用于生成API文档）

### 6.3 参数验证示例
```go
type CreateReq struct {
    g.Meta    `path:"/user/create" method:"post" tags:"用户管理" summary:"创建用户"`
    Username  string `v:"required|length:5,30#用户名不能为空|用户名长度应当在:min到:max之间" dc:"用户名"`
    Password  string `v:"required|length:6,30#密码不能为空|密码长度应当在:min到:max之间" dc:"密码"`
    Mobile    string `v:"required|phone#手机号不能为空|手机号格式不正确" dc:"手机号"`
    Email     string `v:"email#邮箱格式不正确" dc:"邮箱"`
    Status    int    `v:"in:0,1,2#状态值错误" dc:"状态"`
}
```

## 7. 错误处理规范

### 7.1 错误定义
```go
// internal/consts/error_code.go
const (
    CodeOK             = 0      // 成功
    CodeNotAuthorized  = 403    // 未授权
    CodeNotFound       = 404    // 资源不存在
    CodeServerError    = 500    // 服务器内部错误
    CodeBusinessError  = 10000  // 业务错误起始码
    
    // 用户模块错误码 (10100-10199)
    CodeUserNotFound   = 10100  // 用户不存在
    CodePasswordError  = 10101  // 密码错误
    // ...其他错误码
)
```

### 7.2 错误处理方式
1. 返回通用错误:
```go
return nil, gerror.New("用户名已存在")
```

2. 返回带错误码的错误:
```go
return nil, gerror.NewCode(consts.CodeUserNotFound, "用户不存在")
```

3. 包装错误并添加上下文:
```go
if err := doSomething(); err != nil {
    return nil, gerror.Wrap(err, "处理XXX失败")
}
```

4. 在控制器中处理错误:
```go
func (c *cUser) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
    data, err := service.User().Create(ctx, &req.CreateUserInp)
    if err != nil {
        return nil, err
    }
    return &v1.CreateRes{Id: data.Id}, nil
}
```

## 8. 缓存使用规范

### 8.1 缓存操作示例
```go
// 获取缓存
value, err := g.Redis().Do(ctx, "GET", "user:"+gconv.String(userId))

// 设置缓存 (带过期时间)
_, err := g.Redis().Do(ctx, "SETEX", "user:"+gconv.String(userId), 3600, gconv.String(userInfo))

// 删除缓存
_, err := g.Redis().Do(ctx, "DEL", "user:"+gconv.String(userId))
```

### 8.2 缓存最佳实践
1. 缓存键名使用模块前缀，便于管理: `user:profile:1`, `order:list:2`
2. 缓存时间根据数据更新频率设置，频繁更新的数据缓存时间应短
3. 更新数据时主动清除相关缓存
4. 缓存穿透防护:
```go
// 缓存穿透处理示例
func GetUserInfo(ctx context.Context, userId uint64) (*model.UserModel, error) {
    cacheKey := "user:info:" + gconv.String(userId)
    
    // 尝试从缓存获取
    value, err := g.Redis().Do(ctx, "GET", cacheKey)
    if err == nil && !value.IsEmpty() {
        var user *model.UserModel
        if err = gconv.Struct(value, &user); err == nil {
            return user, nil
        }
    }
    
    // 缓存未命中，从数据库查询
    user, err := dao.User.Ctx(ctx).Where(dao.User.Columns.Id, userId).One()
    if err != nil {
        return nil, err
    }
    
    // 数据不存在时也缓存空值，防止缓存穿透（设置较短过期时间）
    if user == nil {
        g.Redis().Do(ctx, "SETEX", cacheKey, 60, "{}")
        return nil, gerror.NewCode(consts.CodeUserNotFound, "用户不存在")
    }
    
    // 存入缓存
    userJson, _ := gjson.Encode(user)
    g.Redis().Do(ctx, "SETEX", cacheKey, 3600, userJson)
    
    return user, nil
}
```

## 9. 日志规范

### 9.1 日志级别使用
1. `Debug`: 开发调试信息
```go
g.Log().Debug(ctx, "处理请求参数", g.Map{"req": req})
```

2. `Info`: 正常流程信息
```go
g.Log().Info(ctx, "用户登录成功", g.Map{"uid": user.Id, "ip": ip})
```

3. `Warn`: 警告信息
```go
g.Log().Warning(ctx, "接口调用频率过高", g.Map{"uid": user.Id, "api": req.URL})
```

4. `Error`: 错误信息
```go
g.Log().Error(ctx, "调用第三方接口失败", g.Map{"api": "xxx", "error": err.Error()})
```

### 9.2 日志上下文
确保日志包含足够的上下文信息:
```go
g.Log().Error(ctx, "订单创建失败", 
    g.Map{
        "userId": userId,
        "orderId": orderId,
        "amount": amount,
        "error": err.Error(),
        "stack": gerror.Stack(err),
    },
)
```

## 10. 安全编码规范

### 10.1 SQL注入防护
正确做法:
```go
// 使用参数绑定方式，防止SQL注入
user, err := dao.User.Ctx(ctx).Where(dao.User.Columns.Username, username).One()

// 复杂条件使用WhereBuilder
user, err := dao.User.Ctx(ctx).WhereBuilder(
    dao.User.Builder().
        Where(dao.User.Columns.Username, username).
        WhereOr(dao.User.Columns.Email, email)
)
```

错误做法:
```go
// 错误：直接拼接SQL
user, err := dao.User.Ctx(ctx).Where("username = '"+username+"'").One()
```

### 10.2 XSS防护
在输出到HTML之前进行转义:
```go
// 使用 ghtml 包进行HTML转义
safeContent := ghtml.Entities(content)
```

### 10.3 CSRF防护
在关键操作时验证Token:
```go
// 生成CSRF Token
token := gsha1.Encrypt(gctx.GetSessionId(ctx) + gtime.Now().String())
g.Session().Set(ctx, "csrf_token", token)

// 验证CSRF Token
func checkCSRFToken(ctx context.Context, token string) bool {
    sessionToken := g.Session().Get(ctx, "csrf_token")
    return sessionToken == token
}
```

### 10.4 敏感数据处理
1. 密码加密:
```go
// 密码加盐哈希
salt := grand.S(10)
passwordHash := gsha256.Encrypt(password + salt)

// 密码验证
calcHash := gsha256.Encrypt(inputPassword + userSalt)
if calcHash != userPasswordHash {
    return false
}
```

2. 数据脱敏:
```go
// 手机号脱敏
maskedMobile := gregex.ReplaceString(`(\d{3})\d{4}(\d{4})`, "$1****$2", mobile)

// 身份证脱敏
maskedID := gregex.ReplaceString(`(\d{6})\d{8}(\w{4})`, "$1********$2", idCard)
```

## 11. 测试规范

### 11.1 单元测试
```go
// internal/logic/user/user_test.go
func Test_User_GetInfo(t *testing.T) {
    // 初始化测试上下文
    ctx := context.Background()
    
    // 测试数据准备
    userId := uint64(1)
    
    // 执行测试逻辑
    user, err := logic.User.GetInfo(ctx, userId)
    
    // 断言结果
    assert.Nil(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, userId, user.Id)
}
```

### 11.2 Mock测试
使用Monkey补丁进行依赖模拟:
```go
// 模拟数据库层
monkey.Patch(dao.User.GetOne, func(ctx context.Context, id uint64) (*entity.User, error) {
    return &entity.User{
        Id: id,
        Username: "test_user",
        Status: 1,
    }, nil
})

// 测试逻辑
user, err := logic.User.GetInfo(ctx, 1)
assert.Nil(t, err)
assert.Equal(t, "test_user", user.Username)

// 恢复原始方法
monkey.UnpatchAll()
```

### 11.3 接口测试
```go
func TestUserApi(t *testing.T) {
    // 创建客户端
    client := ghttp.NewClient()
    
    // 设置请求头
    client.SetHeader("Authorization", "Bearer "+token)
    
    // 发送请求
    resp, err := client.Get("http://localhost:8000/api/v1/user/info?id=1")
    
    // 断言响应
    assert.Nil(t, err)
    assert.Equal(t, 200, resp.StatusCode)
    
    // 解析响应数据
    var respData struct {
        Code    int             `json:"code"`
        Message string          `json:"message"`
        Data    *model.UserInfo `json:"data"`
    }
    err = resp.ParseJson(&respData)
    assert.Nil(t, err)
    assert.Equal(t, 0, respData.Code)
    assert.Equal(t, uint64(1), respData.Data.Id)
}
```

## 12. 性能优化建议

### 12.1 数据库优化
1. 使用索引:
   - 对频繁查询的字段添加索引
   - 避免过度索引，每个索引都会增加写操作的开销
   
2. 大数据量分页优化:
```go
// 使用主键优化大数据量分页
query := dao.User.Ctx(ctx).Where(dao.User.Columns.Status, 1)
if lastId > 0 {
    query = query.Where(dao.User.Columns.Id+">", lastId)
}
list, err := query.Order(dao.User.Columns.Id+" ASC").Limit(pageSize).All()
```

3. 批量操作:
```go
// 批量插入
_, err := dao.User.Ctx(ctx).Data(userList).Batch(100).Insert()

// 批量更新
_, err := dao.User.Ctx(ctx).WhereIn(dao.User.Columns.Id, ids).Data("status", 1).Update()
```

### 12.2 缓存优化
1. 合理使用缓存层级:
   - 本地缓存: 适用于不频繁更新的数据
   - Redis缓存: 适用于需要跨服务共享的数据
   
2. 实体缓存与列表缓存分离处理

### 12.3 并发处理
使用goroutine和channel处理并发任务:
```go
func ProcessUserTasks(ctx context.Context, userIds []uint64) error {
    var (
        wg      sync.WaitGroup
        errChan = make(chan error, len(userIds))
    )
    
    // 并发限制
    limit := 10
    ch := make(chan struct{}, limit)
    
    for _, uid := range userIds {
        wg.Add(1)
        ch <- struct{}{}
        
        go func(userId uint64) {
            defer wg.Done()
            defer func() { <-ch }()
            
            err := processUserTask(ctx, userId)
            if err != nil {
                errChan <- err
            }
        }(uid)
    }
    
    wg.Wait()
    close(errChan)
    
    // 处理错误
    for err := range errChan {
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

## 13. 完整工作流示例

### 13.1 新增功能完整流程

1. **定义模型**:
```go
// internal/model/input/user.go
type CreateUserInp struct {
    Username  string `v:"required|length:5,30" json:"username" dc:"用户名"`
    Password  string `v:"required|length:6,30" json:"password" dc:"密码"`
    Nickname  string `v:"required" json:"nickname" dc:"昵称"`
    Mobile    string `v:"required|phone" json:"mobile" dc:"手机号"`
    Email     string `v:"email" json:"email" dc:"邮箱"`
    Status    int    `json:"status" dc:"状态"`
}

type CreateUserModel struct {
    Id        uint64 `json:"id" dc:"用户ID"`
}
```

2. **定义API**:
```go
// api/v1/user.go
type CreateReq struct {
    g.Meta `path:"/user/create" method:"post" tags:"用户管理" summary:"创建用户"`
    model.CreateUserInp
}

type CreateRes struct {
    Id uint64 `json:"id" dc:"用户ID"`
}
```

3. **生成控制器**:
```bash
gf gen ctrl -m
```

4. **实现控制器**:
```go
// internal/controller/user/user.go
func (c *cUser) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
    data, err := service.User().Create(ctx, &req.CreateUserInp)
    if err != nil {
        return nil, err
    }
    return &v1.CreateRes{Id: data.Id}, nil
}
```

5. **实现业务逻辑**:
```go
// internal/logic/user/user.go
func (s *sUser) Create(ctx context.Context, in *model.CreateUserInp) (out *model.CreateUserModel, err error) {
    // 1. 检查用户名是否已存在
    count, err := dao.User.Ctx(ctx).Where(dao.User.Columns.Username, in.Username).Count()
    if err != nil {
        return nil, err
    }
    if count > 0 {
        return nil, gerror.NewCode(consts.CodeBusiness, "用户名已存在")
    }
    
    // 2. 密码加密处理
    salt := grand.S(10)
    passwordHash := gsha256.Encrypt(in.Password + salt)
    
    // 3. 构建插入数据
    data := do.User{
        Username:     in.Username,
        Password:     passwordHash,
        PasswordSalt: salt,
        Nickname:     in.Nickname,
        Mobile:       in.Mobile,
        Email:        in.Email,
        Status:       in.Status,
        CreateAt:     gtime.Now(),
    }
    
    // 4. 数据库插入
    id, err := dao.User.Ctx(ctx).Data(data).InsertAndGetId()
    if err != nil {
        return nil, gerror.Wrap(err, "创建用户失败")
    }
    
    // 5. 返回结果
    return &model.CreateUserModel{Id: uint64(id)}, nil
}
```

6. **生成服务接口**:
```bash
gf gen service -s=internal/logic -d=internal/service
```

## 14. 常用命令参考

### 14.1 代码生成
- DAO层生成: `gf gen dao`
- 控制器生成: `gf gen ctrl`
- 服务层生成: `gf gen service`

### 14.2 核心模块
- 生成Service: `gf gen service -s=internal/logic -d=internal/service`
- 生成控制器: `gf gen ctrl -m`

### 14.3 插件模块
- 生成Service: `gf gen service -s=addons/xxx/logic -d=addons/xxx/service`
- 生成控制器: `gf gen ctrl -s=addons/xxx/api/api -d=addons/xxx/controller/api -m`

## 15. 最佳实践总结
1. 所有业务逻辑必须在 `logic` 层实现
2. 控制器层只负责请求参数处理和响应
3. 数据库操作统一通过 dao 层处理
4. 使用 `gerror` 包装错误信息，带上详细上下文
5. 正确使用日志级别记录关键信息
6. 敏感数据必须加密存储
7. 接口参数必须严格校验
8. 合理使用缓存提高性能
9. 重要业务必须有单元测试覆盖
10. 遵循代码风格一致性

## 16. 其他重要配置和目录
1. 配置文件位于 `manifest/` 目录
2. 自定义插件开发请放置在 `addons/` 目录
3. 公共函数库统一在 `utility/` 或 `library/` 中维护
4. 存储目录位于 `storage/`
5. 开发工具和脚本位于 `hack/` 目录
