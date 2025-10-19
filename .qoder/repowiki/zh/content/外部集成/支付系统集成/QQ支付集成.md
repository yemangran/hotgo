# QQ支付集成

<cite>
**本文档引用的文件**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go)
- [model.go](file://server/internal/library/payment/qqpay/model.go)
- [pay_v1_notify_qq_pay.go](file://server/internal/controller/api/pay/pay_v1_notify_qq_pay.go)
- [payment.go](file://server/internal/library/payment/payment.go)
- [general.go](file://server/internal/model/input/payin/general.go)
- [config.go](file://server/internal/model/config.go)
- [pay.go](file://server/internal/consts/pay.go)
</cite>

## 目录
1. [简介](#简介)
2. [项目结构](#项目结构)
3. [核心组件](#核心组件)
4. [架构概述](#架构概述)
5. [详细组件分析](#详细组件分析)
6. [依赖分析](#依赖分析)
7. [性能考虑](#性能考虑)
8. [故障排除指南](#故障排除指南)
9. [结论](#结论)

## 简介
本文档旨在为开发者提供QQ支付集成的全面指导。由于QQ支付的官方文档相对较少，本文档将作为主要参考，详细阐述`internal/library/payment/qqpay/`包的实现，包括支付请求的发起、响应处理以及异步通知的处理流程。文档将涵盖数据模型定义、控制器逻辑、接入流程和参数配置说明，并提供代码示例与调试技巧，帮助开发者快速、准确地完成集成。

## 项目结构
QQ支付相关的代码主要位于`server/internal/library/payment/qqpay/`目录下，该目录是支付功能模块的一部分，与支付宝、微信支付等其他支付方式并列。支付控制器位于`server/internal/controller/api/pay/`目录，通过统一的支付服务层进行调用。

```mermaid
graph TD
subgraph "支付库"
QQPay[qqpay]
AliPay[alipay]
WxPay[wxpay]
Payment[payment.go]
end
subgraph "控制器"
PayController[pay_v1_notify_qq_pay.go]
end
subgraph "模型与常量"
PayIn[general.go]
Config[config.go]
Consts[pay.go]
end
Payment --> QQPay
Payment --> AliPay
Payment --> WxPay
PayController --> Payment
PayController --> PayIn
QQPay --> Model[model.go]
QQPay --> Handle[handle.go]
Payment --> Config
Payment --> Consts
```

**Diagram sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go)
- [model.go](file://server/internal/library/payment/qqpay/model.go)
- [pay_v1_notify_qq_pay.go](file://server/internal/controller/api/pay/pay_v1_notify_qq_pay.go)
- [payment.go](file://server/internal/library/payment/payment.go)

**Section sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go)
- [model.go](file://server/internal/library/payment/qqpay/model.go)
- [pay_v1_notify_qq_pay.go](file://server/internal/controller/api/pay/pay_v1_notify_qq_pay.go)

## 核心组件
本节分析QQ支付集成的核心组件，包括`qqpay`包的实现、数据模型定义以及异步通知的处理逻辑。

**Section sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L21-L137)
- [model.go](file://server/internal/library/payment/qqpay/model.go#L9-L27)
- [pay_v1_notify_qq_pay.go](file://server/internal/controller/api/pay/pay_v1_notify_qq_pay.go#L1-L27)

## 架构概述
QQ支付集成遵循了统一的支付接口设计模式。`payment`包定义了`PayClient`接口，`qqpay`包实现了该接口。当需要处理支付或通知时，系统通过`payment.New()`工厂方法创建具体的支付客户端实例，然后调用其方法。异步通知由控制器接收，经过统一的服务层处理，最终调用`qqpay`包的`Notify`方法进行验签和数据解析。

```mermaid
sequenceDiagram
participant Client as "客户端"
participant Controller as "PayController"
participant Service as "PayService"
participant Payment as "PaymentFactory"
participant QQPay as "QQPayClient"
Client->>Controller : POST /api/pay/notify/qqpay
Controller->>Service : service.Pay().Notify()
Service->>Payment : payment.New(PayTypeQQPay)
Payment->>QQPay : 返回 *qqPay 实例
Service->>QQPay : qqPay.Notify()
QQPay-->>Service : 返回 NotifyModel
Service->>Service : 更新支付日志
Service->>Service : 调用业务回调
Service-->>Controller : 处理成功
Controller-->>Client : 返回XML成功响应
```

**Diagram sources**
- [pay_v1_notify_qq_pay.go](file://server/internal/controller/api/pay/pay_v1_notify_qq_pay.go#L1-L27)
- [payment.go](file://server/internal/library/payment/payment.go#L23-L30)
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L38-L83)

## 详细组件分析

### QQ支付客户端分析
`qqpay`包的核心是一个实现了`PayClient`接口的结构体`qqPay`，它封装了与QQ支付API的交互逻辑。

#### 对象结构
```mermaid
classDiagram
class qqPay {
+config *model.PayConfig
+New(config *model.PayConfig) *qqPay
+CreateOrder(ctx context.Context, in payin.CreateOrderInp) (res *payin.CreateOrderModel, err error)
+Notify(ctx context.Context, in payin.NotifyInp) (res *payin.NotifyModel, err error)
+Refund(ctx context.Context, in payin.RefundInp) (res *payin.RefundModel, err error)
}
class model.PayConfig {
+Debug bool
+QQPayAppId string
+QQPayMchId string
+QQPayApiKey string
}
class payin.CreateOrderInp {
+Pay *entity.PayLog
}
class payin.CreateOrderModel {
+TradeType string
+PayURL string
+OutTradeNo string
+JsApi *JSAPI
}
class payin.NotifyInp {
}
class payin.NotifyModel {
+OutTradeNo string
+TransactionId string
+PayAt *gtime.Time
+ActualAmount float64
}
qqPay --> model.PayConfig : "依赖"
qqPay --> payin.CreateOrderInp : "输入"
qqPay --> payin.CreateOrderModel : "输出"
qqPay --> payin.NotifyInp : "输入"
qqPay --> payin.NotifyModel : "输出"
```

**Diagram sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L27-L29)
- [config.go](file://server/internal/model/config.go#L158-L161)
- [general.go](file://server/internal/model/input/payin/general.go#L14-L32)

#### 支付请求发起与响应处理
`CreateOrder`方法负责发起支付请求。它首先根据配置创建一个QQ支付客户端，然后根据交易类型（目前支持`qqweb`和`qqwap`）构建请求参数，调用QQ支付的统一下单API，并处理返回结果。

```mermaid
flowchart TD
Start([开始]) --> CheckTradeType["检查交易类型"]
CheckTradeType --> |QQWeb 或 QQWap| BuildParams["构建请求参数<br/>mch_id, body, out_trade_no<br/>notify_url, nonce_str, spbill_create_ip<br/>trade_type=NATIVE, total_fee"]
BuildParams --> CallAPI["调用 client.UnifiedOrder()"]
CallAPI --> CheckReturnCode{"ReturnCode == SUCCESS?"}
CheckReturnCode --> |否| ReturnError["返回 ReturnMsg 错误"]
CheckReturnCode --> |是| CheckResultCode{"ResultCode == SUCCESS?"}
CheckResultCode --> |否| ReturnError
CheckResultCode --> |是| BuildResponse["构建 CreateOrderModel<br/>TradeType, PayURL, OutTradeNo"]
BuildResponse --> End([返回成功])
```

**Diagram sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L86-L126)

**Section sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L86-L126)

#### 数据模型定义
`model.go`文件定义了`NotifyRequest`结构体，用于映射QQ支付异步通知的XML数据。该结构体使用`xml`和`json`标签，确保了数据的正确解析。

```mermaid
erDiagram
NotifyRequest {
string appid
string mch_id
string nonce_str
string sign
string device_info
string trade_type
string trade_state
string bank_type
string fee_type
string total_fee
string cash_fee
string coupon_fee
string transaction_id
string out_trade_no
string attach
string time_end
string openid
}
```

**Diagram sources**
- [model.go](file://server/internal/library/payment/qqpay/model.go#L9-L27)

**Section sources**
- [model.go](file://server/internal/library/payment/qqpay/model.go#L9-L27)

### 异步通知处理分析
`pay_v1_notify_qq_pay.go`控制器负责处理来自QQ支付平台的异步通知。

#### 处理流程
```mermaid
sequenceDiagram
participant QQPay as "QQ支付平台"
participant Controller as "pay_v1_notify_qq_pay.go"
participant Service as "PayService"
participant QQPayClient as "qqpay.Notify"
QQPay->>Controller : POST 请求携带XML数据
Controller->>Service : service.Pay().Notify(PayTypeQQPay)
Service->>QQPayClient : payment.New(QQPay).Notify()
QQPayClient->>QQPayClient : ParseNotifyToBodyMap()
QQPayClient->>QQPayClient : VerifySign() 验签
alt 验签失败
QQPayClient-->>Service : 返回错误
Service-->>Controller : 返回错误
Controller-->>QQPay : 返回XML失败
else 验签成功
QQPayClient->>QQPayClient : Scan 到 NotifyRequest
QQPayClient->>QQPayClient : 检查 TradeState 是否为 SUCCESS
alt 非成功状态
QQPayClient-->>Service : 返回错误或忽略
else 成功状态
QQPayClient-->>Service : 返回 NotifyModel
Service->>Service : 查询支付日志
Service->>Service : 更新支付状态
Service->>Service : 调用 NotifyCall 回调业务
Service-->>Controller : 处理成功
Controller-->>QQPay : 返回XML成功响应
end
end
```

**Diagram sources**
- [pay_v1_notify_qq_pay.go](file://server/internal/controller/api/pay/pay_v1_notify_qq_pay.go#L1-L27)
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L38-L83)
- [payment.go](file://server/internal/library/payment/payment.go#L27-L27)

**Section sources**
- [pay_v1_notify_qq_pay.go](file://server/internal/controller/api/pay/pay_v1_notify_qq_pay.go#L1-L27)

## 依赖分析
QQ支付集成依赖于多个外部库和内部模块。

```mermaid
graph TD
QQPay[qqpay] --> GoPay[github.com/go-pay/gopay]
GoPay --> QQ[github.com/go-pay/gopay/qq]
QQPay --> Gf[github.com/gogf/gf/v2]
Gf --> Errors[errors]
Gf --> Http[ghttp]
Gf --> Time[gtime]
Gf --> Conv[gconv]
Gf --> Grand[grand]
QQPay --> HotGo[hotgo/internal]
HotGo --> Model[model]
HotGo --> PayIn[payin]
HotGo --> Consts[consts]
```

**Diagram sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L1-L25)
- [config.go](file://server/internal/model/config.go#L158-L161)
- [pay.go](file://server/internal/consts/pay.go#L23-L23)

**Section sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L1-L25)

## 性能考虑
- **验签性能**：`VerifySign`操作是通知处理中的关键步骤，应确保`QQPayApiKey`的获取是高效的。
- **数据库查询**：在`Notify`回调中，对`PayLog`表的查询和更新操作应确保有适当的索引（如`out_trade_no`）以保证性能。
- **并发处理**：支付通知可能并发到达，`payment.New()`工厂方法和`qqPay`实例的设计是无状态的，可以安全地处理并发请求。

## 故障排除指南
- **验签不通过**：检查`QQPayApiKey`配置是否正确，确保与QQ支付商户平台上的API密钥完全一致。
- **订单未找到**：检查`out_trade_no`是否正确传递，并确认该订单号在`pay_log`表中存在且状态为“待支付”。
- **交易状态非成功**：`Notify`方法默认只处理`SUCCESS`状态的交易。如果需要处理其他状态（如退款、关闭），需要修改`handle.go`中的逻辑。
- **创建订单失败**：检查`QQPayMchId`、`notify_url`、`total_fee`等参数是否正确，确保网络可以访问QQ支付API。

**Section sources**
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L38-L83)
- [handle.go](file://server/internal/library/payment/qqpay/handle.go#L86-L126)

## 结论
本文档详细阐述了HotGo项目中QQ支付的集成方案。通过`qqpay`包的`qqPay`结构体实现了`PayClient`接口，完成了支付请求的发起和异步通知的处理。整个流程设计清晰，通过统一的工厂模式和接口定义，保证了代码的可扩展性和可维护性。开发者在集成时，应重点关注配置参数的正确性、验签逻辑以及业务回调的实现。