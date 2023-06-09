import {Routes, Route} from 'react-router-dom';
import Header from './components/header.tsx';
import Index from './components/index.tsx'

export default function App() {
    return (
        <>
            <Header />
            <Routes>
                <Route path="/" element={<Index/>} />
            </Routes>
        </>
    );
}