interface InputProps {
  label: string;
  inputType: React.HTMLInputTypeAttribute;
  value?: any;
  setValue: (input: any) => void;
}

interface InputModalProps {
  modalTitle: string;
  modalInputs: InputProps[];
  showModal: boolean;
  setShowModal: (input: boolean) => void;
  confirmAction: () => void;
}

export function InputModal({
  modalTitle,
  modalInputs,
  showModal,
  setShowModal,
  confirmAction,
}: InputModalProps) {
  if (!showModal) return null;

  const handleConfirmAction = () => {
    confirmAction();
    setShowModal(false);
  };

  return (
    <>
      <div className="justify-center items-center flex overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none">
        <div className="relative w-auto my-6 mx-auto max-w-3xl">
          {/*content*/}
          <div className="border-0 rounded-lg shadow-lg relative flex flex-col w-full bg-white outline-none focus:outline-none">
            {/*header*/}
            <div className="flex items-start justify-between p-5 border-b border-solid border-neutral-300 rounded-t">
              <h3 className="text-3xl font-semibold text-neutral-800">
                {modalTitle}
              </h3>
              <button
                className="p-1 ml-auto bg-transparent border-0 text-black opacity-5 float-right text-3xl leading-none font-semibold outline-none focus:outline-none"
                onClick={() => setShowModal(false)}
              >
                <span className="bg-transparent text-black opacity-5 h-6 w-6 text-2xl block outline-none focus:outline-none">
                  ×
                </span>
              </button>
            </div>

            {/*body*/}
            <div className="relative p-6 flex-auto">
              {modalInputs.map((i) => (
                <div className="flex flex-col items-start my-4 text-neutral-800 text-lg leading-relaxed">
                  <label>{i.label}</label>
                  <input
                    className="p-1 border border-neutral-300 rounded-sm outline-none"
                    type={i.inputType}
                    onChange={(e) => i.setValue(e.target.value)}
                    defaultValue={i.value ?? ""}
                  />
                </div>
              ))}
            </div>

            {/*footer*/}
            <div className="flex items-center justify-end p-6 border-t border-solid border-neutral-300 rounded-b">
              <button
                className="text-red-500 background-transparent font-bold uppercase px-6 py-2 text-sm outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
                type="button"
                onClick={() => setShowModal(false)}
              >
                Cancel
              </button>
              <button
                className="bg-emerald-500 text-white active:bg-emerald-600 font-bold uppercase text-sm px-6 py-3 rounded shadow hover:shadow-lg outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
                type="button"
                onClick={handleConfirmAction}
              >
                Confirm
              </button>
            </div>
          </div>
        </div>
      </div>
      <div className="opacity-25 fixed inset-0 z-40 bg-black"></div>
    </>
  );
}
