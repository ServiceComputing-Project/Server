# Jaeger为何物？

Jaeger 是Uber推出的一款开源分布式追踪系统，兼容OpenTracing API。分布式追踪系统用于记录请求范围内的信息。例如，一次远程方法调用的执行过程和耗时。是我们排查系统问题和系统性能的利器。
 分布式追踪系统种类繁多，但是核心步骤有三个：代码埋点，数据存储和查询展示。
 以上几句描述都是我copy的，所以大家想要对Jaeger有更加深入的了解，可以参阅这篇文章[Jaeger 分布式追踪系统模块分析](https://blog.csdn.net/johnhill_/article/details/81111219),能让你对Jaeger有一个简单的认识。
 当然我们还要记得APM的三大模块分别是集中式日志系统，集中式度量系统和分布式全链接追踪系统。
 而`Jaeger`属于的就是追踪系统，度量系统我们则会使用`prometheus`,日志系统一般则是`elk`。

## 选用Jaeger的原因

一个是它兼容OpenTracing API,写起来简单方便，一个是UI相较于Zipkin的更加直观和丰富，还有一个则是sdk比较丰富，go语言编写，上传采用的是udp传输，效率高速度快。
 相比**[Pinpoint](https://github.com/naver/pinpoint)**的缺点，当然是UI差距了，基本上现在流行的追踪系统UI上都远远逊于它。

# 搭建

#### 测试搭建

在个人使用或者测试上，Jaeger的搭建其实较为简单，因为我们使用的存储方式是内存化的，所以我们可以直接使用官方给我们打包好的镜像。



```undefined
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp \
  -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
```

#### 线上搭建

线上搭建我采用的是jaeger搭配`elasticsearch`，效果还不错，而且搭配起来比较简单，官方给的方式是jaeger搭`cassandra`,不过我试了下没有搭起来。。。。。这是官方的地址[jaeger-docker-compose.yml](https://github.com/jaegertracing/jaeger/blob/master/docker-compose/jaeger-docker-compose.yml),下面是我采用的搭建yml



```csharp
version: '2.1'
services:

  elasticsearch:
    image: elasticsearch:5.6.4
    environment:
      - "ES_JAVA_OPTS=-Xms1024m -Xmx1024m"

  collector:
    image: jaegertracing/jaeger-collector
    environment:
      - SPAN_STORAGE_TYPE=elasticsearch
      - ES_SERVER_URLS=http://elasticsearch:9200
      - ES_USERNAME=elastic
      - LOG_LEVEL=debug
    depends_on:
      - elasticsearch

  agent:
    image: jaegertracing/jaeger-agent
    environment:
      - COLLECTOR_HOST_PORT=collector:14267
      - LOG_LEVEL=debug
    ports:
      - "5775:5775/udp"
      - "5778:5778"
      - "6831:6831/udp"
      - "6832:6832/udp"
    depends_on:
      - collector
  query:
    image: jaegertracing/jaeger-query
    environment:
      - SPAN_STORAGE_TYPE=elasticsearch
      - ES_SERVER_URLS=http://elasticsearch:9200
      - ES_USERNAME=elastic
      - LOG_LEVEL=debug
    ports:
      - 16686:16686
    depends_on:
      - elasticsearch

  hotrod:
    image: jaegertracing/example-hotrod:1.6
    command: all --jaeger-agent.host-port=agent:6831
    ports:
      - 8080:8080
    depends_on:
      - agent
```

这里面除了hotrod以外都是必须的，`hotrod`是用来测试我们搭建后是否能成功上传到我们的agent的。
 这里要注意的一点是，虽然有depends_on,但是由于`elasticsearch`启动的原因，导致`query`, `collector`连不上直接挂掉，`agent`虽然没挂但是也连不上，所以我们需要手动重启`query`,`collector`,`agent`,这个时候再查看应该是都正常启动了。

![img](https:////upload-images.jianshu.io/upload_images/12890383-b806e2af3970afd9.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

jaeger



#### 测试 (地址应该是**http://127.0.0.1:8080**)

进入到`hotrod`,随便点击一个按钮，生成调用传到`agent`

![img](https:////upload-images.jianshu.io/upload_images/12890383-e699b7c683e9cbff.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

hotrod



#### 查看我们的`query`

如果是本地就是这个地址**http://127.0.0.1:16686**

![img](https:////upload-images.jianshu.io/upload_images/12890383-66c9166f1edbcb9f.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

query



![img](https:////upload-images.jianshu.io/upload_images/12890383-575c2859d46b6ef6.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

query2


 可以看到，Jaeger的UI还是非常直观，友好，漂亮的。



## sdk 接入

目前官方提供了Java，Python，Go，C++，Node.js。[sdk](https://github.com/jaegertracing/jaeger#instrumentation-libraries),
 php暂时只有第三方的，我现在使用的是[php](https://github.com/jukylin/jaeger-php)这个sdk，还不错，可以扩展开发，我基于这个sdk写了一个`swoft`框架的jaeger组件，默认支持`mysql`,`redis`,`httpClient`监控，并且可灵活添加header，如果有需要，可以使用，地址：[swoft-jaeger](https://github.com/masixun71/swoft-jaeger)

## 最后

无论是什么样的监控系统，对于线上服务都或多或少都有性能损耗，所以在线上一定要采样处理才是最佳使用方式，当然，我的组件支持采样咯。



文章链接：https://www.jianshu.com/p/ffc597bb4ce8

