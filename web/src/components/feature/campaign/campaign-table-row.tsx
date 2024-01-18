import { useState } from "react";

import { deleteCampaignById } from "@/api";
import { ConfirmationModal } from "@/components/ui/modal";
import { ICampaign } from "../../../types";

export const CampaignTableRow = (campaign: ICampaign) => {
  const [showDeleteModal, setShowDeleteModal] = useState(false);

  const handleDeleteCampaign = async () => {
    const res = await deleteCampaignById(campaign.id);
    if (!res) {
      alert("Unable to delete campaign");
    }
    location.reload();
  };

  return (
    <>
      <tr className="hover:bg-neutral-200">
        <td className="px-6 py-4 whitespace-nowrap font-medium text-neutral-800">
          {campaign.id}
        </td>
        <td className="px-6 py-4 whitespace-nowrap font-medium text-neutral-800">
          {campaign.name}
        </td>
        <td className="px-6 py-4 whitespace-nowrap text-neutral-800">
          {new Date(campaign.createdAt).toLocaleDateString("en-US")}
        </td>
        <td className="px-6 py-4 whitespace-nowrap font-medium flex gap-x-2">
          <a
            href={`/campaign/${campaign.id}`}
            className="font-semibold rounded-lg border border-transparent text-blue-600 hover:text-blue-800 disabled:opacity-50 disabled:pointer-events-none"
          >
            Edit
          </a>
          <button
            type="button"
            className="font-semibold rounded-lg border border-transparent text-blue-600 hover:text-blue-800 disabled:opacity-50 disabled:pointer-events-none"
            onClick={() => setShowDeleteModal(!showDeleteModal)}
          >
            Delete
          </button>
        </td>
      </tr>

      {showDeleteModal && (
        <ConfirmationModal
          modalTitle="Delete Account"
          modalText={`Are you sure you want to delete campaign: ${campaign.id} ?`}
          showModal={showDeleteModal}
          setShowModal={setShowDeleteModal}
          confirmAction={handleDeleteCampaign}
        />
      )}
    </>
  );
};
