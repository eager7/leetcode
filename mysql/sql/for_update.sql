BEGIN ; -- 先开始事务A

SELECT * FROM t_assert_info; -- 事务B插入前读取数据为空






SELECT * FROM t_assert_info; -- 事务B提交前读取数据

SELECT * FROM t_assert_info; -- 事务B提交后读取数据
COMMIT ;




BEGIN ; -- 再开始事务B


INSERT INTO
    t_assert_info(`uuid`,`contract`,`balance`,`version`,`chain_id`)
VALUES ('123','con',10,0,1); -- 插入数据

SELECT * FROM t_assert_info; -- 读取数据，可以读到前面插入数据

COMMIT;
