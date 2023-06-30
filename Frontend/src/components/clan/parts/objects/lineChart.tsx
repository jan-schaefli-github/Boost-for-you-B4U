import React, {useEffect, useState} from 'react';
import {Line} from 'react-chartjs-2';
import {CategoryScale, Chart, Legend, LinearScale, LineElement, PointElement, Title} from 'chart.js';

Chart.register(Title, Legend, CategoryScale, LinearScale, PointElement, LineElement);

interface LineChartProps {
    selectedLocation: number;
    selectedChoice: string;
}

interface LocationDataItem {
    clanName: string;
    clanFameHistory: { fame: number; week: number }[];
}

const LineChart: React.FC<LineChartProps> = ({ selectedLocation, selectedChoice }) => {
    const encodedSelectedChoice = encodeURIComponent(selectedChoice);
    const url = `${import.meta.env.VITE_BASE_URL}/api/clan/riverracelog/linechart?clanTag=${encodedSelectedChoice}&locationID=${selectedLocation}`;
    const [userData, setUserData] = useState<{
        labels: number[];
        datasets: { label: string; data: number[]; fill: boolean; borderColor: string; borderWidth: number; }[];
    }>({labels: [], datasets: []});
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [darkMode, setDarkMode] = useState<boolean>(false);

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(url.toString());
                if (!response.ok) {
                    console.error('Network response was not ok');
                    return;
                }
                const data = await response.json();
                const locationData: LocationDataItem[][] = data.linechartRiverRaceLog;

                if (Array.isArray(locationData) && locationData.length > 0) {
                    const weeks = locationData[0][0].clanFameHistory.map((historyItem) => historyItem.week).reverse(); // Reverse the weeks array
                    const datasets = locationData.map((item) => {
                        const clanName = item[0].clanName;
                        const values = item[0].clanFameHistory.map((historyItem) => historyItem.fame);
                        return {
                            label: clanName,
                            data: values,
                            fill: false,
                            borderColor: getRandomColor(darkMode),
                            borderWidth: 3,
                        };
                    });

                    setUserData({
                        labels: weeks,
                        datasets: datasets,
                    });
                } else {
                    console.log('Array not found or empty:', locationData);
                }
            } catch (error) {
                console.error('Error during data fetch:', error);
            } finally {
                setIsLoading(false);
            }
        };

        fetchUserData().then(() => {
            console.log('Fetching LineChart Data Done');
        });
    }, [selectedLocation, selectedChoice, url, darkMode]);

    useEffect(() => {
        // Detect browser preference for dark mode
        const matchMediaDark = window.matchMedia('(prefers-color-scheme: dark)');
        const handleDarkModeChange = (event: MediaQueryListEvent) => {
            setDarkMode(event.matches);
        };
        matchMediaDark.addEventListener('change', handleDarkModeChange);
        setDarkMode(matchMediaDark.matches);
        return () => {
            matchMediaDark.removeEventListener('change', handleDarkModeChange);
        };
    }, []);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (userData === null) {
        return <div>No data available</div>;
    }

    return (
        <div className="lineChart">
            <Line
                data={userData}
                options={{
                    maintainAspectRatio: false,
                    responsive: true,
                    plugins: {
                        legend: {
                            position: 'top',
                            labels: {color: darkMode ? 'white' : 'black'},
                        },
                        title: {
                            display: true,
                            text: 'Comparison of Score over a 10 Week Time',
                            color: darkMode ? 'white' : 'black',
                            font: {size: 20},
                        },
                    },
                    scales: {
                        x: {
                            title: {
                                display: true,
                                text: 'Week',
                                color: darkMode ? 'white' : 'black',
                            },
                            grid: {color: darkMode ? 'white' : 'black'},
                        },
                        y: {
                            title: {
                                display: true,
                                text: 'Fame',
                                color: darkMode ? 'white' : 'black',
                            },
                            grid: {color: darkMode ? 'white' : 'black'},
                            suggestedMin: 0,
                            ticks: {
                                color: darkMode ? 'white' : 'black',
                            },
                        },
                    },
                    interaction: {
                        mode: 'index',
                    },
                    hover: {
                        mode: 'index',
                        intersect: true,
                    },
                }}
            />
        </div>
    );

    // Helper function to generate random colors with controlled brightness
    function getRandomColor(darkMode: boolean) {
        const letters = '0123456789ABCDEF';
        let color = '#';

        for (let i = 0; i < 6; i++) {
            const index = Math.floor(Math.random() * 16);
            const letter = letters[index];
            color += letter;
        }

        // Adjust brightness for light mode
        if (!darkMode) {
            const brightnessOffset = 30; // Increase this value to make the colors lighter
            return color
                .replace(/^#(\w{2})(\w{2})(\w{2})$/, (_, r, g, b) => {
                    const adjustedR = Math.min(parseInt(r, 16) + brightnessOffset, 255);
                    const adjustedG = Math.min(parseInt(g, 16) + brightnessOffset, 255);
                    const adjustedB = Math.min(parseInt(b, 16) + brightnessOffset, 255);
                    return `#${adjustedR.toString(16)}${adjustedG.toString(16)}${adjustedB.toString(16)}`;
                });
        }

        return color;
    }
};

export default LineChart;
