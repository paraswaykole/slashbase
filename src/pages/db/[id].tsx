import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../components/layouts/applayout'
import { DBConnection } from '../../data/models'
import apiService from '../../network/apiService'

const DBPage: NextPage = () => {

  const router = useRouter()

  return (
    <AppLayout title="Home">
      <main className="maincontainer">
        <h1>Connected to DB</h1>
      </main>
    </AppLayout>
  )
}

export default DBPage
