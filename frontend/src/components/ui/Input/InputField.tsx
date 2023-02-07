import { ComponentPropsWithoutRef } from "react";

interface InputProps extends ComponentPropsWithoutRef<"input"> {
  label: string;
}

const InputTextField = ({ label, className, ...props }: InputProps) => {
  return (
    <div className="field">
      <label className="label">{label}</label>
      <div className="control">
        <input className={`input ${className}`} type="text" {...props} />
      </div>
    </div>
  );
};

export default InputTextField;
