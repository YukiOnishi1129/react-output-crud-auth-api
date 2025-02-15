import { FC, ComponentProps } from "react";

import styles from "./style.module.css";

type TextAreaProps = ComponentProps<"textarea">;

export const TextArea: FC<TextAreaProps> = ({
  disabled = false,
  value,
  placeholder,
  onChange,
}) => (
  <textarea
    disabled={disabled}
    className={styles.text}
    value={value}
    placeholder={placeholder}
    onChange={onChange}
  />
);
