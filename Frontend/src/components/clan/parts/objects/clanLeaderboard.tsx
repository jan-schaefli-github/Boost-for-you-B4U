import React, { useState } from 'react';
import '../../../../assets/css/clan/leaderboard.css';
import Leaderboard from "./locationLeaderboard.tsx";
import LocationButton from './locationButton';

const SelectInput: React.FC = () => {
    const [trigger, setTrigger] = useState(true);
    const [selectedLocation, setSelectedLocation] = useState(57000000);

    const handleLocationSelect = (locationId: number) => {
        setSelectedLocation(locationId);
        setTrigger(prevTrigger => !prevTrigger); // Toggle the trigger to force a reload of the Leaderboard
    };

    return (
        <>
            <div className="clan-clans-leaderboards">
                <p className="clan-leaderboard-title">Clans</p>
                <LocationButton onSelectLocation={handleLocationSelect} selectedLocation={selectedLocation} />
                <Leaderboard selectedLocation={selectedLocation} trigger={trigger} setTrigger={setTrigger} />
            </div>
        </>
    );
};

export default SelectInput;
