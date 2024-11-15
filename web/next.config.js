/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'place-hold.it',
      },
      {
        protocol: 'https',
        hostname: '*.ipfs.w3s.link',
      },
    ],
  },
};

module.exports = nextConfig;
