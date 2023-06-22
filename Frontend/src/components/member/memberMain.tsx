import MemberTable from "./parts/memberTable.tsx";
import MemberBox from "./parts/memberBox.tsx";

function MemberMain() {
  const isMobile = window.innerWidth < 768; // Example threshold for mobile

  return (
    <>
      {isMobile ? <MemberBox /> : <MemberTable />}
    </>
  );
}

export default MemberMain;
