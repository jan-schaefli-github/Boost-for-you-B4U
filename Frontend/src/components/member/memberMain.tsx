import MemberTable from "./memberTable.tsx";
import MemberBox from "./memberBox.tsx";

function MemberMain() {
  const isMobile = window.innerWidth < 768; // Example threshold for mobile

  return (
    <>
      {isMobile ? <MemberBox /> : <MemberTable />}
    </>
  );
}

export default MemberMain;
