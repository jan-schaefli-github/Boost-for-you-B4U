import  { useEffect, useState } from "react";
import MemberTable from "./parts/memberTable.tsx";
import MemberBox from "./parts/memberBox.tsx";

function MemberMain() {
  const [isMobileVar, setIsMobileVar] = useState(false);

  function isMobile() {
    let isMobileVar = window.innerWidth < 768;
    setIsMobileVar(isMobileVar);
  }

  useEffect(() => {
    isMobile(); // Initial check when the component mounts

    window.addEventListener('resize', isMobile);

    return () => {
      window.removeEventListener('resize', isMobile);
    };
  }, []);

  return (
    <>
      {isMobileVar ? <MemberBox /> : <MemberTable />}
    </>
  );
}

export default MemberMain;
