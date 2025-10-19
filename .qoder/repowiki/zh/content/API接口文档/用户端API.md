# 用户端API

<cite>
**本文档中引用的文件**  
- [member.go](file://server/api/api/member/v1/member.go)
- [api.go](file://server/internal/router/api.go)
- [token.go](file://server/internal/library/token/token.go)
- [api_auth.go](file://server/internal/logic/middleware/api_auth.go)
- [response.go](file://server/internal/model/response.go)
- [input.go](file://server/internal/model/input/memberin.go)
</cite>

## 目录
1. [简介](#简介)
2. [用户端API概览](#用户端api概览)
3. [认证与会话管理](#认证与会话管理)
4. [公共接口列表](#公共接口列表)
5. [前端调用示例](#前端调用示例)
6. [通用错误码体系](#通用错误码体系)
7. [与管理端API的权限差异](#与管理端api的权限差异)

## 简介
本文档旨在为前端开发者提供完整的用户端API使用指南。涵盖所有面向终端用户的接口定义、认证机制、请求响应格式及调用示例，确保前后端协作高效顺畅。

## 用户端API概览
用户端API主要位于 `server/api/api/` 路径下，遵循RESTful设计规范，使用JSON作为数据交换格式。所有接口均以 `/api` 为前缀，并通过版本号（如 `v1`）进行管理。

核心模块包括：
- 会员系统（登录、注册、信息管理）
- 支付功能
- 内容获取
- 通用工具接口

**Section sources**
- [api.go](file://server/internal/router/api.go#L1-L50)
- [member.go](file://server/api/api/member/v1/member.go#L1-L20)

## 认证与会话管理
系统采用JWT（JSON Web Token）实现无状态认证，保障接口安全。

### JWT实现方式
1. 用户登录成功后，服务端生成包含用户ID和过期时间的JWT令牌。
2. 令牌通过响应体返回前端，前端需将其存储于本地（如localStorage）。
3. 后续请求需在HTTP头中携带 `Authorization: Bearer <token>`。
4. 服务端通过中间件 `api_auth.go` 验证令牌有效性。

令牌有效期由配置文件控制，默认为2小时，支持刷新机制。

### 会话管理
- **无状态会话**：服务端不保存会话信息，依赖JWT自包含特性。
- **自动刷新**：临近过期时，前端可调用刷新接口获取新令牌。
- **强制登出**：通过黑名单机制实现令牌失效（见 `sys_blacklist` 表）。

**Section sources**
- [token.go](file://server/internal/library/token/token.go#L10-L80)
- [api_auth.go](file://server/internal/logic/middleware/api_auth.go#L15-L60)

## 公共接口列表
以下为面向终端用户的核心接口：

### 会员登录
- **方法**：POST
- **路径**：`/api/member/login`
- **请求头**：
  - `Content-Type: application/json`
- **请求体**：
  ```json
  {
    "account": "用户名或邮箱",
    "password": "密码"
  }
  ```
- **响应体**：
  ```json
  {
    "code": 0,
    "message": "success",
    "data": {
      "token": "JWT令牌",
      "expires_in": 7200
    }
  }
  ```

### 会员注册
- **方法**：POST
- **路径**：`/api/member/register`
- **请求体**：
  ```json
  {
    "account": "用户名",
    "email": "邮箱",
    "password": "密码",
    "confirm_password": "确认密码"
  }
  ```

### 获取用户信息
- **方法**：GET
- **路径**：`/api/member/profile`
- **认证要求**：是
- **响应体**：
  ```json
  {
    "code": 0,
    "data": {
      "id": 1,
      "account": "user123",
      "email": "user@example.com",
      "created_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

### 修改密码
- **方法**：PUT
- **路径**：`/api/member/change-password`
- **认证要求**：是
- **请求体**：
  ```json
  {
    "old_password": "旧密码",
    "new_password": "新密码",
    "confirm_password": "确认新密码"
  }
  ```

**Section sources**
- [member.go](file://server/api/api/member/v1/member.go#L25-L150)
- [input.go](file://server/internal/model/input/memberin.go#L5-L40)

## 前端调用示例
使用axios调用登录接口的示例代码：

```javascript
import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json'
  }
});

// 登录请求
async function login(account, password) {
  try {
    const response = await apiClient.post('/member/login', {
      account,
      password
    });
    
    if (response.data.code === 0) {
      const { token, expires_in } = response.data.data;
      localStorage.setItem('authToken', token);
      // 设置过期时间
      setTimeout(() => {
        console.log('令牌已过期，请重新登录');
      }, expires_in * 1000);
    }
    return response.data;
  } catch (error) {
    console.error('登录失败:', error.response?.data || error.message);
    throw error;
  }
}

// 带认证的请求
apiClient.interceptors.request.use(config => {
  const token = localStorage.getItem('authToken');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});
```

**Section sources**
- [member.go](file://server/api/api/member/v1/member.go#L30-L60)

## 通用错误码体系
所有接口遵循统一的响应结构：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

| 错误码 | 说明 | 处理建议 |
|--------|------|----------|
| 0 | 成功 | 正常处理响应数据 |
| 10001 | 参数错误 | 检查请求体字段格式 |
| 10002 | 用户不存在 | 提示用户注册或检查账号 |
| 10003 | 密码错误 | 提示用户重试或找回密码 |
| 10004 | 认证失败 | 跳转至登录页重新登录 |
| 10005 | 令牌过期 | 使用刷新机制或重新登录 |
| 10006 | 权限不足 | 检查当前用户角色 |
| 10007 | 接口不存在 | 检查URL路径和版本号 |
| 10008 | 请求过于频繁 | 降低请求频率 |
| 50000 | 服务器内部错误 | 联系管理员或稍后重试 |

**Section sources**
- [response.go](file://server/internal/model/response.go#L5-L30)

## 与管理端API的权限差异
用户端API与管理端API在权限控制上存在显著差异：

| 特性 | 用户端API | 管理端API |
|------|-----------|-----------|
| 认证方式 | JWT令牌 | JWT令牌 + 角色权限 |
| 访问范围 | 仅限自身数据 | 可访问全系统数据 |
| 权限模型 | 基于用户身份 | 基于RBAC（角色访问控制） |
| 接口粒度 | 业务功能导向 | 管理操作导向 |
| 安全级别 | 标准防护 | 增强防护（IP限制、操作审计） |

管理端接口位于 `/admin` 路径下，需通过Casbin进行细粒度权限校验，而用户端接口主要验证身份合法性。

**Section sources**
- [api_auth.go](file://server/internal/logic/middleware/api_auth.go#L20-L50)
- [admin_auth.go](file://server/internal/logic/middleware/admin_auth.go#L15-L45)