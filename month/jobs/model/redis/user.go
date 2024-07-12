package redis

import "jobs/utils"

const UserPrifx = "USER"
const User = "UserId_"

func DelHashUser(userId string) error {
	err := utils.RDB.HDel(utils.Ctx, UserPrifx, User+userId).Err()
	return err
}
