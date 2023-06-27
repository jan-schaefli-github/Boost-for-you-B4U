import React from 'react';

interface LeaderboardEntry {
    placement: number;
    name: string;
    points: number;
}

interface LeaderboardComponentProps {
    leaderboard: LeaderboardEntry[];
}

const LeaderboardComponent: React.FC<LeaderboardComponentProps> = ({ leaderboard }) => (
    <div className="leaderboard">
        <h1>Leaderboard</h1>
        <table>
            <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Trophies</th>
            </tr>
            </thead>
            <tbody>
            {leaderboard.map((entry) => (
                <tr key={entry.placement}>
                    <td>{entry.placement}</td>
                    <td>{entry.name}</td>
                    <td>{entry.points}</td>
                </tr>
            ))}
            </tbody>
        </table>
    </div>
);

export default LeaderboardComponent;
