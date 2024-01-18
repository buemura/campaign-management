import axios from "axios";

import { CampaignStatusEnum, ICampaign } from "@/types/campaign";

export type CreateCampaignProps = {
  name: string;
  status: CampaignStatusEnum;
};

export type UpdateCampaignProps = {
  id: string;
  name: string;
  status: CampaignStatusEnum;
};

const apiUrl = "http://127.0.0.1:8080/api/campaign";

export async function getCampaignList(): Promise<ICampaign[] | null> {
  try {
    const { data } = await axios.get<ICampaign[]>(apiUrl);
    return data;
  } catch (error) {
    console.log(error);
    return null;
  }
}

export async function createCampaign(
  body: CreateCampaignProps
): Promise<ICampaign | null> {
  try {
    const { data } = await axios.post<ICampaign>(apiUrl, body);
    return data;
  } catch (error) {
    console.log(error);
    return null;
  }
}

export async function updateCampaignById(
  props: UpdateCampaignProps
): Promise<ICampaign | null> {
  try {
    const { data } = await axios.put<ICampaign>(`${apiUrl}/${props.id}`, props);
    return data;
  } catch (error) {
    console.log(error);
    return null;
  }
}

export async function deleteCampaignById(
  id: string
): Promise<ICampaign | boolean> {
  try {
    await axios.delete<ICampaign>(`${apiUrl}/${id}`);
    return true;
  } catch (error) {
    console.log(error);
    return false;
  }
}
