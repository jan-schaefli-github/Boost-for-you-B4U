import SelectLocationChoices from "./objects/locationChoices.tsx";

function topOfTheBoard() {
    return (
        <>
            <section className="clan-slide" id="part-leaderboards">
                <h1><u>The Top of the board</u></h1>
                <div className="two-columns" id="clan-leaderboard-container">
                    <div className="clan-players-leaderboards">
                        <p className="clan-leaderboard-title">Players</p>
                        <input></input>
                    </div>
                    <div className="clan-clans-leaderboards">
                        <p className="clan-leaderboard-title">Clans</p>
                        <SelectLocationChoices/>
                    </div>
                </div>
            </section>
        </>
    );
}

export default topOfTheBoard;