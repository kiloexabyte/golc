
// Leetcode #355
type Tweet struct {
	userId  int
	tweetId int
}

type Twitter struct {
	recentTweet   []Tweet
	followeeDict  map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		recentTweet:  []Tweet{},
		followeeDict: make(map[int]map[int]bool),
	}
}

func (tw *Twitter) PostTweet(userId int, tweetId int) {
	tw.recentTweet = append(tw.recentTweet, Tweet{userId, tweetId})
}

func (tw *Twitter) GetNewsFeed(userId int) []int {
	// Ensure the user has a followee list
	if _, exists := tw.followeeDict[userId]; !exists {
		tw.followeeDict[userId] = make(map[int]bool)
	}

	res := []int{}
	count := 0
	for i := len(tw.recentTweet) - 1; i >= 0 && count < 10; i-- {
		tweet := tw.recentTweet[i]
		if tweet.userId == userId || tw.followeeDict[userId][tweet.userId] {
			res = append(res, tweet.tweetId)
			count++
		}
	}
	return res
}

func (tw *Twitter) Follow(followerId int, followeeId int) {
	if _, exists := tw.followeeDict[followerId]; !exists {
		tw.followeeDict[followerId] = make(map[int]bool)
	}
	tw.followeeDict[followerId][followeeId] = true
}

func (tw *Twitter) Unfollow(followerId int, followeeId int) {
	if followees, exists := tw.followeeDict[followerId]; exists {
		delete(followees, followeeId)
	}
}
