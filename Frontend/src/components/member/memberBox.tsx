import { useEffect, useState } from 'react';
import '../../assets/css/member/box.css';

interface WarData {
  name: string;
  clanStatus: number;
  fame: number;
  missedDecks: number;
  decksUsedToday: number;
  [key: string]: string | number;
}

const SORT_KEYS: (keyof WarData)[] = ['name', 'fame', 'missedDecks', 'decksUsedToday', 'clanStatus'];
const SORT_LABELS: { [key in keyof WarData]: string } = {
  name: 'Name',
  fame: 'Fame',
  missedDecks: 'Missed Decks',
  decksUsedToday: 'Decks Used Today',
  clanStatus: 'Clan Status',
};

function MemberBox() {
  const [warData, setWarData] = useState<WarData[]>([]);
  const [sortKeyIndex, setSortKeyIndex] = useState<number>(0);
  const [sortOrder, setSortOrder] = useState<string>('asc');

  useEffect(() => {
    fetchWarData();
  }, []);

  const fetchWarData = async () => {
    try {
      const url = new URL('http://localhost:3000/database/clan/warlog');
      const response = await fetch(url.toString());

      if (response.ok) {
        const data = await response.json();
        setWarData(data);
      } else {
        console.error('Failed to fetch war data');
      }
    } catch (error) {
      console.error('Error while fetching war data:', error);
    }
  };

  const handleSortKeyChange = () => {
    setSortKeyIndex(prevIndex => (prevIndex + 1) % SORT_KEYS.length);
  };

  const handleSortOrderChange = () => {
    setSortOrder(prevOrder => (prevOrder === 'asc' ? 'desc' : 'asc'));
  };

  const sortData = (data: WarData[]) => {
    const sortedData = [...data];

    sortedData.sort((a, b) => {
      const sortKey = SORT_KEYS[sortKeyIndex];
      const aValue = a[sortKey];
      const bValue = b[sortKey];

      if (sortOrder === 'asc') {
        if (aValue < bValue) return -1;
        if (aValue > bValue) return 1;
      } else if (sortOrder === 'desc') {
        if (aValue > bValue) return -1;
        if (aValue < bValue) return 1;
      }

      return 0;
    });

    return sortedData;
  };

  const renderDataBoxes = () => {
    const sortedData = sortData(warData);

    return sortedData.map((data: WarData) => (
      <div key={data.tag} className="data-box" data-clan-status={data.clanStatus}>
        <h3>
          {data.name} <br />
          <small>{data.tag}</small>
        </h3>
        <p>Fame: {data.fame}</p>
        <p>Missed Decks: {data.missedDecks}</p>
        <p>Decks Used Today: {data.decksUsedToday}</p>
      </div>
    ));
  };

  const handleScrollToTop = () => {
    window.scrollTo({
      top: 0,
      behavior: 'smooth',
    });
  };

  useEffect(() => {
    const handleScroll = () => {
      const scrollButton = document.getElementById('scroll-button');
      if (scrollButton) {
        scrollButton.style.display = window.scrollY > 0 ? 'block' : 'none';
      }
    };

    window.addEventListener('scroll', handleScroll);
    return () => {
      window.removeEventListener('scroll', handleScroll);
    };
  }, []);

  return (
    <div>
      <div>
        <label className=".dropdown-label">
          Sort By:
          <button className="sort-key-button" onClick={handleSortKeyChange}>
            {SORT_LABELS[SORT_KEYS[sortKeyIndex]]}
          </button>
        </label>
        <label>
          Sort Order:
          <button className="sort-order-button" onClick={handleSortOrderChange}>
            {sortOrder === 'asc' ? 'Ascending ▼' : 'Descending ▲'}
          </button>
        </label>
      </div>
      <div className="data-box-container">{renderDataBoxes()}</div>
      <button id="scroll-button" className="scroll-button" onClick={handleScrollToTop}>
        &#9650;
      </button>
    </div>
  );
}

export default MemberBox;
