import { useQuery } from "@tanstack/react-query";

import { getCampaignList } from "@/api";
import {
  CampaignNewButton,
  CampaignTable,
} from "@/components/feature/campaign";
import { TableSkeleton } from "@/components/ui/skeleton";
import { Header } from "./components/layout";

export default function App() {
  const { data, isLoading, error } = useQuery({
    queryKey: ["campaigns"],
    queryFn: getCampaignList,
  });

  if (isLoading) {
    <div className="w-screen h-screen bg-neutral-100">
      <TableSkeleton rowsCount={5} />
    </div>;
  }

  if (!data || error) {
    <div className="w-screen h-screen bg-neutral-100">
      <span>Failed to fetch campaigns. Try again later!</span>
    </div>;
  }

  return (
    <div className="w-screen h-screen bg-neutral-100">
      <Header />

      <div className="p-4 flex flex-col gap-4 mt-20">
        <CampaignNewButton />
        <CampaignTable campaigns={data || []} />
      </div>
    </div>
  );
}
