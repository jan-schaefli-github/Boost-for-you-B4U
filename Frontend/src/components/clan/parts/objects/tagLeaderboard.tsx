import React, { useEffect, useState } from 'react';

interface LeaderboardEntry {
    clanRank: number;
    name: string;
    role: string;
    trophies: number;
}

interface LeaderboardComponentProps {
    selectedChoice: string;
}

const LeaderboardComponent: React.FC<LeaderboardComponentProps> = ({ selectedChoice }) => {
    const [leaderboard, setLeaderboard] = useState<LeaderboardEntry[]>([]);
    const [errorMessage, setErrorMessage] = useState<string>('');

    useEffect(() => {
        const fetchData = async () => {
            const url = `${import.meta.env.VITE_BASE_URL}/api/clan/members/leaderboard?clanID=` + encodeURIComponent(selectedChoice);

            try {
                const response = await fetch(url.toString(), {
                    headers: {
                      'Access-Control-Allow-Origin': '*'
                    }
                  });
                const placementData = await response.json(); // Access the 'items' array
                if (Array.isArray(placementData) && placementData.length > 0) {
                    const leaderboardEntries: LeaderboardEntry[] = placementData.map((item: LeaderboardEntry) => {
                        return {
                            clanRank: item.clanRank,
                            name: item.name + ' (' + item.role + ')',
                            role: item.role,
                            trophies: item.trophies,
                        };
                    });
                    setLeaderboard(leaderboardEntries);
                    setErrorMessage('');
                } else if (placementData.message === 'Keine Clan-Mitglieder gefunden') {
                    setLeaderboard([]);
                    setErrorMessage('Clan not found');
                } else {
                    console.log('Array not found or empty. ' + placementData);
                }
            } catch (error) {
                console.error(error);
                setErrorMessage('Error fetching leaderboard data');
            }
        };

        fetchData().then(() => {
            console.log('Fetching Leaderboard Data Done');
        });
        }, [selectedChoice]);

    return (
        <div className="leaderboard">
            <h1>Leaderboard</h1>
            {errorMessage ? (
                <p className="error-message">{errorMessage}</p>
            ) : (
                <table>
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Trophies</th>
                    </tr>
                    </thead>
                    <tbody>
                    {leaderboard.map((entry) => (
                        <tr key={entry.clanRank}>
                            <td>{entry.clanRank}</td>
                            <td>{entry.name}</td>
                            <td>{entry.trophies}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default LeaderboardComponent;
