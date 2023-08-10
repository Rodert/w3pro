## 案例


```bash
1WETH 兑换多少 BNB

go run main.go -tokenOut 0xB8c77482e45F1F44dE1745F52C74426C631bDD52
```


```bash
1USDT 兑换多少 DAI

go run main.go -tokenIn 0xdAC17F958D2ee523a2206206994597C13D831ec7 -amountIn 1000000
```

#### quoteExactInputSingle 函数的 5 个参数分别代表的意思:

1. address tokenIn - 输入的代币地址

这个参数指定了你要兑换的输入代币的地址。

2. address tokenOut - 输出的代币地址

这个参数指定了你要兑换得到的输出代币的地址。

3. uint24 fee - 交易的费率

这个参数代表了这个交易池的费率。Uniswap v3 支持不同的费率水平,fee 越高获得的流动性越高。

4. uint256 amountIn - 输入的代币数量

这个参数指定了你要交换的输入代币的数量。

5. uint160 sqrtPriceLimitX96 - 限制代币价格的平方根

这个参数用来限制代币价格可以波动的范围。它代表了代币价格平方根的 96 倍(为了增加精度)。

这个函数可以让你按照确切的输入量,获取输出量的估计值