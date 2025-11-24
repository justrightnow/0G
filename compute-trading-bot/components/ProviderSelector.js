import fetch from "node-fetch";

export async function handleInference(messages) {
  const content = messages[0].content;
  const symbol = content.match(/([A-Z]+USDT)/)[1];

  const binanceUrl =
    "https://fapi.binance.com/fapi/v1/ticker/price?symbol=" + symbol;

  const res = await fetch(binanceUrl);
  const json = await res.json();

  const price = parseFloat(json.price);

  // 简单策略：移动平均策略的非常简化版
  let advice;
  if (price > 50000) advice = "Consider SELL or take profit";
  else if (price < 45000) advice = "Consider BUY at discount level";
  else advice = "Market neutral, wait for breakout";

  return `
Symbol: ${symbol}
Price: ${price}
Advice: ${advice}
  `;
}
