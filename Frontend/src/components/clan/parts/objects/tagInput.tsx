import React, { ChangeEvent, KeyboardEvent, useState } from 'react';

interface TagInputProps {
    selectedChoice: string;
    errorMessage: string;
    onSearch: (selectedChoice: string) => void;
}

const TagInput: React.FC<TagInputProps> = ({ selectedChoice, errorMessage, onSearch }) => {
    const [inputValue, setInputValue] = useState(selectedChoice);

    const handleSelectChange = (event: ChangeEvent<HTMLInputElement>) => {
        setInputValue(event.target.value);
    };

    const handleKeyDown = (event: KeyboardEvent<HTMLInputElement>) => {
        if (event.key === 'Enter') {
            event.preventDefault();
            onSearch(inputValue);
        }
    };

    return (
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
                    value={inputValue}
                />
                <label htmlFor="clanTag" className="form__label">
                    Full Name
                </label>
                {errorMessage && <p className="error-message">{errorMessage}</p>}
            </form>
        </div>
    );
};

export default TagInput;
