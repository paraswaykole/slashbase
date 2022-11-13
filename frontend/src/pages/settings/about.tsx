import type { NextPage } from 'next'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../components/layouts/applayout'
import apiService from '../../network/apiService'

const AboutPage: NextPage = () => {

    const [version, setVersion] = useState<string>('')

    useEffect(() => {
        (async () => {
            const result = await apiService.getHealthCheck()
            setVersion(result.version)
        })()
    }, [])

    return (
        <AppLayout title="Settings - About | Slashbase">

            <h1>About Slashbase</h1>
            <br />
            <img src="/logo-icon.svg" width={44} height={50} />
            <h2>Version </h2>
            <p>{version}</p>
            <span className="tinytext">Copyright Slashbase.com.</span><br />
            <span className="tinytext">Licensed under the Apache License 2.0</span>

        </AppLayout>
    )
}

export default AboutPage
