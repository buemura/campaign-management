import { getCampaignList } from "@/api";
import { ICampaign } from "@/types";
import { useEffect, useState } from "react";

export const useFetchCampaigns = () => {
  const [campaigns, setCampaigns] = useState<ICampaign[]>([]);

  const getCampaigns = async () => {
    const res = await getCampaignList();
    setCampaigns(res);
  };

  useEffect(() => {
    first;

    return () => {
      second;
    };
  }, [third]);
};
