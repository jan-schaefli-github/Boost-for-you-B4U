import React, { ChangeEvent, KeyboardEvent } from 'react';

interface InputComponentProps {
    selectedChoice: string;
    errorMessage: string;
    handleSelectChange: (event: ChangeEvent<HTMLInputElement>) => void;
    handleKeyDown: (event: KeyboardEvent<HTMLInputElement>) => void;
}

const InputComponent: React.FC<InputComponentProps> = ({
                                                           selectedChoice,
                                                           errorMessage,
                                                           handleSelectChange,
                                                           handleKeyDown,
                                                       }) => (
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
                value={selectedChoice}
            />
            <label htmlFor="clanTag" className="form__label">
                Full Name
            </label>
            {errorMessage && <p className="error-message">{errorMessage}</p>}
        </form>
    </div>
);

export default InputComponent;
