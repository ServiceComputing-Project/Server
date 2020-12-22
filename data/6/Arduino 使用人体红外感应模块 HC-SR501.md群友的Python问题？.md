通过这篇文章，可以了解到以下内容：

- HC-SR501 传感器的基本信息及接线方法
- HC-SR501 跳线选择的两种时间模式
- HC-SR501 简单功能实验

# HC-SR501 传感器的基本信息及接线方法

HC-SR501 是一款基于热释电效应的人体热释运动传感器，能检测到人体或者动物上发出的红外线。这个传感器模块可以通过两个旋钮调节检测 3 ~ 7 米的范围，5秒至5分钟的延迟时间，还可以通过跳线来选择`单次触发`以及`重复触发模式`。

## HC-SR501 针脚以及控制

HC-SR501 针脚以及调节的细节参考下表，资料来源于 [henrysbench.capnfatz.com](https://link.jianshu.com?t=http://henrysbench.capnfatz.com/henrys-bench/arduino-sensors-and-input/arduino-hc-sr501-motion-sensor-tutorial/)，由本文作者翻译。

![img](https:////upload-images.jianshu.io/upload_images/1638540-24e639b9c07596b8.png?imageMogr2/auto-orient/strip|imageView2/2/w/500/format/webp)

HC-SR501-Motion-Detector-Pin-Outs_zh.png

| 针脚以及控制 | 功能                                                         |
| ------------ | ------------------------------------------------------------ |
| 时间延迟调节 | 用于调节在检测到移动后，维持高电平输出的时间长短，可以调节范围 5秒 ~ 5分钟 |
| 感应距离调节 | 用于调节检测范围，可调节范围 3米 ~ 7米                       |
| 检测模式条件 | 可选择单次检测模式和连续检测模式                             |
| GND          | 接地针脚                                                     |
| VCC          | 接电源针脚                                                   |
| 输出针脚     | 没有检测到移动为低电平，检测到移动输出高电平                 |

## 时间延迟、距离调节方法

*时间延迟调节*
 将菲涅尔透镜朝上，左边旋钮调节时间延迟，顺时针方向增加延迟时间，逆时针方向减少延迟时间。

![img](https:////upload-images.jianshu.io/upload_images/1638540-28ffcf68db502b9e.png?imageMogr2/auto-orient/strip|imageView2/2/w/300/format/webp)

HC-SR501-Time-Delay-Adjustment.png

*距离调节*
 将菲涅尔透镜朝上，右边旋钮调节感应距离长短，顺时针方向减少距离，逆时针方向增加距离。

![img](https:////upload-images.jianshu.io/upload_images/1638540-855b4b9617d9ff07.png?imageMogr2/auto-orient/strip|imageView2/2/w/300/format/webp)

HC-SR501-Sensitivity-Adjust.png

## 检测模式跳线调节

![img](https:////upload-images.jianshu.io/upload_images/1638540-7e5ea743cb2b1815.png?imageMogr2/auto-orient/strip|imageView2/2/w/577/format/webp)

HC-SR501-Trigger-Mode-Selection.png



如上图，旋钮旁边三针脚为检测模式选择跳线，将跳线帽插在如图上方两针脚，即为单次检测模式，下方两针脚为连续检测模式。

*单次检测模式*
 传感器检测到移动，输出高电平后，延迟时间段一结束，输出自动从高电平变成低电平。
 *连续检测模式*
 传感器检测到移动，输出高电平后，如果人体继续在检测范围内移动，传感器一直保持高电平，知道人离开后才延迟将高电平变为低电平。
 **区别**
 两种检测模式的区别，就在检测移动触发后，人体若继续移动，是否持续输出高电平。

# HC-SR501 简单功能实验

接下来，我们将尝试完成一个简单的实验来使用这个传感器。
 首先我们需要以下原件：

| 名称        | 数量 |
| ----------- | ---- |
| Arduino UNO | 1    |
| HC-SR501    | 1    |
| 导线        | 若干 |

然后，将 Arduino 与 传感器按如下图连接：

![img](https:////upload-images.jianshu.io/upload_images/1638540-9b28ad8050502416.png?imageMogr2/auto-orient/strip|imageView2/2/w/264/format/webp)

HC-SR501-Arduino-Tutorial-Hook-Up.png

接下来，将以下程序编译上传到 Arduino 上。



```arduino
int ledPin = 13;
int pirPin = 7;

int pirValue;
int sec = 0;

void setup()
{
    pinMode(ledPin, OUTPUT);
    pinMode(pirPin, INPUT);

    digitalWrite(ledPin, LOW);
    Serial.begin(9600);
}

void loop()
{
    pirValue = digitalRead(pirPin);
    digitalWrite(ledPin, pirValue);
        // 以下注释可以观察传感器输出状态
    // sec += 1;
    // Serial.print("Second: ");
    // Serial.print(sec);
    // Serial.print("PIR value: ");
    // Serial.print(pirValue);
    // Serial.print('\n');
    // delay(1000);
}
```

完成以上步骤后，将 Arduino 通电，如果一切正常，那么在传感器前移动时，Arduino 上的 LED 等会亮，然后可以通过更改跳线接法体验不同检测模式的区别。

# 小结

这篇文章我们了解了 HC-SR501 人体移动感应传感器的用法及调节接线方法，然后分析并了解了两种不同的检测模式的区别，最后完成了一个小实验体验使用人体移动感应传感器的功能。

# 参考资料

[Arduino HC-SR501 Motion Sensor Tutorial](https://link.jianshu.com?t=http://henrysbench.capnfatz.com/henrys-bench/arduino-sensors-and-input/arduino-hc-sr501-motion-sensor-tutorial/)
 [完整版HC-SR501人体感应模块](https://link.jianshu.com?t=https://wenku.baidu.com/view/26ef5a9c49649b6648d747b2.html)



作者：speculatecat
链接：https://www.jianshu.com/p/3f612cb6bf17
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
