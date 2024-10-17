package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	db             *sql.DB
	tx             *sql.Tx
	SqliteFilePath string
}

func (c *Sqlite) OpenDatabase() error {
	return c.OpenDatabasePath(c.SqliteFilePath)
}

// OpenDatabasePath 打开数据库连接
func (c *Sqlite) OpenDatabasePath(databasePath string) error {
	var err error
	c.db, err = sql.Open("sqlite", databasePath)
	if err != nil {
		log.Println("无法打开数据库:", err)
		return err
	}
	// 这里可以选择检查数据库连接是否实际上是可用的
	// 例如，通过尝试一个简单的查询
	if err = c.db.Ping(); err != nil {
		log.Println("数据库连接失败:", err)
		return err
	}
	return nil
}
func (c *Sqlite) Execute(sql string, args ...any) error {
	// 创建一个表
	_, err := c.db.Exec(sql, args...)
	if err != nil {
		log.Println("执行SQL出错:", err, " path=", c.SqliteFilePath)
		return err
	}
	return nil
}
func (c *Sqlite) BeginTransaction() error {
	var err error
	c.tx, err = c.db.Begin()
	if err != nil {
		return err
	}
	return nil
}
func (c *Sqlite) ExecuteTransaction(sql string, args ...any) error {
	_, err := c.tx.Exec(sql, args...)
	if err != nil {
		log.Println("事务执行SQL出错:", err)
		return err
	}
	return nil

}
func (c *Sqlite) CommitTransaction() error {
	if err := c.tx.Commit(); err != nil {
		log.Println("提交事务失败:", err)
		return err
	}
	c.tx = nil // 重置事务指针
	return nil
}
func (c *Sqlite) RollbackTransaction() error {
	if err := c.tx.Rollback(); err != nil {
		log.Println("回滚事务失败:", err)
		return err
	}
	c.tx = nil // 重置事务指针
	return nil
}
func (c *Sqlite) CloseDatabase() {
	if c.db != nil {
		err := c.db.Close()
		if err != nil {
			log.Println("关闭数据库连接时出错:", err)
		}
	}
}
func (c *Sqlite) ExecuteDataAll(sql string, args ...[]any) (bool, error) {
	if err := c.BeginTransaction(); err != nil {
		return false, err
	}
	var ok bool
	var err error
	for _, argSlice := range args {
		if err = c.ExecuteTransaction(sql, argSlice...); err != nil {
			break
		}
	}
	if ok == false {
		_ = c.RollbackTransaction()
		return ok, err
	}
	if err := c.CommitTransaction(); err != nil {
		return false, err
	}
	return true, nil
}

// QueryData 执行SQL查询并返回结果

func (c *Sqlite) QueryData(query string, args ...any) (*sql.Rows, error) {
	rows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (c *Sqlite) QueryRow(query string, args ...any) *sql.Row {
	row := c.db.QueryRow(query, args...)
	return row
}
