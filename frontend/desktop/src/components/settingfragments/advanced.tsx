import React, { FunctionComponent, useEffect, useState } from 'react'
import Constants from '../../constants'
import eventService from '../../events/eventService'

const ModelOptions = [
  { value: "gpt-4-32k-0314" },
  { value: "gpt-4-32k" },
  { value: "gpt-4-0314" },
  { value: "gpt-4" },
  { value: "gpt-3.5-turbo-0301" },
  { value: "gpt-3.5-turbo" },
  { value: "text-davinci-003" },
  { value: "text-davinci-002" },
  { value: "text-curie-001" },
  { value: "text-babbage-001" },
  { value: "text-ada-001" },
  { value: "text-davinci-001" },
  { value: "davinci-instruct-beta" },
  { value: "davinci" },
  { value: "curie-instruct-beta" },
  { value: "curie" },
  { value: "ada" },
  { value: "babbage" },
]

const AdvancedSettings: FunctionComponent<{}> = () => {
  const [openAIKey, setOpenAIKey] = useState<string>("")
  const [openAIModel, setOpenAIModel] = useState<string>("")

  useEffect(() => {
    (async () => {
      let result = await eventService.getSingleSetting(Constants.SETTING_KEYS.OPENAI_KEY)
      setOpenAIKey(result.data)
    })(),
      (async () => {
        let result = await eventService.getSingleSetting(Constants.SETTING_KEYS.OPENAI_MODEL)
        setOpenAIModel(result.data)
      })()
  }, [])


  const updateOpenAIKey = async () => {
    const result = await eventService.updateSingleSetting(Constants.SETTING_KEYS.OPENAI_KEY, openAIKey)
    if (result.success)
      setOpenAIKey(openAIKey)
  }

  const updateOpenAIModel = async () => {
    const result = await eventService.updateSingleSetting(Constants.SETTING_KEYS.OPENAI_MODEL, openAIModel)
    if (result.success)
      setOpenAIModel(openAIModel)
  }

  const handleModelChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const value = e.target.value;
    setOpenAIModel(value);
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
      <h2>OpenAI Model</h2>
      <p>Update OpenAI Model to enable Generate SQL tool.</p>
      <div className="field has-addons">
        <p className="control">
          <span className="select">
            <select value={openAIModel} onChange={e => handleModelChange(e)}>
              {ModelOptions.map((e, idx) => {
                return <option value={e.value} key={idx}> {e.value} </option>
              })}
            </select>
          </span>
        </p>
        <p className="control">
          <button className="button" onClick={() => updateOpenAIModel()}>
            <i className="fas fa-check" />
          </button>
        </p>
      </div>
      <br />
    </React.Fragment>
  )
}

export default AdvancedSettings
