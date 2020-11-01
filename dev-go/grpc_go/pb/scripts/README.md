# FnuoOS_FNMSF_Translator_Tools

## 要求
python3.6 +

## 修改这些参数
```
if __name__ == "__main__":
    main(
        mysql_host='xxxx',
        mysql_user='xxxx',
        mysql_password='xxxx',
        mysql_db = 'xxxx',
        messgae_name='xxxx',
        table_name='xxxx')
```
1.messgae_name是生成proto 的消息体的名字

## 运行
1. 表字段转go结构体
```
python sql2go.py
```
2. 表字段转proto 消息体
```
python sql2proto.py
```