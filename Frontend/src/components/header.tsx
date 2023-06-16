import {useState, useEffect} from 'react';
import '../assets/css/header.css';

import logo from '../assets/img/favicon_black.png';
import Hamburger from '../assets/img/hamburger.png';

function Header() {
    const [showLinks, setShowLinks] = useState("none");

    function toggleMenu() {
        if (updateShowLinks("") == "none") {
            updateShowLinks("block");
        } else {
            updateShowLinks("none");
        }

    }

    function updateShowLinks(attribute: string) {
        if (attribute === "" || null) {
            return showLinks;
        } else {
            setShowLinks(attribute);
        }
    }


    function handleScreenWidthChange() {
        const screenWidth = window.innerWidth;
        if (screenWidth <= 480) {
            updateShowLinks("none");
        } else {
            updateShowLinks("flex");
        }
    }

    useEffect(() => {

        window.addEventListener('resize', handleScreenWidthChange);
        handleScreenWidthChange();

        return () => {
            window.removeEventListener('resize', handleScreenWidthChange);
        };
    }, []);


    return (
        <>
            <header>
                <div className="logo-name">
                    <a href="/">
                        <img className="logo-img" src={logo} alt="Logo"/>
                    </a>
                    <a href="/" className="company-name">
                        Boost for you
                    </a>
                </div>
                <nav className="menu">
                    <ul className="nav-links" style={{"display": updateShowLinks("")}}>
                        <li>
                            <a href="/">Home</a>
                        </li>
                        <br/>
                        <li>
                            <a href="/member-tracking">Member-Tracking</a>
                        </li>
                        <br/>
                        <li>
                            <a href="/clan-tracking">Clan-Tracking</a>
                        </li>
                        <br/>
                        <li>
                            <a href="/about">About</a>
                        </li>
                        <br/>
                    </ul>
                    <button className="burger-menu" aria-label="Menü anzeigen" onClick={toggleMenu}>
                        <img src={Hamburger} alt="Burger-Menü"/>
                    </button>
                </nav>
            </header>
        </>
    );
}

export default Header;
