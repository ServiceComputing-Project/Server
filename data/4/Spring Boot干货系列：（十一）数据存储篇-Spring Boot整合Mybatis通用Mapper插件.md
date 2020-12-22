## 前言
上次介绍了Spring Boot中Mybatis的简单整合，本篇深入来结合通用Mapper、Mybatis Geneator以及分页PageHelper来打造适合企业开发的模板框架。
## 正文
项目框架还是跟上一篇一样使用Spring Boot的ace后端模板，不过最近在使用vue，所以前端引用了vue进来改写，代码变得更加简洁。
## 项目配置：
Spring Boot： 1.5.9.RELEASE
Maven： 3.5
Java： 1.8
Thymeleaf： 3.0.7.RELEASE
Vue.js： v2.5.11
## 数据源依赖
这里我们还是使用阿里巴巴的druid来当数据库连接池，发现这个有对应的监控界面，我们可以开启。

druid官方文档：https://github.com/alibaba/druid/wiki/常见问题

```
<dependency>
    <groupId>mysql</groupId>
    <artifactId>mysql-connector-java</artifactId>
</dependency>

<dependency>
    <groupId>com.alibaba</groupId>
    <artifactId>druid</artifactId>
    <version>1.0.19</version>
</dependency>
```









对应的application.properties配置：
```
## 数据库访问配置
spring.datasource.type=com.alibaba.druid.pool.DruidDataSource
spring.datasource.driver-class-name = com.mysql.jdbc.Driver
spring.datasource.url = jdbc:mysql://localhost:3306/spring?useUnicode=true&characterEncoding=utf-8
spring.datasource.username = root
spring.datasource.password = root


# 下面为连接池的补充设置，应用到上面所有数据源中
# 初始化大小，最小，最大
spring.datasource.initialSize=5
spring.datasource.minIdle=5
spring.datasource.maxActive=20
# 配置获取连接等待超时的时间
spring.datasource.maxWait=60000
# 配置间隔多久才进行一次检测，检测需要关闭的空闲连接，单位是毫秒
spring.datasource.timeBetweenEvictionRunsMillis=60000
# 配置一个连接在池中最小生存的时间，单位是毫秒
spring.datasource.minEvictableIdleTimeMillis=300000
spring.datasource.validationQuery=SELECT 1 FROM DUAL
spring.datasource.testWhileIdle=true
spring.datasource.testOnBorrow=false
spring.datasource.testOnReturn=false
# 打开PSCache，并且指定每个连接上PSCache的大小
spring.datasource.poolPreparedStatements=true
spring.datasource.maxPoolPreparedStatementPerConnectionSize=20
# 配置监控统计拦截的filters，去掉后监控界面sql无法统计，'wall'用于防火墙
spring.datasource.filters=stat,wall,log4j
# 合并多个DruidDataSource的监控数据
#spring.datasource.useGlobalDataSourceStat=true
```


对应的bean配置：
```
package com.dudu.config;

/**
 * Druid配置
 *
 * @author dudu
 * @date 2017-12-11 0:00
 */
@Configuration
public class DruidConfig {
    private Logger logger = LoggerFactory.getLogger(DruidConfig.class);

    @Value("${spring.datasource.url:#{null}}")
    private String dbUrl;
    @Value("${spring.datasource.username: #{null}}")
    private String username;
    @Value("${spring.datasource.password:#{null}}")
    private String password;
    @Value("${spring.datasource.driverClassName:#{null}}")
    private String driverClassName;
    @Value("${spring.datasource.initialSize:#{null}}")
    private Integer initialSize;
    @Value("${spring.datasource.minIdle:#{null}}")
    private Integer minIdle;
    @Value("${spring.datasource.maxActive:#{null}}")
    private Integer maxActive;
    @Value("${spring.datasource.maxWait:#{null}}")
    private Integer maxWait;
    @Value("${spring.datasource.timeBetweenEvictionRunsMillis:#{null}}")
    private Integer timeBetweenEvictionRunsMillis;
    @Value("${spring.datasource.minEvictableIdleTimeMillis:#{null}}")
    private Integer minEvictableIdleTimeMillis;
    @Value("${spring.datasource.validationQuery:#{null}}")
    private String validationQuery;
    @Value("${spring.datasource.testWhileIdle:#{null}}")
    private Boolean testWhileIdle;
    @Value("${spring.datasource.testOnBorrow:#{null}}")
    private Boolean testOnBorrow;
    @Value("${spring.datasource.testOnReturn:#{null}}")
    private Boolean testOnReturn;
    @Value("${spring.datasource.poolPreparedStatements:#{null}}")
    private Boolean poolPreparedStatements;
    @Value("${spring.datasource.maxPoolPreparedStatementPerConnectionSize:#{null}}")
    private Integer maxPoolPreparedStatementPerConnectionSize;
    @Value("${spring.datasource.filters:#{null}}")
    private String filters;
    @Value("{spring.datasource.connectionProperties:#{null}}")
    private String connectionProperties;

    @Bean
    @Primary
    public DataSource dataSource(){
        DruidDataSource datasource = new DruidDataSource();

        datasource.setUrl(this.dbUrl);
        datasource.setUsername(username);
        datasource.setPassword(password);
        datasource.setDriverClassName(driverClassName);
        //configuration
        if(initialSize != null) {
            datasource.setInitialSize(initialSize);
        }
        if(minIdle != null) {
            datasource.setMinIdle(minIdle);
        }
        if(maxActive != null) {
            datasource.setMaxActive(maxActive);
        }
        if(maxWait != null) {
            datasource.setMaxWait(maxWait);
        }
        if(timeBetweenEvictionRunsMillis != null) {
            datasource.setTimeBetweenEvictionRunsMillis(timeBetweenEvictionRunsMillis);
        }
        if(minEvictableIdleTimeMillis != null) {
            datasource.setMinEvictableIdleTimeMillis(minEvictableIdleTimeMillis);
        }
        if(validationQuery!=null) {
            datasource.setValidationQuery(validationQuery);
        }
        if(testWhileIdle != null) {
            datasource.setTestWhileIdle(testWhileIdle);
        }
        if(testOnBorrow != null) {
            datasource.setTestOnBorrow(testOnBorrow);
        }
        if(testOnReturn != null) {
            datasource.setTestOnReturn(testOnReturn);
        }
        if(poolPreparedStatements != null) {
            datasource.setPoolPreparedStatements(poolPreparedStatements);
        }
        if(maxPoolPreparedStatementPerConnectionSize != null) {
            datasource.setMaxPoolPreparedStatementPerConnectionSize(maxPoolPreparedStatementPerConnectionSize);
        }

        if(connectionProperties != null) {
            datasource.setConnectionProperties(connectionProperties);
        }

        List<Filter> filters = new ArrayList<>();
        filters.add(statFilter());
        filters.add(wallFilter());
        datasource.setProxyFilters(filters);

        return datasource;
    }

    @Bean
    public ServletRegistrationBean druidServlet() {
        ServletRegistrationBean servletRegistrationBean = new ServletRegistrationBean(new StatViewServlet(), "/druid/*");

        //控制台管理用户，加入下面2行 进入druid后台就需要登录
        //servletRegistrationBean.addInitParameter("loginUsername", "admin");
        //servletRegistrationBean.addInitParameter("loginPassword", "admin");
        return servletRegistrationBean;
    }

    @Bean
    public FilterRegistrationBean filterRegistrationBean() {
        FilterRegistrationBean filterRegistrationBean = new FilterRegistrationBean();
        filterRegistrationBean.setFilter(new WebStatFilter());
        filterRegistrationBean.addUrlPatterns("/*");
        filterRegistrationBean.addInitParameter("exclusions", "*.js,*.gif,*.jpg,*.png,*.css,*.ico,/druid/*");
        filterRegistrationBean.addInitParameter("profileEnable", "true");
        return filterRegistrationBean;
    }

    @Bean
    public StatFilter statFilter(){
        StatFilter statFilter = new StatFilter();
        statFilter.setLogSlowSql(true); //slowSqlMillis用来配置SQL慢的标准，执行时间超过slowSqlMillis的就是慢。
        statFilter.setMergeSql(true); //SQL合并配置
        statFilter.setSlowSqlMillis(1000);//slowSqlMillis的缺省值为3000，也就是3秒。
        return statFilter;
    }

    @Bean
    public WallFilter wallFilter(){
        WallFilter wallFilter = new WallFilter();
        //允许执行多条SQL
        WallConfig config = new WallConfig();
        config.setMultiStatementAllow(true);
        wallFilter.setConfig(config);
        return wallFilter;
    }
}
```
## mybatis相关依赖
```
<!--mybatis-->
<dependency>
    <groupId>org.mybatis.spring.boot</groupId>
    <artifactId>mybatis-spring-boot-starter</artifactId>
    <version>1.3.1</version>
</dependency>
<!--通用mapper-->
<dependency>
    <groupId>tk.mybatis</groupId>
    <artifactId>mapper-spring-boot-starter</artifactId>
    <version>1.1.5</version>
</dependency>
<!--pagehelper 分页插件-->
<dependency>
    <groupId>com.github.pagehelper</groupId>
    <artifactId>pagehelper-spring-boot-starter</artifactId>
    <version>1.2.3</version>
</dependency>

<build>
    <plugins>
        <plugin>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-maven-plugin</artifactId>
        </plugin>

        <plugin>
            <groupId>org.mybatis.generator</groupId>
            <artifactId>mybatis-generator-maven-plugin</artifactId>
            <version>1.3.5</version>
            <dependencies>
                <!--配置这个依赖主要是为了等下在配置mybatis-generator.xml的时候可以不用配置classPathEntry这样的一个属性，避免代码的耦合度太高-->
                <dependency>
                    <groupId>mysql</groupId>
                    <artifactId>mysql-connector-java</artifactId>
                    <version>5.1.44</version>
                </dependency>
                <dependency>
                    <groupId>tk.mybatis</groupId>
                    <artifactId>mapper</artifactId>
                    <version>3.4.0</version>
                </dependency>
            </dependencies>
            <executions>
                <execution>
                    <id>Generate MyBatis Artifacts</id>
                    <phase>package</phase>
                    <goals>
                        <goal>generate</goal>
                    </goals>
                </execution>
            </executions>
            <configuration>
                <!--允许移动生成的文件 -->
                <verbose>true</verbose>
                <!-- 是否覆盖 -->
                <overwrite>true</overwrite>
                <!-- 自动生成的配置 -->
                <configurationFile>src/main/resources/mybatis-generator.xml</configurationFile>
            </configuration>
        </plugin>
    </plugins>
</build>
```

上面引入了一些依赖以及generator的配置，这里generator配置文件指向
src/main/resources/mybatis-generator.xml文件，具体一会贴出。

对应的application.properties配置：
```
#指定bean所在包
mybatis.type-aliases-package=com.dudu.domain
#指定映射文件
mybatis.mapperLocations=classpath:mapper/*.xml

#mapper
#mappers 多个接口时逗号隔开
mapper.mappers=com.dudu.util.MyMapper
mapper.not-empty=false
mapper.identity=MYSQL

#pagehelper
pagehelper.helperDialect=mysql
pagehelper.reasonable=true
pagehelper.supportMethodsArguments=true
pagehelper.params=count=countSql
```
## 通用Mapper配置
通用Mapper都可以极大的方便开发人员,对单表封装了许多通用方法，省掉自己写增删改查的sql。

通用Mapper插件网址：https://github.com/abel533/Mapper

```
package com.dudu.util;

import tk.mybatis.mapper.common.Mapper;
import tk.mybatis.mapper.common.MySqlMapper;

/**
 * 继承自己的MyMapper
 *
 * @author
 * @since 2017-06-26 21:53
 */
public interface MyMapper<T> extends Mapper<T>, MySqlMapper<T> {
    //FIXME 特别注意，该接口不能被扫描到，否则会出错
}
```

这里实现一个自己的接口,继承通用的mapper，关键点就是这个接口不能被扫描到，不能跟dao这个存放mapper文件放在一起。

最后在启动类中通过MapperScan注解指定扫描的mapper路径：
```
package com.dudu;
@SpringBootApplication
//启注解事务管理
@EnableTransactionManagement  // 启注解事务管理，等同于xml配置方式的 <tx:annotation-driven />
@MapperScan(basePackages = "com.dudu.dao", markerInterface = MyMapper.class)
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
```

## MyBatis Generator配置
这里配置一下上面提到的mybatis-generator.xml文件,该配置文件用来自动生成表对应的Model,Mapper以及xml,该文件位于src/main/resources下面
Mybatis Geneator 详解: http://blog.csdn.net/isea533/article/details/42102297

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE generatorConfiguration
        PUBLIC "-//mybatis.org//DTD MyBatis Generator Configuration 1.0//EN"
        "http://mybatis.org/dtd/mybatis-generator-config_1_0.dtd">
<generatorConfiguration>
    <!--加载配置文件，为下面读取数据库信息准备-->
    <properties resource="application.properties"/>

    <context id="Mysql" targetRuntime="MyBatis3Simple" defaultModelType="flat">

        <plugin type="tk.mybatis.mapper.generator.MapperPlugin">
            <property name="mappers" value="com.dudu.util.MyMapper" />
            <!--caseSensitive默认false，当数据库表名区分大小写时，可以将该属性设置为true-->
          <property name="caseSensitive" v
          alue="true"/>
        </plugin>

        <!-- 阻止生成自动注释 -->
        <commentGenerator>
            <property name="javaFileEncoding" value="UTF-8"/>
            <property name="suppressDate" value="true"/>
            <property name="suppressAllComments" value="true"/>
        </commentGenerator>

        <!--数据库链接地址账号密码-->
        <jdbcConnection driverClass="${spring.datasource.driver-class-name}"
                        connectionURL="${spring.datasource.url}"
                        userId="${spring.datasource.username}"
                        password="${spring.datasource.password}">
        </jdbcConnection>

        <javaTypeResolver>
            <property name="forceBigDecimals" value="false"/>
        </javaTypeResolver>

        <!--生成Model类存放位置-->
        <javaModelGenerator targetPackage="com.dudu.domain" targetProject="src/main/java">
            <property name="enableSubPackages" value="true"/>
            <property name="trimStrings" value="true"/>
        </javaModelGenerator>

        <!--生成映射文件存放位置-->
        <sqlMapGenerator targetPackage="mapper" targetProject="src/main/resources">
            <property name="enableSubPackages" value="true"/>
        </sqlMapGenerator>

        <!--生成Dao类存放位置-->
        <!-- 客户端代码，生成易于使用的针对Model对象和XML配置文件 的代码
                type="ANNOTATEDMAPPER",生成Java Model 和基于注解的Mapper对象
                type="XMLMAPPER",生成SQLMap XML文件和独立的Mapper接口
        -->
       <javaClientGenerator type="XMLMAPPER" targetPackage="com.dudu.dao" targetProject="src/main/java">
            <property name="enableSubPackages" value="true"/>
       </javaClientGenerator>

        <!--生成对应表及类名
        去掉Mybatis Generator生成的一堆 example
        -->
        <table tableName="LEARN_RESOURCE" domainObjectName="LearnResource" enableCountByExample="false" enableUpdateByExample="false" enableDeleteByExample="false" enableSelectByExample="false" selectByExampleQueryId="false">
            <generatedKey column="id" sqlStatement="Mysql" identity="true"/>
        </table>
    </context>
</generatorConfiguration>
```

其中，我们通过<properties resource="application.properties"/>引入了配置文件，这样下面指定数据源的时候不用写死。

其中tk.mybatis.mapper.generator.MapperPlugin很重要，用来指定通用Mapper对应的文件，这样我们生成的mapper都会继承这个通用Mapper。
```
<plugin type="tk.mybatis.mapper.generator.MapperPlugin">
    <property name="mappers" value="com.dudu.util.MyMapper" />
  <!--caseSensitive默认false，当数据库表名区分大小写时，可以将该属性设置为true-->
  <property name="caseSensitive" value="true"/>
</plugin>
```

这样就可以通过mybatis-generator插件生成对应的文件啦


文章链接：https://www.jianshu.com/p/3d1185e8f6d4
