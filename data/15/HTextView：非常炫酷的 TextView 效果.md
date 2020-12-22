这是一款非常炫酷的TextView，绝对给你惊喜。

Github地址：https://github.com/hanks-zyh/HTextView

先放几张效果图：

Default : Scale：
![](https://upload-images.jianshu.io/upload_images/1099585-e434f6bb8c7fe16f.gif?imageMogr2/auto-orient/strip|imageView2/2/format/webp)

EvaporateText：
![](https://upload-images.jianshu.io/upload_images/1099585-a0f96e62371f2043.gif?imageMogr2/auto-orient/strip|imageView2/2/format/webp)

Fall：
![](https://upload-images.jianshu.io/upload_images/1099585-19d63295fa37b431.gif?imageMogr2/auto-orient/strip|imageView2/2/w/470/format/webp)

Line：
![](https://upload-images.jianshu.io/upload_images/1099585-4d4a068b5b806699.gif?imageMogr2/auto-orient/strip|imageView2/2/w/470/format/webp)

Sparkle：
![](https://upload-images.jianshu.io/upload_images/1099585-82934a5b5a4ab2e2.gif?imageMogr2/auto-orient/strip|imageView2/2/w/470/format/webp)

Anvil：
![](https://upload-images.jianshu.io/upload_images/1099585-fbdf3082ce2e5bc3.gif?imageMogr2/auto-orient/strip|imageView2/2/w/548/format/webp)

再说怎样使用：

在Module的build.gradle#dependencies添加:
```
compile 'hanks.xyz:htextview-library:0.1.2'
```
在布局文件的根节点中添加命名空间：
```
xmlns:htext="http://schemas.android.com/apk/res-auto"
```
布局文件中添加HTextView：
```
<com.hanks.htextview.HTextView
       android:id="@+id/htext"
       android:layout_width="match_parent"
       android:layout_height="100dp"
       android:background="#000000"
       android:gravity="center"
       android:textColor="#FFFFFF"
       android:textSize="30sp"
       htext:animateType="anvil"
       />
```
在java代码中使用：
```
hTextView = (HTextView) findViewById(R.id.text);
hTextView.setAnimateType(HTextViewType.LINE);
hTextView.animateText("new simple string"); // animate
```
Note：仅支持sdk版本15以上。

即，在Module的build.gradle#defaultConfig#minSdkVersion值为15
```
defaultConfig {
        ...
        minSdkVersion 15
       ...
    }
```
