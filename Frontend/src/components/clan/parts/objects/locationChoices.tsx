import React, { useState, useEffect, ChangeEvent } from 'react';

interface Location {
    id: number;
    isCountry: boolean;
    name: string;
}

const url = 'http://localhost:3000/api/clan/locations';

const SelectInput: React.FC = () => {
    const [selectedChoice, setSelectedChoice] = useState('');
    const [choices, setChoices] = useState<string[]>([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(url);
                const data = await response.json();
                const locationData = data.items; // Access the 'items' array
                if (Array.isArray(locationData) && locationData.length > 0) {
                    setChoices(locationData.map((location: Location) => location.name));
                } else {
                    console.log('Array not found or empty.');
                }
            } catch (error) {
                console.error(error);
            }
        };

        fetchData().then(function () {
            return console.log('Done');
        });
    }, []);

    const handleSelectChange = (event: ChangeEvent<HTMLSelectElement>) => {
        const selectedValue = event.target.value;
        setSelectedChoice(selectedValue);
    };

    console.log(selectedChoice);

    return (
        <select onChange={handleSelectChange} value={selectedChoice}>
            {choices.map((choice, index) => (
                <option key={index} value={choice}>
                    {choice}
                </option>
            ))}
        </select>
    );
};

export default SelectInput;
