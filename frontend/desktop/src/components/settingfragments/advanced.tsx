import React, { FunctionComponent, useEffect, useState } from "react";
import OutsideClickHandler from "react-outside-click-handler";
import Constants from "../../constants";
import eventService from "../../events/eventService";

const AdvancedSettings: FunctionComponent<{}> = () => {
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
  ];
  const [openAIKey, setOpenAIKey] = useState<string>("");
  const [openAIModel, setOpenAIModel] = useState<string>("");
  const [isShowingNavDropDown, setIsShowingNavDropDown] = useState<boolean>(false);

  useEffect(() => {
    (async () => {
      let result = await eventService.getSingleSetting(Constants.SETTING_KEYS.OPENAI_KEY);
      setOpenAIKey(result.data);
    })(),
    (async () => {
    let result = await eventService.getSingleSetting(Constants.SETTING_KEYS.OPENAI_MODEL);
    setOpenAIModel(result.data);
    })()
  }, []);
  

  const updateOpenAIKey = async () => {
    const result = await eventService.updateSingleSetting(Constants.SETTING_KEYS.OPENAI_KEY,openAIKey);
    result.error ? setOpenAIKey(openAIKey) : ""
  };
  const updateOpenAIModel = async () => {
    const result = await eventService.updateSingleSetting(Constants.SETTING_KEYS.OPENAI_MODEL,openAIModel);
    result.success? setOpenAIModel(openAIModel) : ""
  };

  return (
    <React.Fragment>
      <h1>Advanced Settings</h1>
      <br />
      <h2>OpenAI Key</h2>
      <p>Update OpenAI API key to enable Generate SQL tool.</p>
      <div className="buttons has-addons">
        <div className="field has-addons">
          <p className="control is-expanded">
            <input className="input" type="text" value={openAIKey}
              onChange={(e: React.ChangeEvent<HTMLInputElement>) => {setOpenAIKey(e.target.value)}} placeholder="Enter API key"/>
          </p>
          <p id="api_key_update" className="control">
            <a className="button" onClick={()=>{updateOpenAIKey()}}>
              <i className="fas fa-check" />
            </a>
          </p>
        </div>
      </div>
      <h2>OpenAI Model</h2>
      <p>Update OpenAI Model to enable Generate SQL tool.</p>
      <div className="buttons has-addons">
        <div className="field has-addons">
          <div className={`dropdown${isShowingNavDropDown ? " is-active" : ""}`}>
            <div className={`dropdown-trigger`}>
              <button className={"button"} aria-haspopup="true" aria-controls="dropdown-menu" onClick={() => {setIsShowingNavDropDown(!isShowingNavDropDown)}}>
                <span>{openAIModel}</span>
                <span className="icon">
                  <i className="fas fa-angle-down" aria-hidden="true"></i>
                </span>
              </button>
            </div>
            {isShowingNavDropDown && (
              <OutsideClickHandler onOutsideClick={() => {setIsShowingNavDropDown(!isShowingNavDropDown)}}>
                <div className="dropdown-menu" id="dropdown-menu" role="menu">
                  <div className="dropdown-content scrollable">
                    {ModelOptions.map((x) => {
                      return (
                        <React.Fragment key={x.value}>
                          <a onClick={() => {setOpenAIModel(x.value),setIsShowingNavDropDown(false)}} className={`dropdown-item${ x.value === openAIModel ? " is-active" : ""}`}>
                            {x.value}
                          </a>
                          {x.value === "home" && (
                            <hr className="dropdown-divider" />
                          )}
                        </React.Fragment>
                      );
                    })}
                  </div>
                </div>
              </OutsideClickHandler>
            )}
          </div>
          <p id="model_update" className="control">
            <a className="button" onClick={()=>updateOpenAIModel()}>
              <i className="fas fa-check" />
            </a>
          </p>
        </div>
      </div>
      <br />
    </React.Fragment>
  );
};

export default AdvancedSettings;
