### 介绍
容器内开发flutter 免除一切下载困扰，舒舒服服做开发  
市面上的大多国内安卓模拟器不支持windows的hyper 虚拟机，
这里推荐一款支持hyper虚拟化安卓模拟器BlueStacks 4
这是国内蓝叠的国际版

### hyper的安卓模拟器
http://www.bluestacks.com/
下载蓝叠国际版BlueStacks 4

##### 修改设备和打开adb
设定 -> 设备 选择三星 S20 ultra  
设定 -> 偏好设置  平台设置，打开adb 同时记住端口

### 在容器里面连接设备


```
adb connect 宿主机ip:adb端口
```

### 修改flutter.gradle 的镜像源
```
code /sdks/flutter/packages/flutter_tools/gradle/flutter.gradle
```

将repositories 里面的内容替换成这个
```
    repositories {
        maven{url 'https://maven.aliyun.com/repository/google'}
        maven{url 'https://maven.aliyun.com/repository/gradle-plugin'}
        maven{url 'https://maven.aliyun.com/repository/public'}
        maven{url 'https://maven.aliyun.com/repository/jcenter'}
        jcenter(){url 'http://maven.aliyun.com/nexus/content/groups/public/'}
        maven { url 'https://mirrors.tuna.tsinghua.edu.cn/flutter/download.flutter.io' }
    }
```

### 最后
先flutter run -v  尝试是否能跑通，如果可以的话就开启debug模式开始玩耍了

### 自行下载安卓对应版本sdk
https://developer.android.google.cn/studio/releases/platforms#10

如有需要代理，请自行添加 例如
````
sdkmanager --install  --no_https --proxy=http --proxy_host=192.168.2.109 --proxy_port=4780 "platforms;android-23"
````