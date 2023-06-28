import ClanBanner from './parts/clanBanner.tsx';
import ClanTopOfTheBoard from './parts/clanLeaderboards.tsx';
import CurrentTracking from './parts/clanCurrentTracking.tsx';
import '../../assets/css/clan/style.css';
import Info from './parts/clanInfo.tsx';
import Progress from "./parts/clanProgress.tsx";

function ClanMain() {
    return (
        <>
            <ClanBanner />
            <main className="clan-main">
                <ClanTopOfTheBoard />
                <Progress />
                <Info />
                <CurrentTracking />
            </main>
        </>
    );
}

export default ClanMain