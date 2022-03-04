package repository

import (
	"database/sql"
	"fmt"
	
)


type UserDetailEntity struct {
	Id 			int64 	`json: "id"`
	FirstName 	string 	`json: "firstname"`
	LastName 	string 	`json: "lastname"`
	Email 		string 	`json: "email"`
	Age 		int64 	`json: "age"`
	

}

type UserDetail interface{
	Insert(entity UserDetailEntity) error //insert to db
	// FindById(id int64) (*UserDetailEntity, error) //select * from ...
	FindAllData() ([]UserDetailEntity, error)
	Update(enttity UserDetailEntity) error //upsate to db
}


type userDetail struct {
	dbConnection *sql.DB
	// dbConnection string
}

func NewUserDetail(dbConnection *sql.DB) UserDetail {
	return &userDetail{
		dbConnection: dbConnection,
	}
}

func (repo userDetail) Insert(entity UserDetailEntity) error {
	//implement db 
	fmt.Println("insert enity ....")
	
	cmd := "INSERT INTO golangtest.userdeatail(firstname, lastname, email, age) VALUES(?,?,?,?)"
	stmt, _ := repo.dbConnection.Prepare(cmd)
	_, err := stmt.Exec(entity.FirstName, entity.LastName, entity.Email, entity.Age)

	if err != nil {
		panic(err)
	}
	
	return nil
}

// func (repo userDetail) FindById(id int64) (*UserDetailEntity,error){
// 	return &UserDetailEntity{
		
// 	}, nil
func (repo userDetail) FindAllData() ([]UserDetailEntity, error){

	result, _ := repo.dbConnection.Query("SELECT * FROM golangtest.userdeatail")
	var userList []UserDetailEntity
	for result.Next() {
		var userEntity UserDetailEntity
		err := result.Scan(
			&userEntity.Id,
			&userEntity.FirstName,
			&userEntity.LastName,
			&userEntity.Email,
			&userEntity.Age,
		)
		if err != nil {
			panic(err.Error())
		}
		userList = append(userList, userEntity)
	}
	return userList, nil
}

func(repo userDetail) Update(entity UserDetailEntity) error {
	cmd := "UPDATE golangtest.userdeatail SET firstname = ?, lastname = ? , age = ? WHERE email = ?"
	stmt, _ := repo.dbConnection.Prepare(cmd)

	_, err := stmt.Exec(entity.FirstName, entity.LastName,  entity.Age,entity.Email,)
	if err != nil {
		panic(err)
	}
	
	return nil
}