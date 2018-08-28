lsfmonitor 是一个小规模的golang项目

项目主要是使用mongodb作为存储数据库,grpc提供服务mongodb数据创建，更新，查询

1. 利用protobuf 生产golang接口代码
2. grpc使用protobuf接口代码提供grpc服务,monitor_job_grpc.go 
               提供create/update 记录
3. job_search_grpc.go  groc提供对mongodb search 服务
4. webmain.go提供基于gin框架web服务,用户名采用ldap用户认证,经过jwt生产token

