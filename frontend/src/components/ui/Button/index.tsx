import { ButtonHTMLAttributes, ReactNode, forwardRef } from "react";

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  icon?: ReactNode;
  text?: string;
};

const Button = forwardRef<HTMLButtonElement, ButtonProps>((props, ref) => {
  const { icon, children, text, className, ...rest } = props;
  return (
    <button className={`button ${className ?? ""}`} {...rest} ref={ref}>
      {icon && <span className="icon is-small">{icon}</span>}
      {text && <span>{text}</span>}
      {children}
    </button>
  );
});

export default Button;