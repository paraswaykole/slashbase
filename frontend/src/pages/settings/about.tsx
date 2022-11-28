import type { NextPage } from 'next'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../components/layouts/applayout'
import Constants from '../../constants'
import apiService from '../../network/apiService'

const AboutPage: NextPage = () => {

    const [version, setVersion] = useState<string>('')
    const [appId, setAppId] = useState<string>('')

    useEffect(() => {
        (async () => {
            let result = await apiService.getHealthCheck()
            setVersion(result.version)
            result = await apiService.getSingleSetting(Constants.SETTING_KEYS.APP_ID)
            setAppId(result.data)
        })()
    }, [])

    return (
        <AppLayout title="Settings - About | Slashbase">

            <h1>About Slashbase</h1>
            <br />
            <h2>Version </h2>
            <p>{version}</p>
            <br />
            <h2>App ID </h2>
            <p>{appId}</p>
            <hr />
            <img src="/logo-icon.svg" width={44} height={50} /><br />
            <span className="tinytext">Copyright Slashbase.com.</span><br />
            <span className="tinytext">Licensed under the Apache License 2.0</span>

        </AppLayout>
    )
}

export default AboutPage
