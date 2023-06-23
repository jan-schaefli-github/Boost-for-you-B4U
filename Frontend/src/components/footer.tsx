import "../assets/css/footer.css";

function Footer() {
  return (
    <footer>
      <div className="footer-container">
        <div className="footer-left">
          <p>
            &copy; {new Date().getFullYear()} B4U. All rights reserved.
          </p>
        </div>
        <div className="footer-right">
          <ul className="footer-links">
            <li>
              <a href="/privacy-policy">Privacy Policy</a>
            </li>
            <li>
              <a href="/terms-of-service">Terms of Service</a>
            </li>
            <li>
              <a href="/contact">Contact</a>
            </li>
          </ul>
        </div>
      </div>
    </footer>
  );
}

export default Footer;
