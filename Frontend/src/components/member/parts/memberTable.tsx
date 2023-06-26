import { useEffect, useState } from 'react';
import '../../../assets/css/member/table.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSort, faSortUp, faSortDown } from '@fortawesome/free-solid-svg-icons';

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

  const getSortIcon = (key: string) => {
    if (sortConfig.key === key) {
      return sortConfig.direction === 'asc' ? (
        <FontAwesomeIcon icon={faSortUp} />
      ) : (
        <FontAwesomeIcon icon={faSortDown} />
      );
    } else {
      return <FontAwesomeIcon icon={faSort} />;
    }
  };

  return (
    <div className='memberTable'>
      <table>
        <thead>
          <tr>
            <th onClick={() => sortTable('name')}>
              Name {getSortIcon('name')}
            </th>
            <th onClick={() => sortTable('fame')}>
              Fame {getSortIcon('fame')}
            </th>
            <th onClick={() => sortTable('missedDecks')}>
              Missed Decks {getSortIcon('missedDecks')}
            </th>
            <th onClick={() => sortTable('decksUsedToday')}>
              Decks Used Today {getSortIcon('decksUsedToday')}
            </th>
          </tr>
        </thead>
        <tbody>
          {warData.map((data, index) => (
            <tr
              key={index}
              className={data.clanStatus === 1 ? 'gray-row' : ''}
            >
              <td>{data.name}<br /> <small>{data.tag}</small></td>
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
