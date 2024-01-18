import { useState } from "react";

import { deleteCampaignById, updateCampaignById } from "@/api";
import { ConfirmationModal, InputModal } from "@/components/ui/modal";
import { CampaignStatusEnum, ICampaign } from "@/types";

export const CampaignTableRow = (campaign: ICampaign) => {
  const [name, setName] = useState(campaign.name);
  const [enabled, setEnabled] = useState(
    campaign.status === CampaignStatusEnum.ENABLED ? true : false
  );
  const [showUpdateModal, setShowUpdateModal] = useState(false);
  const [showDeleteModal, setShowDeleteModal] = useState(false);

  const handleUpdateCampaign = async () => {
    const res = await updateCampaignById({
      id: campaign.id,
      name,
      status: enabled ? CampaignStatusEnum.ENABLED : CampaignStatusEnum.PAUSED,
    });
    if (!res) {
      alert("Unable to create campaign");
    }
    location.reload();
  };

  const handleDeleteCampaign = async () => {
    const isSuccess = await deleteCampaignById(campaign.id);
    if (!isSuccess) {
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
        <td className="px-6 py-4 whitespace-nowrap font-medium text-neutral-800">
          {campaign.status}
        </td>
        <td className="px-6 py-4 whitespace-nowrap text-neutral-800">
          {new Date(campaign.createdAt).toLocaleDateString("en-US")}
        </td>
        <td className="px-6 py-4 whitespace-nowrap font-medium flex gap-x-2">
          <button
            type="button"
            className="font-semibold rounded-lg border border-transparent text-blue-600 hover:text-blue-800 disabled:opacity-50 disabled:pointer-events-none"
            onClick={() => setShowUpdateModal(!showUpdateModal)}
          >
            Edit
          </button>
          <button
            type="button"
            className="font-semibold rounded-lg border border-transparent text-blue-600 hover:text-blue-800 disabled:opacity-50 disabled:pointer-events-none"
            onClick={() => setShowDeleteModal(!showDeleteModal)}
          >
            Delete
          </button>
        </td>
      </tr>

      {showUpdateModal && (
        <InputModal
          modalTitle="Update Campaign"
          modalInputs={[
            {
              label: "Name",
              inputType: "Text",
              value: name,
              setValue: setName,
            },
            {
              label: "Enabled",
              inputType: "checkbox",
              value: enabled,
              setValue: setEnabled,
            },
          ]}
          showModal={showUpdateModal}
          setShowModal={setShowUpdateModal}
          confirmAction={handleUpdateCampaign}
        />
      )}

      {showDeleteModal && (
        <ConfirmationModal
          modalTitle="Delete Campaign"
          modalText={`Are you sure you want to delete campaign: ${campaign.id} ?`}
          showModal={showDeleteModal}
          setShowModal={setShowDeleteModal}
          confirmAction={handleDeleteCampaign}
        />
      )}
    </>
  );
};
