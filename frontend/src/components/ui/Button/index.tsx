import { ButtonHTMLAttributes, ReactNode, forwardRef } from "react";

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  icon?: ReactNode;
  children?: ReactNode;
};

const Button = forwardRef<HTMLButtonElement, ButtonProps>((props, ref) => {
  const { icon, children, className, ...rest } = props;
  return (
    <button className={`button ${className ?? ""}`} {...rest} ref={ref}>
      {icon && <span className="icon is-small">{icon}</span>}
      {children && <span>{children}</span>}
    </button>
  );
});

export default Button;
