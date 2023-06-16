import {Routes, Route} from 'react-router-dom';
import Header from './components/header.tsx';
import Index from './components/index.tsx'
import ClanMain from "./components/clan/clanMain.tsx";

export default function App() {
    return (
        <>
            <Header />
            <Routes>
                <Route path="/" element={<Index/>} />
                <Route path="/about" element={<Index/>} />
                <Route path="/clan-tracking" element={<ClanMain/>} />
            </Routes>
        </>
    );
}