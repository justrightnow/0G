import React, { useState } from "react";

export default function TradingBot({ broker, providerAddress }) {
  const [symbol, setSymbol] = useState("BTCUSDT");
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);

  async function runBot() {
    setLoading(true);
    try {
      // 1. 构建 messages（类似 demo）
      const messages = [
        {
          role: "user",
          content: `Fetch Binance price for ${symbol} and give a trading advice.`,
        },
      ];

      // 2. 构建 headers（包含认证信息）
      const headers = await broker.inference.getRequestHeaders(
        providerAddress,
        JSON.stringify(messages)
      );

      // 3. Provider 的 endpoint（假设 Provider 提供 REST Chat 接口）
      const endpoint = "https://your-provider-endpoint/chat/completions";

      // 4. 发送请求
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          ...headers,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          model: "compute-trading-model",
          stream: false,
          messages,
        }),
      });

      const responseContent = await response.text();

      // 5. 进行内容验证（非常重要）
      const isValid = await broker.inference.processResponse(
        providerAddress,
        responseContent,
        "chat-123"
      );

      if (!isValid) {
        setResult("❌ 内容验证失败，请重试。");
      } else {
        setResult(responseContent);
      }
    } catch (e) {
      console.error(e);
      setResult("Error: " + e.message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="p-4 border rounded">
      <h2 className="text-xl font-bold">Trading Bot</h2>

      <div className="mt-3">
        <label>Symbol:</label>
        <input
          className="border p-2 rounded ml-2"
          value={symbol}
          onChange={(e) => setSymbol(e.target.value.toUpperCase())}
        />
      </div>

      <button
        className="mt-3 bg-blue-600 text-white p-2 rounded"
        onClick={runBot}
        disabled={loading}
      >
        {loading ? "Running..." : "Run Bot"}
      </button>

      {result && (
        <div className="mt-4 p-3 bg-gray-100 rounded border">
          <pre>{result}</pre>
        </div>
      )}
    </div>
  );
}
