package Week02
import (
	"github.com/pkg/errors"
)

func queryUserBiz(id int)(user,error){
	user ,err := queryUser(id);
	if err != nil{
		errors.Wrap(err, "dao: failed query all users")
	}
	return user,err
	
}