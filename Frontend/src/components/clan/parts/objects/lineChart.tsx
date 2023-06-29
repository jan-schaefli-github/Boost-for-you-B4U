import React, { useEffect, useState } from 'react';
import { Line } from 'react-chartjs-2';
import { Chart, Title, Legend, CategoryScale, LinearScale, PointElement, LineElement } from 'chart.js';

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
    const url = `http://localhost:3000/api/clan/riverracelog/linechart?clanTag=${encodedSelectedChoice}&locationID=${selectedLocation}`;
    const [userData, setUserData] = useState<{
        labels: number[];
        datasets: { label: string; data: number[]; fill: boolean; borderColor: string; borderWidth: number; }[];
    }>({ labels: [], datasets: [] });
    const [isLoading, setIsLoading] = useState<boolean>(true);

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(url);
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
                            borderColor: getRandomColor(),
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
    }, [selectedLocation, selectedChoice, url]);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    if (userData === null) {
        return <div>No data available</div>;
    }
    return (
        <div style={{ marginLeft: 'auto', marginRight: 'auto', width: '80vw' }}>
            <Line
                data={userData}
                options={{
                    responsive: true,
                    plugins: {
                        legend: {
                            position: 'top',
                        },
                        title: {
                            display: true,
                            text: 'Comparison of Score over a 10 Week Time',
                        },
                    },
                    scales: {
                        x: {
                            title: {
                                display: true,
                                text: 'Week',
                            },
                        },
                        y: {
                            title: {
                                display: true,
                                text: 'Fame',
                            },
                            suggestedMin: 0,
                        },
                    },
                    interaction: {
                        mode: 'index', // Display intersecting items for all data points
                    },
                    hover: {
                        mode: 'index', // Enable hover mode for displaying intersecting items
                        intersect: true, // Enable hover interaction only when the cursor intersects the item
                    },
                }}
            />
        </div>
    );

    // Helper function to generate random colors
    function getRandomColor() {
        const letters = '89ABCDEF'; // Remove 0-7 to increase brightness
        let color = '#';
        for (let i = 0; i < 6; i++) {
            color += letters[Math.floor(Math.random() * letters.length)];
        }
        return color;
    }
};

export default LineChart;
