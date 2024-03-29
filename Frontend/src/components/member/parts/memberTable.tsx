import { useEffect, useState } from 'react';
import '../../../assets/css/member/table.css';
import Tooltip from '../../toolTip';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSort, faSortUp, faSortDown } from '@fortawesome/free-solid-svg-icons';

const today = new Date().toISOString().split('T')[0];

const offset = 0;

interface WarData {
  clanRank: number;
  name: string;
  tag: string;
  role: string;
  trophies: number;
  clanStatus: number;
  fame: number;
  missedDecks: number;
  decksUsed: number;
  boatAttacks: number;
  repairPoints: number;
  [key: string]: string | number;
}

const fetchUrls = [`day-log`, `week-log`, `whole-log`];

function MemberTable() {
  const [warData, setWarData] = useState<WarData[]>([]);
  const [sortConfig, setSortConfig] = useState<{ key: string; direction: string }>({
    key: 'clanRank',
    direction: 'asc',
  });
  const [clanTag, setClanTag] = useState('#P9UVQCJV');
  const [fetchUrlIndex, setFetchUrlIndex] = useState(0);

  useEffect(() => {
    fetchWarData();
  }, [clanTag]);

  const fetchWarData = async () => {
    try {
      const formattedClanTag = clanTag.replace('#', '');
      const url = new URL(
        `${import.meta.env.VITE_BASE_URL}/database/clan/${
          fetchUrls[fetchUrlIndex]
        }/${formattedClanTag}/${offset}`
      );
      const response = await fetch(url.toString());

      if (response.ok) {
        const WholeData = await response.json();
        const data = WholeData.items;
        const sortedDataAboveZero = [...data].filter(item => item.clanStatus != 0);
        const sortedDataBelowZero = [...data].filter(item => item.clanStatus != 1);

        sortedDataAboveZero.sort((a, b) => {
          if (a.clanRank < b.clanRank) return -1;
          if (a.clanRank > b.clanRank) return 1;
          return 0;
        });

        sortedDataBelowZero.sort((a, b) => {
          if (a.clanRank < b.clanRank) return -1;
          if (a.clanRank > b.clanRank) return 1;
          return 0;
        });

        const sortedData = [...sortedDataAboveZero, ...sortedDataBelowZero];
        setWarData(sortedData);
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
  
    const sortedDataAboveZero = [...warData].filter(item => item.clanStatus != 0);
    const sortedDataBelowZero = [...warData].filter(item => item.clanStatus != 1);
  
    sortedDataAboveZero.sort((a, b) => {
      if (key === 'role') {
        const roleOrder = ['leader', 'coLeader', 'elder', 'member'];
        return roleOrder.indexOf(a.role) - roleOrder.indexOf(b.role);
      } else {
        if (a[key] < b[key]) return -1;
        if (a[key] > b[key]) return 1;
        return 0;
      }
    });
  
    sortedDataBelowZero.sort((a, b) => {
      if (key === 'role') {
        const roleOrder = ['leader', 'coLeader', 'elder', 'member'];
        return roleOrder.indexOf(a.role) - roleOrder.indexOf(b.role);
      } else {
        if (a[key] < b[key]) return -1;
        if (a[key] > b[key]) return 1;
        return 0;
      }
    });
  
    if (direction === 'desc') {
      sortedDataAboveZero.reverse();
      sortedDataBelowZero.reverse();
    }
  
    const sortedData = [...sortedDataAboveZero, ...sortedDataBelowZero];
  
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
      <div>
        <input
          className='clanTagInput'
          type='text'
          value={clanTag}
          onChange={(e) => setClanTag(e.target.value)}
          placeholder='Enter clan tag here (e.g. #P9UVQCJV)'
        />
      </div>
      <table>
        <thead>
          <tr>
            <th onClick={() => sortTable('clanRank')}>
              # {getSortIcon('clanRank')}
            </th>
            <th onClick={() => sortTable('name')}>
              Name {getSortIcon('name')}
            </th>
            <th onClick={() => sortTable('role')}>
              Role {getSortIcon('role')}
            </th>
            <th onClick={() => sortTable('trophies')}>
              <Tooltip position={{ top: '-45px', left: '-150%' }} text='Trophies '>
                <img src="./clashIcon/icon-trophy.png" alt="Trophies " />
              </Tooltip>
              {getSortIcon('trophies')}
            </th>
            <th onClick={() => sortTable('fame')}>
              <Tooltip position={{ top: '-45px', left: '-190%' }} text='Fame'>
                <img src="./clashIcon/icon-fame.png" alt="Fame" />
              </Tooltip>
              {getSortIcon('fame')}
            </th>
            <th onClick={() => sortTable('missedDecks')}>
              <Tooltip position={{ top: '-45px', left: '-150%' }} text='Missed Decks'>
                <img src="./clashIcon/icon-decks-missed.png" alt="Missed Decks" />
              </Tooltip>
              {getSortIcon('missedDecks')}
            </th>
            <th onClick={() => sortTable('decksUsed')}>
              <Tooltip position={{ top: '-45px', left: '-150%' }} text='Decks Used'>
                <img src="./clashIcon/icon-decks-used-to-day.png" alt="Decks Used" />
              </Tooltip>
              {getSortIcon('decksUsed')}
            </th>
            <th onClick={() => sortTable('boatAttacks')}>
              <Tooltip position={{ top: '-45px', left: '-150%' }} text='Boat Attacks'>
                <img src="./clashIcon/icon-boat-attack.png" alt="Boat Attacks" />
              </Tooltip>
              {getSortIcon('boatAttacks')}
            </th>
            <th onClick={() => sortTable('repairPoints')}>
              <Tooltip position={{ top: '-45px', left: '-190%' }} text='Repair Points'>
                <img src="./clashIcon/icon-repair-hammer.png" alt="Repair Points " />
              </Tooltip>
              {getSortIcon('repairPoints')}
            </th>
          </tr>
        </thead>
        <tbody>
          {warData.map((data, index) => (
            <tr key={index} data-clan-status={data.clanStatus}>
              <td>{data.clanRank}</td>
              <td>
                {data.name}{data.joinDate === today && <img src="./clashIcon/icon_new.png" alt="New Player" />}
                <br />
                <span className='tag'>{data.tag}</span>
              </td>
              <td>{data.role !== "" ? data.role : "--"}</td>
              <td>{data.trophies}</td>
              <td>{data.fame}</td>
              <td>{data.missedDecks}</td>
              <td>{data.decksUsed}</td>
              <td>{data.boatAttacks}</td>
              <td>{data.repairPoints}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default MemberTable;
