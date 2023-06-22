import '../../../assets/css/clan/banner.css'

function ClanBanner() {
    return (
        <>
            <div className="clan-banner">
                <div className="clan-banner-img"></div>
                <section className="primary-wrapper">
                    <h1 className="clan-site-title">HELLO, WELCOME TO THE CLANS Surveillance BY B4u</h1>
                    <h3 className="clan-site-tagline">Track ANY CLAN BY SIGN UP LIVE AND SIMPLE</h3>
                    <div className="clan-banner-signup">
                        <p className="clan-site-signup-tagline">You Clan doesn’t get tracked yet sign up here</p>
                        <button className="signup-button-banner">SignUp Today</button>
                    </div>
                </section>
            </div>
        </>
    );
}

export default ClanBanner;