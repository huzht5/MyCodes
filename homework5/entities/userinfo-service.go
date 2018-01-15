package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := mydb.Insert(u)
	checkErr(err)
	return err
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	as := []UserInfo{}
	err := mydb.Desc("i_d").Find(&as)
	checkErr(err)
	return as
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	a := &UserInfo{}
	_, err := mydb.Id(id).Get(a)
	checkErr(err)
	return a
}

// DeleteByID .
func (*UserInfoAtomicService) DeleteByID(id int) error {
	// 软删除
	mydb.Id(id).Delete(&UserInfo{})
	// 真正删除
	mydb.Id(id).Unscoped().Delete(&UserInfo{})
	return nil
}
