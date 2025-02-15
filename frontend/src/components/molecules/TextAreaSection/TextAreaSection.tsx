import { FC, ComponentProps } from "react";
import { TextArea } from "../../atoms";
import styles from "./style.module.css";

type TextAreaProps = ComponentProps<"textarea"> & {
  errorMessage?: string;
};

export const TextAreaSection: FC<TextAreaProps> = (props) => (
  <>
    <TextArea placeholder={"Content"} {...props} />
    {props?.errorMessage && (
      <p className={styles.error}>{props.errorMessage}</p>
    )}
  </>
);
