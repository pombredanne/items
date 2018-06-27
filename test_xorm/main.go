package main

import (
	_ "github.com/lib/pq"
	"github.com/xormplus/xorm"
	"fmt"
	"time"
	"strconv"
)
//用户表结构
type User struct {
	Id      int       `xorm:"not null pk autoincr INTEGER"`
	Name    string    `xorm:"VARCHAR(20)"`
	Created time.Time `xorm:"default 'now()' DATETIME"`
	ClassId int       `xorm:"default 1 INTEGER"`
}

//Class表结构
type Class struct {
	Id   int    `xorm:"not null pk autoincr INTEGER"`
	Name string `xorm:"VARCHAR(20)"`
}

//中间表结构
type UserClass struct{
	User `xorm:"extends"`
	Name string
}

//此方法仅用于orm查询时，查询表认定
func (UserClass) TableName() string {
	return "public.user"
}
func main() {
	//1.创建db引擎
	db, err := xorm.NewPostgreSQL("postgres://postgres:123@localhost:5432/test?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	//2.显示sql语句
	db.ShowSQL(true)

	//3.设置连接数
	db.SetMaxIdleConns(2000)
	db.SetMaxOpenConns(1000)

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 500)//缓存的条数
	db.SetDefaultCacher(cacher)

	//4.执行插入语句的几种方式
		//4.1 orm插入方式:不好控制，如果仅仅插入的对象的属性是name=ft3,那么id无法同步到数据库自动生成
		if false {
			user := new(User)
			user.Name = "ftq"
			_, err = db.Insert(user)
			if err != nil {
				fmt.Println(err)
			}
		}
	    //4.2 命令插入方式
			//4.2.1 db.Exec():单事务单次提交
			if false{
			sql:="insert into public.user(name) values(?)"
			db.Exec(sql,"ft4")
			}

			//4.2.2 db.SQL().Execute():单事务准备了Statement处理sql语句
			if false{
				sql:="insert into public.user(name) values(?)"
				db.SQL(sql,"ft5").Execute()
			}

			//4.2.3使用sql配置文件管理语句,两种载入配置的方式LoadSqlMap()和RegisterSqlMap(),以及SqlMapClient()替代SQL()
			if false {
				err = db.LoadSqlMap("./sql.xml")
				//err = db.RegisterSqlMap(xorm.Xml("./","sql.xml"))
				if err != nil {
					fmt.Println(err)
				}
				db.SqlMapClient("insert_1","ft7").Execute()
			}

	//5.执行查询的几种方式
		//5.1 orm查询:在user初始化的时候,该orm查询直接通过扫描user类型确定表名;组合使用Where(),Get()
		if false {
			user := new(User)
			boolget ,err2 :=db.Where("id=?",5).Get(user)
			fmt.Println(boolget,err2,user)
		}

		//5.2 orm查询:组合使用Where(),Get(),And()
		if false {
			user := new(User)
			boolget ,err2 :=db.Where("id=?",5).And("name=?","ft7").And("id>?",3).Get(user)
			fmt.Println(boolget,err2,user)
		}
      	//5.3 orm查询: AllCols()查询所有列，Cols()查询部分列，Find()解析多行结果,Get()解析单行结果
        if false {
        	users := new ([]User)
        	err = db.AllCols().Find(users)
        	//err = db.Cols("id","name").Find(users)
        	if err !=nil {
        		fmt.Println(err)
			}
        	fmt.Println(users)
        }
        //5.4 orm查询:连接查询Join()
        if false {
        	users := new([]UserClass)
		    db.Join("INNER","class","user.class_id=class.id").Find(users)
		    //db.SQL("select u.id,u.name,c.name from public.user as u left join public.class as c on u.class_id=c.id").Find(users)
		    fmt.Println(users)

		}
		//5.5 sql查询略
		if false {
			//和insert类似,Find查找多行结果，Get获取 单行结果
			users := new([]UserClass)
			db.SQL("select u.id,u.name,c.name from public.user as u left join public.class as c on u.class_id=c.id").Find(users)
			fmt.Println(users)
			}
		//5.6 链式查找
		if false {
			//值得一提的是，支持查找某行的某个字段，不过一般在sql语句中就可以完成过滤，如果sql语句过于复杂，可以链式查找过滤
			id := db.SQL("select * from public.user").Query().Results[0]["id"]
			fmt.Println(id)
		}
	//6.执行更新
		//6.1 ORM方式: 只有非0值的属性会被更新，user的id和created都是默认零值，不被处理
		if false {
			user :=new(User)
			user.Name="ftx"
			//[xorm] [info]  2018/02/08 12:04:01.330624 [SQL] UPDATE "user" SET "name" = $1 WHERE "id"=$2 []interface {}{"ftx", 4}
			db.Id(4).Update(user)
		}
		//6.2 SQL方式略,和insert类似

	//7.事务
		//7.1简单事务
		if false {
			session :=db.NewSession()
			defer session.Close()

			session.Begin()
			//业务:新添加学生，并且创建新的班级，如果班级因为主键冲突创建失败，则整个事务回滚
			_,err =session.SQL("insert into public.user(name,class_id) values('ft13',2)").Execute()
			//表中已经有id=3的班级了
			_,err =session.SQL("insert into public.class(id,name) values(3,'高中3班')").Execute()
			if err!=nil {
				session.Rollback()
			}
			session.Commit()

		}
		//7.2嵌套事务
		if false {
			session := db.NewSession()
			defer session.Close()
			session.Begin()
			_,err=session.Exec("insert into public.user(name,class_id) values('ft23',2)")
			if err!=nil {
				session.Rollback()
			}
			_,err=session.Exec("insert into public.user(id,name,class_id) values(1,'ft24',2)")
			if err!=nil {
				session.Rollback()
			}

			tx,_:=session.BeginTrans()
			_,err=tx.Session().Exec("insert into public.user(name,class_id) values('ft25',2)")
			if err!=nil {
				tx.RollbackTrans()
			}
			tx.CommitTrans()
			session.Commit()


		}
	//8.缓存:使用Raw方式修改以后，需要清理缓存
	if true {

			//建立500条数据
			session := db.NewSession()
			defer session.Close()
			if false {
				session.Begin()
				for i := 30; i < 530; i++ {
					value := "ft" + strconv.Itoa(i)
					_, err = session.Exec("insert into public.user(name) values(?)", value)
					if err != nil {
						session.Rollback()
					}
				}
				session.Commit()
			}

			//查询前531条数据，并随意输出其中一条
			users := make([]User,10)
			db.SQL("select * from public.user where id<531 order by id").Find(&users)
			fmt.Println("读第一遍:","id:",users[50].Id,"name:",users[50].Name)

			db.SQL("select * from public.user where id<531 order by id").Find(&users)
			fmt.Println("读第二遍:","id:",users[50].Id,"name:",users[50].Name)

			var step int =1
		    stepString := users[50].Name + strconv.Itoa(step)
			session.Exec("update public.user set name=? where id =45",stepString)

			//清理缓存
			db.ClearCache(new(User))

			time.Sleep(5*time.Second)

			session.SQL("select * from public.user where id<531 order by id").Find(&users)
			fmt.Println("读第三遍:","id:",users[50].Id,"name:",users[50].Name)

			//虽然很不好意思，但是就算开启了缓存数据也是脏了

	}
	//9.读写分离
	if false {
	//假设有多台服务器用来响应客户的读请求
	var dbGroup *xorm.EngineGroup
	conns :=[]string {
		"postgres://postgres:123@localhost:5432/test?sslmode=disable",
		"postgres://postgres:123@localhost:5432/test?sslmode=disable",
		"postgres://postgres:123@localhost:5432/test?sslmode=disable",
		"postgres://postgres:123@localhost:5432/test?sslmode=disable",
	}

	//负载均衡策略:(特性自行百度)
		// 1.xorm.RandomPolicy()随机访问负载均衡,
		// 2.xorm.WeightRandomPolicy([]int{2, 3,4})权重随机负载均衡
		// 3.xorm.RoundRobinPolicy() 轮询访问负载均衡
		// 4.xorm.WeightRoundRobinPolicy([]int{2, 3,4}) 权重轮训负载均衡
		// 5.xorm.LeastConnPolicy()最小连接数负载均衡
	dbGroup, err = xorm.NewEngineGroup("postgres", conns, xorm.RoundRobinPolicy())
	//dbGroup使用方法和db一致
		 //简单查询
		 dbGroup.SQL("inser into public.users(name) values('ft2000')").Execute()
		 dbGroup.Exec("inser into public.users(name) values('ft2001')")

		 //事务查询
		 session :=dbGroup.NewSession()
	 	 defer session.Close()
	 	 session.Begin()
 	 	 _,err = session.Exec("inser into public.users(name) values('ft2001')")
 	 	 if err!=nil {
		   session.Rollback()
	 	 }
	 	 session.Commit()
	}
}

//注意:
//1.好像不会默认按id增长排序，所以书写sql语句要提前写好order by id ，楼主没怎么写，咳咳
//2. [5.4] postgresql建表会建在public策略的table里，所以查询语句表明写的是public.xxxx,这也造成了连表orm查询会发生前缀报错
//         比如变成了"SELECT * FROM "public"."user" INNER JOIN "class" ON user.class_id=class.id
//		   这和内部的split有关，
//3.[8.]带的缓存好像很容易失效,在创建500个数据后，经过查查改查的操作，查询到的结果是一样的始终是一样的，本来改值之后应该最后一遍查会变化，
//      然而并没有,缓存功能即使清理了缓存，还是会读到脏的