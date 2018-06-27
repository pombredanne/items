# 在 sqlx 基础上进行 拼接sql
```
初步思想是能够直接直观的看到整个SQL和
通过接收get绑定的请求参数struct直接拼接where
```
- 默认 关键词 map 的 key
    - select sql_where sql_groupby sql_order sql_paging join_groupby
    - insert sql_columns sql_values returning table_name
    - update sql_set sql_where table_name returning
