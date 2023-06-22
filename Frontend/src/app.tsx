import {Routes, Route} from 'react-router-dom';
import Header from './components/header.tsx';
import Index from './components/index.tsx'
import ClanMain from "./components/clan/clanMain.tsx";
import MemberMain from './components/member/memberMain.tsx';

export default function App() {
    return (
        <>
            <Header />
            <main>
            <Routes>
                <Route path="/" element={<Index/>} />
                <Route path="/about" element={<Index/>} />
                <Route path="/clan-tracking" element={<ClanMain/>} />
                <Route path="/member-tracking" element={<MemberMain/>} />
            </Routes>
            </main>
        </>
    );
}