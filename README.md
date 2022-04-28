### 分布式定时任务调度平台

#### 项目介绍

- 使用Master-Worker架构开发可视化分布式定时任务平台，支持对服务进行分布式部署。
- 服务支持秒级调度，最长调度时间级别为年级别。
- 支持随时修改任务的调度时间以及任务的调度内容，支持在任务调度的时候对任务进行kill。
- 使用etcd实现服务注册以及服务发现，支持随时查看集群中的机器数量和ip地址。
- 使用mongodb存放日志文件，支持查看任务每一次调度的实际时间和计划时间，支持查看任务运行后的输出，在任务出现bug时自动保存报错时的输出。
- 使用mysql存放用户数据
- 使用redis验证登录时的图像验证码以及验证开通短信服务后的用户输入的验证码



#### 使用技术栈

- gin
- mysql
- redis
- etcd
- mongodb
- cronexpr



#### 架构图

