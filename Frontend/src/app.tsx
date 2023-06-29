import {Routes, Route} from 'react-router-dom';
import {BrowserRouter} from "react-router-dom";
import Header from './components/header.tsx';
import Footer from './components/footer.tsx';
import Index from './components/index.tsx'
import ClanMain from "./components/clan/clanMain.tsx";
import MemberMain from './components/member/memberMain.tsx';
import NotFound from './components/notFound.tsx';
import SignUp from './components/clan/parts/clanSignUp.tsx'

export default function App() {
    return (
        <>
        <BrowserRouter basename="/">
            <Header />
            <main>
                <Routes >
                    <Route path={`/`} element={<Index/>} />
                    <Route path={`/about`} element={<Index/>} />
                    <Route path={`/clan-tracking`} element={<ClanMain/>} />
                    <Route path={`/signup`} element={<SignUp/>} />
                    <Route path={`/member-tracking`} element={<MemberMain/>} />
                    <Route path={`*`} element={<NotFound />} />
                </Routes>
            </main>
            <Footer />
            </BrowserRouter>
        </>
    );
}