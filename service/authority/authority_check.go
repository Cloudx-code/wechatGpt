package authority

type ManageAuthorityService struct {
	UserId   string
	UserName string
}

func NewManageAuthorityService(userId, userName string) *ManageAuthorityService {
	return &ManageAuthorityService{UserId: userId, UserName: userName}
}

func (m *ManageAuthorityService) CheckAuthority() (string, error) {
	//expireTime, err := m.getAuthorityFromCache()
	//if err != nil {
	//	return fmt.Sprintf("获取用户权限失败，请重试,err:%v", err), err
	//}
	//if time.Now().After(expireTime) {
	//	return "超出使用时间，请联系xxx续费", nil
	//}
	return "", nil
}

//// 先查缓存，再查sqlite，都没有则创建sqlite记录。
//func (m *ManageAuthorityService) getAuthorityFromCache() (time.Time, error) {
//	var expireTime time.Time
//	cacheInfo, ok := local_cache.Get(consts.RedisKeyAuthority + m.UserId)
//	if ok {
//		expireTime, _ = cacheInfo.(time.Time)
//		return expireTime, nil
//	}
//	// 从DB获取
//	userInfo, err := sqlite.NewDalUser().GetUserById(m.UserId)
//	if err != nil {
//		logs.Error("fail to GetUserById,err:%v", err)
//		return expireTime, err
//	}
//	if userInfo == nil {
//		err = sqlite.NewDalUser().Register(m.UserName, m.UserId, consts.DefaultFreeTime)
//		if err != nil {
//			logs.Error("fail to Register firstTime,err:%v", err)
//			return expireTime, err
//		}
//		expireTime = time.Now().Add(consts.DefaultFreeTime)
//		return expireTime, nil
//	}
//	// 添加到缓存
//	local_cache.Set(consts.RedisKeyAuthority+m.UserId, userInfo.ExpiresAt)
//	return userInfo.ExpiresAt, nil
//}
