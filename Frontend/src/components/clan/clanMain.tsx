import ClanBanner from './parts/clanBanner.tsx';
import ClanTopOfTheBoard from './parts/clanLeaderboards.tsx';
import CurrentTracking from './parts/clanCurrentTracking.tsx';
import Progress from './parts/clanProgress.tsx';
import '../../assets/css/clan/style.css';
import Info from './parts/clanInfo.tsx';

function ClanMain() {
    return (
        <>
            <ClanBanner />
            <main className="clan-main">
                <ClanTopOfTheBoard />
                <CurrentTracking />
                <Progress />
                <Info />
            </main>
        </>
    );
}

export default ClanMain