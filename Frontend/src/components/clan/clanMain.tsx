import ClanBanner from './parts/clanBanner.tsx';
import ClanTopOfTheBoard from './parts/clanLeaderboards.tsx';
import CurrentTracking from './parts/clanCurrentTracking.tsx';
import '../../assets/css/clan/style.css';
import Info from './parts/clanInfo.tsx';
import Progress from "./parts/clanProgress.tsx";

function ClanMain() {
    return (
        <>
            {/* BANNER of the site - displays Welcome Text, a SignUp button and a background Image */}
            <ClanBanner />
            <main className="clan-main">
                {/*Leaderboards Section - Clan and clan member Leaderboard*/}
                <ClanTopOfTheBoard />
                {/*Line Chart Section - Displaying a comparison line chart of the progress of top clans and costume inserted clans*/}
                <Progress />
                {/*Info Section - Displaying Info about what the Tracking is.*/}
                <Info />
                {/*Current Tracking Section - displaying all currently listed and tracked clans*/}
                <CurrentTracking />
            </main>
        </>
    );
}

export default ClanMain