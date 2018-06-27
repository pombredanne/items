package vo

//1.host/client/game/match_query,POST 对阵信息查询
type MatchQuery struct {
	Game      string `json:"game" binding:"required"`       //竞彩类型FT足球，BT篮球
	LotteryId string `json:"lottery_id" binding:"required"` //投注类型,附录B
}

//result type
type MatchQueryRs struct {
}

//2.host/client/game/order,POST,竞彩下注请求
type Order struct {
	Game      string `json:"game" binding:"required"`
	LotteryId string `json:"lottery_id" binding:"required"`
	CouponId  string  `json:"coupon_id" binding:"required"`  //优惠券id
	UserId    string  `json:"user_id" binding:"required"`    //用户id
	OrderInfo string `json:"order_info" binding:"required"` //订单详情,是一个struct的jsonMarshal,string态
}

type GameOrderInfo struct {
	BetCodes string `json:"bet_codes"` //投注号码
	BetMulti string   `json:"bet_multi"` //投注倍数
	BetMoney string    `json:"bet_money"` //投注金额(分)
	BetType  string `json:"bet_type"`  //投注类型
}
//result type
type OrderRs struct{
}

//3.host/client/game/match_result  比赛结果
type MatchResult struct{
	Game string `json:"game"`
	All  string `json:"all"` //0,只请求最近一场，1，最近3天所有结束的比赛结果
}

//result
type MRs struct{
	Game string `json:"game"`
	TeamName string `json:"team_name"`
	MatchDate string `json:"match_date"`
	MatchTime string `json:"match_time"`
	MatchNumber string `json:"match_number"`
	MatchIndex string `json:"match_index"`
	League string `json:"league"`
	FtLetPointMulti string `json:"ft_let_point_multi"` //是个啥子
	BtLetPointMulti string `json:"bt_let_point_multi"`
	BtBasePointMulti string `json:"bt_base_point_multi"`
}

//4.host/client/user/game_orders 竞彩订单和订单详情
type GameOrders struct{
	UserId string `json:"user_id"`
	Limit string `json:"limit"`
	Page string`json:"page"`
	OrderStatus string `json:"order_status"`
}

type GameOrdersInfoRs struct{
}

//5.host/client/game/match_history 近期战绩详情
type MatchHistory struct{
	Game string `json:"game"`
	MatchDate string `json:"match_date"`
	MatchNumber string `json:"match_number"`
}

type MatchHistoryRs struct{
}

//6.host//client/game/match_point 比赛积分详情
type MatchPoint struct{
	Game string `json:"game"`
	MatchDate string `json:"match_date"`
	MatchNumber string `json:"match_number"`
}

type MatchPointRs struct{
}

//7.host/client/game/match_rate 比赛赔率
type MatchRate struct{
	Game string `json:"game"`
	MatchDate string `json:"match_date"`
	MatchNumber string `json:"match_number"`
}
type MatchRateRs struct{
}