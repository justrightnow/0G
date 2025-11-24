import dynamic from "next/dynamic";
import { use0GBroker } from "../hooks/use0gBroker.ts";

const TradingBot = dynamic(() => import("../components/TradingBot"), {
  ssr: false,
});

export default function Home() {
  const { broker, account, providerAddress } = use0GBroker();

  if (!account) {
    return <div>请先连接钱包</div>;
  }

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold">Compute Network Trading Bot</h1>

      <TradingBot broker={broker} providerAddress={providerAddress} />
    </div>
  );
}
