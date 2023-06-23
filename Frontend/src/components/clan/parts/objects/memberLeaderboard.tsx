import React, {ChangeEvent, KeyboardEvent, useEffect, useState} from 'react';
import '../../../../assets/css/clan/leaderboard.css';

interface LeaderboardEntry {
    placement: number;
    name: string;
    points: number;
}

const leaderboardData: LeaderboardEntry[] = [
];

const SelectInput: React.FC = () => {
    const [selectedChoice, setSelectedChoice] = useState<string>('#P9UVQCJV');
    const [errorMessage, setErrorMessage] = useState<string>('');
    const [leaderboard, setLeaderboard] = useState<LeaderboardEntry[]>(leaderboardData);
    const url = 'http://localhost:3000/api/clan/members/leaderboard?clanID=' + encodeURIComponent(selectedChoice);
    const [trigger, setTrigger] = useState(true);

    const handleSelectChange = (event: ChangeEvent<HTMLInputElement>): void => {
        const selectedValue = event.target.value;
        setSelectedChoice(selectedValue);
        setErrorMessage('');
    };

    const handleKeyDown = (event: KeyboardEvent<HTMLInputElement>): void => {
        if (event.key === 'Enter') {
            event.preventDefault();
            if (!selectedChoice.startsWith('#')) {
                setErrorMessage('Input must start with a hashtag (#)');
            } else {
                // Handle sending the valid input
                console.log('Sending input:', selectedChoice);
                fetchData().then(() => {
                    console.log('Fetching Leaderboard Data Done');
                });
            }
        }
    };

    const fetchData = async () => {
        try {
            if (selectedChoice !== '') {
                const response = await fetch(url);
                const placementData = await response.json(); // Access the 'items' array
                if (Array.isArray(placementData) && placementData.length > 0) {
                    //Map the array to the LeaderboardEntry interface
                    const leaderboardEntries: LeaderboardEntry[] = placementData.map((item: any) => {
                        return {
                            placement: item.clanRank,
                            name: item.name + ' (' + item.role + ')',
                            points: item.trophies,
                        };
                    });
                    setLeaderboard(leaderboardEntries);
                } else if (placementData.message === 'Keine Clan-Mitglieder gefunden') {

                    setErrorMessage('Clan not found');
                } else {
                    console.log('Array not found or empty. ' + placementData);
                }
            } else {
                setErrorMessage('Pls give a clan!');
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
    }, []);

    return (
        <div className="clan-clans-leaderboards">
            <div className="selection-min">
                <p className="clan-leaderboard-title">Clan Member Leaderboard</p>
                <form>
                    <input
                        onChange={handleSelectChange}
                        onKeyDown={handleKeyDown}
                        type="text"
                        name="search"
                        placeholder="Please provide a Clan Tag (#XXXXXXX)"
                        className="form__input"
                        id="clanTag"
                    />
                    <label htmlFor="clanTag" className="form__label">Full Name</label>
                    {errorMessage && <p className="error-message">{errorMessage}</p>}
                </form>
            </div>
            <div className="leaderboard">
                <h1>Leaderboard</h1>
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
                        <tr key={entry.placement}>
                            <td>{entry.placement}</td>
                            <td>{entry.name}</td>
                            <td>{entry.points}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default SelectInput;
