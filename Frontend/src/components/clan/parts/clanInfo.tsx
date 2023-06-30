import "../../../assets/css/clan/info.css";

function Info() {
    return (
        <section className="clan-slide" id="part-info">
            <div className="info-container">
                <div className="info-content">
                    <h1 id="part-info-title">Boost Your Clan with our Tracking</h1>
                    <p className="part-info-text">
                        By signing up for the Tracking service, your clan will be tracked by our system.
                    </p>

                    <p className="part-info-text">
                        After a day of tracking, you can check the member Tracking tab for your clan.
                    </p>
                    <p className="part-info-text">
                        Now you have more insights into the activities of your players, a better overview of your clan and the BOOST you need.
                    </p>
                </div>
                <div className="info-image">
                    <img src="../rocketClashRoyal.avif" alt="Info" />
                </div>
            </div>
        </section>
    );
}

export default Info;
