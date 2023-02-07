import { ComponentPropsWithoutRef, useState } from "react";

interface InputProps extends ComponentPropsWithoutRef<"input"> {
  label: string;
}

const PasswordInputField = ({ label, className, ...props }: InputProps) => {
  const [showPassword, setShowPassword] = useState<boolean>(false);
  return (
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
        >
          <i className={showPassword ? "fas fa-eye" : "fas fa-eye-slash"} />
        </span>
      </div>
    </div>
  );
};

export default PasswordInputField;
