import Link from 'next/link';

import HomeHero from '@/components/features/home/HomeHero';

export default function Home() {
  return (
    <div className="flex flex-col">
      <h1>NextJS Boilerplate</h1>
      <HomeHero />
      <Link href={'/health-check'}>Check</Link>
    </div>
  );
}
