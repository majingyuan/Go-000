package Week02

func queryUserInfo(id int)(user,error){
	user ,err := queryUserBiz(id)
	
	return user,err
}