func maxProfit(prices []int) int {
    maxProfit := 0
    curStartValue := prices[0]
    for i:=1; i<len(prices); i++ {
        curProfit := prices[i] - curStartValue
        if curProfit > maxProfit {
            maxProfit = curProfit
        }
        if prices[i] < curStartValue {
            curStartValue = prices[i]
        }
    }

    return maxProfit
}