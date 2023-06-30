import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import '../assets/css/header.css';

import logo from '../assets/img/favicon_black.png';
import Hamburger from '../assets/img/hamburger.png';

function Header() {
  const [showLinks, setShowLinks] = useState('none');

  function toggleMenu() {
    if (showLinks === 'none') {
      updateShowLinks('block');
    } else {
      updateShowLinks('none');
    }
  }

  function updateShowLinks(attribute: string) {
    if (attribute === '' || attribute === null) {
      return showLinks;
    } else {
      setShowLinks(attribute);
    }
  }

  function handleScreenWidthChange() {
    const screenWidth = window.innerWidth;
    if (screenWidth <= 712) {
      updateShowLinks('none');
    } else {
      updateShowLinks('flex');
    }
  }

  function handleLinkClick() {
    toggleMenu(); // Collapse the menu after a link is clicked
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
            <img className="logo-img" src={logo} alt="Logo" />
          </a>
          <a href="/" className="company-name">
            Boost for you
          </a>
        </div>
        <nav className="menu">
          <ul className="nav-links" style={{ display: showLinks }}>
            <li>
              <Link to={`/`} onClick={handleLinkClick}>
                Home
              </Link>
            </li>
            <br />
            <li>
              <Link to={`/member-tracking`} onClick={handleLinkClick}>
                Member-Tracking
              </Link>
            </li>
            <br />
            <li>
              <Link to={`/clan-tracking`} onClick={handleLinkClick}>
                Clan-Tracking
              </Link>
            </li>
            <br />
            <li>
              <Link to={`/about`} onClick={handleLinkClick}>
                About
              </Link>
            </li>
            <br />
          </ul>
          <button
            className="burger-menu"
            aria-label="Menü anzeigen"
            onClick={toggleMenu}
          >
            <img src={Hamburger} alt="Burger-Menü" />
          </button>
        </nav>
      </header>
    </>
  );
}

export default Header;
