import React, { useState, useEffect, ChangeEvent } from 'react';

interface Location {
    id: number;
    name: string;
}

const url = 'http://localhost:3000/api/clan/locations';

interface LocationButtonProps {
    onSelectLocation: (locationId: number) => void;
    selectedLocation: number;
}

const LocationButton: React.FC<LocationButtonProps> = ({ onSelectLocation, selectedLocation }) => {
    const [selectedChoice, setSelectedChoice] = useState('');
    const [choices, setChoices] = useState<Location[]>([]);

    const handleSelectChange = (event: ChangeEvent<HTMLSelectElement>): void | number => {
        const selectedValue = event.target.value;
        setSelectedChoice(selectedValue);
        // Find the selected location based on the selected ID
        const selectedLocation = choices.find((location) => location.id === parseInt(selectedValue, 10));
        if (selectedLocation) {
            console.log('Selected Location ID:', selectedLocation.id);
            onSelectLocation(selectedLocation.id); // Call the callback function with the selected location ID
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
            console.log('Fetching locationButton Data Done');
        });
    }, [selectedLocation]); // Reload the component when selectedLocation changes

    return (
        <div className="selection-container">
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
