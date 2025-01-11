package main

const UrlArgus = "https://h5.if.qidian.com/argus/api/"
const UrlArgusV1 = UrlArgus + "v1/"
const UrlArgusV2 = UrlArgus + "v2/"
const UrlCheckIn = UrlArgusV2 + "checkin/checkin"
const UrlAdvMainPage = UrlArgusV2 + "video/adv/mainPage"
const UrlFinishWatch = UrlArgusV1 + "video/adv/finishWatch"

type BaseResp struct {
	Data    interface{} `json:"Data"`
	Message string      `json:"Message"`
	Result  int         `json:"Result"`
}

func (r *BaseResp) IsSuccess() bool {
	return r.Result == 0
}

type CheckinResp struct {
	BaseResp
	Data struct {
		CurrentReadingTime int   `json:"CurrentReadingTime"`
		BookId             int   `json:"BookId"`
		CanRecheckIn       int   `json:"CanRecheckIn"`
		CardType           int   `json:"CardType"`
		ChapterId          int   `json:"ChapterId"`
		Date               int64 `json:"Date"`
		FragId             int   `json:"FragId"`
		HasCheckIn         int   `json:"HasCheckIn"`
		HasVideoCheckin    int   `json:"HasVideoCheckin"`
		IsToday            int   `json:"IsToday"`
		NoBrokenTime       int   `json:"NoBrokenTime"`
		Rewards            []struct {
			Count            int    `json:"Count"`
			HasVideoUrge     int    `json:"HasVideoUrge"`
			LotteryCount     int    `json:"LotteryCount"`
			RewardId         int    `json:"RewardId"`
			RewardImage      string `json:"RewardImage"`
			RewardName       string `json:"RewardName"`
			RewardSimpleName string `json:"RewardSimpleName"`
			StrategyId       int    `json:"StrategyId"`
			Type             int    `json:"Type"`
		} `json:"Rewards"`
		ShareStatus        int    `json:"ShareStatus"`
		ShowAd             int    `json:"ShowAd"`
		TwiceRewardText    string `json:"TwiceRewardText"`
		TwiceRewardTextEnd string `json:"TwiceRewardTextEnd"`
		UserNickName       string `json:"UserNickName"`
		VideoButtonTag     string `json:"VideoButtonTag"`
	} `json:"Data"`
}
type AdvMainPage struct {
	BaseResp
	Data struct {
		Avatar      string `json:"Avatar"`
		BaizeModule struct {
			BubbleText      []interface{} `json:"BubbleText"`
			DecorateId      int           `json:"DecorateId"`
			DecorateTimeout string        `json:"DecorateTimeout"`
			DecrateStauts   int           `json:"DecrateStauts"`
			Energy          int           `json:"Energy"`
		} `json:"BaizeModule"`
		DailyBenefitModule struct {
			Desc       string   `json:"Desc"`
			Process    int      `json:"Process"`
			RotateText []string `json:"RotateText"`
			TaskList   TaskList `json:"TaskList"`
			Title      string   `json:"Title"`
		} `json:"DailyBenefitModule"`
		EncryptedGuid   string `json:"EncryptedGuid"`
		Guid            string `json:"Guid"`
		IndexBannerTabs []struct {
			BannerDesc   string `json:"BannerDesc"`
			BannerId     string `json:"BannerId"`
			BusinessType int    `json:"BusinessType"`
			ButtonDesc   string `json:"ButtonDesc"`
			LinkUrl      string `json:"LinkUrl"`
			TaskType     string `json:"TaskType"`
		} `json:"IndexBannerTabs"`
		LastRewardItems []struct {
			LinkUrl   string `json:"LinkUrl"`
			TabDesc   string `json:"TabDesc"`
			TabImgUrl string `json:"TabImgUrl"`
			TabName   string `json:"TabName"`
		} `json:"LastRewardItems"`
		MoreRewardTab struct {
			CurrentTime int64    `json:"CurrentTime"`
			ExpireTime  int64    `json:"ExpireTime"`
			TaskList    TaskList `json:"TaskList"`
			Title       string   `json:"Title"`
		} `json:"MoreRewardTab"`
		NickName         string          `json:"NickName"`
		SurpriseBenefit  SurpriseBenefit `json:"SurpriseBenefit"`
		UserRewardModule struct {
			RewardCount string `json:"RewardCount"`
			RewardTitle string `json:"RewardTitle"`
			ScoreCount  string `json:"ScoreCount"`
			ScoreTitle  string `json:"ScoreTitle"`
		} `json:"UserRewardModule"`
		VideoRewardTab struct {
			CurrentTime int64    `json:"CurrentTime"`
			ExpireTime  int64    `json:"ExpireTime"`
			TaskList    TaskList `json:"TaskList"`
			Title       string   `json:"Title"`
		} `json:"VideoRewardTab"`
	} `json:"Data"`
}

type SurpriseBenefit struct {
	Desc         string `json:"Desc"`
	IntervalTime int    `json:"IntervalTime"`
	IsFinished   int    `json:"IsFinished"`
	TaskId       string `json:"TaskId"`
	TaskRawId    string `json:"TaskRawId"`
	Title        string `json:"Title"`
}
type TaskList []Task
type Task struct {
	CompleteTime  int    `json:"CompleteTime"`
	Desc          string `json:"Desc"`
	DisplaySort   int    `json:"DisplaySort"`
	Icon          string `json:"Icon"`
	IsFinished    int    `json:"IsFinished"`
	IsReceived    int    `json:"IsReceived"`
	MileStoneType int    `json:"MileStoneType"`
	Process       int    `json:"Process"`
	RewardDesc    string `json:"RewardDesc,omitempty"`
	TaskId        string `json:"TaskId"`
	TaskRawId     string `json:"TaskRawId"`
	TaskType      int    `json:"TaskType"`
	Total         int    `json:"Total"`
	Unit          int    `json:"Unit"`
}

func (adv *AdvMainPage) GetSurpriseBenefit() *SurpriseBenefit {
	return &adv.Data.SurpriseBenefit
}
func (adv *AdvMainPage) GetDailyBenefitTaskList() TaskList {
	return adv.Data.DailyBenefitModule.TaskList
}
func (adv *AdvMainPage) GetMoreRewardTabTaskList() TaskList {
	return adv.Data.MoreRewardTab.TaskList
}

// 只有第一个是10点
func (adv *AdvMainPage) GetVideoRewardTabTaskList(onlyFirst bool) TaskList {
	if onlyFirst {
		ret := adv.Data.VideoRewardTab.TaskList[:1]
		return ret
	}
	return adv.Data.VideoRewardTab.TaskList
}

type FinishWatch struct {
	BaseResp
	Data struct {
		RewardList RewardList `json:"RewardList"`
	}
}
type RewardList []struct {
	Desc string `json:"Desc"`
	Icon string `json:"Icon"`
}
