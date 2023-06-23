import ClanLeaderboard from "./objects/clanLeaderboard.tsx";
import MemberLeaderboard from "./objects/memberLeaderboard.tsx";

function topOfTheBoard() {
    return (
        <>
            <section className="clan-slide" id="part-leaderboards">
                <h1><u>The Top of the board</u></h1>
                <div className="two-columns" id="clan-leaderboard-container">
                    <MemberLeaderboard />
                    <ClanLeaderboard/>
                </div>
            </section>
        </>
    );
}

export default topOfTheBoard;