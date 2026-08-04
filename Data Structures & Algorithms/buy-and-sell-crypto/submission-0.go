func maxProfit(prices []int) int {
    maxProfit := 0

    for i := 0; i < len(prices); i++ {
        for j := i + 1; j < len(prices); j++ {
            profit := prices[j] - prices[i]
            if profit > maxProfit {
                maxProfit = profit
            }
        }
    }

    return maxProfit
}
