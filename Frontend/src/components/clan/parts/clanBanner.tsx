import '../../../assets/css/clan/banner.css';

// Component for displaying the clan banner
function ClanBanner() {
    // Function to handle the signup button click
    const handleSignupClick = () => {
        window.location.href = '/signup';
    };

    return (
        <>
            <div className="clan-banner">
                <div className="clan-banner-img"><img src="../Clash-Royale-background.jpg" alt="Clash Royale Banner Image" /></div>
                <section className="primary-wrapper">
                    <h1 className="clan-site-title">HELLO, WELCOME TO THE CLANS SURVEILLANCE BY B4U</h1>
                    <h3 className="clan-site-tagline">Track ANY CLAN BY SIGN UP LIVE AND SIMPLE</h3>
                    <div className="clan-banner-signup">
                        <p className="clan-site-signup-tagline">Your Clan doesn’t get tracked yet sign up here</p>
                        <button className="signup-button-banner" onClick={handleSignupClick}>
                            SignUp Today
                        </button>
                    </div>
                </section>
            </div>
        </>
    );
}

export default ClanBanner;
