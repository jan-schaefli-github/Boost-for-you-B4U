import MemberTable from "./parts/memberTable.tsx";
import MemberBox from "./parts/memberBox.tsx";
import { useEffect, useState } from "react";

function MemberMain() {
  const [isMobileVar, setIsMobileVar] = useState(false);

  function isMobile() {
    let isMobileVar = window.innerWidth < 768;
    setIsMobileVar(isMobileVar)
  }

  useEffect(() => {
    window.addEventListener('load', isMobile);
    window.addEventListener('resize', isMobile);
  }, []);

  return (
    <>
      {isMobileVar ? <MemberBox /> : <MemberTable />}
    </>
  );
}

export default MemberMain;
