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
                {/* HEADER element is seen on all the pages on the Website, because it's it outside the routing */}
                <Header/>
                <main>
                    {/* ROUTING */}
                    <Routes>
                        {/* INDEX on / */}
                        <Route path={`/`} element={<Index/>}/>
                        {/* ABOUT on /about */}
                        <Route path={`/about`} element={<Index/>}/>
                        {/* CLAN-TRACKING on /clan-tracking */}
                        <Route path={`/clan-tracking`} element={<ClanMain/>}/>
                        {/* SIGNUP on /signup */}
                        <Route path={`/signup`} element={<SignUp/>}/>
                        {/* MEMBER-TRACKING on /member-tracking */}
                        <Route path={`/member-tracking`} element={<MemberMain/>}/>
                        {/* NOT FOUND on All other addresses*/}
                        <Route path={`*`} element={<NotFound/>}/>
                    </Routes>
                </main>
                {/* FOOTER element is seen on all the pages on the Website, because it's it outside the routing */}
                <Footer/>
            </BrowserRouter>
        </>
    );
}