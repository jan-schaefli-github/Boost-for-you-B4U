import { useState } from 'react';
import LocationButton from "./objects/locationButton.tsx";
import TagInput from "./objects/tagInput.tsx";
import LineChart from "./objects/lineChart.tsx";

function Progress() {
    const [selectedLocation, setSelectedLocation] = useState<number>(57000000);
    const [errorMessage, setErrorMessage] = useState<string>('');
    const [selectedChoice, setSelectedChoice] = useState<string>('');

    const handleLocationSelect = (selectedLocation: number) => {
        setSelectedLocation(selectedLocation);
    };

    const handleSearch = (selectedChoice: string) => {
        setSelectedChoice(selectedChoice);
        setErrorMessage('');
    };

    return (
        <>
            <section className="clan-slide" id="part-progress">
                <u>
                    <h1>Progress</h1>
                </u>
                <br />
                <LocationButton onSelectLocation={handleLocationSelect} selectedLocation={selectedLocation} />
                <TagInput onSearch={handleSearch} errorMessage={errorMessage} selectedChoice={selectedChoice}/>
                <LineChart selectedLocation={selectedLocation} selectedChoice={selectedChoice} />
                {errorMessage && <p className="error-message">{errorMessage}</p>}
            </section>
        </>
    );
}

export default Progress;
