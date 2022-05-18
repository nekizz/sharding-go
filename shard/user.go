package shard

import "github.com/go-pg/sharding"

const CreateTableUser = `CREATE TABLE ?shard.users (id bigint DEFAULT ?shard.next_id(), account_id int, name text, emails jsonb)`

type User struct {
	tableName string `sql:"?shard.users"`

	Id        int64
	AccountId int64
	Name      string
	Emails    []string
}

// CreateUser picks shard by account id and creates user in the shard.
func CreateUser(cluster *sharding.Cluster, user *User) error {
	return cluster.Shard(user.AccountId).Insert(user)
}

// GetUser splits shard from user id and fetches user from the shard.
func GetUser(cluster *sharding.Cluster, id int64) (*User, error) {
	var user User
	err := cluster.SplitShard(id).Model(&user).Where("id = ?", id).Select()
	return &user, err
}

// GetUsers picks shard by account id and fetches users from the shard.
func GetUsers(cluster *sharding.Cluster, accountId int64) ([]User, error) {
	var users []User
	err := cluster.Shard(accountId).Model(&users).Where("account_id = ?", accountId).Select()
	return users, err
}
