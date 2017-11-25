// Package solution contains solutions to the problem described at https://www.interviewcake.com/question/stock-price.
package solution

// Solution0 solves the problem in O(n2).
func Solution0(stockPricesYesterday []int) int {
	maxProfit := stockPricesYesterday[1] - stockPricesYesterday[0]

	for i := 0; i < len(stockPricesYesterday); i++ {
		buyPrice := stockPricesYesterday[i]
		for k := i + 1; k < len(stockPricesYesterday); k++ {
			sellPrice := stockPricesYesterday[k]
			profit := sellPrice - buyPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// Solution1 solves the problem in O(n) time and O(1) space.
func Solution1(stockPricesYesterday []int) int {
	if len(stockPricesYesterday) < 2 {
		// there is no profit if we can't buy and sell
		return 0
	}

	maxProfit := stockPricesYesterday[1] - stockPricesYesterday[0]
	lowestBuyPrice := stockPricesYesterday[0]

	for i := 1; i < len(stockPricesYesterday); i++ {
		price := stockPricesYesterday[i]
		if price-lowestBuyPrice > maxProfit {
			maxProfit = price - lowestBuyPrice
		}
		if price < lowestBuyPrice {
			lowestBuyPrice = price
		}
	}

	return maxProfit
}
