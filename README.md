# id-generator

## 介绍
ID 生成器

## 接口
1. 获取通用ID，使用雪花算法（snowflake）
   + 请求地址：/snowflake_id
   + 返回值：
    ```json
    {
        "code": 0,
        "data": {
            "number": 564625288823046144
        }
    }
    ```
        |参数名|参数类型|参数说明|是否必须|
        |---|---|---|---|
        |number|int64|生成值|true|
   
1. 设置项目自增ID信息
   + 请求地址：/set_auto_id
   + 请求参数
    ```json
        {
            "id_type":"string",
            "filling":"0",
            "project_id":"0001",
            "table_name":"test_01",
            "column_name":"demo_01",
            "st_prefix":"bt",
            "n_length":11,
            "st_start":1,
            "n_increment":500
        }
    ```
       + 参数说明
        
        |参数名|参数类型|参数说明|是否必须|默认值|
        |---|---|---|---|---|
        |id_type|string|id类型,int,string|true|int|
        |filling|string|填充值，string类型可用，只一位|false|0|
        |project_id|string|项目ID|true|-|
        |table_name|string|表名|false|-|
        |column_name|string|列名|false|-|
        |st_prefix|string|前缀,string类型可用|false|-|
        |n_length|int|长度，string类型可用|false|-|
        |st_start|int|起始值|false|1|
        |n_increment|int|每次加载步长，性能参数，可忽略|false|500|    
    
    + 返回参数

    ```json
    {
        "code": 0,
        "data": {
            "key_number": "564614232419598336"
        }
    }
    ```
   
        |参数名|参数类型|参数说明|是否必须|
        |---|---|---|---|    
        |key_number|string|标记key|true|

1. 获取自增ID
   + 请求地址：/get_auto_number




1. 根据project_id，table_name，column_name获取标记Key
   + 请求地址：/get_auto_id_key
   + 请求参数  
    ```json
    {
        "project_id":"0001",
        "table_name":"test_01",
        "column_name":"demo_01"
    }
    ```
       + 参数说明
        
        |参数名|参数类型|参数说明|是否必须|默认值|
        |---|---|---|---|---|
        |project_id|string|项目ID|true|-|
        |table_name|string|表名|false|-|
        |column_name|string|列名|false|-|   

    + 返回参数

    ```json
    {
        "code": 0,
        "data": {
            "key_number": "564614232419598336"
        }
    }
    ```
   
        |参数名|参数类型|参数说明|是否必须|
        |---|---|---|---|    
        |key_number|string|标记key|true|


## 数据库文件
+ resources/id_generator.sql
