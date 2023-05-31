const express = require('express');
const axios = require('axios');
require('dotenv').config();

const app = express();
const port = process.env.PORT || 3000;

const accessToken = process.env.ACCESS_TOKEN;
const clanTag = process.env.CLAN_TAG;
const encodedClanTag = encodeURIComponent(clanTag);
const headers = { Authorization: `Bearer ${accessToken}`, 'Access-Control-Allow-Origin': '*' };

app.get('/clan', async (req, res) => {
    try {
        const url = 'https://api.clashroyale.com/v1/clans/' + encodedClanTag;
        const response = await axios.get(url, { headers });
        res.json(response.data);
    } catch (error) { 
        res.status(500).json({ error: 'An error occurred' });
    }
});

app.get('/members', async (req, res) => {
    try {
        const url = 'https://api.clashroyale.com/v1/clans/' + encodedClanTag + '/members';
        const response = await axios.get(url, { headers });
        res.json(response.data);
    } catch (error) { 
        res.status(500).json({ error: 'An error occurred' });
    }
});

app.get('/currentriverrace', async (req, res) => {
    try {
        const url = 'https://api.clashroyale.com/v1/clans/' + encodedClanTag + '/currentriverrace';
        const response = await axios.get(url, { headers });
        res.json(response.data);
    } catch (error) { 
        res.status(500).json({ error: 'An error occurred' });
    }
});

app.get('/riverracelog', async (req, res) => {
    try {
        const url = 'https://api.clashroyale.com/v1/clans/' + encodedClanTag + '/riverracelog';
        const response = await axios.get(url, { headers });
        res.json(response.data);
    } catch (error) { 
        res.status(500).json({ error: 'An error occurred' });
    }
});

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
  });