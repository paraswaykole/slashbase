import React, { FunctionComponent, useEffect, useState } from 'react'
import Constants from '../../constants'
import eventService from '../../events/eventService'

const AdvancedSettings: FunctionComponent<{}> = () => {


    const [openAIKey, setOpenAIKey] = useState<string>("")

    useEffect(() => {
        (async () => {
            let result = await eventService.getSingleSetting(Constants.SETTING_KEYS.OPENAI_KEY)
            setOpenAIKey(result.data)
        })()
    }, [])

    const updateOpenAIKey = async () => {
        const result = await eventService.updateSingleSetting(Constants.SETTING_KEYS.OPENAI_KEY, openAIKey)
        if (result.success)
            setOpenAIKey(openAIKey)
    }

    return (
        <React.Fragment>
            <h1>Advanced Settings</h1>
            <br />
            <h2>OpenAI Key</h2>
            <p>Update OpenAI API key to enable Generate SQL tool.</p>
            <div className="buttons has-addons">
                <div className="field has-addons">
                    <p className="control is-expanded">
                        <input
                            className="input"
                            type="text"
                            value={openAIKey}
                            onChange={(e: React.ChangeEvent<HTMLInputElement>) => { setOpenAIKey(e.target.value) }}
                            placeholder="Enter API key" />
                    </p>
                    <p className="control">
                        <a className="button" onClick={updateOpenAIKey}>
                            <i className="fas fa-check" />
                        </a>
                    </p>
                </div>
            </div>
            <br />
        </React.Fragment>
    )
}

export default AdvancedSettings
