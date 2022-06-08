### 项目介绍

##### 1.1 项目背景

~~~shell
针对于告警分组，把不同业务的告警信息发送至指定钉钉群
~~~

##### 1.2 项目配置

* etc/app.toml

~~~go
# 该项目配置文件
# 定义钉钉token、应用端口等信息
# ## 样例
[app]
host = "0.0.0.0"
port = "18080"

[log]
level = "debug"
dir = "logs"
format = "text"
to = "stdout"

# 钉钉token
# 多钉钉告警信息发送，是根据Prometheus配置的tag，alertmanager根据不同tag发送至不同的webhook地址来实现
# 具体配置参看。。。
[dingding]
	# erp项目组钉钉token
    # 请求示例地址: http://127.0.0.1:18080/dingding/erp
    [dingding.erp]
        token = "******"
	# pos项目组钉钉token
    # 请求示例地址: http://127.0.0.1:18080/dingding/pos
    [dingding.pos]
        token = "******"


~~~

* alertmanager配置

~~~shell
# 针对以上app.toml配置文件dingding参数配置，对应的alertmanager配置如下
global:
  resolve_timeout: 1m

route:
  receiver: 'default-receiver'
  group_by: ['alertname']
  group_wait: 10s
  group_interval: 2m
  repeat_interval: 1h

  routes: # 定义路由
  - receiver: "web.hook.erp"
    match:
     # 匹配含有app: erp标签的告警项,把搞告警信息发送至下面对应的接收者
     # 该标签 Prometheus自定义，或者自行根据已有标签来区分不同业务木模块可
      app: erp 

receivers:
- name: 'default-receiver'
  webhook_configs:
  - url: 'http://10.10.1.70:18080/dingding/pos'
    send_resolved: true
  - name: 'web.hook.erp'  # 上面路由匹配到后 发送到对应的webhoook地址
   webhook_configs:
   # 该webhhok地址需要与dingding告警项目中配置文件的dingding.erp对应
   # http://10.10.1.70:18080/dingding 固定接口，erp可变，不同项目名称，需要与钉钉告警配置文件中对应
   # [dingding.erp]
   #    token = "*****"
   - url: 'http://10.10.1.70:18080/dingding/erp' 
     send_resolved: true
inhibit_rules:
  - source_match:
      severity: 'critical'
    target_match:
      severity: 'warning'
    equal: ['alertname', 'dev', 'instance']
~~~

##### 1.3 部署方式

###### 1.3.1 构建运行

~~~go
# 拉取代码
https://github.com/tchuaxiaohua/prometheus-dingding.git

# 启动(依赖go环境)
方式一、直接启动
cd prometheus-dingding
go run main.go
方式二、编译启动
cd prometheus-dingding
go build -o dinghook
./dinghook

# 钉钉token配置文件参考etc/app.toml说明
~~~

###### 1.3.2 docker

~~~shell
# 拉去代码
https://github.com/tchuaxiaohua/prometheus-dingding.git
# 构建镜像
cd prometheus-dingding
cp doc/Dockerfile .
# ## 构建执行脚本
go mod tidy
go build  -ldflags "-s -w" -o dingtalk  main.go
docker build -t huahua5404/prometheus-dingding:v1 .

# 启动
docker run -it -d -p 18080:18080 --name dinghook huahua5404/prometheus-dingding:v1

~~~

###### 1.3.3 k8s

~~~shell
# 创建configmap 指定配置文件
kubectl create configmap dingtalk-conf --from-file=etc/app.toml -n monitoring

# 创建pod dingtaLk-k8s.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: dingtalk
  namespace: monitoring
spec:
  replicas: 1
  selector:
    matchLabels:
      app: dingtalk
  template:
    metadata:
      labels:
        app: dingtalk
    spec:
      containers:
      - name: dingtalk-hook
        image: huahua5404/prometheus-dingding:v1 
        imagePullPolicy: Always 
        ports:
        - containerPort: 18080
          name: http
		volumeMounts:
          - name: dingtalk-conf
            mountPath: /apps/etc/
            readOnly: true    
        resources:
          requests:
            cpu: 50m
            memory: 100Mi
          limits:
            cpu: 50m
            memory: 100Mi
      volumes:
        - name: dingtalk-conf
          configMap:
            name: dingtalk-conf
---
apiVersion: v1
kind: Service
metadata:
  name: dingtalk-svc
  namespace: monitoring
spec:
  selector:
    app: dingtalk
  ports:
  - name: hook
    port: 18080
    targetPort: http
~~~
##### 1.4 关于模板
##### 1.4.1 模板介绍
~~~shell
1) 模板文件: template/alert.tmpl
2) 模板取值优先级说明(根据标签以此查找):
    故障主机: pod > container > hostname > instance
    告警详情: Description > Summary
~~~
##### 1.4.2 模板自定义
~~~shell
告警模板内容都是基于标签进行赋值，如果想定义更丰富的模板则需要自己调整prometheus监控指标的标签值。
.Labels 获取的是指标的标签,如果自定义，直接`.Labels.tag`即可，按照MD文档格式，即可完成自定义告警模板。
~~~

