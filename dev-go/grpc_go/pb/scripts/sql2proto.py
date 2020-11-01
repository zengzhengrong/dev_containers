import pymysql

def main(mysql_host,mysql_user,mysql_password,mysql_db,messgae_name,table_name):
    conn = pymysql.connect(host=mysql_host,user=mysql_user,password=mysql_password,db=mysql_db,autocommit=True,charset='utf8')

    cur = conn.cursor()
    sql = f"select COLUMN_NAME,ORDINAL_POSITION,COLUMN_TYPE,COLUMN_COMMENT from information_schema.COLUMNS where table_name = '{table_name}'"
    cur.execute(sql)
    with open(f'{messgae_name}.proto','w+') as protofile:

        content_list = []
        filter_fields = []
        for field in cur.fetchall():
            # filter
            if field[0] == 'id':
                filter_fields.append(field[0])
                continue
            elif field[0] == 'created_at':
                filter_fields.append(field[0])
                continue
            elif field[0] == 'updated_at':
                filter_fields.append(field[0])
                continue
            elif field[0] == 'deleted_at':
                filter_fields.append(field[0])
                continue
            # type 
            type_name = ''
            if field[2].startswith('varchar'):
                type_name = 'string'
            elif field[2].startswith('char'):
                type_name = 'string'
            elif field[2].startswith('int'):
                type_name = 'int32'
            elif field[2].startswith('tinyint'):
                type_name = 'int32'
            elif field[2].startswith('bigint'):
                type_name = 'int32'
            elif field[2].startswith('text'):
                type_name = 'string'
            elif field[2].startswith('double'):
                type_name = 'double'
            elif field[2].startswith('float'):
                type_name = 'double'
            content_list.append(f'// {field[0]} is {field[3]}\n')
            content_list.append(f'{type_name} {field[0]} = ; \n')
        if len(filter_fields) > 0:
            for i in range(0,len(content_list)):
                if i % 2 == 0:
                    # print(i)
                    # print(content_list[i])
                    index = content_list[i+1].find(';') - 1
                    update_list = list(content_list[i+1])
                    update_list.insert(index,str(int(i/2) + 1))
                    # print(update_list)
                    content = ''.join(update_list)
                    content_list[i+1] = content
                    # [index] = str(int(i/2))
                    # print(content_list[i+1])
                    # number = int(i/2)
                    # print(number)
        head = [
            "// generated table fields to proto message by sql2proto.py\n",
            "// author:zzr\n",
            "// 1.注意：这会将表的字段全部输出，需要手动精简和重新规范命名\n",
            "// 2.注意：会过滤 主键id 、 创建时间created_at 、 更新时间updated_at 、 删除时间deleted_at\n",
            "// 3.已知问题：int64 类型 在swagger ui 显示的是string 目前使用int32 替代\n",
            "\n",
            "message %s {\n" % (messgae_name),
            ]
        trailer_format = ["}\n"]

        compose = head + content_list + trailer_format
        protofile.writelines(compose)
        
    cur.close()
    conn.close()

if __name__ == "__main__":
    main(
        mysql_host='xxx',
        mysql_user='xx',
        mysql_password='xxx',
        mysql_db = 'xxxx',
        messgae_name='xxx',
        table_name='xxx')