import React, { ChangeEvent } from 'react';

interface Location {
    id: number;
    name: string;
}

interface LocationButtonProps {
    choices: Location[];
    selectedChoice: string;
    handleSelectChange: (event: ChangeEvent<HTMLSelectElement>) => void;
}

const LocationButton: React.FC<LocationButtonProps> = ({ choices, selectedChoice, handleSelectChange }) => {
    return (
        <div className="selection-min">
            <p className="clan-leaderboard-title">Clans</p>
            <select onChange={handleSelectChange} value={selectedChoice}>
                {choices.map((choice, index) => {
                    return (
                        <option key={index} value={choice.id}>
                            {choice.name}
                        </option>
                    );
                })}
            </select>
        </div>
    );
};

export default LocationButton;
