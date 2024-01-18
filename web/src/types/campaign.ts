export enum CampaignStatusEnum {
  ENABLED = "ENABLED",
  PAUSED = "PAUSED",
}

export interface ICampaign {
  id: string;
  name: string;
  status: CampaignStatusEnum;
  createdAt: Date;
}
