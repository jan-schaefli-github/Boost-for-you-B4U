import React, {useState, useEffect, ChangeEvent} from 'react';
import '../../../../assets/css/clan/leaderboard.css'

interface Location {
    id: number;
    name: string;
}

const url = 'http://localhost:3000/api/clan/locations';

const SelectInput: React.FC = () => {
    const [selectedChoice, setSelectedChoice] = useState('');
    const [choices, setChoices] = useState<Location[]>([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch(url);
                const data = await response.json();
                const locationData = data.items; // Access the 'items' array
                if (Array.isArray(locationData) && locationData.length > 0) {
                    setChoices(locationData.map((location: Location) => ({id: location.id, name: location.name})));
                } else {
                    console.log('Array not found or empty.');
                }
            } catch (error) {
                console.error(error);
            }
        };

        fetchData().then(() => {
            console.log('Fetching Leaderboard Data Done');
        });
    }, []);

    const handleSelectChange = (event: ChangeEvent<HTMLSelectElement>) => {
        const selectedValue = event.target.value;
        setSelectedChoice(selectedValue);

        // Find the selected location based on the selected ID
        const selectedLocation = choices.find((location) => location.id === parseInt(selectedValue, 10));
        if (selectedLocation) {
            console.log('Selected Location ID for Leaderboard:', selectedLocation.id);
        }
    };

    return (
        <>
            <select onChange={handleSelectChange} value={selectedChoice}>
                {choices.map((choice, index) => {
                    return (
                        <option key={index} value={choice.id}>
                            {choice.name}
                        </option>
                    );
                })}
            </select>
            <div className="leaderboard">
                <h1>
                    <svg className="ico-cup">
                        <use xlinkHref="#cup"></use>
                    </svg>
                    Most active Players
                </h1>
                <ol>
                    <li>
                        <mark>Jan Schääääääääääääääääääääääfli</mark>
                        <small>0 = noob</small>
                    </li>
                    <li>
                        <mark>Brandon Barnes</mark>
                        <small>301</small>
                    </li>
                    <li>
                        <mark>Raymond Knight</mark>
                        <small>292</small>
                    </li>
                    <li>
                        <mark>Trevor McCormick</mark>
                        <small>245</small>
                    </li>
                    <li>
                        <mark>Andrew Fox</mark>
                        <small>203</small>
                    </li>
                </ol>
            </div>
        </>
    )
};

export default SelectInput;
