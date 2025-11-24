import { useEffect, useState } from "react";
import { createZGComputeNetworkBroker } from "@0glabs/0g-serving-broker";

export function use0GBroker() {
  const [broker, setBroker] = useState(null);
  const [account, setAccount] = useState(null);
  const providerAddress = "0xYourProviderAddress";

  useEffect(() => {
    async function init() {
      if (!window.ethereum) return;

      const accounts = await window.ethereum.request({
        method: "eth_requestAccounts",
      });

      setAccount(accounts[0]);

      const signer = new ethers.providers.Web3Provider(
        window.ethereum
      ).getSigner();

      const brokerInstance = await createZGComputeNetworkBroker(signer);

      setBroker(brokerInstance);
    }

    init();
  }, []);

  return { broker, account, providerAddress };
}
