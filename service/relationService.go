package service

import (
	"TikTokLite/log"
	"TikTokLite/proto/pkg"
	"TikTokLite/repository"
	"errors"
)

func RelationAction(toUserId, tokenUserId int64, action string) error {
	//这是关注的第一个业务逻辑，即用户不能关注他自己
	//这里的tokenUserId为将token解析后的user_id，即此时的登陆者，那么toUserId则表示为被关注的用户

	if tokenUserId == toUserId {
		return errors.New("you can't follow yourself") //返回一个错误，表示此时的用户id和将要关注的用户一样
	}

	if action == "1" {
		log.Infof("follow action id:%v,toid:%v", tokenUserId, toUserId)
		err := repository.FollowAction(tokenUserId, toUserId)
		if err != nil {
			return err
		}
	} else {
		//这里类似于点赞操作，其中会有点赞与取消点赞的操作
		log.Infof("unfollow action id:%v,toid:%v", tokenUserId, toUserId)
		err := repository.UnFollowAction(tokenUserId, toUserId)
		if err != nil {
			return err
		}
	}

	return nil
}

func RelationFollowList(userId int64, tokenUserId int64) (*message.DouyinRelationFollowListResponse, error) {
	followList, err := repository.GetFollowList(userId, "follow")
	if err != nil {
		return nil, err
	}
	log.Infof("user:%v, followList:%+v", userId, followList)
	//list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}

	followListResponse := message.DouyinRelationFollowListResponse{
		UserList: make([]*message.User, len(followList)),
		/*ToDo:make([]*message.User, len(followList))：make函数是Go语言的内置函数，用于创建切片、映射或通道。在这个例子中，make函数创建了一个message.User类型的指针切片。切片的长度和容量都是len(followList)，即followList的长度
		使用指针切片
		当我们说指针切片时，我们是指一个切片，它的元素是指针12。这种数据结构在Go语言中非常常见，因为它允许我们创建一个动态的指针数组，这些指针可以指向任何类型的变量12。
		例如，[]*int是一个指针切片，它的元素是指向整数的指针12。我们可以使用这个切片来存储一组整数的内存地址，然后通过这些地址来直接操作这些整数12。
		指针切片在许多场景中都非常有用。例如，当我们需要在函数间共享大量数据时，我们可以使用指针切片，而不是复制整个数据集4。这样可以提高程序的性能，因为复制指针通常比复制数据要快得多4。希望这个解释对您有所帮助！
		*/
	}

	for i, u := range followList {
		follow := messageUserInfo(u)
		/*		if _, ok := list[follow.Id]; ok {
				follow.IsFollow = true
			}*/
		followListResponse.UserList[i] = follow
	}
	/*
		用于处理返回数据类型
	*/
	return &followListResponse, nil
}

func RelationFollowerList(userId int64, tokenUserId int64) (*message.DouyinRelationFollowerListResponse, error) {
	followList, err := repository.GetFollowList(userId, "follower")
	if err != nil {
		return nil, err
	}
	log.Infof("user:%v, followerList:%+v", userId, followList)
	//list, err := tokenFollowList(tokenUserId)
	if err != nil {
		return nil, err
	}
	followListResponse := message.DouyinRelationFollowerListResponse{
		UserList: make([]*message.User, len(followList)),
	}

	for i, u := range followList {
		follow := messageUserInfo(u)
		/*		if _, ok := list[follow.Id]; ok {
				follow.IsFollow = true
			}*/
		followListResponse.UserList[i] = follow
	}

	return &followListResponse, nil
}

func tokenFollowList(userId int64) (map[int64]struct{}, error) {
	m := make(map[int64]struct{}) //创建了一个空的结构体类型映射
	list, err := repository.GetFollowList(userId, "follow")
	if err != nil {
		return nil, err
	}
	for _, u := range list {
		m[u.Id] = struct{}{} //该函数De目的只是为了获得符合条件的键值，即为用户关注的对象的id
		//struct{}表示一个空的结构体，这是一种特殊的数据类型，他没有任何字段
		//struct{}{}是一个空结构体实例，当我们只关心映射的键，而不关心值的时候，我们可以使用空结构体作为映射的值
	}
	return m, nil
}
