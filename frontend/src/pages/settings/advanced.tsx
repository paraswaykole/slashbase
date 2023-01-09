import React, { FunctionComponent, useEffect, useState } from 'react'
import Constants from '../../constants'
import apiService from '../../network/apiService'

const AdvancedSettingsPage: FunctionComponent<{}> = () => {

    const [logsExpire, setLogsExpire] = useState<number | undefined>(undefined)

    useEffect(() => {
        (async () => {
            const result = await apiService.getSingleSetting(Constants.SETTING_KEYS.LOGS_EXPIRE)
            setLogsExpire(result.data === undefined ? 0 : result.data)
        })()
    }, [])

    const updateLogsExpire = async (days: number) => {
        const result = await apiService.updateSingleSetting(Constants.SETTING_KEYS.LOGS_EXPIRE, days.toString())
        if (result.success)
            setLogsExpire(days)
    }

    return (
        <React.Fragment>
            <h1>Advanced Settings</h1>
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

export default AdvancedSettingsPage
