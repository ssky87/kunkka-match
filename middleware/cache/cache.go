package cache

import (
	"fmt"
	"github.com/shopspring/decimal"
	"kunkka-match/engine"
	"kunkka-match/middleware"
)

func GetSymbols() []string {
	symbols, _ := middleware.RedisClient.LRange("kunkka:match:symbol", 0, -1).Result()
	fmt.Printf("Redis 查询交易对列表: [%v]\n", symbols)
	return symbols
}

// 保存交易标的缓存
func SaveSymbol(symbol string) {
	middleware.RedisClient.LPush("kunkka:match:symbol", symbol)
}

//移除交易对
func RemoveSymbol(symbol string) {

	//lrem 移除等于 symbol的元素
	//当 count>0 时，从表头开始查找移除count个
	//当 count=0 时，从表头开始查找，移除所有等于value的
	middleware.RedisClient.LRem("kunkka:match:symbol", 0, symbol)
}

// 删除所有交易标的： 删除key
func RemoveAllSymbol() {
	middleware.RedisClient.Del("kunkka:match:symbol")
}

//保存交易标的以及价格
func SavePrice(symbol string, price string) {
	middleware.RedisClient.Set("kunkka:match:price:"+symbol, price, 0)
}

// 根据交易标的查询价格
func GetPrice(symbol string) decimal.Decimal {
	str := middleware.RedisClient.Get("kunkka:match:price:" + symbol).Val()
	val, _ := decimal.NewFromString(str)
	return val
}

// 从缓存查询订单
func GetOrder(symbol, orderId string) engine.Order {
	return engine.Order{}
}
