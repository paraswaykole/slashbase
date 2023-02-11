import { ComponentPropsWithoutRef, useState } from "react";
import { Tooltip } from "react-tooltip";

interface InputProps extends ComponentPropsWithoutRef<"input"> {
  label: string;
}

const PasswordInputField = ({ label, className, ...props }: InputProps) => {
  const [showPassword, setShowPassword] = useState<boolean>(false);
  return (
    <>
      <div className="field">
        <label className="label">{label}</label>
        <div className="control has-icons-right">
          <input
            className="input"
            type={showPassword ? "text" : "password"}
            {...props}
          />
          <span
            className="control icon is-clickable is-small is-right"
            onClick={() => setShowPassword((prev) => !prev)}
            id="toggleShowPassword"
            data-tooltip-content={showPassword ? "Hide Password" : "Show Password"}
          >
            <i className={showPassword ? "fas fa-eye-slash" : "fas fa-eye"} />
          </span>
        </div>
      </div>
      <Tooltip anchorId="toggleShowPassword" place="right" variant="dark" />
    </>
  );
};

export default PasswordInputField;
