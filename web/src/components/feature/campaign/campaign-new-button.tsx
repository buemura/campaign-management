import { createCampaign } from "@/api";
import { InputModal } from "@/components/ui/modal";
import { CampaignStatusEnum } from "@/types";
import { useState } from "react";

export function CampaignNewButton() {
  const [name, setName] = useState("");
  const [enabled, setEnabled] = useState(true);
  const [showNewModal, setShowNewModal] = useState(false);

  const handleNewCampaign = async () => {
    console.log(enabled);

    const res = await createCampaign({
      name,
      status: enabled ? CampaignStatusEnum.ENABLED : CampaignStatusEnum.PAUSED,
    });
    if (!res) {
      alert("Unable to create campaign");
    }
    // location.reload();
  };

  return (
    <div>
      <button
        className="bg-blue-600 text-neutral-100 p-2 rounded-lg w-40"
        onClick={() => setShowNewModal(!showNewModal)}
      >
        + New Campaign
      </button>

      {showNewModal && (
        <InputModal
          modalTitle="Create Campaign"
          modalInputs={[
            {
              label: "Name",
              inputType: "Text",
              setValue: setName,
            },
            {
              label: "Enabled",
              inputType: "checkbox",
              value: enabled,
              setValue: setEnabled,
            },
          ]}
          showModal={showNewModal}
          setShowModal={setShowNewModal}
          confirmAction={handleNewCampaign}
        />
      )}
    </div>
  );
}
