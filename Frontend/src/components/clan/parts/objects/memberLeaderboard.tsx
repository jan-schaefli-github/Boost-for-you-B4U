import React, {useState} from 'react';
import InputComponent from './tagInput.tsx';
import LeaderboardComponent from './tagLeaderboard.tsx';
import '../../../../assets/css/clan/leaderboard.css';

const SelectInput: React.FC = () => {
    const [errorMessage, setErrorMessage] = useState<string>('');
    const [selectedChoice, setSelectedChoice] = useState<string>('#P9UVQCJV');

    const handleSearch = (selectedTag: string) => {
        setSelectedChoice(selectedTag);
        setErrorMessage('');
    };

    return (
        <div className="clan-clans-leaderboards">
            <p className="clan-leaderboard-title">Clan Member Leaderboard</p>
            <InputComponent onSearch={handleSearch} errorMessage={errorMessage}/>
            <LeaderboardComponent selectedTag={selectedChoice}/>
        </div>
    );
};

export default SelectInput;
