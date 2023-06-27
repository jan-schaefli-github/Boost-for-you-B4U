import { useEffect, useState } from "react";
import "../../../assets/css/member/box.css";
import Tooltip from "../../toolTip";

const today = new Date().toISOString().split("T")[0];

interface WarData {
  clanRank: number;
  name: string;
  role: string;
  fame: number;
  decksUsedToday: number;
  missedDecks: number;
  boatAttacks: number;
  clanStatus: number;
  [key: string]: string | number;
}

const SORT_KEYS: (keyof WarData)[] = [
  "clanRank",
  "name",
  "role",
  "fame",
  "decksUsedToday",
  "missedDecks",
  "boatAttacks",
];
const SORT_LABELS: { [key in keyof WarData]: string } = {
  clanRank: "Rank",
  name: "Name",
  role: "Role",
  fame: "Fame",
  decksUsedToday: "Decks Used Today",
  missedDecks: "Missed Decks",
  boatAttacks: "Boat Attacks",
  clanStatus: "Clan Status",
};

function MemberBox() {
  const [warData, setWarData] = useState<WarData[]>([]);
  const [sortKeyIndex, setSortKeyIndex] = useState<number>(0);
  const [sortOrder, setSortOrder] = useState<string>("asc");

  useEffect(() => {
    fetchWarData();
  }, []);

  const fetchWarData = async () => {
    try {
      const url = new URL(
        "http://localhost:3000/database/clan/warlog/standard"
      );
      const response = await fetch(url.toString());

      if (response.ok) {
        const data = await response.json();
        setWarData(data);
      } else {
        console.error("Failed to fetch war data");
      }
    } catch (error) {
      console.error("Error while fetching war data:", error);
    }
  };

  const handleSortKeyChange = () => {
    setSortKeyIndex((prevIndex) => (prevIndex + 1) % SORT_KEYS.length);
  };

  const handleSortOrderChange = () => {
    setSortOrder((prevOrder) => (prevOrder === "asc" ? "desc" : "asc"));
  };

  const sortData = (data: WarData[]) => {
    const sortedDataAboveZero = data.filter((item) => item.clanStatus >= 1);
    const sortedDataBelowZero = data.filter((item) => item.clanStatus <= 0);

    sortedDataAboveZero.sort((a, b) => {
      if (SORT_KEYS[sortKeyIndex] === "role") {
        if (SORT_KEYS[sortKeyIndex] === "role") {
          const roleOrder = ["leader", "coLeader", "elder", "member", ""];
          if (sortOrder === "asc") {
            return roleOrder.indexOf(a.role) - roleOrder.indexOf(b.role);
          } else if (sortOrder === "desc") {
            return roleOrder.indexOf(b.role) - roleOrder.indexOf(a.role);
          }
        }        
      } else {
        const sortKey = SORT_KEYS[sortKeyIndex];
        const aValue = a[sortKey];
        const bValue = b[sortKey];

        if (sortOrder === "asc") {
          if (aValue < bValue) return -1;
          if (aValue > bValue) return 1;
        } else if (sortOrder === "desc") {
          if (aValue > bValue) return -1;
          if (aValue < bValue) return 1;
        }
      }
      return 0;
    });

    sortedDataBelowZero.sort((a, b) => {
      if (SORT_KEYS[sortKeyIndex] === "role") {
        const roleOrder = ["leader", "coLeader", "elder", "member", ""];
        if (sortOrder === "asc") {
          return roleOrder.indexOf(a.role) - roleOrder.indexOf(b.role);
        } else if (sortOrder === "desc") {
          return roleOrder.indexOf(b.role) - roleOrder.indexOf(a.role);
        }
      } else {
        const sortKey = SORT_KEYS[sortKeyIndex];
        const aValue = a[sortKey];
        const bValue = b[sortKey];
    
        if (sortOrder === "asc") {
          if (aValue < bValue) return -1;
          if (aValue > bValue) return 1;
        } else if (sortOrder === "desc") {
          if (aValue > bValue) return -1;
          if (aValue < bValue) return 1;
        }
      }
    
      return 0;
    });
    

    return [...sortedDataAboveZero, ...sortedDataBelowZero];
  };

  const renderDataBoxes = () => {
    const sortedData = sortData(warData);

    return sortedData.map((data: WarData) => (
      <div
        key={data.tag}
        className="data-box"
        data-clan-status={data.clanStatus}
      >
        <h3>
          {data.name}
          {data.joinDate === today && (
            <img src="./clashIcon/icon_new.png" alt="New Player" />
          )}
          <i>{data.role !== "" ? data.role : "--"}</i> <br />
          <small>{data.tag}</small>
        </h3>
        <div className="stats-container">
          <Tooltip position={{ top: "-45px", left: "-10px" }} text="Fame">
            <p>
              <img src="./clashIcon/icon-fame.png" alt="Fame" />
              {data.fame}
            </p>
          </Tooltip>
          {data.boatAttacks !== 0 ? (
            <Tooltip
              position={{ top: "-45px", left: "-10px" }}
              text="Decks Used Today, Made Boat Attack!"
            >
              <p>
                <img
                  src="./clashIcon/icon_decks_used_to_day_boat_attack.png"
                  alt="Decks Used Today, Made Boat Attack"
                />
                {data.decksUsedToday}
              </p>
            </Tooltip>
          ) : (
            <Tooltip
              position={{ top: "-45px", left: "-10px" }}
              text="Decks Used Today"
            >
              <p>
                <img
                  src="./clashIcon/icon_decks_used_to_day.png"
                  alt="Decks Used Today"
                />
                {data.decksUsedToday}
              </p>
            </Tooltip>
          )}
          <Tooltip
            position={{ top: "-45px", left: "-10px" }}
            text="Missed Decks"
          >
            <p>
              <img src="./clashIcon/icon_decks_missed.png" alt="Missed Decks" />
              {data.missedDecks}
            </p>
          </Tooltip>
        </div>
      </div>
    ));
  };

  const handleScrollToTop = () => {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  };

  useEffect(() => {
    const handleScroll = () => {
      const scrollButton = document.getElementById("scroll-button");
      if (scrollButton) {
        scrollButton.style.display = window.scrollY > 0 ? "block" : "none";
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <div className="member-box">
      <div className="sort-nav">
        <label className=".dropdown-label">
          <button className="sort-key-button" onClick={handleSortKeyChange}>
            {SORT_LABELS[SORT_KEYS[sortKeyIndex]]}
          </button>
        </label>
        <label>
          <button className="sort-order-button" onClick={handleSortOrderChange}>
            {sortOrder === "asc" ? "▲" : "▼"}
          </button>
        </label>
      </div>
      <div className="data-box-container">{renderDataBoxes()}</div>
      <button
        id="scroll-button"
        className="scroll-button"
        onClick={handleScrollToTop}
      >
        &#9650;
      </button>
    </div>
  );
}

export default MemberBox;
