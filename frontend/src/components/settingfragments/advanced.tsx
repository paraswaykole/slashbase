import React, { FunctionComponent, useEffect, useState } from 'react'
import Constants from '../../constants'
import apiService from '../../network/apiService'
import toast from 'react-hot-toast'

const AdvancedSettings: FunctionComponent<{}> = () => {

  const [openAIKey, setOpenAIKey] = useState<string>("")
  const [openAIModel, setOpenAIModel] = useState<string>("")
  const [modelOptions, setModelOptions] = useState<{ value: string }[]>([])

  useEffect(() => {
    (async () => {
      const result = await apiService.listSupportedAIModels()
      setModelOptions(result.data.map(model => ({ value: model })))
    })();
    (async () => {
      const result = await apiService.getSingleSetting(Constants.SETTING_KEYS.OPENAI_KEY)
      setOpenAIKey(result.data)
    })();
    (async () => {
      const result = await apiService.getSingleSetting(Constants.SETTING_KEYS.OPENAI_MODEL)
      setOpenAIModel(result.data)
    })();
  }, [])


  const updateOpenAIKey = async () => {
    const result = await apiService.updateSingleSetting(Constants.SETTING_KEYS.OPENAI_KEY, openAIKey)
    if (result.success) {
      setOpenAIKey(openAIKey)
      toast.success("saved")
    }
  }

  const updateOpenAIModel = async () => {
    const result = await apiService.updateSingleSetting(Constants.SETTING_KEYS.OPENAI_MODEL, openAIModel)
    if (result.success)
      setOpenAIModel(openAIModel)
    toast.success("saved")
  }

  const handleModelChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const value = e.target.value
    setOpenAIModel(value)
  }

  return (
    <React.Fragment>
      <h1>Advanced Settings</h1>
      <br />
      <h2>OpenAI Key</h2>
      <p>Update OpenAI API key to enable Generate SQL tool.</p>
      <div className="buttons has-addons">
        <div className="field has-addons" style={{ minWidth: 550 }}>
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
              {modelOptions.map((e, idx) => {
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
