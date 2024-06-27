package testhelper

import (
	"context"
	"log"
	"testcontainer_demo/dao"
	
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func init() {
	if dao.DB != nil {
		return
	}
	err, mysqlTestUrl := CreateTestMySQLContainer(context.Background())
	if err != nil {
		panic(err)
	}
	dao.DB, err = dao.OpenDB(mysqlTestUrl)
	if err != nil {
		panic(err)
	}
}
func CreateTestMySQLContainer(ctx context.Context) (error, string) {
	container, err := mysql.RunContainer(ctx,
		testcontainers.WithImage("mysql:8.0"),
		mysql.WithDatabase("test_db"),
		mysql.WithUsername("root"),
		mysql.WithPassword("root@123"),
		//也可以使用sql脚本初始化数据库
		//mysql.WithScripts(filepath.Join("..", "testdata", "init-db.sql")
	)
	if err != nil {
		return err, ""
	}
	//容器端口3306暴露到外面
	//port, err := container.MappedPort(ctx, "3306")
	//if err != nil {
	//	return err
	//}
	str, err := container.ConnectionString(ctx)
	if err != nil {
		return err, ""
	}
	log.Printf("can use this connecting string to login in db:%s", str)
	return nil, str
}

//需要其他依赖容器可以类似创建
//func CreateTestRedisContainer(ctx context.Context) error {}
//func CreateTestZKContainer(ctx context.Context) error {}
