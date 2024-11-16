const { createCanvas, registerFont } = require('canvas');
const path = require('path');

// Register Fonts
const fontsPath = path.join(process.cwd(), 'public', 'fonts');
console.log(fontsPath)
registerFont(path.join(fontsPath, 'Space_Grotesk', 'SpaceGrotesk-VariableFont_wght.ttf'), {
    family: 'Space Grotesk',
});
registerFont(path.join(fontsPath, 'Space_Grotesk', 'SpaceGrotesk-Medium.ttf'), {
    family: 'Space Grotesk-Medium',
});
registerFont(path.join(fontsPath, 'Inter', 'Inter-Bold.ttf'), {
    family: 'Inter-Bold',
});
registerFont(path.join(fontsPath, 'Inter', 'Inter-Regular.ttf'), {
    family: 'Inter',
});

// Helper function for text wrapping
function wrapText(ctx, text, maxWidth) {
  const words = text.split(' ');
  let line = '';
  const lines = [];

  for (let i = 0; i < words.length; i++) {
      const testLine = line + words[i] + ' ';
      const metrics = ctx.measureText(testLine);
      const testWidth = metrics.width;

      if (testWidth > maxWidth && i > 0) {
          lines.push(line);
          line = words[i] + ' ';
      } else {
          line = testLine;
      }
  }
  lines.push(line);
  return lines;
}

// List of gradient options
const gradients = [
  ['#B3FFAB', '#12FFF7'], // Aqua to Light Green
  ['#FF9A8B', '#FF6A88'], // Pink to Orange
  ['#A18CD1', '#FBC2EB'], // Purple to Blue
  ['#FDCB82', '#F6A6FF'], // Sunset
  ['#A6FFCB', '#12D8FA'], // Mint to Sky Blue
  ['#00FFA3', '#00CFFF'], // Lime to Teal
];

// Function to randomly select a gradient
function getRandomGradient() {
  const randomIndex = Math.floor(Math.random() * gradients.length);
  return gradients[randomIndex];
}

// Define canvas dimensions
const width = 370;
const height = 550;
const canvas = createCanvas(width, height);
const context = canvas.getContext('2d');

// Main function to generate the ticket
function generateTicket({ eventName, eventLocation, eventDate, attendeeName }) {
  // Background setup
  context.fillStyle = '#FFFFFF';
  context.fillRect(0, 0, width, height);

  // Randomly select a gradient
  const [startColor, endColor] = getRandomGradient();

  // Header gradient
  const gradient = context.createLinearGradient(0, 0, width, 80);
  gradient.addColorStop(0, startColor);
  gradient.addColorStop(1, endColor);
  context.fillStyle = gradient;
  context.fillRect(0, 0, width, 80);

  // "ATTENDEE" text
  const headerHeight = 55.64;
  context.font = '400 26px "Space Grotesk"';
  context.fillStyle = '#0C1527';
  context.textAlign = 'center';
  context.fillText('ATTENDEE', width / 2, headerHeight / 2 + 18);

  // Event Logo Section
  const logoBoxWidth = 300;
  const logoBoxHeight = 70;
  const logoBoxX = (width - logoBoxWidth) / 2;
  const logoBoxY = headerHeight + 50;
  const borderRadius = 10;

  context.fillStyle = '#0C1527';
  context.beginPath();
  context.moveTo(logoBoxX + borderRadius, logoBoxY);
  context.lineTo(logoBoxX + logoBoxWidth - borderRadius, logoBoxY);
  context.quadraticCurveTo(logoBoxX + logoBoxWidth, logoBoxY, logoBoxX + logoBoxWidth, logoBoxY + borderRadius);
  context.lineTo(logoBoxX + logoBoxWidth, logoBoxY + logoBoxHeight - borderRadius);
  context.quadraticCurveTo(logoBoxX + logoBoxWidth, logoBoxY + logoBoxHeight, logoBoxX + logoBoxWidth - borderRadius, logoBoxY + logoBoxHeight);
  context.lineTo(logoBoxX + borderRadius, logoBoxY + logoBoxHeight);
  context.quadraticCurveTo(logoBoxX, logoBoxY + logoBoxHeight, logoBoxX, logoBoxY + logoBoxHeight - borderRadius);
  context.lineTo(logoBoxX, logoBoxY + borderRadius);
  context.quadraticCurveTo(logoBoxX, logoBoxY, logoBoxX + borderRadius, logoBoxY);
  context.closePath();
  context.fill();

  // "EVENT LOGO HERE" text
  context.font = '400 18px "Inter"';
  context.fillStyle = '#FFFFFF';
  context.textAlign = 'center';
  context.fillText('EVENT LOGO HERE', width / 2, logoBoxY + logoBoxHeight / 2);

  // Event Place and Date Section
  const eventTextY = logoBoxY + logoBoxHeight + 40;
  context.font = '500 15px "Space Grotesk-Medium"';
  context.fillStyle = '#0C1527';
  context.fillText(`${eventLocation} • ${eventDate}`, width / 2, eventTextY);

  // Event Title Section
  context.font = '28px "Inter-Bold"';
  const titleMaxWidth = 280;
  const titleY = eventTextY + 50;
  const titleLines = wrapText(context, eventName, titleMaxWidth);

  titleLines.forEach((line, index) => {
      context.fillText(line, width / 2, titleY + index * 34);
  });

  // Footer Gradient
  const footerY = titleY + titleLines.length * 34 + 50;
  const footerGradient = context.createLinearGradient(0, footerY, width, footerY + 150);
  footerGradient.addColorStop(0, startColor);
  footerGradient.addColorStop(1, endColor);
  context.fillStyle = footerGradient;
  context.fillRect(0, footerY, width, height - footerY);

  context.font = '400 18px "Inter"';
  context.fillStyle = '#0C1527';
  context.fillText(`[ ${attendeeName} ]`, width / 2, footerY + 50);

  context.font = '400 14px "Inter"';
  context.fillText('EVENT.COM', width / 2, footerY + 80);

  return canvas.toBuffer('image/png');
}

module.exports = { generateTicket };
