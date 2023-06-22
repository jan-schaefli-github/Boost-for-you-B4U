import { useEffect, useState } from 'react';
import '../../assets/css/member-table.css';

interface WarData {
  tag: string;
  name: string;
  clanStatus: number;
  fame: number;
  missedDecks: number;
  decksUsedToday: number;
  [key: string]: string | number;
}

function MemberTable() {
  const [warData, setWarData] = useState<WarData[]>([]);
  const [sortConfig, setSortConfig] = useState<{ key: string; direction: string }>({
    key: '',
    direction: '',
  });

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

  const sortTable = (key: string) => {
    let direction = 'asc';

    if (sortConfig.key === key && sortConfig.direction === 'asc') {
      direction = 'desc';
    }

    const sortedData = [...warData].sort((a, b) => {
      if (a[key] < b[key]) return -1;
      if (a[key] > b[key]) return 1;
      return 0;
    });

    if (direction === 'desc') {
      sortedData.reverse();
    }

    setWarData(sortedData);
    setSortConfig({ key, direction });
  };

  return (
    <div className="container">
      <table>
        <thead>
          <tr>
            <th onClick={() => sortTable('tag')}>Tag</th>
            <th onClick={() => sortTable('name')}>Name</th>
            <th onClick={() => sortTable('fame')}>Fame</th>
            <th onClick={() => sortTable('missedDecks')}>Missed Decks</th>
            <th onClick={() => sortTable('decksUsedToday')}>Decks Used Today</th>
          </tr>
        </thead>
        <tbody>
          {warData.map((data, index) => (
            <tr key={index} className={data.clanStatus === 0 ? 'gray-row' : ''}>
              <td>{data.tag}</td>
              <td>{data.name}</td>
              <td>{data.fame}</td>
              <td>{data.missedDecks}</td>
              <td>{data.decksUsedToday}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default MemberTable;
