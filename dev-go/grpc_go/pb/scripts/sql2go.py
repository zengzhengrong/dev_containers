import pymysql

def main(mysql_host,mysql_user,mysql_password,mysql_db,struct_name,table_name):
    conn = pymysql.connect(host=mysql_host,user=mysql_user,password=mysql_password,db=mysql_db,autocommit=True,charset='utf8')

    cur = conn.cursor()
    sql = f"select COLUMN_NAME,COLUMN_TYPE,COLUMN_DEFAULT,COLUMN_COMMENT from information_schema.COLUMNS where table_name = '{table_name}'"
    cur.execute(sql)

    filter_list = []
    out_put_list = []
    out_put_lines = []
    for field in cur.fetchall():

        filter_colum_dict = {}
        filter_colum_dict['COLUMN_NAME'] = field[0]
        filter_colum_dict['COLUMN_TYPE'] = field[1]
        filter_colum_dict['COLUMN_DEFAULT'] = field[2]
        filter_colum_dict['COLUMN_COMMENT'] = field[3]
        # name
        output_dict = {}
        split_name = field[0].split("_")
        new_split_name = []
        for i  in split_name:
            new_split_name.append(i.title())
        name :str = "".join(new_split_name)
        if name.find('Id') != -1:
            name = name.replace('Id','ID')
        # type
        type_name = ''
        if field[1].startswith('varchar'):
            type_name = 'string'
        elif field[1].startswith('char'):
            type_name = 'string'
        elif field[1].startswith('int'):
            type_name = 'int64'
        elif field[1].startswith('tinyint'):
            type_name = 'int64'
        elif field[1].startswith('text'):
            type_name = 'string'
        elif field[1].startswith('double'):
            type_name = 'float64'
        elif field[1].startswith('float'):
            type_name = 'float64'
        elif field[1].startswith('bigint'):
            type_name = 'int64'
        elif field[1].startswith('datetime'):
            type_name = 'time.Time'
        elif field[1].startswith('datetime') and name == 'DeletedAt':
            type_name = '*time.Time'
        # comment
        default = filter_colum_dict['COLUMN_DEFAULT']
        if default == None:
            default = ''
        if default == '' and type_name == 'int64':
            default = '0'

        comment = f"""`json:"{filter_colum_dict['COLUMN_NAME']}" gorm:"type:{filter_colum_dict['COLUMN_TYPE']};default:'{default}';comment:'{filter_colum_dict['COLUMN_COMMENT']}'"`"""
        if field[1].startswith('text'):
            comment = f"""`json:"{filter_colum_dict['COLUMN_NAME']}" gorm:"type:{filter_colum_dict['COLUMN_TYPE']};comment:'{filter_colum_dict['COLUMN_COMMENT']}'"`"""
        # Output dict  
       
        output_dict['name'] = name
        output_dict['type'] = type_name
        output_dict['comment'] = comment

        # print(filter_colum_dict)
        filter_list.append(filter_colum_dict)
        # print(output_dict)
        out_put_list.append(output_dict)

        # Output line 
        line = f"{output_dict['name']} {output_dict['type']} {output_dict['comment']} \n"
        out_put_lines.append(line)
    head_format = ["// generated table fields to go model by sql2go.py\n",
        "// author:zzr\n",
        "// 含有id或者Id的字段 将变成 ID ，int32 类型的默认值都是0\n",
        "// 注意：这会将表的字段全部输出，需要手动精简和重新规范命名\n",
        "package model\n",
        "\n",
        """import "time"\n""",
        "type %s struct {\n" % (struct_name),
        ]
    common_columns = [
        "// 前4个为通用字段\n",
        """ID uint `gorm:"primary_key;comment:'主键ID'"`\n"""
        """CreatedAt time.Time `gorm:"comment:'创建时间'"`\n"""
        """UpdatedAt time.Time `gorm:"comment:'更新时间'"`\n"""
        """DeletedAt *time.Time `sql:"index" gorm:"comment:'删除时间'"`\n"""
    ]

    trailer_format = ["}"]
    compose = head_format + common_columns + out_put_lines + trailer_format
    # print(compose)
    with open(f'{table_name}.go','w+') as gofile:
        gofile.writelines(compose)

    cur.close()
    conn.close()

if __name__ == "__main__":
    main(
        mysql_host='xxx',
        mysql_user='xxx',
        mysql_password='xx',
        mysql_db = 'xx',
        struct_name='xxx',
        table_name='xxx')