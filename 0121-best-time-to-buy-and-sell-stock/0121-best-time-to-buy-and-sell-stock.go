func maxProfit(prices []int) int {

    maxProfit := 0
    
    bestBuy := prices[0]

    for i:=1;i<len(prices);i++ {

        if prices[i] < bestBuy {
            bestBuy = prices[i]
        }

        profit := prices[i] - bestBuy

        if profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}