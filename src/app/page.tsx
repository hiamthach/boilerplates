import Link from 'next/link';

export default function Home() {
  return (
    <div className="flex flex-col">
      <h1>NextJS Boilerplate</h1>
      <Link href={'/health-check'}>Check</Link>
    </div>
  );
}
