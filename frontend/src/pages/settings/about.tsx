import type { NextPage } from 'next'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../components/layouts/applayout'
import Constants from '../../constants'
import apiService from '../../network/apiService'

const AboutPage: NextPage = () => {

    const [version, setVersion] = useState<string>('')
    const [appId, setAppId] = useState<string>('')
    const [telemetryEnabled, setTelemetryEnabled] = useState<boolean | undefined>(undefined)

    useEffect(() => {
        (async () => {
            let result = await apiService.getHealthCheck()
            setVersion(result.version)
            result = await apiService.getSingleSetting(Constants.SETTING_KEYS.APP_ID)
            setAppId(result.data)
            result = await apiService.getSingleSetting(Constants.SETTING_KEYS.TELEMETRY_ENABLED)
            setTelemetryEnabled(result.data)
        })()
    }, [])

    const toggleTelemetry = async () => {
        if (telemetryEnabled === undefined) {
            return
        }
        const result = await apiService.updateSingleSetting(Constants.SETTING_KEYS.TELEMETRY_ENABLED, (!telemetryEnabled).toString())
        if (result.success)
            setTelemetryEnabled(!telemetryEnabled)
    }

    return (
        <AppLayout title="Settings - About | Slashbase">

            <h1>About Slashbase</h1>
            <br />
            <h2>Version </h2>
            <p>{version}</p>
            <br />
            <h2>App ID </h2>
            <p>{appId}</p>
            <br />
            <h2>Telemetry</h2>
            <p>Send anonymous usage data to Slashbase?</p>
            {telemetryEnabled !== undefined &&
                <div className="buttons has-addons">
                    <button className={`button is-small${telemetryEnabled ? ' is-success is-selected' : ''}`} onClick={toggleTelemetry}>Yes</button>
                    <button className={`button is-small${telemetryEnabled ? '' : ' is-danger is-selected'}`} onClick={toggleTelemetry}>No</button>
                </div>
            }
            <br />
            <img src="/logo-icon.svg" width={44} height={50} /><br />
            <span className="tinytext">Copyright Slashbase.com.</span><br />
            <span className="tinytext">Licensed under the Apache License 2.0</span>

        </AppLayout>
    )
}

export default AboutPage
