import { FC, ComponentProps, ReactNode } from "react";
import styles from "./style.module.css";

type CommonButtonProps = ComponentProps<"button"> & {
  children: ReactNode;
};

export const CommonButton: FC<CommonButtonProps> = ({
  type,
  children,
  onClick,
}) => (
  <button className={styles.button} type={type} onClick={onClick}>
    {children}
  </button>
);
