import { useEffect, useState } from 'react';
import '../../assets/css/member-box.css';

interface WarData {
  tag: string;
  name: string;
  clanStatus: number;
  fame: number;
  missedDecks: number;
  decksUsedToday: number;
  [key: string]: string | number;
}

function MemberBox() {
  const [warData, setWarData] = useState<WarData[]>([]);
  const [sortKey, setSortKey] = useState<string>('name');
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

  const handleSortKeyChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSortKey(event.target.value);
  };

  const handleSortOrderChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSortOrder(event.target.value);
  };

  const sortData = (data: WarData[]) => {
    const sortedData = [...data];

    sortedData.sort((a, b) => {
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
      <div key={data.tag} className="data-box">
        <h3>{data.name}</h3>
        <p>Tag: {data.tag}</p>
        <p>Clan Status: {data.clanStatus}</p>
        <p>Fame: {data.fame}</p>
        <p>Missed Decks: {data.missedDecks}</p>
        <p>Decks Used Today: {data.decksUsedToday}</p>
      </div>
    ));
  };

  return (
    <div>
      <div>
        <label className='.dropdown-label '>
          Sort By:
          <select className='dropdown-select' value={sortKey} onChange={handleSortKeyChange}>
            <option value="name">Name</option>
            <option value="clanStatus">Clan Status</option>
            <option value="fame">Fame</option>
            <option value="missedDecks">Missed Decks</option>
            <option value="decksUsedToday">Decks Used Today</option>
          </select>
        </label>
        <label>
          Sort Order:
          <select className='dropdown-select' value={sortOrder} onChange={handleSortOrderChange}>
            <option value="asc">Ascending</option>
            <option value="desc">Descending</option>
          </select>
        </label>
      </div>
      <div className="data-box-container">{renderDataBoxes()}</div>
    </div>
  );
}

export default MemberBox;
