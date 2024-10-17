package dbSqlite

import (
	"database/sql"
	"github.com/WangaduoApi/ad-api-gin/utility/db"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type DbUser struct {
	sqlite db.Sqlite
}

// fileExists 检查给定路径的文件是否存在。如果无法确定（例如，由于权限问题或IO错误），则返回错误。

func (c *DbUser) fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil // 文件存在
	}
	if os.IsNotExist(err) {
		return false, nil // 文件确实不存在
	}
	return false, err // 无法确定（可能是权限问题、IO错误等）
}

// ContentDatabase 链接数据库
func (c *DbUser) ContentDatabase(ID string) error {
	DatabaseFilePath := GetUserIDSqliteFilePath(ID)
	isExits, err := c.fileExists(DatabaseFilePath)
	if err != nil {
		return err
	}
	if isExits == false {
		_ = EnsureDir(DatabaseFilePath)
		//这里进行初始化一个全新的数据库
		if err := c.sqlite.OpenDatabasePath(DatabaseFilePath); err != nil {
			return err
		}
		_ = c.sqlite.Execute(`CREATE TABLE "son_fans" (
  			  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			  "chat_name" text,
			  "chat_platform" TEXT,
			  "chat_img" TEXT,
			  "chat_phone" TEXT,
			  "chat_id" INTEGER,
			  "app_version" TEXT,
			  "enable" integer,
			  "son_id" TEXT,
			  "created_at" DATE,
			  "last_login_time" DATE,
			  "updated_at" DATE,
			  PRIMARY KEY ( "id" ) 
			);
		CREATE INDEX "子账号ID" ON "son_fans" ( "son_id" );`)
		_ = c.sqlite.Execute(`CREATE TABLE "son_fans_info" (
			  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			  "type" TEXT,
			  "chat_id" INTEGER,
			  "chat_img" TEXT,
			  "chat_platform" TEXT,
			  "chat_phone" TEXT,
			  "chat_name" TEXT,
			  "fans_id" INTEGER,
			  "created_at" DATE,
			  "updated_at" DATE 
			);
		CREATE INDEX "粉丝主表ID" ON "son_fans_info" ( "fans_id" );`)
		c.sqlite.CloseDatabase()
	}
	return c.sqlite.OpenDatabasePath(DatabaseFilePath)
}

// ContentDatabaseAndBeginTransaction 链接数据库 并且开启事务
func (c *DbUser) ContentDatabaseAndBeginTransaction(bookID string) error {
	if err := c.ContentDatabase(bookID); err != nil {
		return err
	}
	if err := c.BeginTransaction(); err != nil {
		return err
	}
	return nil
}
func (c *DbUser) BeginTransaction() error {
	return c.sqlite.BeginTransaction()
}
func (c *DbUser) Execute(sql string, args ...any) error {
	return c.sqlite.Execute(sql, args...)
}
func (c *DbUser) RollbackTransaction() error {
	return c.sqlite.RollbackTransaction()
}
func (c *DbUser) CloseDatabase() {
	c.sqlite.CloseDatabase()
}
func (c *DbUser) CommitTransaction() error {
	return c.sqlite.CommitTransaction()
}
func (c *DbUser) ExecuteTransaction(sql string, args ...any) error {
	return c.sqlite.ExecuteTransaction(sql, args...)
}
func (c *DbUser) QueryData(sql string, args ...any) (*sql.Rows, error) {
	return c.sqlite.QueryData(sql, args...)
}

// GetBookIDSqliteFilePath 获取ID 所关联数据库的路径 只返回路径 不判断文件是否存在

func GetUserIDSqliteFilePath(ID string) string {
	return filepath.Join(viper.GetString("dbSqlite.userSqliteRootPath"), ID+"main.db")
}

// EnsureDir 确保目录存在，如果不存在则创建它
func EnsureDir(filePath string) error {
	// 从文件路径中提取目录路径
	dirPath := filepath.Dir(filePath)
	// 使用 Stat 方法检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，使用 MkdirAll 创建目录及其所有父目录
		return os.MkdirAll(dirPath, os.ModePerm) // os.ModePerm 通常是 0777，表示最大可能的访问权限
	}
	// 如果目录已经存在或者在创建过程中发生错误，返回该错误
	return nil
}
