import {useEffect, useState} from "react";
import "../../../assets/css/clan/clanGrid.css";

interface Clan {
    tag: string;
    // Add other properties of the clan object if available
}

function CurrentTracking() {
    const [trackedClans, setTrackedClans] = useState<ClanWithNames[]>([]);

    // This useEffect hook will be called only once on component load
    useEffect(() => {
        fetchClans();
    }, []);

    // Function to fetch the initial clan data
    const fetchClans = async () => {
        try {
            const response = await fetch(`${import.meta.env.VITE_BASE_URL}/database/clan`);
            const data = await response.json();
            setTrackedClans(data);
        } catch (error) {
            console.error("Error fetching clans:", error);
        }
    };

    // Render the tracked clans
    return (
        <section className="clan-slide" id="part-tracking">
            <h1 className="clan-tracking-title">Join these already Tracked clans</h1>
            <div className="clan-grid">
                {trackedClans.map((clan) => (
                    <div className="clan-card" key={clan.tag}>
                        <p>{clan.tag}</p>
                        <p>{clan.name}</p>
                    </div>
                ))}
            </div>
        </section>
    );
}

export default CurrentTracking;
