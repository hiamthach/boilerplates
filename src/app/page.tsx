import Link from 'next/link';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1>NextJS Boilerplate</h1>
      <Link href={'/health-check'}>Check</Link>
    </main>
  );
}
