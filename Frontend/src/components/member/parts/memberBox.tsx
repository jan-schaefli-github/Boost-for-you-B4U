import { useEffect, useState } from "react";
import "../../../assets/css/member/box.css";
import Tooltip from "../../toolTip";

const today = new Date().toISOString().split("T")[0];

interface WarData {
  clanRank: number;
  name: string;
  role: string;
  fame: number;
  decksUsed: number;
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
  "decksUsed",
  "missedDecks",
  "boatAttacks",
];
const SORT_LABELS: { [key in keyof WarData]: string } = {
  clanRank: "Rank",
  name: "Name",
  role: "Role",
  fame: "Fame",
  decksUsed: "Decks Used",
  missedDecks: "Missed Decks",
  boatAttacks: "Boat Attacks",
  clanStatus: "Clan Status",
};

const fetchUrls = [
  `day-log`,
  `week-log`,
  `whole-log`,
];

const fetchUrlLabels = [
  "Day Log",
  "Week Log",
  "Whole Log",
];

function MemberBox() {
  const [filterStage, setFilterStage] = useState(0);
  const [warData, setWarData] = useState<WarData[]>([]);
  const [sortKeyIndex, setSortKeyIndex] = useState<number>(0);
  const [sortOrder, setSortOrder] = useState<string>("asc");
  const [fetchUrlIndex, setFetchUrlIndex] = useState(0);
  const [clanTag, setClanTag] = useState("#P9UVQCJV");
  const [offset, setOffset] = useState(0);
  const [message, setMessage] = useState("Loading...");

  useEffect(() => {
    fetchWarData();
  }, [clanTag, fetchUrlIndex, offset]);

  const fetchWarData = async () => {

    setMessage("Loading...");

    try {
      const formattedClanTag = clanTag.replace("#", "");
      const url = new URL(
        `${import.meta.env.VITE_BASE_URL}/database/clan/${fetchUrls[fetchUrlIndex]}/${formattedClanTag}/${offset}`
      );
      const response = await fetch(url.toString());

      if (response.ok) {
        const data = await response.json();

        // If the data has an error, set the message and clear the data
        if (data.error) {
  
          if (data.error === "notFound") {
            setMessage("Data not found");
            setWarData([]);
          } else {
            console.error("Failed to fetch war data:", data.error);
            setMessage("Error while fetching war data");
            setWarData([]);
          }
        } else {
          setMessage("");
          setWarData(data);
        }
      } else {
        console.error("Failed to fetch war data:", response.status);
        setMessage("Error while fetching war data");
        setWarData([]);
      }
    } catch (error) {
      console.error("Error while fetching war data:", error);
      setMessage("Error while fetching war data");
      setWarData([]);
    }
  };

  const handleFilterClick = () => {
    setFilterStage((prevStage) => (prevStage + 1) % 3);
  };

  const handleSortKeyChange = () => {
    setSortKeyIndex((prevIndex) => (prevIndex + 1) % SORT_KEYS.length);
  };

  const handleSortOrderChange = () => {
    setSortOrder((prevOrder) => (prevOrder === "asc" ? "desc" : "asc"));
  };

  const handleFetchClick = () => {
    setOffset(0);
    setFetchUrlIndex((prevIndex) => (prevIndex + 1) % fetchUrls.length);
  };

  const handleRemoveOffset = () => {
    setOffset(offset - 1);
    if (offset <= 0) {
      setOffset(0);
    }
  };

  const handleAddOffset = () => {
    setOffset(offset + 1);
  };

  const sortData = (data: WarData[]) => {
    if (!data) {
      return [];
    }
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
            <img src="./clashIcon/icon-new.png" alt="New Player" />
          )}
          <i>{data.role != "" ? data.role : "--"}</i> <br />
          <small>{data.tag}</small>
        </h3>
        <div className="stats-container">
          <Tooltip position={{ top: "-45px", left: "-10px" }} text="Fame">
            <p>
              <img src="./clashIcon/icon-fame.png" alt="Fame" />
              {data.fame}
            </p>
          </Tooltip>
            <Tooltip
              position={{ top: "-45px", left: "-10px" }}
              text="Decks Used"
            >
              <p>
                <img
                  src="./clashIcon/icon-decks-used-to-day.png"
                  alt="Decks Used"
                />
                {data.decksUsed}
              </p>
            </Tooltip>
            {data.boatAttacks === 0 ? (
          <Tooltip
            position={{ top: "-45px", left: "-10px" }}
            text="Missed Decks"
          >
            <p>
              <img src="./clashIcon/icon-decks-missed.png" alt="Missed Decks" />
              {data.missedDecks}
            </p>
          </Tooltip>
          ) : (
            <Tooltip
              position={{ top: "-45px", left: "-10px" }}
              text="Missed Decks"
            >
              <p>
                <img
                  src="./clashIcon/icon-decks-missed-boat-attack.png"
                  alt="Missed Decks"
                />
                {data.decksUsed}
              </p>
            </Tooltip>
          )}
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
        {filterStage === 0 ? (
          <input
            className="filter-input"
            placeholder="Enter clan tag ..."
            value={clanTag}
            onChange={(e) => setClanTag(e.target.value)}
          />
        ) : filterStage === 1 ? (
          <>
            <label>
              <button className="backwards" onClick={handleAddOffset} offset-data={offset}>◄</button>
            </label>
            <label>
              <button className="fetch-url-button" onClick={handleFetchClick}>
                {fetchUrlLabels[fetchUrlIndex]}
              </button>
            </label>
            <label>
              <button className="forwards" onClick={handleRemoveOffset} offset-data={offset}>►</button>
            </label>
          </>
        ) : (
          <>
            <label className=".dropdown-label">
              <button className="sort-key-button" onClick={handleSortKeyChange}>
                {SORT_LABELS[SORT_KEYS[sortKeyIndex]]}
              </button>
            </label>
            <label>
              <button
                className="sort-order-button"
                onClick={handleSortOrderChange}
              >
                {sortOrder === "asc" ? "▲" : "▼"}
              </button>
            </label>
          </>
        )}
        <div className="filter-button" onClick={handleFilterClick}>
          {filterStage === 0 ? (
            <span className="icon">
              <img src="./icon/calendar.svg" alt="calendar" />
            </span>
          ) : filterStage === 1 ? (
            <span className="icon">
              <img src="./icon/filter.svg" alt="filter" />
            </span>
          ) : (
            <span className="icon">
              <img src="./icon/search.svg" alt="search" />
            </span>
          )}
        </div>
      </div>
      <div className="message">{message}</div>
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
