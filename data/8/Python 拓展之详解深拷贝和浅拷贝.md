**写在之前**

在昨天的文章里 （[零基础学习 Python 之字典](https://mp.weixin.qq.com/s?__biz=MzUxMTc3NTI4Ng==&mid=2247483697&idx=1&sn=3aa7d991bdff25950d8184ad33eaeab7&chksm=f96fc668ce184f7eabdbcfefc38d0162882d2683db8e8113c2ea3c0cb29a543fb5522353f492&scene=21#wechat_redirect)），写字典的方法的时候留了一个小尾巴，那就是 copy() 方法还没讲。一是因为 copy 这个方法比较特殊，不单单是它表面的意思；二是以为昨天的文章写得比较长，可能你看到那的时候就没啥耐心去仔细思考了，但是这个知识点又比较重要，也是面试过程中会被长问起的题，我之前在面试的时候（[干货满满--亲身经历的 Python 面试题](https://mp.weixin.qq.com/s?__biz=MzUxMTc3NTI4Ng==&mid=2247483657&idx=1&sn=6e26aa93d6ad3805d049de21838f0521&chksm=f96fc650ce184f4656d3c9e3aeacf6281ceb01f6f8d0f5c26c01cab73b0741e41fefb390e3fb&scene=21#wechat_redirect)）就被问起过。所以我把 copy 单独摘出来今天单讲。

**正式开始**

首先我在这介绍两个新的小知识，要在下面用到。一个是函数 id() ，另一个是运算符 is。id() 函数就是返回对象的内存地址；is 是比较两个变量的对象引用是否指向同一个对象，在这里请不要和 == 混了，== 是比较两个变量的值是否相等。



```php
>>> a = [1,2,3]
>>> b = [1,2,3]
>>> id(a)
38884552L
>>> a is b
False
>>> a == b
True

</pre>
```

copy 这个词有两种叫法，一种是根据它的发音音译过来的，叫拷贝；另一种就是标准的翻译，叫复制。

其实单从表面意思来说，copy 就是将某件东西再复制一份，但是在很多编程语言中，比如 Python，C++中，它就不是那么的简单了。



```ruby
>>> a = 1
>>> b = a
>>> b
1
```

看到上面的例子，从表面上看我们似乎是得到了两个 1，但是如果你看过我之前写的文章，你应该对一句话有印象，那就是 “变量无类型”， Python 中变量就是一个标签，这里我们有请 id() 闪亮登场，看看它们在内存中的位置。



```ruby
>>> a = 1
>>> b = a
>>> b
1
>>> id(a)
31096808L
>>> id(b)
31096808L
```

看出来了吗，id(a) 和 id(b) 相等，所以并没有两个 1，只是一个 1 而已，只不过是在 1 上贴了两张标签，名字是 a 和 b 罢了，这种现象普遍存在于 Python 之中，这种赋值的方式实现了 “假装” 拷贝，真实的情况还是两个变量和同一个对象之间的引用关系。

我们再来看 copy() 方法：



```ruby
>>> a = {'name':'rocky','like':'python'}
>>> b = a.copy()
>>> b
{'name': 'rocky', 'like': 'python'}
>>> id(a)
31036280L
>>> id(b)
38786728L
```

咦，果然这次得到的 b 和原来的 a 不同，它是在内存中又开辟了一个空间。那么我们这个时候就来推理了，虽然它们两个是一样的，但是它们在两个不同的内存空间里，那么肯定彼此互不干扰，如果我们去把 b 改了，那么 a 肯定不变。



```ruby
>>> b['name'] = 'leey'
>>> b
{'name': 'leey', 'like': 'python'}
>>> a
{'name': 'rocky', 'like': 'python'}
```

结果和我们上面推理的一模一样，所以理解了**对象有类型，变量无类型，变量是对象的标签**，就能正确推断出 Python 提供的结果。

我们接下来在看一个例子，请你在往下看的时候保证上面的你已经懂了，不然容易晕车。



```ruby
>>> a = {'name':'rocky','like':'python'}
>>> b = a
>>> b
{'name': 'rocky', 'like': 'python'}
>>> b['name'] = 'leey'
>>> b
{'name': 'leey', 'like': 'python'}
>>> a
{'name': 'leey', 'like': 'python'}
```

上面的例子看出什么来了吗？修改了 b 对应的字典类型的对象，a 的对象也变了。也就是说， b = a 得到的结果是两个变量引用了同一个对象，但是事情真的这么简单吗？请睁大你的眼睛往下看，重点来了。



```ruby
>>> first = {'name':'rocky','lanaguage':['python','c++','java']}
>>> second = first.copy()
>>> second
{'name': 'rocky', 'lanaguage': ['python', 'c++', 'java']}
>>> id(first)
31036280L
>>> id(second)
38786728L
```

在这里的话没有问题，和我们之前说的一样，second 是从 first 拷贝过来的，它们分别引用的是两个对象。



```ruby
>>> second['lanaguage'].remove('java')
>>> second
{'name': 'rocky', 'lanaguage': ['python', 'c++']}
>>> first
{'name': 'rocky', 'lanaguage': ['python', 'c++']}
```

发现什么了吗？按理说上述例子中 second 的 lanaguage 对应的是一个列表，我删除这个列表里的值，也只应该改变的是 second 啊，为什么连 first 的也会改，不是应该互不干扰吗？是不是很意外？是我们之前说的不对吗？那我们再试试另一个键：



```ruby
>>> second['name'] = 'leey'
>>> second
{'name': 'leey', 'lanaguage': ['python', 'c++']}
>>> first
{'name': 'rocky', 'lanaguage': ['python', 'c++']}
```

前面说的原理是有效的，那这到底是为什么啊，来来来，有请我们的 id() 再次闪亮登场。



```ruby
>>> id(first['name'])
38829152L
>>> id(second['name'])
38817544L
>>> id(first['lanaguage'])
38754120L
>>> id(second['lanaguage'])
38754120L
```

其实这里深层次的原因是和 Python 的存储数据的方式有关，这里不做过多的说明（其实是我也不懂。。 在这里，我们只需要知道的是，当 copy() 的时候，列表这类由字符串，数字等复合而成的对象仍然是复制了引用，也就是贴标签，并没有建立一个新的对象，我们把这种拷贝方式叫做浅拷贝（唉呀妈呀，终于把这个概念引出来了。。，言外之意就是并没有解决深层次的问题，再言外之意就是还有能够解决深层次问题的方法。

确实，在 Python 中还有一个深拷贝（deep copy），在使用它之前要引入一个 copy 模块，我们来试一下。



```ruby
>>> import copy
>>> first = {'name':'rocky','lanaguage':['python','c++','java']}
>>> second = copy.deepcopy(first)
>>> second
{'name': 'rocky', 'lanaguage': ['python', 'c++', 'java']}
>>> second['lanaguage'].remove('java')
>>> second
{'name': 'rocky', 'lanaguage': ['python', 'c++']}
>>> first
{'name': 'rocky', 'lanaguage': ['python', 'c++', 'java']}
```

用了深拷贝以后，果然就不是引用了。

**写在最后**

深拷贝和浅拷贝到这里就讲完了，花了一番功夫总算写的还令自己满意，不知道朋友们看到这里的时候是否是觉得对这一部分豁然开朗，我尽力了。这个拓展也可能是成为一个系列，补充一些我觉得理解起来比较困难或者平时面试求职或者工作中常见的知识点，希望您多捧场。

最后感谢你能看到这里，希望我写的东西能够让你有到收获，但是我还是希望我在文章里插入的代码，你们能自己动手试一下，都很简单。原创不易，每一个字，每一个标点都是自己手敲的，所以希望大家能多给点支持，该关注关注，该点赞点赞，该转发转发，有什么问题欢迎在后台联系我，也可以在公众号 -- Python空间 找到我的微信加我。

The end。



文章链接：https://www.jianshu.com/p/9ed9b5ce7bb0

