import React, { FunctionComponent, useEffect, useState } from 'react'
import Constants from '../../constants'
import eventService from '../../events/eventService'

const GeneralSettings: FunctionComponent<{}> = () => {


    const [telemetryEnabled, setTelemetryEnabled] = useState<boolean | undefined>(undefined)
    const [logsExpire, setLogsExpire] = useState<number | undefined>(undefined)

    useEffect(() => {
        (async () => {
            let result = await eventService.getSingleSetting(Constants.SETTING_KEYS.TELEMETRY_ENABLED)
            setTelemetryEnabled(result.data)
            result = await eventService.getSingleSetting(Constants.SETTING_KEYS.LOGS_EXPIRE)
            console.log(result.data)
            setLogsExpire(result.data === undefined ? 0 : result.data)
        })()
    }, [])

    const toggleTelemetry = async () => {
        if (telemetryEnabled === undefined) {
            return
        }
        const result = await eventService.updateSingleSetting(Constants.SETTING_KEYS.TELEMETRY_ENABLED, (!telemetryEnabled).toString())
        if (result.success)
            setTelemetryEnabled(!telemetryEnabled)
    }

    const updateLogsExpire = async (days: number) => {
        const result = await eventService.updateSingleSetting(Constants.SETTING_KEYS.LOGS_EXPIRE, days.toString())
        if (result.success)
            setLogsExpire(days)
    }

    return (
        <React.Fragment>
            <h1>General Settings</h1>
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
            <h2>Clear query history</h2>
            <p>Sets the time in days for query history to expire</p>
            {logsExpire !== undefined &&
                <div className="buttons has-addons">
                    <button className={`button is-small${logsExpire === 30 ? ' is-success is-selected' : ''}`} onClick={() => { updateLogsExpire(30) }}>30 days</button>
                    <button className={`button is-small${logsExpire === 60 ? ' is-success is-selected' : ''}`} onClick={() => { updateLogsExpire(60) }}>60 days</button>
                    <button className={`button is-small${logsExpire === 90 ? ' is-success is-selected' : ''}`} onClick={() => { updateLogsExpire(90) }}>90 days</button>
                </div>
            }
            <br />
        </React.Fragment>
    )
}

export default GeneralSettings
