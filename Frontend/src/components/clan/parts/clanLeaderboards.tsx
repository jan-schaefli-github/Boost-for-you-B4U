import SelectLocationChoices from "./objects/locationChoices.tsx";

function topOfTheBoard() {
    return (
        <>
            <section className="clan-slide" id="part-leaderboards">
                <h1><u>The Top of the board</u></h1>
                <div className="clan-players-leaderboards">
                    <p>Players</p>
                    <input></input>
                </div>
                <div className="clan-clans-leaderboards">
                    <p>Clans</p>
                    <SelectLocationChoices />
                </div>
            </section>
        </>
    );
}

export default topOfTheBoard;