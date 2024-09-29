# 基于GIN的云存储网盘
# GIN based cloud storage network disk
参考链接(Reference)<br>
<https://github.com/xiaogao67/gin-cloud-storage><br>
前端模板引擎：Thymeleaf -> GO html/template（自带）<br>
<font color=Red>**声明：本项目仅用作学习用途！<br>
Disclaim: This project is only for learning purposes!**</font><br>
注：已测试阿里云OSS存储，参数已隐藏
# 知识汇总
- [GIN](#gin)
- [数据库类型](#数据库类型)
- [数据库设计三范式](#数据库设计三范式)
- [跨域问题](#跨域问题)
- [CSRF](#csrf)
- [HTTP状态码](#http状态码)
- [会话控制](#会话控制cookie-session-token的区别)
## GIN
Gin是一个轻量级的Go语言Web框架，它具有高性能和简洁的设计。由于其快速的路由匹配和处理性能，Gin成为Go语言中最受欢迎的Web框架之一。以下是一些Gin框架的主要特点：<br>
快速和轻量级、路由和中间件、JSON解析、支持插件<br>
Github地址: <https://github.com/gin-gonic/gin><br>
*gin.Context: Context 代表的是父子协程的一个上下文对象，主要作用是共享数据、以及对子协程做一些超时控制等。gin 里面的 Context 除了 context.Context 的基本功能外，还提供了很多 HTTP 请求处理相关的一些功能，比如获取请求数据、处理响应等<br>

## 数据库类型
关系型数据库（RDBMS）:使用表格形式组织数据。数据以行和列的形式存储在表中。
支持复杂的查询和数据关系，通过SQL（结构化查询语言）进行操作。
强调ACID（原子性、一致性、隔离性、持久性）特性，以保证数据的一致性和可靠性。
通常是垂直扩展（即通过增加单台服务器的硬件能力来提升性能）。
适合数据量较小但关系复杂的应用场景。
架构通常是固定的，表结构需要事先定义，并且变更表结构可能需要复杂的迁移操作。
适合结构化数据和复杂查询的需求。
使用SQL进行数据操作和查询，支持复杂的JOIN操作。
强调对事务的支持，以保证数据操作的完整性和可靠性。<br>
非关系型数据库（NoSQL）:不使用表格形式存储数据，而是使用多种不同的数据模型，如文档、键值对、列族、图等。
可以支持不同的数据结构和查询方式。
通常强调可用性和分布式系统的伸缩性，可能在一致性上做出妥协。例如，许多NoSQL数据库采用最终一致性模型。
通常是水平扩展（即通过增加更多的服务器来分担负载）。
更适合大数据量和分布式应用的场景。
数据模型较为灵活，可以存储结构化、半结构化或非结构化的数据。
适合需要频繁变更数据结构的应用或需要存储大量多样化数据的场景。
使用各自特定的查询语言或API，不一定支持SQL标准的复杂查询操作。
一些NoSQL数据库可能不支持传统的事务模型，或者只支持部分事务操作。<br>
常用的关系型数据库：MySQL、PostgreSQL、Oracle Database、Microsoft SQL Server、SQLite<br>
常用的非关系型数据库：MongoDB、Redis、Cassandra、Couchbase、Neo4j、DynamoDB<br>

## 数据库设计三范式
第一范式（1NF）：原子性（存储的数据应该具有“不可再分性”）<br>
第二范式（2NF）：唯一性 (消除非主键部分依赖联合主键中的部分字段)（一定要在第一范式已经满足的情况下）<br>
第三范式（3NF）：独立性，消除传递依赖(非主键值不依赖于另一个非主键值)<br>

## 跨域问题
跨域（Cross-Origin）指的是浏览器阻止前端网页从一个域名（Origin）向另一个域名的服务器发送请求。具体来说，一个页面的协议、域名、端口三者任意一个与请求的目标地址不同，就被视为跨域请求。<br>
同源策略（Same-Origin Policy）是浏览器的一个重要安全机制，防止恶意网站通过跨域方式窃取敏感数据。该策略限制了从一个源加载的文档或脚本如何与另一个源的资源进行交互。
同源策略的定义：如果两个URL的协议、域名和端口都相同，则这两个URL具有相同的源。<br>
解决方案：JSONP、CORS、代理服务器、nginx反向代理、后端设置允许跨域<br>

## CSRF
CSRF (Cross-site request forgery，跨站请求伪造)也被称为One Click Attack或者Session Riding，通常缩写为CSRF或者XSRF，是一种对网站的恶意利用。尽管听起来像跨站脚本(XSS)，但它与XSS非常不同，XSS利用站点内的信任用户，而CSRF则通过伪装成受信任用户请求受信任的网站。
简单的说，是攻击者通过一些技术手段欺骗用户的浏览器去访问一个自己以前认证过的站点并运行一些操作（如发邮件，发消息，甚至财产操作（如转账和购买商品）。因为浏览器之前认证过，所以被访问的站点会觉得这是真正的用户操作而去运行。<br>
*防护方法：只使用JSON API、验证HTTP Referer字段、在请求地址中添加token验证

## HTTP状态码
**消息<br>**
100 Continue: 客户端应当继续发送请求<br>
101 Switching Protocols: 服务器已经理解了客户端的请求，并将通过Upgrade消息头通知客户端采用不同的协议来完成这个请求<br>
102 Processing: 处理将被继续执行<br>
**成功<br>**
200 OK: 请求已成功，请求所希望的响应头或数据体将随此响应返回<br>
201 Created: 请求已经被实现，而且有一个新的资源已经依据请求的需要而建立，且其 URI 已经随Location 头信息返回<br>
202 Accepted: 服务器已接受请求，但尚未处理<br>
203 Non-Authoritative Information: 服务器已成功处理了请求，但返回的实体头部元信息不是在原始服务器上有效的确定集合，而是来自本地或者第三方的拷贝<br>
204 No Content: 服务器成功处理了请求，但不需要返回任何实体内容，并且希望返回更新了的元信息<br>
205 Reset Content: 服务器成功处理了请求，且没有返回任何内容<br>
206 Partial Content: 服务器已经成功处理了部分 GET 请求<br>
207 Multi-Status: 之后的消息体将是一个XML消息，并且可能依照之前子请求数量的不同，包含一系列独立的响应代码<br>
**重定向<br>**
300 Multiple Choices: 被请求的资源有一系列可供选择的回馈信息，每个都有自己特定的地址和浏览器驱动的商议信息<br>
301 Moved Permanently: 被请求的资源已永久移动到新位置，并且将来任何对此资源的引用都应该使用本响应返回的若干个 URI 之一<br>
302 Move Temporarily: 请求的资源临时从不同的 URI响应请求<br>
303 See Other: 对应当前请求的响应可以在另一个 URL 上被找到，而且客户端应当采用 GET 的方式访问那个资源<br>
304 Not Modified: 如果客户端发送了一个带条件的 GET 请求且该请求已被允许，而文档的内容（自上次访问以来或者根据请求的条件）并没有改变，则服务器应当返回这个状态码<br>
305 Use Proxy: 被请求的资源必须通过指定的代理才能被访问<br>
306 Switch Proxy: 在最新版的规范中，306状态码已经不再被使用<br>
307 Temporary Redirect: 请求的资源临时从不同的URI 响应请求<br>
**请求错误<br>**
400 Bad Request: 语义有误，当前请求无法被服务器理解或请求参数有误<br>
401 Unauthorized: 当前请求需要用户验证<br>
402 Payment Required: 该状态码是为了将来可能的需求而预留的<br>
403 Forbidden: 服务器已经理解请求，但是拒绝执行它<br>
404 Not Found: 请求失败，请求所希望得到的资源未被在服务器上发现<br>
405 Method Not Allowed: 请求行中指定的请求方法不能被用于请求相应的资源<br>
406 Not Acceptable: 请求的资源的内容特性无法满足请求头中的条件，因而无法生成响应实体<br>
407 Proxy Authentication Required: 客户端必须在代理服务器上进行身份验证<br>
408 Request Timeout: 请求超时。客户端没有在服务器预备等待的时间内完成一个请求的发送<br>
409 Conflict: 由于和被请求的资源的当前状态之间存在冲突，请求无法完成<br>
410 Gone: 被请求的资源在服务器上已经不再可用，而且没有任何已知的转发地址<br>
411 Length Required: 服务器拒绝在没有定义 Content-Length头的情况下接受请求<br>
412 Precondition Failed: 服务器在验证在请求的头字段中给出先决条件时，没能满足其中的一个或多个<br>
413 Request Entity Too Large: 服务器拒绝处理当前请求，因为该请求提交的实体数据大小超过了服务器愿意或者能够处理的范围<br>
414 Request-URI Too Long: 请求的URI 长度超过了服务器能够解释的长度，因此服务器拒绝对该请求提供服务<br>
415 Unsupported Media Type: 对于当前请求的方法和所请求的资源，请求中提交的实体并不是服务器中所支持的格式，因此请求被拒绝<br>
416 Requested Range Not Satisfiable: 如果请求中包含了 Range 请求头，并且 Range 中指定的任何数据范围都与当前资源的可用范围不重合，同时请求中又没有定义 If-Range 请求头，那么服务器就应当返回416状态码<br>
417 Expectation Failed: 在请求头 Expect 中指定的预期内容无法被服务器满足，或者这个服务器是一个代理服务器，它有明显的证据证明在当前路由的下一个节点上，Expect 的内容无法被满足<br>
418 I'm a teapot: 愚人节玩笑<br>
421 Misdirected Request: 请求被指向到无法生成响应的服务器（比如由于连接重复使用）<br>
422 Unprocessable Entity: 请求格式正确，但是由于含有语义错误，无法响应<br>
423 Locked: 当前资源被锁定<br>
424 Failed Dependency: 由于之前的某个请求发生的错误，导致当前请求失败，例如 PROPPATCH<br>
425 Too Early: 服务器不愿意冒风险来处理该请求，原因是处理该请求可能会被“重放”，从而造成潜在的重放攻击<br>
426 Upgrade Required: 客户端应当切换到TLS/1.0<br>
449 Retry With: 由微软扩展，代表请求应当在执行完适当的操作后进行重试<br>
451 Unavailable For Legal Reasons: 该请求因法律原因不可用<br>
**服务器错误<br>**
500 Internal Server Error: 服务器遇到了一个未曾预料的状况，导致了它无法完成对请求的处理<br>
501 Not Implemented: 服务器不支持当前请求所需要的某个功能<br>
502 Bad Gateway: 作为网关或者代理工作的服务器尝试执行请求时，从上游服务器接收到无效的响应<br>
503 Service Unavailable: 由于临时的服务器维护或者过载，服务器当前无法处理请求<br>
504 Gateway Timeout: 作为网关或者代理工作的服务器尝试执行请求时，未能及时从上游服务器（URI标识出的服务器，例如HTTP、FTP、LDAP）或者辅助服务器（例如DNS）收到响应<br>
505 HTTP Version Not Supported: 服务器不支持，或者拒绝支持在请求中使用的 HTTP 版本<br>
506 Variant Also Negotiates: 服务器存在内部配置错误：被请求的协商变元资源被配置为在透明内容协商中使用自己，因此在一个协商处理中不是一个合适的重点<br>
507 Insufficient Storage: 服务器无法存储完成请求所必须的内容。这个状况被认为是临时的<br>
509 Bandwidth Limit Exceeded: 服务器达到带宽限制<br>
510 Not Extended: 获取资源所需要的策略并没有被满足<br>
600 Unparseable Response Headers: 源站没有返回响应头部，只返回实体内容<br>

## 会话控制：cookie session token的区别
cookie: HTTP服务器发送到用户浏览器并保存在本地的一小块数据<br>
校验通过下发cookie，后续向服务器发送请求时自动携带cookie<br>
session: 保存在服务器端的一块数据，保存当前访问用户的相关信息<br>
校验通过后创建session信息，然后将session_id的值通过响应头返回给浏览器<br>
token: 服务器端生成并返回给HTTP客户端的一串加密字符串，token中保存着用户信息<br>
校验通过后相应token，token一般在响应体中返回给客户端<br>
*token的特点：<br>
服务器端压力更小：数据存储在客户端<br>
相对更安全：数据加密，可以避免CSRF（跨站请求伪造）<br>
扩展性更强：服务间可以共享，增加服务节点更简单<br>
*JWT(JSON Web Token): 目前最流行的跨域认证解决方案，可用于基于token的身份认证