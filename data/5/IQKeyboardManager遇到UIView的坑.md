# 踩坑

最新很偶然的机会了解到IQKeyboardManager第三方库，感觉从各方面比较，都比我现在使用的TPKeyboardAvoiding好。所以果断的替换。整个APP测试了一边，感觉不错，完全符合要求。直到5小时前，我开始编写一个UITableView的UITableViewCell中嵌套UITextView，进行调试时，出现了崩溃的现象。

Section header view 好像没有受到UITableView的改变而改变，一直停留在原来的位置。什么鬼啊，直接无语。开始怀疑IQKeyboardManager没有支持这样的场景。下载Demo并在Demo修改UITableView的样例，发现并不存在问题。

- 又开始百度，Google的搜索问题，发现没有人碰到过这个问题。
- 并且在IQKeyboardManager的官网也没有人描写，那么说明这个问题是我特有的。
- 删除设计到UIScrollview相关的第三方库。 问题依然存在。
- 删除所有第三方库。问题依然存在。
- 清空工程，reset模拟器。 问题依然存在。
- 新建一个空工程，复制相同的代码，问题不在了。。。

# 原因

应该还是我的代码的问题，开始筛选原因吧。


![img](https:////upload-images.jianshu.io/upload_images/1693553-4b58b2bea517797d.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

UITableViewWrapperView


 在查看UI界面时，发现了一个奇怪的现象，UITableViewWrapperView为何在键盘弹出后没有改变。这个应该是问题所在。
 但是面对上千个文件，怎么进行查找啊。这个时候才现在组件化真的是很有好处，直接按组件进行排查。_
 最后将问题定位在UIViewExtensions.h中，删除这个文件后，问题被修复了。



# 解决

发现原来UIViewExtensions类中的方法



```objectivec
- (UIView*) superviewOfClassType: (Class) classType
{
    UIView* view = self.superview;
    
    while (view != nil)
    {
        if ([view isKindOfClass: classType])
        {
            return view;
        }
        
        view = view.superview;
    }
    
    return nil;
}
```

与IQKeyboardManager库中IQUIView+Hierarchy类重名了，也就是所说的覆盖了。



```objectivec
-(UIView*)superviewOfClassType:(Class)classType
{
    UIView *superview = self.superview;
    
    while (superview)
    {
        if ([superview isKindOfClass:classType])
        {
            //If it's UIScrollView, then validating for special cases
            if ([superview isKindOfClass:[UIScrollView class]])
            {
                NSString *classNameString = NSStringFromClass([superview class]);

                //  If it's not UITableViewWrapperView class, this is internal class which is actually manage in UITableview. The speciality of this class is that it's superview is UITableView.
                //  If it's not UITableViewCellScrollView class, this is internal class which is actually manage in UITableviewCell. The speciality of this class is that it's superview is UITableViewCell.
                //If it's not _UIQueuingScrollView class, actually we validate for _ prefix which usually used by Apple internal classes
                if ([superview.superview isKindOfClass:[UITableView class]] == NO &&
                    [superview.superview isKindOfClass:[UITableViewCell class]] == NO &&
                    [classNameString hasPrefix:@"_"] == NO)
                {
                    return superview;
                }
            }
            else
            {
                return superview;
            }
        }
        
        superview = superview.superview;
    }
    
    return nil;
}
```

**其实分类的方法重名并不存在覆盖的问题，只是在编译的时候谁的方法在前，那么谁的方法将会被执行。**

修改UIViewExtensions类中的方法后，显示正常。



![img](https:////upload-images.jianshu.io/upload_images/1693553-01054911e3f44758.gif?imageMogr2/auto-orient/strip|imageView2/2/w/326/format/webp)

正常喽

查看键盘弹出的UI界面可以清楚的看到，UITableViewWrapperView在键盘弹出后向上缩进，避免键盘所在的部分被覆盖。

![img](https:////upload-images.jianshu.io/upload_images/1693553-e3ca5e22d446b2a2.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

UITableViewWrapperView正常

还是希望IQKeyboardManager能够模仿AFNetworking或SDWebImage使用自己的前缀，避免这样的问题再次出现。

// END



文章链接：https://www.jianshu.com/p/583116a4a5c1

