## git 问题汇总

1. **QA:提交代码到gitlab/github，没有提交记录，但是代码已提交上去**
引用：https://www.jianshu.com/p/4c38fe9f1a8d
问题原因：
- 资料不对
```
git show
Author:xxx <youremail@email.com>
```
看到里面的信息和gitlab 上配置的不一样，确认问题
```
git config --global user.email "youremail@email.com"
```
然后一定要再次确认，也就是输入
```
git config --global user.email
```
之后直接回车就好
最后检查下邮箱是否正确

2. 