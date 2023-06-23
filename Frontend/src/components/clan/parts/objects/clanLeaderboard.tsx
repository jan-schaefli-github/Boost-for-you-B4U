import React, { useState, useEffect, ChangeEvent } from 'react';
import '../../../../assets/css/clan/leaderboard.css';
import Leaderboard from "./locationLeaderboard.tsx";

interface Location {
    id: number;
    name: string;
}

const url = 'http://localhost:3000/api/clan/locations';

const SelectInput: React.FC = () => {
    const [selectedChoice, setSelectedChoice] = useState('');
    const [choices, setChoices] = useState<Location[]>([]);
    const [trigger, setTrigger] = useState(true);
    const [selectedLocation, setSelectedLocation] = useState(57000000);

    const handleSelectChange = (event: ChangeEvent<HTMLSelectElement>): void | number => {
        const selectedValue = event.target.value;
        setSelectedChoice(selectedValue);
        setTrigger(true);
        // Find the selected location based on the selected ID
        const selectedLocation = choices.find((location) => location.id === parseInt(selectedValue, 10));
        if (selectedLocation) {
            console.log('Selected Location ID for Leaderboard:', selectedLocation.id);
            setSelectedLocation(selectedLocation.id);
        } else {
            return 0;
        }
    };

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(url);
                const data = await response.json();
                const locationData = data.items; // Access the 'items' array
                if (Array.isArray(locationData) && locationData.length > 0) {
                    setChoices(locationData.map((location: Location) => ({ id: location.id, name: location.name })));
                } else {
                    console.log('Array not found or empty. ' + locationData);
                }
            } catch (error) {
                console.error(error);
            }
        };

        fetchData().then(() => {
            console.log('Fetching Location Data Done');
        });
    }, []);

    return (
        <>
            <div className="clan-clans-leaderboards">
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
                <Leaderboard selectedLocation={selectedLocation}  trigger={trigger} setTrigger={setTrigger}/>
            </div>
        </>
    );
};

export default SelectInput;
