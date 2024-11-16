require('dotenv').config();
const express = require('express');
const { generateTicket } = require('./ticketGenerator');

const app = express();
app.use(express.json());

app.post('/generate-ticket', (req, res) => {
    const { eventName, eventLocation, eventDate, attendeeName } = req.body;

    if (!eventName || !eventLocation || !eventDate || !attendeeName) {
        return res.status(400).json({ error: 'Missing required fields' });
    }

    try {
        const imageBuffer = generateTicket({ eventName, eventLocation, eventDate, attendeeName });
        res.setHeader('Content-Type', 'image/png');
        res.send(imageBuffer);
    } catch (error) {
        console.error('Error generating ticket:', error);
        res.status(500).json({ error: 'Failed to generate ticket image' });
    }
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Ticket generator service running on http://localhost:${PORT}`);
});
