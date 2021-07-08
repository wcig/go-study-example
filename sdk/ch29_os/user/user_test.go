package user

import (
	"fmt"
	"os/user"
	"testing"
)

// os/user: 允许按name或id查找用户账号。

// 类型
// 1.Group: 表示一用户组
// type Group struct {
//    Gid  string // group ID
//    Name string // group name
// }
// func LookupGroup(name string) (*Group, error)  // 按名称查找用户组，没有则返回UnknownGroupError类型错误
// func LookupGroupId(gid string) (*Group, error) // 按组id查询用户组，没有则返回UnknownGroupIdError类型错误

// 2.UnknownGroupError: LookupGroup查找不到用户组返回
// type UnknownGroupError string
// func (e UnknownGroupError) Error() string

// 3.UnknownGroupIdError: LookupGroupId查找不到用户组返回
// type UnknownGroupIdError string
// func (e UnknownGroupIdError) Error() string

// 4.UnknownUserError: Lookup查找不到用户返回
// type UnknownUserError
//    func (e UnknownUserError) Error() string

// 5.UnknownUserIdError: LookupId查找不到用户返回
// type UnknownUserIdError
//    func (e UnknownUserIdError) Error() string

// 6.User: 表示用户账号
// type User struct {
//    Uid string
//    Gid string
//    Username string
//    Name string
//    HomeDir string
// }
// func Current() (*User, error)               // 返回当前用户
// func Lookup(username string) (*User, error) // 按用户名称查找用户，没有则返回UnknownUserError类型错误
// func LookupId(uid string) (*User, error)    // 按用户id查找用户，没有则返回UnknownUserIdError类型错误
// func (u *User) GroupIds() ([]string, error) // 返回用户所属的用户组id列表

func TestLookup(t *testing.T) {
	u, err := user.Lookup("yangbo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user uid:%s, gid:%s, username:%s, name:%s, homedir:%s\n",
		u.Uid, u.Gid, u.Username, u.Name, u.HomeDir)
	// output:
	// user uid:501, gid:20, username:yangbo, name:yangbo, homedir:/Users/yangbo
}

func TestLookupId(t *testing.T) {
	u, err := user.LookupId("501")
	if err != nil {
		panic(err)
	}
	fmt.Printf("user uid:%s, gid:%s, username:%s, name:%s, homedir:%s\n",
		u.Uid, u.Gid, u.Username, u.Name, u.HomeDir)
	// output:
	// user uid:501, gid:20, username:yangbo, name:yangbo, homedir:/Users/yangbo
}

func TestLookupGroup(t *testing.T) {
	g, err := user.LookupGroup("staff")
	if err != nil {
		panic(err)
	}
	fmt.Printf("group gid:%s, name:%s\n", g.Gid, g.Name)
	// output:
	// group gid:20, name:staff
}

func TestLookupGroupId(t *testing.T) {
	g, err := user.LookupGroupId("20")
	if err != nil {
		panic(err)
	}
	fmt.Printf("group gid:%s, name:%s\n", g.Gid, g.Name)
	// output:
	// group gid:20, name:staff
}
