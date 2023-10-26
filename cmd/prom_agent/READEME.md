# prometheus-manager_agent

> 作为直接操作prometheus server的agent，负责prometheus server的配置文件生成、reload、prometheus server的启动、停止等操作

## 功能

1. prometheus server配置文件生成, 通过获取来自server端推送的配置信息，生成prometheus server的配置文件, 推送方式可以采用http或者grpc
2. prometheus server配置文件reload, 通过prometheus server的reload接口，实现prometheus server的配置文件reload(支持k8s的prometheus server)
3. prometheus server的启动、停止, 通过prometheus server的启动、停止脚本，实现prometheus server的启动、停止
4. alert manager告警hook, 通过alert manager的告警hook接口，实现告警信息的接收和处理

## 关于接收配置信息的接口

### 报警规则

1. loadRules: 加载报警规则(加载该agent负责的prometheus server的报警规则), 并解析出规则ID和规则hash
2. 获取从prometheus server推送的报警规则hash数据, 并与本地的报警规则hash数据进行比较, 如果不一致, 则向server发起reload请求, server端接收到reload请求后, 会重新计算新的报警规则, 并通过推送接口推送给agent
3. 接收从prometheus server推送的报警规则, 并解析出规则ID和规则hash存储在本地, 供下次比较使用, 同时把新的报警规则写入prometheus server的配置文件中, 并reload prometheus server