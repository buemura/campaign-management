import { ICampaign } from "../../../types";
import { CampaignTableRow } from "./campaign-table-row";

const tableHeader = ["Id", "Name", "Status", "Created At"];

type CampaignTableProps = {
  campaigns: ICampaign[];
};

export const CampaignTable = ({ campaigns }: CampaignTableProps) => {
  return (
    <div className="flex flex-col bg-white">
      <div className="-m-1.5 overflow-x-auto">
        <div className="p-1.5 min-w-full inline-block align-middle">
          <div className="overflow-hidden">
            <table className="min-w-full divide-y divide-neutral-300">
              <thead>
                <tr>
                  {tableHeader?.map((h) => (
                    <th
                      key={h}
                      scope="col"
                      className="px-6 py-3 text-start text-lg font-medium text-neutral-800"
                    >
                      {h}
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody className="divide-y divide-neutral-200">
                {campaigns?.map((campaign) => (
                  <CampaignTableRow key={campaign.id} {...campaign} />
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
};
