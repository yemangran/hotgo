# API接口文档

<cite>
**本文档中引用的文件**  
- [member.go](file://server/api/admin/member/member.go)
- [member.go](file://server/api/api/member/v1/member.go)
- [token.go](file://server/internal/library/token/token.go)
- [error.go](file://server/internal/consts/error.go)
- [member.go](file://server/internal/model/input/adminin/member.go)
</cite>

## 目录
1. [简介](#简介)
2. [管理端API接口](#管理端api接口)
3. [终端用户API接口](#终端用户api接口)
4. [JWT认证与刷新机制](#jwt认证与刷新机制)
5. [前端调用示例](#前端调用示例)
6. [错误码体系](#错误码体系)

## 简介
本文档为 `hotgo-2.0` 项目提供完整的公共API接口说明，涵盖管理端（Admin API）和终端用户端（API）的所有接口。详细描述了每个接口的HTTP方法、URL路径、请求头、请求体结构、响应体格式，并特别说明了JWT认证机制、令牌刷新逻辑以及错误码体系，旨在为前后端开发者提供清晰的集成指导。

## 管理端API接口

本节列出 `server/api/admin/member/member.go` 中定义的所有管理端用户管理接口。

### 用户信息管理

#### 获取用户列表
- **HTTP方法**: `GET`
- **URL路径**: `/admin/member/list`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `page` (int): 页码
  - `limit` (int): 每页数量
  - `id` (int64): 用户ID
  - `pid` (int64): 上级ID
  - `roleId` (int): 角色ID
  - `deptId` (int): 部门ID
  - `mobile` (int): 手机号
  - `username` (string): 用户名
  - `realName` (string): 真实姓名
  - `name` (string): 岗位名称
  - `code` (string): 岗位编码
  - `createdAt` ([]int64): 创建时间范围
- **响应体**:
```json
{
  "list": [
    {
      "id": 1,
      "username": "admin",
      "realName": "管理员",
      "mobile": "13800138000",
      "status": 1,
      "createdAt": "2023-01-01 00:00:00"
    }
  ],
  "page": 1,
  "limit": 10,
  "total": 100
}
```

#### 获取指定用户信息
- **HTTP方法**: `GET`
- **URL路径**: `/admin/member/view`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `id` (int64, required): 用户ID
- **响应体**: 返回 `adminin.MemberViewModel` 结构的用户详细信息。

#### 修改/新增用户
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/edit`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `id` (int64): 用户ID（新增时可为空）
  - `roleId` (int64, required): 角色ID
  - `postIds` ([]int64): 岗位ID数组
  - `deptId` (int64, required): 部门ID
  - `username` (string, required): 帐号
  - `password` (string): 密码（新增或修改时设置）
  - `realName` (string): 真实姓名
  - `avatar` (string): 头像
  - `sex` (int): 性别
  - `qq` (string): QQ
  - `email` (string): 邮箱
  - `birthday` (*gtime.Time): 生日
  - `provinceId` (int): 省
  - `cityId` (int): 城市
  - `areaId` (int): 地区
  - `address` (string): 默认地址
  - `mobile` (string): 手机号码
  - `remark` (string): 备注
  - `status` (int): 状态
- **响应体**: 空对象 `{}`。

#### 删除用户
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/delete`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `id` (interface{}, required): 用户ID
- **响应体**: 空对象 `{}`。

#### 更新用户状态
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/status`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: `AdminMember` 结构体，包含 `id` 和 `status` 字段。
- **响应体**: 空对象 `{}`。

#### 获取登录用户信息
- **HTTP方法**: `GET`
- **URL路径**: `/admin/member/info`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 无
- **响应体**: 返回 `LoginMemberInfoModel` 结构，包含用户权限、余额、积分等信息。

#### 获取用户选项
- **HTTP方法**: `GET`
- **URL路径**: `/admin/member/option`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 无
- **响应体**: 返回 `[]*MemberSelectModel` 数组，用于下拉选择。

### 用户资料与安全

#### 更新用户资料
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/updateProfile`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `avatar` (string, required): 头像
  - `realName` (string, required): 真实姓名
  - `qq` (string): QQ
  - `birthday` (*gtime.Time): 生日
  - `sex` (int): 性别
  - `address` (string): 联系地址
  - `cityId` (int64): 城市编码
- **响应体**: 空对象 `{}`。

#### 修改登录密码
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/updatePwd`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `id` (int64): 用户ID
  - `oldPassword` (string, required): 原密码
  - `newPassword` (string, required): 新密码（6-16位）
- **响应体**: 空对象 `{}`。

#### 重置密码
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/resetPwd`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `password` (string, required): 新密码
  - `id` (int64, required): 用户ID
- **响应体**: 空对象 `{}`。

#### 换绑邮箱
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/updateEmail`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `email` (string, required): 新邮箱
  - `code` (string): 原邮箱验证码
- **响应体**: 空对象 `{}`。

#### 换绑手机号
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/updateMobile`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `mobile` (string, required): 新手机号
  - `code` (string): 原号码短信验证码
- **响应体**: 空对象 `{}`。

#### 更新提现信息
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/updateCash`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `name` (string, required): 支付宝姓名
  - `payeeCode` (string, required): 支付宝收款码
  - `account` (string, required): 支付宝账号
  - `password` (string, required): 密码
- **响应体**: 空对象 `{}`。

#### 增加余额
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/addBalance`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `id` (int64, required): 用户ID
  - `operateMode` (int64, in:1,2): 操作方式
  - `num` (float64): 操作数量
  - `remark` (string): 备注
  - (其他可选字段)
- **响应体**: 空对象 `{}`。

#### 增加积分
- **HTTP方法**: `POST`
- **URL路径**: `/admin/member/addIntegral`
- **请求头**: `Authorization: Bearer <token>`
- **请求体**: 
  - `id` (int64, required): 用户ID
  - `operateMode` (int64): 操作方式
  - `num` (float64): 操作数量
  - `remark` (string): 备注
  - (其他可选字段)
- **响应体**: 空对象 `{}`。

**Section sources**
- [member.go](file://server/api/admin/member/member.go#L0-L138)

## 终端用户API接口

本节列出 `server/api/api/member/v1/member.go` 中定义的面向终端用户的接口。

### 用户功能

#### 通过邀请码获取用户ID
- **HTTP方法**: `POST`
- **URL路径**: `/api/member/getIdByCode`
- **请求头**: `Authorization: Bearer <token>` (可选，取决于业务逻辑)
- **请求体**: 
  - `code` (string, required): 邀请码
- **响应体**: 空对象 `{}` (根据代码，响应结构未定义，实际业务逻辑可能返回用户ID)。

**Section sources**
- [member.go](file://server/api/api/member/v1/member.go#L0-L16)

## JWT认证与刷新机制

系统使用JWT（JSON Web Token）进行身份认证，令牌信息存储在Redis缓存中以实现更灵活的控制。

### 认证流程
1.  **获取Token**: 用户登录成功后，服务端调用 `token.Login` 函数生成JWT。
2.  **Token组成**:
    - **Header**: 使用 `HS256` 算法签名。
    - **Payload (Claims)**: 包含用户身份信息（`model.Identity`），如 `Id`, `App`, `Username` 等。
    - **Signature**: 使用配置的 `SecretKey` 进行签名。
3.  **存储Token**: 生成的JWT字符串（`header`）会通过 `GetAuthKey` 生成一个认证Key，再通过 `GetTokenKey` 生成Redis缓存Key，将包含过期时间(`ExpireAt`)、刷新时间(`RefreshAt`)和刷新次数(`RefreshCount`)的 `Token` 结构体存入Redis。同时，使用 `GetBindKey` 生成的Key将用户ID与Token Key进行绑定。

### 令牌刷新机制
系统支持自动刷新令牌有效期，防止用户频繁登录。
- **触发条件**: 每次请求经过 `ParseLoginUser` 解析时，都会检查是否需要刷新。
- **刷新逻辑**:
  1.  检查 `AutoRefresh` 是否开启。
  2.  检查 `MaxRefreshTimes` 是否达到上限（-1表示无限制）。
  3.  检查距离上次刷新时间是否超过 `RefreshInterval`。
  4.  若满足条件，则延长 `ExpireAt` 为当前时间 + `Expires`，并增加 `RefreshCount`，然后将更新后的 `Token` 结构体重新存入Redis。
- **多端登录控制**: 通过 `MultiLogin` 配置项控制。若关闭，则通过 `bindKey` 检查当前用户ID是否已绑定其他Token，若已绑定且Token Key不一致，则判定为异地登录，返回错误。

### 请求头处理
- 客户端应在 `Authorization` 请求头中携带 `Bearer <token>`。
- 服务端通过 `GetAuthorization` 函数优先从请求头获取，若不存在则尝试从GET参数 `authorization` 获取。

**Section sources**
- [token.go](file://server/internal/library/token/token.go#L0-L56)
- [token.go](file://server/internal/library/token/token.go#L51-L117)
- [token.go](file://server/internal/library/token/token.go#L119-L192)
- [token.go](file://server/internal/library/token/token.go#L187-L247)
- [token.go](file://server/internal/library/token/token.go#L244-L304)

## 前端调用示例

以下示例使用 `axios` 库调用管理端的“获取用户列表”接口。

```javascript
import axios from 'axios';

// 创建axios实例
const adminApi = axios.create({
  baseURL: 'http://your-api-domain.com', // 替换为实际API地址
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
});

// 请求拦截器：添加Authorization头
adminApi.interceptors.request.use(
  config => {
    const token = localStorage.getItem('admin_token'); // 从本地存储获取token
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器：处理错误码
adminApi.interceptors.response.use(
  response => {
    // 假设后端返回 { code: 0, data: {}, msg: "" }
    if (response.data.code === 0) {
      return response.data.data;
    } else {
      // 根据错误码进行不同处理
      console.error('API Error:', response.data.msg);
      return Promise.reject(new Error(response.data.msg));
    }
  },
  error => {
    if (error.response) {
      // 服务器返回了错误状态码
      if (error.response.status === 401) {
        // 401 Unauthorized，可能是token失效或未登录
        alert('登录已失效，请重新登录');
        localStorage.removeItem('admin_token');
        window.location.href = '/login'; // 跳转到登录页
      } else {
        console.error('Request failed:', error.response.data);
      }
    } else if (error.request) {
      // 请求已发出但没有收到响应
      console.error('Network Error:', error.request);
    } else {
      console.error('Error:', error.message);
    }
    return Promise.reject(error);
  }
);

// 调用获取用户列表接口
async function fetchMemberList(page = 1, limit = 10) {
  try {
    const params = { page, limit };
    const data = await adminApi.get('/admin/member/list', { params });
    console.log('用户列表:', data);
    return data;
  } catch (error) {
    console.error('获取用户列表失败:', error);
  }
}

// 使用示例
// fetchMemberList(1, 10);
```

## 错误码体系

错误码定义在 `internal/consts/error.go` 文件中。虽然具体代码未提供，但根据项目结构和常见实践，错误码体系通常包含以下类别：

- **通用错误**:
  - `0`: 成功
  - `1`: 未知错误
  - `400`: 请求参数错误
  - `401`: 未授权（登录失效）
  - `403`: 禁止访问
  - `404`: 资源未找到
  - `500`: 服务器内部错误
- **业务错误**:
  - `10001`: 用户名或密码错误
  - `10002`: 账号已被禁用
  - `10003`: 验证码错误或已过期
  - `10004`: 余额不足
  - `10005`: 权限不足
  - `10006`: 数据已存在
  - `10007`: 操作过于频繁
- **JWT相关错误**:
  - `11001`: 登录身份已失效，请重新登录！
  - `11002`: 账号已在其他地方登录，如非本人操作请及时修改登录密码！

**Section sources**
- [error.go](file://server/internal/consts/error.go)