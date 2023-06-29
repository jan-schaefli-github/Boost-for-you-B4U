import React, {useEffect, useState} from 'react';
import '../../../../assets/css/clan/leaderboard.css';

// Leaderboard Information Interface
interface LeaderboardEntry {
    rank: number;
    name: string;
    clanScore: number;
}

interface SelectInputProps {
    setTrigger: (value: boolean) => void;
    trigger: boolean;
    selectedLocation: number;
}

const SelectInput: React.FC<SelectInputProps> = ({
                                                     setTrigger,
                                                     trigger,
                                                     selectedLocation,
                                                 }) => {
    // Prop for leaderboard Data
    const [leaderboard, setLeaderboard] = useState<LeaderboardEntry[]>([]);
    const url = `${import.meta.env.VITE_BASE_URL}/api/clan/leaderboard?locationID=` + selectedLocation;
    const [errorMessage, setErrorMessage] = useState<string>('')

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(url.toString(), {
                    headers: {
                      'Access-Control-Allow-Origin': '*'
                    }
                  });
                const data = await response.json();
                const placementData = data.items; // Access the 'items' array
                if (Array.isArray(placementData) && placementData.length > 0) {
                    //Map the array to the LeaderboardEntry interface
                    const leaderboardEntries: LeaderboardEntry[] = placementData.map((item: LeaderboardEntry) => {
                        return {
                            rank: item.rank,
                            name: item.name,
                            clanScore: item.clanScore,
                        };
                    });
                    setLeaderboard(leaderboardEntries);
                } else {
                    setErrorMessage('Error fetching leaderboard data')
                    console.log('Array not found or empty. ' + placementData);
                }
            } catch (error) {
                setErrorMessage('Error fetching leaderboard data')
                console.error(error);
            }
        };
        // Check if the trigger prop has changed
        if (trigger) {
            setTrigger(false);
            fetchData().then(() => {
                console.log('Fetching locationClanLeaderboard Data Done');
            });
        }
    }, [setTrigger, trigger, url]);

    return (
        <>
            <div className="leaderboard">
                <h1>Leaderboard</h1>
                {errorMessage ? (
                    <p className="error-message">{errorMessage}</p>
                ) : (
                <table>
                    <thead>
                    <tr>
                        <th>Placement</th>
                        <th>Name</th>
                        <th>Score</th>
                    </tr>
                    </thead>
                    <tbody>
                    {leaderboard.map((entry) => (
                        <tr key={entry.rank}>
                            <td>{entry.rank}</td>
                            <td>{entry.name}</td>
                            <td>{entry.clanScore}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
                    )}
            </div>
        </>
    );
};

export default SelectInput;
