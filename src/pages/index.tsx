import type { NextPage } from 'next'
import AppLayout from '../components/layouts/applayout'

const HomePage: NextPage = () => {
  return (
    <AppLayout title="Home">
      <main>
        <h1>Welcome to Slashbase</h1>
      </main>
    </AppLayout>
  )
}

export default HomePage
