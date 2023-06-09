import {useState, useEffect} from 'react';
import '../assets/css/header.css';

// Images
import LLogo from '../assets/img/favicon_black.ico';
import DLogo from '../assets/img/favicon_white.ico';
import Hamburger from '../assets/img/hamburger.png';

function darkModeLogo(): string {
    let favicon;
    // Dark Mode / Light Mode
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        // Dark Mode
        favicon = DLogo;
        console.log('Dark Mode');
    } else {
        // Light Mode
        favicon = LLogo;
        console.log('Light Mode');
    }
    return favicon;
}

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
        // Cleanup the event listener when the component unmounts
        return () => {
            window.removeEventListener('resize', handleScreenWidthChange);
        };
    }, []);

    return (
        <>
            <header>
                <div className="logo-name">
                    <a href="/">
                        <img className="logo-img" src={darkModeLogo()} alt="Logo"/>
                    </a>
                    <a href="/" className="company-name">
                        Boost for you
                    </a>
                </div>
                <nav className="menu">
                    <ul className="nav-links" style={{"display": updateShowLinks("")}}>
                        <li>
                            <a href="#">Link 1</a>
                        </li>
                        <br/>
                        <li>
                            <a href="#">Link 2</a>
                        </li>
                        <br/>
                        <li>
                            <a href="#">Link 3</a>
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
