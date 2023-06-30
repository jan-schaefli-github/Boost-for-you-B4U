import React, { useState, useEffect } from "react";
import "../../../assets/css/clan/signup.css";

function RegisterForm() {
    const [clanTag, setClanTag] = useState("");
    const [errorMessage, setErrorMessage] = useState("");
    const [successMessage, setSuccessMessage] = useState("");

    useEffect(() => {
        const urlParams = new URLSearchParams(window.location.search);
        const tag = urlParams.get("clanTag");
        if (tag) {
            setClanTag(tag);
        }
    }, []);

    const handleSubmit = (event: React.FormEvent) => {
        event.preventDefault();

        // Perform validation and submit the form to the tag endpoint
        // using the provided clanTag value
        // You can use fetch or any other library to make the HTTP request

        // Example using fetch
        fetch(
            `${import.meta.env.VITE_BASE_URL}/database/clan/create?clanTag=${encodeURIComponent(
                clanTag
            )}`
        )
            .then((response) => {
                if (response.ok) {
                    return response.json();
                } else {
                    throw new Error("Failed to store clanTag in YAML file");
                }
            })
            .then((data) => {
                console.log(data); // Handle the response from the server
                setSuccessMessage("Clan tag stored successfully!");
                setErrorMessage("");
            })
            .catch((data) => {
                const displayError = data;
                console.log(displayError)
                const errorMessage = `${displayError}`;
                setErrorMessage(errorMessage); // Set error message
                setSuccessMessage("");
            });
    };

    return (
        <>
            <section className="clan-slide" id="part-signup">
                <h1>Signup form</h1>
                <form onSubmit={handleSubmit}>
                    <div className="form-group">
                        <label htmlFor="clanTag">Clan Tag:</label>
                        <input
                            type="text"
                            id="clanTag"
                            value={clanTag}
                            onChange={(event: React.ChangeEvent<HTMLInputElement>) =>
                                setClanTag(event.target.value)
                            }
                            required
                        />
                    </div>
                    <div className="feedback">
                        {successMessage && (
                            <div className="success-message">{successMessage}</div>
                        )}
                        {errorMessage && (
                            <div className="error-message">{errorMessage}</div>
                        )}
                    </div>
                    <button className="submit-button" type="submit">
                        Submit
                    </button>
                </form>
            </section>
        </>
    );
}

export default RegisterForm;
