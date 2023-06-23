import React, {useEffect, useState} from 'react';
import '../../../../assets/css/clan/leaderboard.css';

// Leaderboard Information Interface
interface LeaderboardEntry {
    placement: number;
    name: string;
    points: number;
}

const leaderboardData: LeaderboardEntry[] = [
    {placement: 1, name: 'Select a Country', points: 1000},
    {placement: 2, name: 'Player 2', points: 90},
    {placement: 3, name: 'Player 3', points: 80},
    // Add more entries as needed
];

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
    const [leaderboard, setLeaderboard] = useState<LeaderboardEntry[]>(leaderboardData);
    const url = 'http://localhost:3000/api/clan/leaderboard?locationID=' + selectedLocation;

    const fetchData = async () => {
        try {
            const response = await fetch(url);
            const data = await response.json();
            const placementData = data.items; // Access the 'items' array
            if (Array.isArray(placementData) && placementData.length > 0) {
                //Map the array to the LeaderboardEntry interface
                const leaderboardEntries: LeaderboardEntry[] = placementData.map((item: any) => {
                    return {
                        placement: item.rank,
                        name: item.name,
                        points: item.clanScore,
                    };
                });
                setLeaderboard(leaderboardEntries);
            } else {
                console.log('Array not found or empty. ' + placementData);
            }
        } catch (error) {
            console.error(error);
        }
    };

    useEffect(() => {
        // Check if the trigger prop has changed
        if (trigger) {
            setTrigger(false);
            fetchData().then(() => {
                console.log('Fetching Leaderboard Data Done');
            });
        }
    }, [trigger]);

    return (
        <>
            <div className="leaderboard">
                <h1>Leaderboard</h1>
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
                        <tr key={entry.placement}>
                            <td>{entry.placement}</td>
                            <td>{entry.name}</td>
                            <td>{entry.points}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>
        </>
    );
};

export default SelectInput;
